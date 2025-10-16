package constant

// SystemRole represents the role of a user in the system.
type SystemRole string

const (
	SystemRoleOwner SystemRole = "owner"
	SystemRoleAdmin SystemRole = "admin"
	SystemRoleStaff SystemRole = "staff"
)
