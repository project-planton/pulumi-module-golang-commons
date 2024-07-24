package pulumiazureprovider

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/azurecredential/model"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Get(ctx *pulumi.Context, azureCredential *model.AzureCredential,
	nameSuffixes ...string) (*azure.Provider, error) {
	awsNative, err := azure.NewProvider(ctx, ProviderResourceName(nameSuffixes),
		&azure.ProviderArgs{
			ClientId:       pulumi.String(azureCredential.Spec.ClientId),
			ClientSecret:   pulumi.String(azureCredential.Spec.ClientSecret),
			SubscriptionId: pulumi.String(azureCredential.Spec.SubscriptionId),
			TenantId:       pulumi.String(azureCredential.Spec.TenantId),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create azure provider")
	}
	return awsNative, nil
}

func ProviderResourceName(suffixes []string) string {
	name := englishword.EnglishWord_azure.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}
