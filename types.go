package rfkill

// Device is the type listed in rfkill columns
type Device struct {
	ID          string     `json:"id"`
	Type        Type       `json:"type"`
	DeviceName  string     `json:"device"`
	SoftBlocked BlockState `json:"soft"`
	HardBlocked BlockState `json:"hard"`
}

// BlockState is the state of the type of block
type BlockState string

// the two block states
const (
	Blocked   BlockState = "blocked"
	Unblocked BlockState = "unblocked"
)

// Type is the device type
type Type string

// various types
const (
	BluetoothType Type = "bluetooth"
	WlanType      Type = "wlan"
)
