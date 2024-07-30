package he

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnshe "github.com/libdns/he"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"he testapikey01",
			"testapikey01",
		},
		{
			`he {
			    api_key testapikey02
			}`,
			"testapikey02",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			dispenser := caddyfile.NewTestDispenser(tc.input)
			p := Provider{&libdnshe.Provider{}}

			err := p.UnmarshalCaddyfile(dispenser)

			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			if tc.output != p.Provider.APIKey {
				t.Errorf(
					"Expected APIKey to be '%s' but got '%s'",
					tc.output, p.Provider.APIKey,
				)
			}
		})
	}
}
