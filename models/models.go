package models

import "math/big"

type Result struct {
	Result *big.Int `json:"result"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
