package pulumikubernetesprovider

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/kubernetesclustercredential/model"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/pulumi/pulumioutput"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	_ExecPluginPath = "/usr/local/bin/kube-client-go-google-credential-plugin"
)

// GetWithKubernetesClusterCredential returns kubernetes provider for the kubernetes cluster credential
func GetWithKubernetesClusterCredential(ctx *pulumi.Context,
	kubernetesClusterCredential *model.KubernetesClusterCredential,
	nameSuffixes ...string) (*kubernetes.Provider, error) {

	kubeConfigString := buildGkeKubeConfigWithCredentialPlugin(kubernetesClusterCredential)

	provider, err := kubernetes.NewProvider(ctx, ProviderResourceName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig:            pulumi.String(kubeConfigString),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

// buildGkeKubeConfigWithCredentialPlugin generates a base64 encoded kubeconfig from the GKE cluster details.
func buildGkeKubeConfigWithCredentialPlugin(kubernetesClusterCredential *model.KubernetesClusterCredential) string {
	kubeconfigFormatString := `apiVersion: v1
kind: Config
current-context: kube-context
contexts:
- name: kube-context
  context: {cluster: gke-cluster, user: kube-user}
clusters:
- name: gke-cluster
  cluster:
    server: https://%s
    certificate-authority-data: %s
users:
- name: kube-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1
      interactiveMode: Never
      command: %s
      args:
        - %s
`
	gkeClusterSpec := kubernetesClusterCredential.Spec.GkeClusterSpec
	return fmt.Sprintf(
		kubeconfigFormatString,
		gkeClusterSpec.ClusterEndpoint,
		gkeClusterSpec.ClusterCaData,
		_ExecPluginPath,
		gkeClusterSpec.ServiceAccountKeyBase64,
	)
}

func ProviderResourceName(suffixes []string) string {
	name := englishword.EnglishWord_kubernetes.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	outputName := fmt.Sprintf("%s_%s",
		englishword.EnglishWord_kubernetes.String(),
		pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		outputName = fmt.Sprintf("%s_%s", outputName, s)
	}
	return outputName
}
