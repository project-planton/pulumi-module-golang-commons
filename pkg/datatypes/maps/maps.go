package maps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"sort"
)

func SortMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

// ConvertToPulumiMap takes string map and converts it into pulumi.Map
func ConvertToPulumiMap(input map[string]string) pulumi.Map {
	resp := pulumi.Map{}
	for k, v := range input {
		resp[k] = pulumi.String(v)
	}
	return resp
}
