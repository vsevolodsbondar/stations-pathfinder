package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenEndStartStationAreSame(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"../smallAndLarge.map",
		"small",
		"small",
		"10",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Starting and ending locations should not be the same") {
		t.Fatalf("expected error message\n%s", output)
	}
}
