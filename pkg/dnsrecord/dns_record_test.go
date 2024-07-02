package dnsrecord

import "testing"

func TestPulumiResourceNameForDnsRec(t *testing.T) {
	type input struct {
		recName string
		suffix  []string
	}
	testCases := []struct {
		name     string
		input    *input
		expected string
	}{
		{
			name: "* replaced with wildcard word",
			input: &input{
				recName: "*.dev.planton.cloud",
				suffix:  nil,
			},
			expected: "wildcard-dev-planton-cloud",
		},
		{
			name: "dot at the end should be ignore",
			input: &input{
				recName: "*.dev.planton.cloud.",
				suffix:  nil,
			},
			expected: "wildcard-dev-planton-cloud",
		},
		{
			name: "dot at the beginning should be ignore",
			input: &input{
				recName: ".*.dev.planton.cloud",
				suffix:  nil,
			},
			expected: "wildcard-dev-planton-cloud",
		},
		{
			name: "one suffix should be appended",
			input: &input{
				recName: "*.dev.planton.cloud",
				suffix:  []string{"a"},
			},
			expected: "wildcard-dev-planton-cloud-a",
		},
		{
			name: "multiple suffixes should be appended",
			input: &input{
				recName: "*.dev.planton.cloud",
				suffix:  []string{"a", "b"},
			},
			expected: "wildcard-dev-planton-cloud-a-b",
		},
		{
			name: "empty suffix should be ignored",
			input: &input{
				recName: "*.dev.planton.cloud",
				suffix:  []string{""},
			},
			expected: "wildcard-dev-planton-cloud",
		},
		{
			name: "empty record name should be ignored",
			input: &input{
				recName: "",
				suffix:  []string{"dev.planton.cloud"},
			},
			expected: "dev-planton-cloud",
		},
	}
	t.Run("test pulumi resource name for dns rec", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := PulumiResourceName(tc.input.recName, tc.input.suffix...)
				if result != tc.expected {
					t.Errorf("expected: %s got %s", tc.expected, result)
				}
			})
		}
	})
}
