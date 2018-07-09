package rfkill

import (
	"fmt"
	"testing"
)

func TestListAll(t *testing.T) {
	devices, err := ListAll()
	if err != nil {
		t.Fatalf("failed to get ListAll: %s", err)
	}
	if len(devices) == 0 {
		t.Error("expected devices, but found none")
	}
	fmt.Printf("%#v\n", devices)
}

func TestRfkillDev(t *testing.T) {
	rd := newRfkillDev()
	err := rd.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer rd.Close()
	//for {
	buf, err := rd.Next()
	if err != nil {
		t.Error(err)
		//break
	}
	fmt.Printf("%#v\n", buf)
	//}
}

func TestRfkillSysPaths(t *testing.T) {
	matches, err := sysPaths()
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) == 0 {
		t.Errorf("expected more than 0 sys paths for rfkill devices")
	}
	fmt.Printf("%#v\n", matches)
}
