package rfkill

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

// ListAll collects the state of all devices present
func ListAll() ([]Device, error) {
	cmd := exec.Command("rfkill", "--json")
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return nil, errors.New(errOut.String())
	}
	list := map[string][]Device{}
	if err := json.Unmarshal(out.Bytes(), &list); err != nil {
		return nil, err
	}
	return list[""], nil
}

/*
for now this is a wrapper, but might could be a bit more direct ...

openat(AT_FDCWD, "/dev/rfkill", O_RDONLY) = 3
fcntl(3, F_SETFL, O_RDONLY|O_NONBLOCK)  = 0
read(3, "\0\0\0\0\2\0\0\0", 8)          = 8
openat(AT_FDCWD, "/sys/class/rfkill/rfkill0/name", O_RDONLY) = 4
fstat(4, {st_mode=S_IFREG|0444, st_size=4096, ...}) = 0
read(4, "tpacpi_bluetooth_sw\n", 4096)  = 20
close(4)                                = 0
fstat(1, {st_mode=S_IFCHR|0620, st_rdev=makedev(136, 4), ...}) = 0
write(1, "0: tpacpi_bluetooth_sw: Bluetoot"..., 34) = 34
write(1, "\tSoft blocked: no\n", 18)    = 18
write(1, "\tHard blocked: no\n", 18)    = 18
read(3, "\1\0\0\0\2\0\0\0", 8)          = 8
openat(AT_FDCWD, "/sys/class/rfkill/rfkill1/name", O_RDONLY) = 4
fstat(4, {st_mode=S_IFREG|0444, st_size=4096, ...}) = 0
read(4, "hci0\n", 4096)                 = 5
close(4)                                = 0
write(1, "1: hci0: Bluetooth\n", 19)    = 19
write(1, "\tSoft blocked: no\n", 18)    = 18
write(1, "\tHard blocked: no\n", 18)    = 18
read(3, "\2\0\0\0\1\0\1\0", 8)          = 8
openat(AT_FDCWD, "/sys/class/rfkill/rfkill2/name", O_RDONLY) = 4
fstat(4, {st_mode=S_IFREG|0444, st_size=4096, ...}) = 0
read(4, "phy0\n", 4096)                 = 5

*/
