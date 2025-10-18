package socket

import (
	"encoding/json"
	"fmt"
)



func MarshalEnvelope(t string, v interface{}) ([]byte, error) {
    b, err := json.Marshal(v)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }

    envelope := TEnvelope{
        Type:    t,
        Payload: json.RawMessage(b),
    }

    msg, err := json.Marshal(envelope)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal envelope: %w", err)
    }

    return msg, nil
}

func UnmarshalEnvelope(data []byte) (*TEnvelope, error) {
    var envelope TEnvelope
    if err := json.Unmarshal(data, &envelope); err != nil {
        return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
    }
    return &envelope, nil
}

func DecodePayload[T any](envelope *TEnvelope) (*T, error) {
    var payload T
    if err := json.Unmarshal(envelope.Payload, &payload); err != nil {
        return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
    }
    return &payload, nil
}
