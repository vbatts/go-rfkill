package rfkill

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Block device with ID of `id`
func Block(id string) error {
	name := filepath.Join(rfkillSysPath, "rfkill"+id, "soft")
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return errors.New("index not found")
	}
	return ioutil.WriteFile(name, []byte(activeBlock), os.FileMode(0644))
}

// Unblock device with ID of `id`
func Unblock(id string) error {
	name := filepath.Join(rfkillSysPath, "rfkill"+id, "soft")
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return errors.New("index not found")
	}
	return ioutil.WriteFile(name, []byte(inactiveBlock), os.FileMode(0644))
}

// Block this device
func (d Device) Block() error {
	return Block(d.ID)
}

// Unblock this device
func (d Device) Unblock() error {
	return Unblock(d.ID)
}
