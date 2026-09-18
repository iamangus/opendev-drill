// Package greeting provides the repository's canonical greeting text.
//
// The greeting is derived from the repository itself: README.md documents the
// project with the heading "# opendev-drill", and that heading is pinned by
// TestREADMEHeading in readme_test.go, which requires README.md to begin with
// "# opendev-drill". Greet returns that greeting text.
package greeting

// Greet returns the repository greeting, "opendev-drill", matching the
// "# opendev-drill" heading that README.md must begin with per
// TestREADMEHeading in readme_test.go.
func Greet() string {
	return "opendev-drill"
}
