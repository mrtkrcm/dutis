package util

import (
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
)

func InstallDeps() {
	installHomebrew()
	installDuti()
}

func installHomebrew() {
	fmt.Println("Check Homebrew Environment")
	if !commandExists("brew") {
		fmt.Println("Homebrew not exists, installing ...")
		cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
		_, _ = cmd.Output()
		updatePathForHomebrew()
	}

	cmd := exec.Command("brew", "--help")
	_, err := cmd.Output()

	if err != nil {
		err.Error()
	} else {
		fmt.Println("Homebrew works fine")
	}
}

func updatePathForHomebrew() {
	homebrewPath := "/usr/local/bin/brew"
	if _, err := os.Stat(homebrewPath); os.IsNotExist(err) {
		homebrewPath = "/opt/homebrew/bin/brew"
	}

	if _, err := os.Stat(homebrewPath); err == nil {
		brewDir := filepath.Dir(homebrewPath)
		os.Setenv("PATH", brewDir+":"+os.Getenv("PATH"))
	}
}

func installDuti() {
	fmt.Println("Check Duti Environment")
	if !commandExists("duti") {
		fmt.Println("Duti not exists, installing ...")
		cmd := exec.Command("brew", "install", "duti")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			fmt.Println("Error installing duti:", err)
			return
		}
	}

	cmd := exec.Command("man", "duti")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		fmt.Println("Error checking duti installation:", err)
		return
	} else {
		fmt.Println("Duti works fine")
	}
}
