package mysqlstorage

import (
	"errors"

	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/terror"
	goctx "golang.org/x/net/context"
)

// Snapshot defines the interface for the snapshot fetched from KV store.
type Snapshot struct {
	kv.Retriever
}

// BatchGet gets a batch of values from snapshot.
func (Snapshot) BatchGet(keys []kv.Key) (map[string][]byte, error) {
	terror.Log(errors.New("No implemented"))
	return make(map[string][]byte, 0), nil
}

// Transaction defines the interface for operations inside a Transaction.
// This is not thread safe.
type Transaction struct {
	kv.MemBuffer
}

// Commit commits the transaction operations to KV store.
func (transaction Transaction) Commit(goctx.Context) error {
	terror.Log(errors.New("No implemented"))
	return nil
}

// Rollback undoes the transaction operations to KV store.
func (transaction Transaction) Rollback() error {
	terror.Log(errors.New("No implemented"))
	return nil
}

// String implements fmt.Stringer interface.
func (transaction Transaction) String() string {
	terror.Log(errors.New("No implemented"))
	return ""
}

// LockKeys tries to lock the entries with the keys in KV store.
func (transaction Transaction) LockKeys(keys ...kv.Key) error {
	terror.Log(errors.New("No implemented"))
	return nil
}

// SetOption sets an option with a value, when val is nil, uses the default
// value of this option.
func (transaction Transaction) SetOption(opt kv.Option, val interface{}) {
	terror.Log(errors.New("No implemented"))
}

// DelOption deletes an option.
func (transaction Transaction) DelOption(opt kv.Option) {
	terror.Log(errors.New("No implemented"))
}

// IsReadOnly checks if the transaction has only performed read operations.
func (transaction Transaction) IsReadOnly() bool {
	terror.Log(errors.New("No implemented"))
	return true
}

// StartTS returns the transaction start timestamp.
func (transaction Transaction) StartTS() uint64 {
	terror.Log(errors.New("No implemented"))
	return 0
}

// Valid returns if the transaction is valid.
// A transaction become invalid after commit or rollback.
func (transaction Transaction) Valid() bool {
	terror.Log(errors.New("No implemented"))
	return true
}
