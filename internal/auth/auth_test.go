package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	ApiKey := "1234qwerty"
	validHeader := http.Header{}
	validHeader.Set("Authorization", "ApiKey 1234qwerty")

	emptyHeader := http.Header{}

	emptyHeaderValue := http.Header{}
	emptyHeaderValue.Set("Authorization", "")

	malformedHeader1 := http.Header{}
	malformedHeader1.Set("Authorization", "ApiKey1234qwerty")

	malformedHeader2 := http.Header{}
	malformedHeader2.Set("Authorization", "gibberish 1234qwerty")

	type args struct {
		headers http.Header
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid header",
			args:    args{validHeader},
			want:    ApiKey,
			wantErr: false,
		},
		{
			name:    "Empty header",
			args:    args{emptyHeader},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty header value",
			args:    args{emptyHeaderValue},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Malformed header 1",
			args:    args{malformedHeader1},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Malformed header 2",
			args:    args{malformedHeader2},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
