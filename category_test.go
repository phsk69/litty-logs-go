package litty

import "testing"

// TestShortenCategory — make sure we yeet the namespace bloat correctly fr fr 🔥
func TestShortenCategory(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"dots get yeeted", "Microsoft.Hosting.Lifetime", "Lifetime"},
		{"single dot", "pkg.Service", "Service"},
		{"no dots no slashes stays the same", "Program", "Program"},
		{"slashes get yeeted", "github.com/user/pkg", "pkg"},
		{"empty string stays empty", "", ""},
		{"just a dot", ".", ""},
		{"just a slash", "/", ""},
		{"dots before slashes — dots win", "github.com/user/pkg.Service", "Service"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortenCategory(tt.input)
			if got != tt.expected {
				t.Errorf("ShortenCategory(%q) = %q, want %q — namespace bloat survived 💀", tt.input, got, tt.expected)
			}
		})
	}
}
