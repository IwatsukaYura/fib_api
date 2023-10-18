package models


type Result struct {
	Result int	`json:"result"`
}

type Error struct {
	Status int	`json:"status"`
	Message string	`json:"message"`
}