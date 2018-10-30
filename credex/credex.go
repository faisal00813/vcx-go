package credex

// import "fmt"

import (
	vcx "github.com/faisal00813/vcx-go/vcx"
)

type callbackResult struct {
	commandHandle, errorCode uint32
	intValue                 uint32
	stringValue              string
}

// Connection struct which represents a connection
type Connection struct {
	state, Handle uint32
	SourceID      string
}

// CreateConnection creates a new connection with the given sourceID/DID
func (connection Connection) CreateConnection(sourceID string) (Connection, error) {
	c := make(chan callbackResult, 1)
	callback := cbWithIntResponse(c)
	vcx.Connection_create(createCommandHandle(), sourceID, callback)
	result := <-c
	conn := Connection{0, result.intValue, sourceID}
	return conn, nil
}
