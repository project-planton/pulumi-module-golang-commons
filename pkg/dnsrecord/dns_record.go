package dnsrecord

import (
	"strings"
)

func PulumiResourceName(recName string, suffix ...string) string {
	recName = strings.TrimSuffix(recName, ".")
	recName = strings.TrimPrefix(recName, ".")
	var respBuilder strings.Builder
	if recName != "" {
		respBuilder.WriteString(recName)
	}
	if recName != "" && suffix != nil && len(suffix) > 0 && suffix[0] != "" {
		respBuilder.WriteString(".")
	}
	for i, s := range suffix {
		if s == "" {
			continue
		}
		respBuilder.WriteString(strings.ToLower(s))
		if i != len(suffix)-1 {
			respBuilder.WriteString(".")
		}
	}
	r := respBuilder.String()
	r = strings.ReplaceAll(r, "*", "wildcard")
	r = strings.ReplaceAll(r, ".", "-")
	return r
}
