package types

import "time"

// Types - data structures for auth
// These structs represent requests, responses, and entities
// that are only meaningful within this api.

type Session struct {
    ID        ID                `json:"id"`
    UserID    int64             `json:"user_id"`
    Key       SessionKey        `json:"key"`
    IP        string            `json:"ip,omitempty"`
    OS        string            `json:"os,omitempty"`
    Device    string            `json:"device,omitempty"`
    Browser   string            `json:"browser,omitempty"`
    IsValid   bool              `json:"is_valid"`
    Expiry    time.Time         `json:"expiry,omitempty"`
    CreatedAt time.Time         `json:"created_at"`
}



type LoginPayload struct {
    Email string                `json:"email"`
    Password string             `json:"password"`
}


type AuthUser struct {
    ID ID                       `json:"id"`
    DisplayName string          `json:"display_name"`
    Email string                `json:"email"`
    Avatar string               `json:"avatar"`
    SystemRole string           `json:"system_role"`
    IsActive bool               `json:"is_active"`
    Preference map[string]any   `json:"preference"`
}


type SessionWithUser struct {
    Session Session
    AuthUser    AuthUser
}