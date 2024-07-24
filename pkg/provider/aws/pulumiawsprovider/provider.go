package pulumiawsprovider

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/awscredential/model"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/pulumi/pulumioutput"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"reflect"
)

func GetNative(ctx *pulumi.Context, awsCredential *model.AwsCredential,
	region string, nameSuffixes ...string) (*aws.Provider, error) {
	awsNative, err := aws.NewProvider(ctx, ProviderResourceName(nameSuffixes), &aws.ProviderArgs{
		AccessKey: pulumi.String(awsCredential.Spec.AccessKeyId),
		SecretKey: pulumi.String(awsCredential.Spec.SecretAccessKey),
		Region:    pulumi.String(region),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws provider")
	}
	return awsNative, nil
}

func GetClassic(ctx *pulumi.Context, awsCredential *model.AwsCredential,
	region string, nameSuffixes ...string) (*awsclassic.Provider, error) {

	awsClassic, err := awsclassic.NewProvider(ctx, ProviderResourceName(nameSuffixes), &awsclassic.ProviderArgs{
		AccessKey: pulumi.String(awsCredential.Spec.AccessKeyId),
		SecretKey: pulumi.String(awsCredential.Spec.SecretAccessKey),
		Region:    pulumi.String(region),
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws classic provider")
	}
	return awsClassic, nil
}

func ProviderResourceName(suffixes []string) string {
	name := englishword.EnglishWord_aws.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	outputName := fmt.Sprintf("%s_%s", englishword.EnglishWord_aws.String(),
		pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		outputName = fmt.Sprintf("%s_%s", outputName, s)
	}
	return outputName
}
