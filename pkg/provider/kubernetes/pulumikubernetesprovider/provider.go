package pulumikubernetesprovider

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/kubernetesclustercredential"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/kubernetesclustercredential/enums/kubernetesprovider"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigkekubernetesprovider"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithKubernetesClusterCredential returns kubernetes provider for the kubernetes cluster credential
func GetWithKubernetesClusterCredential(ctx *pulumi.Context,
	kubernetesClusterCredentialSpec *kubernetesclustercredential.KubernetesClusterCredentialSpec,
	providerName string) (*kubernetes.Provider, error) {

	kubeConfigString := ""

	if kubernetesClusterCredentialSpec.KubernetesProvider == kubernetesprovider.KubernetesProvider_gcp_gke {
		c := kubernetesClusterCredentialSpec.GkeClusterSpec

		kubeConfigString = fmt.Sprintf(pulumigkekubernetesprovider.GcpExecPluginKubeConfigTemplate,
			c.ClusterEndpoint,
			c.ClusterCaData,
			pulumigkekubernetesprovider.GcpExecPluginPath,
			c.ServiceAccountKeyBase64)
	}

	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig:            pulumi.String(kubeConfigString),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
