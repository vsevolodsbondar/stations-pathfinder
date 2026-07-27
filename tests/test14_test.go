package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenTooManyArgsIsUsed(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"smallAndLarge.map",
		"small",
		"large",
		"10",
		"11",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Usage") {
		t.Fatalf("expected usage message\n%s", output)
	}
}
