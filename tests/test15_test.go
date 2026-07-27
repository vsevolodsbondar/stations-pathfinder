package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenStartStationDoesNotExist(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"../smallAndLarge.map",
		"smalls",
		"large",
		"10",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Starting station does not exist.") {
		t.Fatalf("expected error message\n%s", output)
	}
}
