package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name        string
		input       http.Header
		want        string
		expectedErr error
	}{
		{
			name:        "valid header",
			input:       http.Header{"Authorization": []string{"ApiKey token123"}},
			want:        "token123",
			expectedErr: nil,
		},
		{
			name:        "invalid field",
			input:       http.Header{"Authorization": []string{"Bearer token123"}},
			want:        "",
			expectedErr: ErrMalformedAuthHeader,
		},
		{
			name:        "invalid header",
			input:       http.Header{"Content-Type": []string{"application/json"}},
			want:        "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)

			if got != tt.want {
				t.Errorf("GetAPIKey() got = %q, want %q", got, tt.want)
			}

			if err != tt.expectedErr {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedErr)
			}
		})
	}
}
