package google

import (
	"encoding/base64"
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	iacv1sjmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"

	"github.com/pkg/errors"
	base642 "github.com/plantoncloud-inc/go-commons/encoding/base64"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Get(ctx *pulumi.Context, googleProviderCredential *iacv1sjmodel.GoogleProviderCredential, nameSuffixes ...string) (*gcp.Provider, error) {
	gsaCredentialString, err := base64.StdEncoding.DecodeString(base642.CleanString(googleProviderCredential.ServiceAccountKeyBase64))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 encoded google service account credential")
	}
	provider, err := gcp.NewProvider(ctx, getName(nameSuffixes), &gcp.ProviderArgs{
		Credentials: pulumi.String(gsaCredentialString),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

func getName(suffixes []string) string {
	name := englishword.EnglishWord_google.String()
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}
