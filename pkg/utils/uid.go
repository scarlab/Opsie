package utils

import "github.com/sony/sonyflake"

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GenerateID() uint64 {
    id, err := flake.NextID()
    if err != nil {
        panic(err)
    }
    return id
}
