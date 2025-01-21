package edge

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	gqlclient "github.com/pluralsh/console/go/client"
	"github.com/pluralsh/plural-cli/pkg/utils"
	"github.com/urfave/cli"
	"sigs.k8s.io/yaml"
)

const (
	cloudConfigURL     = "https://raw.githubusercontent.com/pluralsh/edge/main/cloud-config.yaml"
	pluralConfigURL    = "https://raw.githubusercontent.com/pluralsh/edge/main/plural-config.yaml"
	buildDir           = "build"
	cloudConfigFile    = "cloud-config.yaml"
	volumeName         = "edge-rootfs"
	wifiConfigTemplate = `
stages:
  boot:
    - name: Enable wireless
      commands:
        - connmanctl enable wifi
        - wpa_passphrase '@WIFI_SSID@' '@WIFI_PASSWORD@' > /etc/wpa_supplicant/wpa_supplicant.conf
        - wpa_supplicant -B -i wlan0 -c /etc/wpa_supplicant/wpa_supplicant.conf
        - udhcpc -i wlan0 &`
)

type Configuration struct {
	Image   string            `json:"image"`
	Bundles map[string]string `json:"bundles"`
}

func (p *Plural) handleEdgeImage(c *cli.Context) error {
	project := c.String("project")
	user := c.String("user")
	username := c.String("username")
	password := c.String("password")
	wifiSsid := c.String("wifi-ssid")
	wifiPassword := c.String("wifi-password")
	outputDir := c.String("output-dir")
	cloudConfig := c.String("cloud-config")
	pluralConfig := c.String("plural-config")

	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	utils.Highlight("creating bootstrap token for %s project\n", project)
	token, err := p.createBootstrapToken(project, user)
	if err != nil {
		return err
	}

	utils.Highlight("reading configuration\n")
	config, err := p.readConfig(pluralConfig)
	if err != nil {
		return err
	}

	utils.Highlight("preparing output directory\n")
	currentDir, err := os.Getwd()
	outputDirPath := filepath.Join(currentDir, outputDir)
	if err = os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		return err
	}

	buildDirPath := filepath.Join(outputDirPath, buildDir)
	if err = os.MkdirAll(buildDirPath, os.ModePerm); err != nil {
		return err
	}

	utils.Highlight("writing configuration\n")
	cloudConfigPath := filepath.Join(outputDirPath, cloudConfigFile)
	if err = p.writeCloudConfig(token.Token, username, password, wifiSsid, wifiPassword, cloudConfigPath, cloudConfig); err != nil {
		return err
	}

	utils.Highlight("preparing %s volume\n", volumeName)
	if err = utils.Exec("docker", "volume", "create", volumeName); err != nil {
		return err
	}
	defer func() {
		utils.Highlight("removing %s volume\n", volumeName)
		_ = utils.Exec("docker", "volume", "rm", volumeName)
	}()

	for bundle, image := range config.Bundles {
		utils.Highlight("writing %s bundle\n", bundle)
		if err = utils.Exec(
			"docker", "run", "-i", "--rm", "--user", "root", "--mount", "source=edge-rootfs,target=/rootfs",
			"gcr.io/go-containerregistry/crane:latest", "--platform=linux/arm64",
			"pull", image, fmt.Sprintf("/rootfs/%s.tar", bundle)); err != nil {
			return err
		}
	}

	utils.Highlight("unpacking image contents\n")
	if err = utils.Exec("docker", "run", "-i", "--rm", "--privileged",
		"--mount", "source=edge-rootfs,target=/rootfs", "quay.io/luet/base",
		"util", "unpack", config.Image, "/rootfs"); err != nil {
		return err
	}

	utils.Highlight("building image\n")
	if err = utils.Exec("docker", "run", "-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-v", buildDirPath+":/tmp/build",
		"-v", cloudConfigPath+":/cloud-config.yaml",
		"--mount", "source=edge-rootfs,target=/rootfs",
		"--privileged", "-i", "--rm",
		"--entrypoint=/build-arm-image.sh", "quay.io/kairos/auroraboot:v0.4.3",
		"--model", "rpi4",
		"--directory", "/rootfs",
		"--config", "/cloud-config.yaml", "/tmp/build/kairos.img"); err != nil {
		return err
	}

	utils.Success("image saved to %s directory\n", outputDir)
	return nil
}

func (p *Plural) createBootstrapToken(project, user string) (*gqlclient.BootstrapTokenBase, error) {
	attrributes := gqlclient.BootstrapTokenAttributes{}

	if user != "" {
		usr, err := p.ConsoleClient.GetUser(user)
		if err != nil {
			return nil, err
		}
		if usr == nil {
			return nil, fmt.Errorf("cannot find %s user", user)
		}
		attrributes.UserID = &usr.ID
	}

	proj, err := p.ConsoleClient.GetProject(project)
	if err != nil {
		return nil, err
	}
	if proj == nil {
		return nil, fmt.Errorf("cannot find %s project", project)
	}
	attrributes.ProjectID = proj.ID

	return p.ConsoleClient.CreateBootstrapToken(attrributes)
}

func (p *Plural) readConfig(override string) (*Configuration, error) {
	var content []byte
	var err error
	if override == "" {
		content, err = p.readDefaultConfig()
	} else {
		content, err = p.readFile(override)
	}
	if err != nil {
		return nil, err
	}

	var config *Configuration
	if err = yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func (p *Plural) readDefaultConfig() ([]byte, error) {
	response, err := http.Get(pluralConfigURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	buffer := new(bytes.Buffer)
	if _, err = buffer.ReadFrom(response.Body); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *Plural) readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func (p *Plural) writeCloudConfig(token, username, password, wifiSsid, wifiPassword, path, override string) error {
	if override != "" {
		config, err := os.ReadFile(override)
		if err != nil {
			return err
		}

		return os.WriteFile(path, config, os.ModePerm)
	}

	response, err := http.Get(cloudConfigURL)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	buffer := new(bytes.Buffer)
	if _, err = buffer.ReadFrom(response.Body); err != nil {
		return err
	}

	template := buffer.String()
	template = strings.ReplaceAll(template, "@USERNAME@", username)
	template = strings.ReplaceAll(template, "@PASSWORD@", password)
	template = strings.ReplaceAll(template, "@URL@", consoleURL)
	template = strings.ReplaceAll(template, "@TOKEN@", token)

	if wifiSsid != "" && wifiPassword != "" {
		wifiConfig := strings.ReplaceAll(wifiConfigTemplate, "@WIFI_SSID@", wifiSsid)
		wifiConfig = strings.ReplaceAll(wifiConfig, "@WIFI_PASSWORD@", wifiPassword)
		template += "\n" + wifiConfig
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(template)
	return err
}
