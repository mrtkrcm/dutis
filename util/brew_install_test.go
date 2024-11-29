package util

import (
	"os/exec"
	"testing"
)

func TestBrewInstall(t *testing.T) {
	cmd := exec.Command("brew", "--version")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Homebrew is not installed or not functional: %v", err)
	}
}
