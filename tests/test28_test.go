package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenTooManyStations(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_too_many_stations.txt",
		"s0",
		"s1",
		"1",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "Error: map contains more than 10000 stations") {
		t.Fatalf("expected error message\n%s", output)
	}
}
