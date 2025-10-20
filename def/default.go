package def

// SystemRole represents the role of a user in the system.
type SystemRole string

const (
	SystemRoleOwner SystemRole = "owner"
	SystemRoleAdmin SystemRole = "admin"
	SystemRoleStaff SystemRole = "staff"
)



// Define a private type for context keys
type ContextKey string

const (
    ContextKeySession ContextKey = "session"
    ContextKeyUser    ContextKey = "user"
)

// Define a private type for context keys

const (
    CookieNameSession string = "session"
)
