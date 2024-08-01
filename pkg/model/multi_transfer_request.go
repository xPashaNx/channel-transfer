package model

import (
	"encoding/json"

	"github.com/anoideaopen/channel-transfer/pkg/data"
)

// MultiTransferResult is the result of processing a request to transfer funds from
// one channel to another. The statuses are symmetrical to [proto].
// TODO: Custom statuses should be added later.
type MultiTransferResult struct {
	Status  string
	Message string
}

// MultiTransferItemRequest contains token data for a MultiTransferRequest
type MultiTransferItemRequest struct {
	Token  string
	Amount string
}

// MultiTransferRequest contains the internal representation of a request to transfer
// funds from one channel to another. This structure is filled from the request
// and enters the queue for processing.
type MultiTransferRequest struct {
	TransferResult

	Request   ID
	Transfer  ID
	User      ID
	Method    string
	Chaincode string
	Channel   string
	Nonce     string
	PublicKey string
	Sign      string
	To        string
	Items     []*MultiTransferItemRequest
}

func (tr *MultiTransferRequest) MarshalBinary() (data []byte, err error) {
	return json.Marshal(tr)
}

func (tr *MultiTransferRequest) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, tr)
}

// Clone should create an exact copy of the object, located in a different
// memory location from the original. This is necessary in case of cache
// optimization to avoid marshalling the object in some cases. Clone is also
// used as a template for finding an object in the repository.
func (tr *MultiTransferRequest) Clone() data.Object {
	trCopy := *tr
	return &trCopy
}

// Instance should return a unique object type to share namespace between
// the stored data. In the simplest case, you can return the type name via
// InstanceOf, but keep in mind that you need to preserve compatibility or
// provide for migration when refactoring.
func (tr *MultiTransferRequest) Instance() data.Type {
	return data.InstanceOf(tr)
}
