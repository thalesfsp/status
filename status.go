package status

//////
// Const, vars, and types.
//////

// Status is the status of something.
type Status string

const (
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
	Completed   Status = "completed"
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
)

//////
// Methods.
//////

// Implement the Stringer interface.
func (s Status) String() string {
	return string(s)
}
