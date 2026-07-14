package applicationset

import "testing"

func TestCheckExactlyOneSource(t *testing.T) {
	tests := []struct {
		name         string
		sourceCount  int
		sourcesCount int
		wantErr      bool
	}{
		{
			name:         "single source only is valid",
			sourceCount:  1,
			sourcesCount: 0,
			wantErr:      false,
		},
		{
			name:         "multiple sources only is valid",
			sourceCount:  0,
			sourcesCount: 2,
			wantErr:      false,
		},
		{
			name:         "neither source nor sources is invalid",
			sourceCount:  0,
			sourcesCount: 0,
			wantErr:      true,
		},
		{
			name:         "both source and sources is invalid",
			sourceCount:  1,
			sourcesCount: 2,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkExactlyOneSource(tt.sourceCount, tt.sourcesCount)
			if tt.wantErr && err == nil {
				t.Fatalf("expected an error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
		})
	}
}
