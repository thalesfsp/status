package status

//////
// Const, vars, and types.
//////

// Status is the status of something.
type Status string

const (
	None Status = "none"

	// Common on displaying.
	Active   Status = "active"
	Hidden   Status = "hidden"
	Inactive Status = "inactive"

	// Common on CRUD.
	Counted   Status = "counted"
	Created   Status = "created"
	Deleted   Status = "deleted"
	Listed    Status = "listed"
	Retrieved Status = "retreived"
	Updated   Status = "updated"

	// Common on state machines.
	Canceled    Status = "canceled"
	Completed   Status = "completed"
	Done        Status = "done"
	Failed      Status = "failed"
	Initialized Status = "initialized"
	Paused      Status = "paused"
	Runnning    Status = "running"
	Stopped     Status = "stopped"

	// Common for message brokers (pubsub).
	Published  Status = "published"
	Subscribed Status = "subscribed"

	// Metrics.
	Retried   Status = "retried"
	Succeeded Status = "succeeded"
	Total     Status = "total"

	// Errors.
	Invalid  Status = "invalid"
	Missing  Status = "missed"
	NotFound Status = "not found"
	Required Status = "required"

	// Connection
	Connected    Status = "connected"
	Disconnected Status = "disconnected"
)

//////
// Methods.
//////

// Implement the Stringer interface.
func (s Status) String() string {
	return string(s)
}
