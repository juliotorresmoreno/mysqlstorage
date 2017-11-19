package mysqlstorage

import "github.com/pingcap/tidb/kv"

// Iterator is the interface for a iterator on KV store.
type Iterator struct {
}

func (iterator Iterator) Valid() bool {
	return true
}

func (iterator Iterator) Key() kv.Key {
	return kv.Key{}
}

func (iterator Iterator) Value() []byte {
	return kv.Key{}
}

func (iterator Iterator) Next() error {
	return nil
}

func (iterator Iterator) Close() {

}

// Retriever is the interface wraps the basic Get and Seek methods.
type Retriever struct {
}

// Get gets the value for key k from kv store.
// If corresponding kv pair does not exist, it returns nil and ErrNotExist.
func (retriever Retriever) Get(k kv.Key) ([]byte, error) {
	return []byte{}, nil
}

// Seek creates an Iterator positioned on the first entry that k <= entry's key.
// If such entry is not found, it returns an invalid Iterator with no error.
// The Iterator must be Closed after use.
func (retriever Retriever) Seek(k kv.Key) (Iterator, error) {
	return Iterator{}, nil
}

// SeekReverse creates a reversed Iterator positioned on the first entry which key is less than k.
// The returned iterator will iterate from greater key to smaller key.
// If k is nil, the returned iterator will be positioned at the last key.
func (retriever Retriever) SeekReverse(k kv.Key) (Iterator, error) {
	return Iterator{}, nil
}
