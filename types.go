package rfkill

// Device is the type listed in rfkill columns
type Device struct {
	ID          string     `json:"id"`
	Type        Type       `json:"type"`
	DeviceName  string     `json:"device"`
	SoftBlocked BlockState `json:"soft"`
	HardBlocked BlockState `json:"hard"`
}

// IsBlocked is convinience of comparing whether the device is either soft or
// hard blocked.
func (d Device) IsBlocked() bool {
	return d.SoftBlocked == Blocked || d.HardBlocked == Blocked
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
	NfcType       Type = "nfc"
	FmType        Type = "fm"
	GpsType       Type = "gps"
	UwbType       Type = "uwb"
	WimaxType     Type = "wimax"
	WwanType      Type = "wwan"
)
