// Package models.
package models

// Status enum.
type Status int

// Create Status values.
const (
	StatusPending Status = iota
	StatusAccepted
	StatusRejected
	StatusReady
	StatusFinished
	StatusCancelled
	StatusInvalid // This is a special status that is used to indicate that the status is invalid.
)

// StatusMap This struct is used for keeping Status data.
type StatusMap struct {
	Id          int
	Code        Status
	Description string
}

// InvalidStatusMap This struct is used for keeping invalid Status data.
var InvalidStatusMap = StatusMap{
	Id:          0,
	Code:        StatusInvalid,
	Description: "Invalid Status",
}

// IsValid check for StatusMap.
func (status StatusMap) IsValid() bool {
	return status.Id > 0 && status.Code != StatusInvalid
}
