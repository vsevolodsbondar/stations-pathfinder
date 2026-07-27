package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenNegCoordinates(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_neg_coords.txt",
		"bond_square",
		"space_port",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: Invalid station (bond_square,-20,-6), line 7: Axis contains not valid symbol: -.") {
		t.Fatalf("expected error message\n%s", output)
	}
}
