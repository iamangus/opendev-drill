package greeting

import "testing"

// TestGreet verifies that Greet returns the repository greeting text
// "opendev-drill", the greeting documented by README.md's heading
// "# opendev-drill" and pinned by TestREADMEHeading in readme_test.go.
func TestGreet(t *testing.T) {
	const want = "opendev-drill"

	got := Greet()
	if got != want {
		t.Fatalf("Greet() = %q, want %q", got, want)
	}
}
