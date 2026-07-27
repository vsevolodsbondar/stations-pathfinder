package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenNoConnectionsBlock(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_no_connection_word.txt",
		"bond_square",
		"space_port",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: No connections block in file") {
		t.Fatalf("expected error message\n%s", output)
	}
}
