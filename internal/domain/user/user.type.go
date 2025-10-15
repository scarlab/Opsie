package user

// Types - data structures for user
// These structs represent requests, responses, and entities
// that are only meaningful within this domain.

type TNewOwnerPayload struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}
