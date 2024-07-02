package aws

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	iacv1sjmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNative(ctx *pulumi.Context, awsProviderCredential *iacv1sjmodel.AwsProviderCredential,
	region string, nameSuffixes ...string) (*aws.Provider, error) {
	awsNative, err := aws.NewProvider(ctx, getName(nameSuffixes), &aws.ProviderArgs{
		AccessKey: pulumi.String(awsProviderCredential.AwsAccessKeyId),
		SecretKey: pulumi.String(awsProviderCredential.AwsSecretAccessKey),
		Region:    pulumi.String(region),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws provider")
	}
	return awsNative, nil
}

func GetClassic(ctx *pulumi.Context, awsProviderCredential *iacv1sjmodel.AwsProviderCredential,
	region string, nameSuffixes ...string) (*awsclassic.Provider, error) {

	awsClassic, err := awsclassic.NewProvider(ctx, getName(nameSuffixes), &awsclassic.ProviderArgs{
		AccessKey: pulumi.String(awsProviderCredential.AwsAccessKeyId),
		SecretKey: pulumi.String(awsProviderCredential.AwsSecretAccessKey),
		Region:    pulumi.String(region),
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws classic provider")
	}
	return awsClassic, nil
}

func getName(suffixes []string) string {
	name := englishword.EnglishWord_aws.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}
