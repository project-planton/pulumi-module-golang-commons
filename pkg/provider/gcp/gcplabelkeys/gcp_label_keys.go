package gcplabelkeys

import (
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/labels/labelkeys"
)

var (
	Resource     = labelkeys.WithNormalizedDomainPrefix("resource")
	Organization = labelkeys.WithNormalizedDomainPrefix("organization")
	Environment  = labelkeys.WithNormalizedDomainPrefix("environment")
	ResourceKind = labelkeys.WithNormalizedDomainPrefix("resource-kind")
	ResourceId   = labelkeys.WithNormalizedDomainPrefix("resource-id")
)
