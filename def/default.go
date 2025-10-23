package def

// SystemRole represents the role of a user in the system.
type SystemRole string
func (k SystemRole) ToString() string  {
	return string(k)
}
const (
	SystemRoleOwner SystemRole = "owner"
	SystemRoleAdmin SystemRole = "admin"
	SystemRoleStaff SystemRole = "staff"
)



// Define a private type for context keys
type ContextKey string
func (k ContextKey) ToString() string  {
	return string(k)
}
const (
    ContextKeySession ContextKey = "session"
    ContextKeyUser    ContextKey = "user"
)

// Define a private type for context keys

const (
    CookieNameSession string = "session"
)



// ---
type ResourceStatus string
const (
	StatusStopped    ResourceStatus = "stopped"
	StatusStarting   ResourceStatus = "starting"
	StatusRunning    ResourceStatus = "running"
	StatusRestarting ResourceStatus = "restarting"
	StatusDegraded   ResourceStatus = "degraded"
	StatusFailed     ResourceStatus = "failed"
)