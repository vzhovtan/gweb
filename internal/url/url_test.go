package url

// TBD - refactor to use interfaces
// TBD - use interfaces for DB mockup
import (
	"testing"
)

func TestShortenURL(t *testing.T) {
	tests := []struct {
		name     string
		original string
		want     string
	}{
		{
			name:     "given the original valid url, getting back the shorten one, 1st case",
			original: "www.google.com",
			want:     "191347bf",
		},
		{
			name:     "given the original valid url, getting back the shorten one, 2nd case",
			original: "www.yahoo.com",
			want:     "d61a0fed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortenURL(tt.original)
			if got != tt.want {
				t.Errorf("ShortenURL  = %v, wanted %v", got, tt.want)
			}
		})
	}
}
