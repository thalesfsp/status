package status

//////
// Const, vars, and types.
//////

// Name of the entity.
const Name = "status"

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
	Emitted    Status = "emitted"
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

	// Connection.
	Connected    Status = "connected"
	Disconnected Status = "disconnected"

	// Instantiation.
	Instantiated   Status = "instantiated"
	Uninstantiated Status = "uninstantiated"

	// ETL
	Converted   Status = "converted"
	Exported    Status = "exported"
	Extracted   Status = "extracted"
	Imported    Status = "imported"
	Loaded      Status = "loaded"
	Processed   Status = "processed"
	Transformed Status = "transformed"

	// Common on authentication.
	Authenticated   Status = "authenticated"
	Authentication  Status = "authentication"
	Authorization   Status = "authorization"
	Authorized      Status = "authorized"
	Unauthenticated Status = "unauthenticated"
	Unauthorized    Status = "unauthorized"

	// Common on authorization.
	Granted  Status = "granted"
	Revoked  Status = "revoked"
	Verified Status = "verified"
)

//////
// Methods.
//////

// Implement the Stringer interface.
func (s Status) String() string {
	return string(s)
}
