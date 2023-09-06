package bootstrap

const storageClassManifest = `
apiVersion: v1
kind: Namespace
metadata:
  name: local-path-storage

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: local-path-provisioner-service-account
  namespace: local-path-storage

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: local-path-provisioner-role
rules:
  - apiGroups: [ "" ]
    resources: [ "nodes", "persistentvolumeclaims", "configmaps" ]
    verbs: [ "get", "list", "watch" ]
  - apiGroups: [ "" ]
    resources: [ "endpoints", "persistentvolumes", "pods" ]
    verbs: [ "*" ]
  - apiGroups: [ "" ]
    resources: [ "events" ]
    verbs: [ "create", "patch" ]
  - apiGroups: [ "storage.k8s.io" ]
    resources: [ "storageclasses" ]
    verbs: [ "get", "list", "watch" ]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: local-path-provisioner-bind
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: local-path-provisioner-role
subjects:
  - kind: ServiceAccount
    name: local-path-provisioner-service-account
    namespace: local-path-storage

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: local-path-provisioner
  namespace: local-path-storage
spec:
  replicas: 1
  selector:
    matchLabels:
      app: local-path-provisioner
  template:
    metadata:
      labels:
        app: local-path-provisioner
    spec:
      serviceAccountName: local-path-provisioner-service-account
      containers:
        - name: local-path-provisioner
          image: rancher/local-path-provisioner:v0.0.24
          imagePullPolicy: IfNotPresent
          command:
            - local-path-provisioner
            - --debug
            - start
            - --config
            - /etc/config/config.json
          volumeMounts:
            - name: config-volume
              mountPath: /etc/config/
          env:
            - name: POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
      volumes:
        - name: config-volume
          configMap:
            name: local-path-config
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
  name: standard
provisioner: rancher.io/local-path
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: local-path-config
  namespace: local-path-storage
data:
  config.json: |-
    {
            "nodePathMap":[
            {
                    "node":"DEFAULT_PATH_FOR_NON_LISTED_NODES",
                    "paths":["/opt/local-path-provisioner"]
            }
            ]
    }
  setup: |-
    #!/bin/sh
    set -eu
    mkdir -m 0777 -p "$VOL_DIR"
  teardown: |-
    #!/bin/sh
    set -eu
    rm -rf "$VOL_DIR"
  helperPod.yaml: |-
    apiVersion: v1
    kind: Pod
    metadata:
      name: helper-pod
    spec:
      containers:
      - name: helper-pod
        image: busybox
        imagePullPolicy: IfNotPresent
`

const azureIdentityManifest = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: unapproved
    controller-gen.kubebuilder.io/version: v0.5.0
    meta.helm.sh/release-name: bootstrap
    meta.helm.sh/release-namespace: bootstrap
  name: azureassignedidentities.aadpodidentity.k8s.io
  labels:
    app.kubernetes.io/name: aad-pod-identity
    app.kubernetes.io/instance: aad-pod-identity
    app.kubernetes.io/managed-by: Helm
    helm.sh/chart: aad-pod-identity
spec:
  group: aadpodidentity.k8s.io
  names:
    kind: AzureAssignedIdentity
    listKind: AzureAssignedIdentityList
    plural: azureassignedidentities
    singular: azureassignedidentity
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: AzureAssignedIdentity contains the identity <-> pod mapping which is matched.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AzureAssignedIdentitySpec contains the relationship between an AzureIdentity and an AzureIdentityBinding.
            properties:
              azureBindingRef:
                description: AzureBindingRef is an embedded resource referencing the AzureIdentityBinding used by the AzureAssignedIdentity, which requires x-kubernetes-embedded-resource fields to be true
                properties:
                  apiVersion:
                    description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
                    type: string
                  kind:
                    description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  metadata:
                    type: object
                  spec:
                    description: AzureIdentityBindingSpec matches the pod with the Identity. Used to indicate the potential matches to look for between the pod/deployment and the identities present.
                    properties:
                      azureIdentity:
                        type: string
                      metadata:
                        type: object
                      selector:
                        type: string
                      weight:
                        description: Weight is used to figure out which of the matching identities would be selected.
                        type: integer
                    type: object
                  status:
                    description: AzureIdentityBindingStatus contains the status of an AzureIdentityBinding.
                    properties:
                      availableReplicas:
                        format: int32
                        type: integer
                      metadata:
                        type: object
                    type: object
                type: object
                x-kubernetes-embedded-resource: true
              azureIdentityRef:
                description: AzureIdentityRef is an embedded resource referencing the AzureIdentity used by the AzureAssignedIdentity, which requires x-kubernetes-embedded-resource fields to be true
                properties:
                  apiVersion:
                    description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
                    type: string
                  kind:
                    description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  metadata:
                    type: object
                  spec:
                    description: AzureIdentitySpec describes the credential specifications of an identity on Azure.
                    properties:
                      adEndpoint:
                        type: string
                      adResourceID:
                        description: For service principal. Option param for specifying the  AD details.
                        type: string
                      auxiliaryTenantIDs:
                        description: Service principal auxiliary tenant ids
                        items:
                          type: string
                        nullable: true
                        type: array
                      clientID:
                        description: Both User Assigned MSI and SP can use this field.
                        type: string
                      clientPassword:
                        description: Used for service principal
                        properties:
                          name:
                            description: Name is unique within a namespace to reference a secret resource.
                            type: string
                          namespace:
                            description: Namespace defines the space within which the secret name must be unique.
                            type: string
                        type: object
                      metadata:
                        type: object
                      replicas:
                        format: int32
                        nullable: true
                        type: integer
                      resourceID:
                        description: User assigned MSI resource id.
                        type: string
                      tenantID:
                        description: Service principal primary tenant id.
                        type: string
                      type:
                        description: UserAssignedMSI or Service Principal
                        type: integer
                    type: object
                  status:
                    description: AzureIdentityStatus contains the replica status of the resource.
                    properties:
                      availableReplicas:
                        format: int32
                        type: integer
                      metadata:
                        type: object
                    type: object
                type: object
                x-kubernetes-embedded-resource: true
              metadata:
                type: object
              nodename:
                type: string
              pod:
                type: string
              podNamespace:
                type: string
              replicas:
                format: int32
                nullable: true
                type: integer
            type: object
          status:
            description: AzureAssignedIdentityStatus contains the replica status of the resource.
            properties:
              availableReplicas:
                format: int32
                type: integer
              metadata:
                type: object
              status:
                type: string
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: unapproved
    controller-gen.kubebuilder.io/version: v0.5.0
    meta.helm.sh/release-name: bootstrap
    meta.helm.sh/release-namespace: bootstrap
  name: azureidentities.aadpodidentity.k8s.io
  labels:
    app.kubernetes.io/name: aad-pod-identity
    app.kubernetes.io/instance: aad-pod-identity
    app.kubernetes.io/managed-by: Helm
    helm.sh/chart: aad-pod-identity
