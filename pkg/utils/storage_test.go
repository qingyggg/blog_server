package utils

import (
	"context"
	"testing"
)

func TestUrlConvertReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "src path",
			in:   "/src/imagebucket/demo.jpg",
			want: "imagebucket/demo.jpg",
		},
		{
			name: "absolute nginx src url",
			in:   "http://localhost:8080/src/imagebucket/demo.jpg",
			want: "imagebucket/demo.jpg",
		},
		{
			name: "absolute preview src url",
			in:   "http://localhost:4173/src/imagebucket/demo.jpg",
			want: "imagebucket/demo.jpg",
		},
		{
			name: "bucket object path",
			in:   "imagebucket/demo.jpg",
			want: "imagebucket/demo.jpg",
		},
		{
			name: "invalid path",
			in:   "/assets/demo.jpg",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UrlConvertReverse(context.Background(), tt.in); got != tt.want {
				t.Fatalf("UrlConvertReverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
