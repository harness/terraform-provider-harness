package convert

import (
	"testing"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]interface{}
		key      string
		required bool
		want     string
		wantErr  bool
	}{
		{
			name: "Key exists",
			m:    map[string]interface{}{"test": "value"},
			key:  "test",
			want: "value",
		},
		{
			name:     "Key missing, not required",
			m:        map[string]interface{}{},
			key:      "missing",
			required: false,
			want:     "",
		},
		{
			name:     "Key missing, required",
			m:        map[string]interface{}{},
			key:      "missing",
			required: true,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getString(tt.m, tt.key, tt.required)
			if (err != nil) != tt.wantErr {
				t.Errorf("getString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name    string
		m       map[string]interface{}
		key     string
		want    int
		wantErr bool
	}{
		{
			name: "Int value",
			m:    map[string]interface{}{"test": 42},
			key:  "test",
			want: 42,
		},
		{
			name: "String value",
			m:    map[string]interface{}{"test": "42"},
			key:  "test",
			want: 42,
		},
		{
			name:    "Invalid type",
			m:       map[string]interface{}{"test": []string{"not", "an", "int"}},
			key:     "test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getInt(tt.m, tt.key, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("getInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("getInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
