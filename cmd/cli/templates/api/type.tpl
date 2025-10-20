package types

import "time"

// Types - data structures for {{.Name}}
// These structs represent requests, responses, and entities
// that are only meaningful within this api.


type {{.Name}} struct {
    ID              ID              `json:"id"`
    Name            string          `json:"name"`
    UpdatedAt       time.Time       `json:"updated_at"`
    CreatedAt       time.Time       `json:"created_at"`
}

type New{{.Name}}Payload struct {
    Name            string      `json:"name"`
}