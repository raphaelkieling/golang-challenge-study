package dto

import (
	"encoding/json"
)

type OrderProductCreate struct {
	Name     string
	Quantity float64
}

type OrderCreate struct {
	Products []OrderProductCreate
}

func UnmarshalOrderCreate(data []byte) (OrderCreate, error) {
	orderCreate := OrderCreate{}
	err := json.Unmarshal(data, &orderCreate)

	return orderCreate, err
}
