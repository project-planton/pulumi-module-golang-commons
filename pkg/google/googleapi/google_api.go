package googleapi

import (
	"fmt"
	"github.com/plantoncloud-inc/go-commons/cloud/gcp/apis"
)

func PulumiResourceName(projectName string, api apis.Api) string {
	return fmt.Sprintf("%s-%s", projectName, api)
}
