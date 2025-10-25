package service

import (
    "crypto/rand"
    "encoding/hex"
)

func GenerateID() string {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        panic(err)
    }
    return hex.EncodeToString(bytes)
}

