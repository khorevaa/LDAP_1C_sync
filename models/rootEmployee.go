package models

import (
	"time"
)

//RootEmployee ...
type RootEmployee struct {
	Name            string
	Surname         string
	Birthday        time.Time
	Position        string
	Department      string
	Company         string
	Mail            string
	TelephoneNumber string
	Office          string
	OfficePosition  string
	ObjectGUID      string
}
