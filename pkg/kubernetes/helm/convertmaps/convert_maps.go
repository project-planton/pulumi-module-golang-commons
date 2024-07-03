package convertmaps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ConvertGoMapToPulumiMap converts a Golang map to a Pulumi map.
func ConvertGoMapToPulumiMap(goMap map[string]string) pulumi.Map {
	pulumiMap := make(pulumi.Map)
	for k, v := range goMap {
		pulumiMap[k] = pulumi.String(v)
	}
	return pulumiMap
}
