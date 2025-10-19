package types

// Types - data structures for {{.Name}}
// These structs represent requests, responses, and entities
// that are only meaningful within this api.


type Example{{.Name}} struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
