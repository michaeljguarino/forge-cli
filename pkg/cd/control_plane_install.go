package cd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/osteele/liquid"
	"github.com/pluralsh/plural-cli/pkg/api"
	"github.com/pluralsh/plural-cli/pkg/bundle"
	"github.com/pluralsh/plural-cli/pkg/config"
	"github.com/pluralsh/plural-cli/pkg/crypto"
	"github.com/pluralsh/plural-cli/pkg/manifest"
	"github.com/pluralsh/plural-cli/pkg/provider"
	"github.com/pluralsh/plural-cli/pkg/template"
	"github.com/pluralsh/plural-cli/pkg/utils"
	"github.com/pluralsh/plural-cli/pkg/utils/git"
)

var (
	liquidEngine = liquid.NewEngine()
)

const (
	templateUrl = "https://raw.githubusercontent.com/pluralsh/console/master/templates/values.yaml.liquid"
	tplUrl      = "https://raw.githubusercontent.com/pluralsh/console/master/templates/values.yaml.tpl"
)

func ControlPlaneValues(conf config.Config, domain, dsn, name string) (string, error) {
	consoleDns := fmt.Sprintf("console.%s", domain)
	kasDns := fmt.Sprintf("kas.%s", domain)
	randoms := map[string]string{}
	for _, key := range []string{"jwt", "erlang", "adminPassword", "kasApi", "kasPrivateApi", "kasRedis"} {
		rand, err := crypto.RandStr(32)
		if err != nil {
			return "", err
		}
		randoms[key] = rand
	}

	client := api.FromConfig(&conf)
	me, err := client.Me()
	if err != nil {
		return "", fmt.Errorf("you must run `plural login` before installing")
	}

	root, err := git.Root()
	if err != nil {
		return "", err
	}

	project, err := manifest.ReadProject(filepath.Join(root, "workspace.yaml"))
	if err != nil {
		return "", err
	}

	prov, err := provider.FromManifest(project)
	if err != nil {
		return "", err
	}

	configuration := map[string]interface{}{
		"consoleDns":    consoleDns,
		"kasDns":        kasDns,
		"aesKey":        utils.GenAESKey(),
		"adminName":     me.Email,
		"adminEmail":    me.Email,
		"clusterName":   name,
		"pluralToken":   conf.Token,
		"postgresUrl":   dsn,
		"provider":      prov.Name(),
		"clusterIssuer": "plural",
	}
	for k, v := range randoms {
		configuration[k] = v
	}

	cryptos, err := cryptoVals()
	if err != nil {
		return "", err
	}

	for k, v := range cryptos {
		configuration[k] = v
	}

	clientId, clientSecret, err := ensureInstalledAndOidc(client, consoleDns)
	if err != nil {
		return "", err
	}
	configuration["pluralClientId"] = clientId
	configuration["pluralClientSecret"] = clientSecret

	tpl, err := fetchTemplate(tplUrl)
	if err != nil {
		return "", err
	}

	return template.RenderString(string(tpl), configuration)
}

func cryptoVals() (map[string]string, error) {
	res := make(map[string]string)
	keyFile, err := config.PluralDir("key")
	if err != nil {
		return res, err
	}

	aes, err := utils.ReadFile(keyFile)
	if err != nil {
		return res, err
	}
	res["key"] = aes

	identityFile, err := config.PluralDir("identity")
	if err != nil {
		return res, nil
	}

	identity, err := utils.ReadFile(identityFile)
	if err != nil {
		return res, nil
	}
	res["identity"] = identity
	return res, nil
}

func CreateControlPlane(conf config.Config) (string, error) {
	client := api.FromConfig(&conf)
	me, err := client.Me()
	if err != nil {
		return "", fmt.Errorf("you must run `plural login` before installing")
	}

	azureSurvey := []*survey.Question{
		{
			Name:   "console",
			Prompt: &survey.Input{Message: "Enter a dns name for your installation of the console (eg console.your.domain):"},
		},
		{
			Name:   "kubeProxy",
			Prompt: &survey.Input{Message: "Enter a dns name for the kube proxy (eg kas.your.domain), this is used for dashboarding functionality:"},
		},
		{
			Name:   "clusterName",
			Prompt: &survey.Input{Message: "Enter a name for this cluster:"},
		},
		{
			Name:   "postgresDsn",
			Prompt: &survey.Input{Message: "Enter a postgres connection string for the underlying database (should be postgres://<user>:<password>@<host>:5432/<database>):"},
		},
	}
	var resp struct {
		Console     string
		KubeProxy   string
		ClusterName string
		PostgresDsn string
	}
	if err := survey.Ask(azureSurvey, &resp); err != nil {
		return "", err
	}

	randoms := map[string]string{}
	for _, key := range []string{"jwt", "erlang", "adminPassword", "kasApi", "kasPrivateApi", "kasRedis"} {
		rand, err := crypto.RandStr(32)
		if err != nil {
			return "", err
		}
		randoms[key] = rand
	}

	configuration := map[string]string{
		"consoleDns":  resp.Console,
		"kasDns":      resp.KubeProxy,
		"aesKey":      utils.GenAESKey(),
		"adminName":   me.Email,
		"adminEmail":  me.Email,
		"clusterName": resp.ClusterName,
		"pluralToken": conf.Token,
		"postgresUrl": resp.PostgresDsn,
	}
	for k, v := range randoms {
		configuration[k] = v
	}

	clientId, clientSecret, err := ensureInstalledAndOidc(client, resp.Console)
	if err != nil {
		return "", err
	}
	configuration["pluralClientId"] = clientId
	configuration["pluralClientSecret"] = clientSecret

	tpl, err := fetchTemplate(templateUrl)
	if err != nil {
		return "", err
	}

	bindings := map[string]interface{}{
		"configuration": configuration,
	}

	res, err := liquidEngine.ParseAndRender(tpl, bindings)
	return string(res), err
}

func fetchTemplate(url string) (res []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, resp.Body)
	return out.Bytes(), err
}

func ensureInstalledAndOidc(client api.Client, dns string) (clientId string, clientSecret string, err error) {
	inst, err := client.GetInstallation("console")
	if err != nil || inst == nil {
		repo, err := client.GetRepository("console")
		if err != nil {
			return "", "", err
		}
		_, err = client.CreateInstallation(repo.Id)
		if err != nil {
			return "", "", err
		}
	}

	redirectUris := []string{fmt.Sprintf("https://%s/oauth/callback", dns)}
	err = bundle.SetupOIDC("console", client, redirectUris, "POST")
	if err != nil {
		return
	}

	inst, err = client.GetInstallation("console")
	if err != nil {
		return
	}

	return inst.OIDCProvider.ClientId, inst.OIDCProvider.ClientSecret, nil
}
