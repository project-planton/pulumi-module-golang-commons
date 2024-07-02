package pulumikubernetesprovider

import (
	"encoding/base64"
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	iacv1sjmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"

	"github.com/pkg/errors"
	base642 "github.com/plantoncloud-inc/go-commons/encoding/base64"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithStackCredentials returns kubernetes provider for the kubernetes credential in the stack credential
func GetWithStackCredentials(ctx *pulumi.Context,
	kubernetesProviderCredential *iacv1sjmodel.KubernetesProviderCredential,
	nameSuffixes ...string) (*kubernetes.Provider, error) {
	kubeConfigString, err := base64.StdEncoding.DecodeString(
		base642.CleanString(kubernetesProviderCredential.KubeconfigBase64))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 encoded kube-config")
	}
	provider, err := kubernetes.NewProvider(ctx, GetPulumiResourceName(nameSuffixes), &kubernetes.ProviderArgs{
		EnableServerSideApply: pulumi.Bool(true),
		Kubeconfig:            pulumi.String(kubeConfigString),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

func GetPulumiResourceName(suffixes []string) string {
	name := englishword.EnglishWord_kubernetes.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}