spec:
  group: aadpodidentity.k8s.io
  names:
    kind: AzureIdentity
    listKind: AzureIdentityList
    plural: azureidentities
    singular: azureidentity
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.type
      name: Type
      type: string
    - jsonPath: .spec.clientID
      name: ClientID
      type: string
    - description: CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
      jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: AzureIdentity is the specification of the identity data structure.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AzureIdentitySpec describes the credential specifications of an identity on Azure.
            properties:
              adEndpoint:
                type: string
              adResourceID:
                description: For service principal. Option param for specifying the  AD details.
                type: string
              auxiliaryTenantIDs:
                description: Service principal auxiliary tenant ids
                items:
                  type: string
                nullable: true
                type: array
              clientID:
                description: Both User Assigned MSI and SP can use this field.
                type: string
              clientPassword:
                description: Used for service principal
                properties:
                  name:
                    description: Name is unique within a namespace to reference a secret resource.
                    type: string
                  namespace:
                    description: Namespace defines the space within which the secret name must be unique.
                    type: string
                type: object
              metadata:
                type: object
              replicas:
                format: int32
                nullable: true
                type: integer
              resourceID:
                description: User assigned MSI resource id.
                type: string
              tenantID:
                description: Service principal primary tenant id.
                type: string
              type:
                description: UserAssignedMSI or Service Principal
                type: integer
            type: object
          status:
            description: AzureIdentityStatus contains the replica status of the resource.
            properties:
              availableReplicas:
                format: int32
                type: integer
              metadata:
                type: object
            type: object
        type: object
    served: true
    storage: true
    subresources: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: unapproved
    controller-gen.kubebuilder.io/version: v0.5.0
    meta.helm.sh/release-name: bootstrap
    meta.helm.sh/release-namespace: bootstrap
  name: azureidentitybindings.aadpodidentity.k8s.io
  labels:
    app.kubernetes.io/name: aad-pod-identity
    app.kubernetes.io/instance: aad-pod-identity
    app.kubernetes.io/managed-by: Helm
    helm.sh/chart: aad-pod-identity
spec:
  group: aadpodidentity.k8s.io
  names:
    kind: AzureIdentityBinding
    listKind: AzureIdentityBindingList
    plural: azureidentitybindings
    singular: azureidentitybinding
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.azureIdentity
      name: AzureIdentity
      type: string
    - jsonPath: .spec.selector
      name: Selector
      type: string
    - description: CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
      jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: AzureIdentityBinding brings together the spec of matching pods and the identity which they can use.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AzureIdentityBindingSpec matches the pod with the Identity. Used to indicate the potential matches to look for between the pod/deployment and the identities present.
            properties:
              azureIdentity:
                type: string
              metadata:
                type: object
              selector:
                type: string
              weight:
                description: Weight is used to figure out which of the matching identities would be selected.
                type: integer
            type: object
          status:
            description: AzureIdentityBindingStatus contains the status of an AzureIdentityBinding.
            properties:
              availableReplicas:
                format: int32
                type: integer
              metadata:
                type: object
            type: object
        type: object
    served: true
    storage: true
    subresources: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: unapproved
    controller-gen.kubebuilder.io/version: v0.5.0
    meta.helm.sh/release-name: bootstrap
    meta.helm.sh/release-namespace: bootstrap
  name: azurepodidentityexceptions.aadpodidentity.k8s.io
  labels:
    app.kubernetes.io/name: aad-pod-identity
    app.kubernetes.io/instance: aad-pod-identity
    app.kubernetes.io/managed-by: Helm
    helm.sh/chart: aad-pod-identity
spec:
  group: aadpodidentity.k8s.io
  names:
    kind: AzurePodIdentityException
    listKind: AzurePodIdentityExceptionList
    plural: azurepodidentityexceptions
    singular: azurepodidentityexception
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: AzurePodIdentityException contains the pod selectors for all pods that don't require NMI to process and request token on their behalf.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
            properties:
              metadata:
                type: object
              podLabels:
                additionalProperties:
                  type: string
                type: object
            type: object
          status:
            description: AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
            properties:
              metadata:
                type: object
              status:
                type: string
            type: object
        type: object
    served: true
    storage: true
`
