package pulumiakskubernetesprovider

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/azurecredential"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/containerservice"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithAddedClusterWithAzureCredentials returns kubernetes provider for the added AKS cluster based on the azure provider
func GetWithAddedClusterWithAzureCredentials(ctx *pulumi.Context,
	addedAksCluster *containerservice.KubernetesCluster,
	azureCredential *azurecredential.AzureCredential,
	dependencies []pulumi.Resource,
	providerName string) (*kubernetes.Provider, error) {

	clusterCaCert := addedAksCluster.KubeConfigs.ApplyT(
		func(kubeConfigs []containerservice.KubernetesClusterKubeConfig) string {
			return *kubeConfigs[0].ClusterCaCertificate
		})

	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(AzureExecPluginKubeConfigTemplate,
				addedAksCluster.Fqdn,
				clusterCaCert,
				azureCredential.Spec.ClientId,
				azureCredential.Spec.ClientSecret,
				azureCredential.Spec.TenantId,
			),
		}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
