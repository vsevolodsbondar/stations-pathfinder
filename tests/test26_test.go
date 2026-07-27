package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenNoStationsBlock(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_no_station_word.txt",
		"bond_square",
		"space_port",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: No stations block in file") {
		t.Fatalf("expected error message\n%s", output)
	}
}
