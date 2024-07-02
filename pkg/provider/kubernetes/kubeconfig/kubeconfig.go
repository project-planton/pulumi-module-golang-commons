package kubeconfig

// GoogleCredentialPluginKubeconfigTemplate requires the following inputs for rendering a kubeconfig that works
// 1. cluster endpoint ip
// 2. cluster cert-authority data
// 3. base64 encoded google service account key
const GoogleCredentialPluginKubeconfigTemplate = `
apiVersion: v1
kind: Config
current-context: kube-context
contexts: [{name: kube-context, context: {cluster: kube-cluster, user: kube-user}}]
clusters:
- name: kube-cluster
  cluster:
    server: https://%s
    certificate-authority-data: %s
users:
- name: kube-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1
      interactiveMode: Never
      command: /usr/local/bin/kube-client-go-google-credential-plugin
      args:
       - %s
`

// AwsCredentialPluginKubeconfigTemplate requires the following inputs for rendering a kubeconfig that works
// 1. cluster endpoint ip
// 2. cluster cert-authority data
// 3. aws acces key id
// 4. aws secret acces key
// 5. aws region
const AwsCredentialPluginKubeconfigTemplate = `
apiVersion: v1
kind: Config
current-context: kube-context
contexts: [{name: kube-context, context: {cluster: kube-cluster, user: kube-user}}]
clusters:
- name: kube-cluster
  cluster:
    server: %s
    certificate-authority-data: %s
users:
- name: kube-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1
      interactiveMode: Never
      command: /usr/local/bin/kube-client-go-aws-credential-plugin
      args:
       - %s
       - %s
       - %s
`
