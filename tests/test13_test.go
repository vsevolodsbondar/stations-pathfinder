package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenTofewArgsIsUsed(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"smallAndLarge.map",
		"small",
		"large",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Usage") {
		t.Fatalf("expected usage message\n%s", output)
	}

}
