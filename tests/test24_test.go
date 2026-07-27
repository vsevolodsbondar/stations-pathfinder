package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenStationNameDupl(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_dupl_names.txt",
		"bond_square",
		"space_port",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: Station is duplicated: bond_square, line: 8") {
		t.Fatalf("expected error message\n%s", output)
	}
}
