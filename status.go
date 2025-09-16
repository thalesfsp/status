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

	// Common on configuration.
	Enabled  Status = "enabled"
	Disabled Status = "disabled"

	// Common on displaying.
	Active   Status = "active"
	Hidden   Status = "hidden"
	Inactive Status = "inactive"

	// Common on CRUD.
	Counted   Status = "counted"
	Created   Status = "created"
	Deleted   Status = "deleted"
	Listed    Status = "listed"
	Retrieved Status = "retrieved"
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
	Notified   Status = "notified"
	Published  Status = "published"
	Received   Status = "received"
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
	Offline      Status = "offline"
	Online       Status = "online"

	// Instantiation.
	Instantiated   Status = "instantiated"
	Setup          Status = "setup"
	Uninstantiated Status = "not instantiated"

	// ETL
	Converted   Status = "converted"
	Exported    Status = "exported"
	Extracted   Status = "extracted"
	Finished    Status = "finished"
	Imported    Status = "imported"
	Interrupted Status = "interrupted"
	Loaded      Status = "loaded"
	Processed   Status = "processed"
	Started     Status = "started"
	Transformed Status = "transformed"

	// Common on authentication.
	Authenticated   Status = "authenticated"
	Authentication  Status = "authentication"
	Authorization   Status = "authorization"
	Authorized      Status = "authorized"
	Unauthenticated Status = "unauthenticated"
	Unauthorized    Status = "unauthorized"

	// Common on authorization.
	Approved Status = "approved"
	Denied   Status = "denied"
	Granted  Status = "granted"
	Rejected Status = "rejected"
	Revoked  Status = "revoked"
	Verified Status = "verified"
	Waiting  Status = "waiting"
)

//////
// Methods.
//////

// Implement the Stringer interface.
func (s Status) String() string {
	return string(s)
}
