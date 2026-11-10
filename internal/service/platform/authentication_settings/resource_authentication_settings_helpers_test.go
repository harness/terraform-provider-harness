package authentication_settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeWhitelistedDomain(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"plain domain", "example.com", "example.com"},
		{"www prefix", "www.example.com", "example.com"},
		{"http scheme", "http://example.com", "example.com"},
		{"https scheme", "https://example.com", "example.com"},
		{"http scheme with www", "http://www.example.com", "example.com"},
		{"https scheme with www", "https://www.example.com", "example.com"},
		{"uppercase", "EXAMPLE.COM", "example.com"},
		{"uppercase www", "WWW.EXAMPLE.COM", "example.com"},
		{"surrounding whitespace", "  example.com  ", "example.com"},
		{"trailing slash", "example.com/", "example.com"},
		{"scheme, www and trailing slash", "https://www.example.com/", "example.com"},
		{"subdomain preserved", "sub.example.co.uk", "sub.example.co.uk"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, normalizeWhitelistedDomain(tc.input))
		})
	}
}

func TestValidateWhitelistedDomain(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"plain domain", "example.com", false},
		{"www prefix allowed", "www.example.com", false},
		{"subdomain", "sub.example.co.uk", false},
		{"hyphenated domain", "example-domain.com", false},
		{"http scheme rejected", "http://example.com", true},
		{"https scheme rejected", "https://example.com", true},
		{"https scheme with www rejected", "https://www.example.com", true},
		{"non-http scheme rejected", "ftp://example.com", true},
		{"IPv4 rejected", "10.0.0.1", true},
		{"IPv6 rejected", "2001:db8::1", true},
		{"CIDR range rejected", "10.0.0.0/8", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			warnings, errors := validateWhitelistedDomain(tc.input, "whitelisted_domains.0")
			assert.Empty(t, warnings)
			if tc.wantErr {
				assert.NotEmpty(t, errors, "expected %q to be rejected", tc.input)
			} else {
				assert.Empty(t, errors, "expected %q to be accepted", tc.input)
			}
		})
	}
}

func TestHashNormalizedDomain(t *testing.T) {
	base := hashNormalizedDomain("example.com")

	// Entries that the API normalizes to the same stored value must hash equal,
	// so a config value matches the API-returned value without a spurious diff.
	equivalent := []string{
		"www.example.com",
		"http://example.com",
		"https://example.com",
		"http://www.example.com",
		"https://www.example.com/",
		"EXAMPLE.COM",
		" example.com ",
	}
	for _, v := range equivalent {
		assert.Equal(t, base, hashNormalizedDomain(v), "expected hash of %q to equal hash of example.com", v)
	}

	assert.NotEqual(t, base, hashNormalizedDomain("other.com"))
	assert.NotEqual(t, hashNormalizedDomain("a.example.com"), hashNormalizedDomain("b.example.com"))
}
