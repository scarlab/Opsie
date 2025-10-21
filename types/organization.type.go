package types

import "time"

// Types - data structures for Organization
// These structs represent requests, responses, and entities
// that are only meaningful within this api.


type Organization struct {
    ID              ID          `json:"id"`
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Logo            string      `json:"logo"`
    UpdatedAt       time.Time   `json:"updated_at"`
    CreatedAt       time.Time   `json:"created_at"`
}

type UserOrganization struct {
    ID              ID          `json:"id"`
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Logo            string      `json:"logo"`
    UpdatedAt       time.Time   `json:"updated_at"`
    CreatedAt       time.Time   `json:"created_at"`

    IsDefault       bool        `json:"is_default"`
    JoinedAt        time.Time   `json:"joined_at"`
}

type NewOrganizationPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Logo            string      `json:"logo"`
}

type UpdateOrganizationPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
}
