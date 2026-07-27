package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenNumbOfTrainsInvalid(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_dupl_names.txt",
		"bond_square",
		"space_port",
		"0",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: Trains must be a positive integer") {
		t.Fatalf("expected error message\n%s", output)
	}
}
