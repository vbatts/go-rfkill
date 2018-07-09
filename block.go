package rfkill

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

var sanitaryChars = " ;&|"

// Block device with ID of `id`
func Block(id string) error {
	if strings.Contains(id, sanitaryChars) {
		return errors.New("unallowed characters in id")
	}
	cmd := exec.Command("rfkill", "block", id)
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return errors.New(errOut.String())
	}
	return nil
}

// Unblock device with ID of `id`
func Unblock(id string) error {
	if strings.Contains(id, sanitaryChars) {
		return errors.New("unallowed characters in id")
	}
	cmd := exec.Command("rfkill", "unblock", id)
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return errors.New(errOut.String())
	}
	return nil
}
