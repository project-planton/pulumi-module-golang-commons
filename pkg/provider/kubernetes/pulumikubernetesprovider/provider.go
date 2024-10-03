package pulumikubernetesprovider

import (
	kubernetesclustercredentialv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/credential/kubernetesclustercredential/v1"
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigkekubernetesprovider"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithKubernetesClusterCredential returns kubernetes provider for the kubernetes cluster credential
func GetWithKubernetesClusterCredential(ctx *pulumi.Context,
	kubernetesClusterCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec,
	providerName string) (*kubernetes.Provider, error) {

	if kubernetesClusterCredentialSpec == nil {
		provider, err := kubernetes.NewProvider(ctx,
			providerName,
			&kubernetes.ProviderArgs{
				EnableServerSideApply: pulumi.Bool(true),
			})
		if err != nil {
			return nil, errors.Wrap(err, "failed to get new provider")
		}
		return provider, nil
	}

	kubeConfigString := ""

	if kubernetesClusterCredentialSpec.KubernetesProvider == kubernetesclustercredentialv1.KubernetesProvider_gcp_gke {
		c := kubernetesClusterCredentialSpec.GkeCluster

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
