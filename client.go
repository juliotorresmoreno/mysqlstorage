package mysqlstorage

import (
	"errors"

	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/terror"
	goctx "golang.org/x/net/context"
)

// Response represents the response returned from KV layer.
type Response struct {
}

// Next returns a resultSubset from a single storage unit.
// When full result set is returned, nil is returned.
// TODO: Find a better interface for resultSubset that can avoid allocation and reuse bytes.
func (response Response) Next() (resultSubset []byte, err error) {
	terror.Log(errors.New("No implemented"))
	return []byte{}, nil
}

// Close response.
func (response Response) Close() error {
	terror.Log(errors.New("No implemented"))
	return nil
}

// Client is used to send request to KV layer.
type Client struct {
}

// Send sends request to KV layer, returns a Response.
func (client Client) Send(ctx goctx.Context, req *kv.Request) Response {
	terror.Log(errors.New("No implemented"))
	return Response{}
}

// IsRequestTypeSupported checks if reqType and subType is supported.
func (client Client) IsRequestTypeSupported(reqType, subType int64) bool {
	terror.Log(errors.New("No implemented"))
	return true
}
