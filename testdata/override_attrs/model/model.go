package model

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type Inner struct {
	X int `json:"x"`
}
