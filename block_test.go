package rfkill

import "testing"

func TestBlockUnblock(t *testing.T) {
	devices, err := ListAll()
	if err != nil {
		t.Fatalf("failed to get ListAll: %s", err)
	}
	var wlanDevice Device
	for _, dev := range devices {
		if dev.Type == WlanType {
			wlanDevice = dev
		}
	}
	// record initial state
	startedBlocked := wlanDevice.IsBlocked()

	err = Block(wlanDevice.ID)
	if err != nil {
		t.Errorf("failed to block %q: %s", wlanDevice.DeviceName, err)
	}
	err = Unblock(wlanDevice.ID)
	if err != nil {
		t.Errorf("failed to unblock %q: %s", wlanDevice.DeviceName, err)
	}

	if startedBlocked {
		err = Block(wlanDevice.ID)
		if err != nil {
			t.Errorf("failed to block %q: %s", wlanDevice.DeviceName, err)
		}
	}
}
