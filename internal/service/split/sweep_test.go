package split

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolveFMESweepBasePath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		fmeEndpoint     string
		harnessEndpoint string
		want            string
		wantErr         string
	}{
		{
			name:        "override wins",
			fmeEndpoint: "https://fme.example.com/custom/",
			want:        "https://fme.example.com/custom/",
		},
		{
			name: "default gateway",
			want: "https://app.harness.io/fme",
		},
		{
			name:            "trailing slash",
			harnessEndpoint: "https://qa.harness.io/gateway/",
			want:            "https://qa.harness.io/fme",
		},
		{
			name:            "non gateway endpoint",
			harnessEndpoint: "https://qa.harness.io/api",
			wantErr:         "must end in /gateway",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := resolveFMESweepBasePath(tt.fmeEndpoint, tt.harnessEndpoint)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
