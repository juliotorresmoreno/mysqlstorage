package mysqlstorage

import (
	"time"

	goctx "golang.org/x/net/context"
)

// Oracle is the interface that provides strictly ascending timestamps.
type Oracle struct {
}

func (oracle Oracle) GetTimestamp(ctx goctx.Context) (uint64, error) {
	return 0, nil
}

func (oracle Oracle) GetTimestampAsync(ctx goctx.Context) Future {
	return Future{}
}

func (oracle Oracle) IsExpired(lockTimestamp uint64, TTL uint64) bool {
	return true
}

func (oracle Oracle) Close() {

}

// Future is a future which promises to return a timestamp.
type Future struct {
}

func (future Future) Wait() (uint64, error) {
	return 0, nil
}

const physicalShiftBits = 18

// ComposeTS creates a ts from physical and logical parts.
func ComposeTS(physical, logical int64) uint64 {
	return uint64((physical << physicalShiftBits) + logical)
}

// ExtractPhysical returns a ts's physical part.
func ExtractPhysical(ts uint64) int64 {
	return int64(ts >> physicalShiftBits)
}

// GetPhysical returns physical from an instant time with millisecond precision.
func GetPhysical(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
