package models

import (
	"time"
)

//Company ...
type Company struct {
	ID   uint64 `sql:"type:bigserial; primaryKey"`
	Name string `sql:"type:text; unique"`
}

//Department ...
type Department struct {
	ID   uint64 `sql:"type:bigserial; primaryKey"`
	Name string `sql:"type:text; unique"`
}

//Position ...
type Position struct {
	ID   uint64 `sql:"type:bigserial; primaryKey"`
	Name string `sql:"type:text; unique"`
}

//Employee ...
type Employee struct {
	ID              uint64    `sql:"type:bigserial; primaryKey"`
	LdapID          string    `sql:"type:char(32); unique_index:idx_employeeNatural"`
	Name            string    `sql:"type:text; index:idx_fullname"`
	Surname         string    `sql:"type:text; index:idx_fullname"`
	Birthday        time.Time `sql:"type:date"`
	PositionID      uint64    `sql:"type:bigserial REFERENCES positions(id)"`
	DepartmentID    uint64    `sql:"type:bigserial REFERENCES departments(id)"`
	CompanyID       uint64    `sql:"type:bigserial REFERENCES companies(id)"`
	Mail            string    `sql:"type:text"`
	TelephoneNumber string    `sql:"type:text"`
	Office          string    `sql:"type:text"`
	OfficePosition  string    `sql:"type:text"`
}
