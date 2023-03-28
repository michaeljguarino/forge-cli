import { ApolloClient } from '@apollo/client'

import {
  AppsIcon,
  InstallIcon,
  WizardInstaller,
  WizardPicker,
  WizardStepConfig,
} from '@pluralsh/design-system'

import {
  CreateQuickStackDocument,
  Datatype,
  GetRecipeDocument,
  InstallStackShellDocument,
  ListRecipesDocument,
  ListRepositoriesDocument,
  Provider,
  Recipe,
  RecipeSection,
  RootQueryType,
} from '../../graphql/generated/graphql'

import { Application } from './Application'

const toPickerItems = (applications: Array<any>, provider: Provider, forcedApps: any): Array<WizardStepConfig> => applications?.map(app => ({
  key: app.id,
  label: app.name,
  imageUrl: app.icon,
  node: <Application
    key={app.id}
    provider={provider}
  />,
  isRequired: Object.keys(forcedApps).includes(app.name),
  tooltip: forcedApps[app.name],
})) || []

const toDefaultSteps = (applications: any, provider: Provider, forcedApps: any): Array<WizardStepConfig> => [{
  key: 'apps',
  label: 'Apps',
  Icon: AppsIcon,
  node: <WizardPicker items={toPickerItems(applications, provider, forcedApps)} />,
  isDefault: true,
},
{
  key: 'placeholder',
  isPlaceholder: true,
},
{
  key: 'install',
  label: 'Install',
  Icon: InstallIcon,
  node: <WizardInstaller />,
  isDefault: true,
}]

const toDependencySteps = (applications: {section: RecipeSection, dependencyOf: Set<string>}[], provider: Provider): Array<WizardStepConfig> => [...applications.map(app => ({
  key: app.section.repository!.id,
  label: app.section.repository!.name,
  imageUrl: app.section.repository!.icon!,
  node: <Application
    key={app.section.repository!.id}
    provider={provider}
  />,
  isDependency: true,
  dependencyOf: app.dependencyOf,
}))]

const buildSteps = async (client: ApolloClient<unknown>, provider: Provider, selectedApplications: Array<WizardStepConfig>) => {
  const dependencyMap = new Map<string, {section: RecipeSection, dependencyOf: Set<string>}>()

  for (const app of selectedApplications) {
    const { data: { recipes } = {} } = await client.query<Pick<RootQueryType, 'recipes'>>({
      query: ListRecipesDocument,
      variables: { repositoryId: app.key },
    })

    const { node: recipeBase } = recipes?.edges?.find(edge => edge!.node!.provider === provider) || { node: undefined }

    if (!recipeBase) continue

    const { data: recipe } = await client.query<{recipe: Recipe}>({
      query: GetRecipeDocument,
      variables: { id: recipeBase?.id },
    })

    const sections = recipe.recipe.recipeSections!.filter(section => section!.repository!.name !== app.label)

    sections.forEach(section => {
      if (selectedApplications.find(app => app.key === section!.repository!.id)) return

      if (!dependencyMap.has(section!.repository!.name)) {
        dependencyMap.set(section!.repository!.name, { section: section!, dependencyOf: new Set([app.label!]) })

        return
      }

      const dep = dependencyMap.get(section!.repository!.name)!
      const dependencyOf: Array<string> = [...Array.from(dep.dependencyOf.values()), app.label!]

      dependencyMap.set(section!.repository!.name, { section: section!, dependencyOf: new Set<string>(dependencyOf) })
    })
  }

  return toDependencySteps(Array.from(dependencyMap.values()), provider)
}

const install = async (client: ApolloClient<unknown>, apps: Array<WizardStepConfig<any>>, provider: Provider) => {
  const toAPIContext = (context: any) => ({ ...Object.keys(context || {}).reduce((acc, key) => ({ ...acc, [key]: context[key].value }), {}) })
  const toDataTypeValues = (context: any, datatype: any) => Object.keys(context || {}).reduce((acc: Array<any>, key) => (context[key].type === datatype ? [...acc, context[key].value] : [...acc]), [])

  const { data: { quickStack } } = await client.mutate({
    mutation: CreateQuickStackDocument,
    variables: { applicationIds: apps.filter(app => !app.isDependency).map(app => app.key), provider },
  })

  const configuration = apps.reduce((acc, app) => ({ ...acc, [app.label!]: toAPIContext(app.data?.context || {}) }), {})
  const domains = apps.reduce((acc: Array<any>, app) => [...acc, ...toDataTypeValues(app.data?.context || {}, Datatype.Domain)], [])
  const buckets = apps.reduce((acc: Array<any>, app) => [...acc, ...toDataTypeValues(app.data?.context || {}, Datatype.Bucket)], [])

  return client.mutate({
    mutation: InstallStackShellDocument,
    variables: { name: quickStack.name, oidc: true, context: { configuration: JSON.stringify(configuration), domains, buckets } },
    refetchQueries: [{ query: ListRepositoriesDocument, variables: { provider } }],
  })
}

export {
  toDependencySteps, toDefaultSteps, buildSteps, toPickerItems, install,
}
