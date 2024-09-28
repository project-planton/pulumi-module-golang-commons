package awstagkeys

import (
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/labels/labelkeys"
	"strings"
)

var (
	Resource     = labelkeys.WithDomainPrefix(englishword.EnglishWord_resource.String())
	Organization = labelkeys.WithDomainPrefix(englishword.EnglishWord_organization.String())
	Environment  = labelkeys.WithDomainPrefix(englishword.EnglishWord_environment.String())
	ResourceKind = labelkeys.WithDomainPrefix(replaceUnderscoresWithHyphens(
		englishword.EnglishWord_resource_kind.String()))
	ResourceId = labelkeys.WithDomainPrefix(replaceUnderscoresWithHyphens(
		englishword.EnglishWord_resource_id.String()))
)

func replaceUnderscoresWithHyphens(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}
