package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenInvalidStationName(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_invalid_names.txt",
		"bond_square",
		"space_port",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: Invalid station (Orange_junction^,6,1), line 11: Not valid symbol in name: O.") {
		t.Fatalf("expected error message\n%s", output)
	}
}
