package mysqlstorage

import (
	"github.com/pingcap/tidb/kv"
)

type Storage struct {
}

// Begin transaction
func (store Storage) Begin() (Transaction, error) {
	return Transaction{}, nil
}

// BeginWithStartTS begins transaction with startTS.
func (store Storage) BeginWithStartTS(startTS uint64) (kv.Transaction, error) {
	return Transaction{}, nil
}

// GetSnapshot gets a snapshot that is able to read any data which data is <= ver.
// if ver is MaxVersion or > current max committed version, we will use current version for this snapshot.
func (store Storage) GetSnapshot(ver kv.Version) (kv.Snapshot, error) {
	return Snapshot{}, nil
}

// GetClient gets a client instance.
func (store Storage) GetClient() Client {
	return Client{}
}

// Close store
func (store Storage) Close() error {
	return nil
}

// UUID return a unique ID which represents a Storage.
func (store Storage) UUID() string {
	return ""
}

// CurrentVersion returns current max committed version.
func (store Storage) CurrentVersion() (kv.Version, error) {
	return kv.Version{}, nil
}

// GetOracle gets a timestamp oracle client.
func (store Storage) GetOracle() Oracle {
	return Oracle{}
}

// SupportDeleteRange gets the storage support delete range or not.
func (store Storage) SupportDeleteRange() (supported bool) {
	return true
}
