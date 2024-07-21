package project

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud-inc/go-commons/cloud/gcp/apis"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/gcp/googleapi"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	BillProjectApis = []apis.Api{
		apis.BigQuery,
	}
)

func AddApi(ctx *pulumi.Context, projectName string, project *organizations.Project, apis []apis.Api) ([]pulumi.Resource, error) {
	resp := make([]pulumi.Resource, 0)
	for _, api := range apis {
		addedProjectService, err := projects.NewService(ctx,
			googleapi.PulumiResourceName(projectName, api),
			&projects.ServiceArgs{
				Project:                  project.ProjectId,
				DisableDependentServices: pulumi.BoolPtr(true),
				Service:                  pulumi.String(api),
			}, pulumi.Parent(project))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to enable %s api for %s project", api, projectName)
		}
		resp = append(resp, addedProjectService)
	}
	return resp, nil
}
