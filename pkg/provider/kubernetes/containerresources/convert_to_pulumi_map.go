package containerresources

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/kubernetes/model"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConvertToPulumiMap(resources *model.ContainerResources) pulumi.Map {
	return pulumi.Map{
		"limits": pulumi.Map{
			"cpu":    pulumi.String(resources.Limits.Cpu),
			"memory": pulumi.String(resources.Limits.Memory),
		},
		"requests": pulumi.Map{
			"cpu":    pulumi.String(resources.Requests.Cpu),
			"memory": pulumi.String(resources.Requests.Memory),
		},
	}
}
