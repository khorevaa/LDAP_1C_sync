package controllers

import (
	"fmt"
	"ldap_client/models"
)

// PrintEmployee ...
func PrintEmployee(emp models.RootEmployee) {
	fmt.Println(emp.Name)
	fmt.Println(emp.Surname)
	fmt.Println(emp.Birthday)
	fmt.Println(emp.Position)
	fmt.Println(emp.Department)
	fmt.Println(emp.Company)
	fmt.Println(emp.Mail)
	fmt.Println(emp.TelephoneNumber)
	fmt.Println(emp.Office)
	fmt.Println(emp.OfficePosition)
	fmt.Println(emp.ObjectGUID)
	fmt.Println()
}
