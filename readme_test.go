package drill

import (
	"os"
	"strings"
	"testing"
)

func TestREADMEHeading(t *testing.T) {
	contents, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("read README.md: %v", err)
	}

	if !strings.HasPrefix(string(contents), "# opendev-drill") {
		t.Fatalf("README.md does not begin with # opendev-drill")
	}
}
