package pulumigkekubernetesprovider

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithAddedClusterWithGsaKey returns kubernetes provider for the added container cluster based on the google provider
// the provider creation should depend on the readiness of the node-pools
func GetWithAddedClusterWithGsaKey(ctx *pulumi.Context, serviceAccountKey *serviceaccount.Key, addedContainerCluster *container.Cluster,
	addedNodePools []pulumi.Resource, nameSuffixes ...string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx, pulumikubernetesprovider.ProviderResourceName(nameSuffixes),
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(GoogleCredentialPluginKubeConfigTemplate,
				addedContainerCluster.Endpoint,
				addedContainerCluster.MasterAuth.ClusterCaCertificate().Elem(),
				serviceAccountKey.PrivateKey),
		}, pulumi.DependsOn(addedNodePools))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
