package rfkill

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/unix"
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

func newRfkillDev() *rfkillDev {
	return &rfkillDev{
		Path: rfkillDevPath,
	}
}

var (
	rfkillDevPath = "/dev/rfkill"
	rfkillSysPath = ""
)

type rfkillDev struct {
	Path string
	File *os.File
}

func (rd *rfkillDev) Open() error {
	var err error
	rd.File, err = os.OpenFile(rd.Path, os.O_RDONLY, os.FileMode(0664))
	if err != nil {
		return fmt.Errorf("failed to open %q: %s", rd.File, err)
	}
	ret, err := unix.FcntlInt(rd.File.Fd(), unix.F_SETFL, unix.O_RDONLY|unix.O_NONBLOCK)
	if err != nil && err != syscall.Errno(0x0) {
		return fmt.Errorf("failed to fcntl %q: %#V %s", rd.File, err, err)
	}
	if ret != 0 {
		return fmt.Errorf("%q fcntl returned non-zero %d", rd.Path, ret)
	}
	return nil
}

func (rd *rfkillDev) Next() ([]byte, error) {
	// err may be unix.EAGAIN which means done reading
	buf := make([]byte, 8)
	_, err := rd.File.Read(buf)
	return buf, err
}

func (rd *rfkillDev) Close() error {
	return rd.File.Close()
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
