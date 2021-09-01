package utils

import (
	"bytes"
	"os/exec"
)

func GetLatestTag() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	tag := out.String()
	return tag[:len(tag)-1], nil
}
