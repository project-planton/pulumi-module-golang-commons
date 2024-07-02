package pulumiekskubernetesprovider

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithAddedClusterWithAwsCredentials returns kubernetes provider for the added eks cluster based on the aws provider
func GetWithAddedClusterWithAwsCredentials(ctx *pulumi.Context, addedEksCluster *eks.Cluster,
	awsProvider *awsclassic.Provider,
	dependencies []pulumi.Resource, nameSuffixes ...string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx, pulumikubernetesprovider.GetPulumiResourceName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig: pulumi.Sprintf(AwsCredentialPluginKubeConfigTemplate,
			addedEksCluster.Endpoint,
			addedEksCluster.CertificateAuthority.Data().Elem(),
			awsProvider.AccessKey.Elem(),
			awsProvider.SecretKey.Elem(),
			awsProvider.Region.Elem(),
		),
	}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
