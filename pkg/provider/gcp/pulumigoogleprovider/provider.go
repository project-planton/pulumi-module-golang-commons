package pulumigoogleprovider

import (
	"encoding/base64"
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/gcpcredential"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/pulumi/pulumioutput"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Get(ctx *pulumi.Context, gcpCredentialSpec *gcpcredential.GcpCredentialSpec, nameSuffixes ...string) (*gcp.Provider, error) {
	serviceAccountKey, err := base64.StdEncoding.DecodeString(gcpCredentialSpec.ServiceAccountKeyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 encoded"+
			" google service account credential")
	}
	provider, err := gcp.NewProvider(ctx, ProviderResourceName(nameSuffixes), &gcp.ProviderArgs{
		Credentials: pulumi.String(serviceAccountKey),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

func ProviderResourceName(suffixes []string) string {
	name := englishword.EnglishWord_google.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	outputName := fmt.Sprintf("gcp_%s", pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		outputName = fmt.Sprintf("%s_%s", outputName, s)
	}
	return outputName
}
