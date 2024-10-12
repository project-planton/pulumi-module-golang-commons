package awstagkeys

import (
	"github.com/project-planton/pulumi-module-golang-commons/pkg/labels/labelkeys"
)

var (
	Resource     = labelkeys.WithDomainPrefix("resource")
	Organization = labelkeys.WithDomainPrefix("organization")
	Environment  = labelkeys.WithDomainPrefix("environment")
	ResourceKind = labelkeys.WithDomainPrefix("resource-kind")
	ResourceId   = labelkeys.WithDomainPrefix("resource-id")
)
