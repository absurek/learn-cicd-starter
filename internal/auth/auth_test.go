package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}

	tests := []test{
		{
			name:    "Valid API Key Authorization Header",
			header:  http.Header{"Authorization": []string{"ApiKey my_api_key"}},
			want:    "my_api_key",
			wantErr: false,
		},
		{
			name:    "Malformed API Key Authorization Header",
			header:  http.Header{"Authorization": []string{"NotApiKey my_api_key"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No Authorization Header",
			header:  http.Header{"Content-Type": []string{"text/plain"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if err != nil && !tt.wantErr {
				t.Fatalf("expected no error, got: %v", err)
			}

			if got != tt.want {
				t.Fatalf("expected: %s, got: %s", tt.want, got)
			}
		})
	}
}
