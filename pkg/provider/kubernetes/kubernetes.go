package kubernetes

import (
	"encoding/base64"
	"fmt"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	iacv1sjmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"

	"github.com/pkg/errors"
	base642 "github.com/plantoncloud-inc/go-commons/encoding/base64"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/provider/kubernetes/kubeconfig"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithStackCredentials returns kubernetes provider for the kubernetes credential in the stack credential
func GetWithStackCredentials(ctx *pulumi.Context, kubernetesProviderCredential *iacv1sjmodel.KubernetesProviderCredential, nameSuffixes ...string) (*kubernetes.Provider, error) {
	kubeconfigString, err := base64.StdEncoding.DecodeString(base642.CleanString(kubernetesProviderCredential.KubeconfigBase64))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 encoded google service account credential")
	}
	provider, err := kubernetes.NewProvider(ctx, getName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig:            pulumi.String(kubeconfigString),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

// GetWithAddedClusterWithGsaKey returns kubernetes provider for the added container cluster based on the google provider
// the provider creation should depend on the readiness of the node-pools
func GetWithAddedClusterWithGsaKey(ctx *pulumi.Context, serviceAccountKey *serviceaccount.Key, addedContainerCluster *container.Cluster,
	addedNodePools []pulumi.Resource, nameSuffixes ...string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx, getName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig: pulumi.Sprintf(kubeconfig.GoogleCredentialPluginKubeconfigTemplate,
			addedContainerCluster.Endpoint,
			addedContainerCluster.MasterAuth.ClusterCaCertificate().Elem(),
			serviceAccountKey.PrivateKey),
	}, pulumi.DependsOn(addedNodePools))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

// GetWithAddedClusterWithAwsCredentials returns kubernetes provider for the added eks cluster based on the aws provider
func GetWithAddedClusterWithAwsCredentials(ctx *pulumi.Context, addedEksCluster *eks.Cluster,
	awsProvider *awsclassic.Provider,
	dependencies []pulumi.Resource, nameSuffixes ...string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx, getName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig: pulumi.Sprintf(kubeconfig.AwsCredentialPluginKubeconfigTemplate,
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

func getName(suffixes []string) string {
	name := englishword.EnglishWord_kubernetes.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}
