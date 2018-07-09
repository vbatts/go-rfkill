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
