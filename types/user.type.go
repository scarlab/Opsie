package types

// Types - data structures for user
// These structs represent requests, responses, and entities
// that are only meaningful within this api.

type User struct {
    ID ID                       `json:"id"`
    DisplayName string          `json:"display_name"`
    Email string                `json:"email"`
    Password string             `json:"password"`
    Avatar string               `json:"avatar"`
    SystemRole string           `json:"system_role"`
    Preference map[string]any   `json:"preference"`
    IsActive bool               `json:"is_active"`
    UpdatedAt string            `json:"updated_at"`
    CreatedAt string            `json:"created_at"`
}



type NewOwnerPayload struct {
    DisplayName string          `json:"display_name"`
    Email string                `json:"email"`
    Password string             `json:"password"`
}


type UpdateAccountNamePayload struct {
    DisplayName string          `json:"display_name"`
}
type UpdateAccountPasswordPayload struct {
    Password string             `json:"password"`
    NewPassword string          `json:"new_password"`
}


