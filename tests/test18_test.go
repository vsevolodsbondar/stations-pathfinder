package tests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShouldHaveErrorWhenNoPathFound(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=seva",
		"./testData/error_nopath.txt",
		"start",
		"end",
		"10",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "There is no route that reaches ending station.") {
		t.Fatalf("expected error message\n%s", output)
	}
}

func TestShouldHaveErrorWhenNoPathFound2(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "../",
		"-algorithm=anatolii",
		"./testData/error_nopath.txt",
		"start",
		"end",
		"10",
	)

	out, _ := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("Output", output)

	if !strings.Contains(output, "There is no route that reaches ending station.") {
		t.Fatalf("expected error message\n%s", output)
	}
}
