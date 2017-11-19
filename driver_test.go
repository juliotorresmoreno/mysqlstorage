package mysqlstorage

import "testing"

func TestRun(t *testing.T) {
	t.Run("Connect", testConnect)
}

func testConnect(t *testing.T) {
	store := Storage{}
	store.Begin()
}
