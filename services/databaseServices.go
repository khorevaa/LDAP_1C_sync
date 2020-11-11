package services

import (
	"ldap_client/models"
)

// AddToDB ...
func AddToDB(emp models.RootEmployee) {
	company := models.Company{Name: emp.Company}
	models.DB.Create(&company)
	if company.ID == 0 {
		models.DB.Where("name = ?", emp.Company).Find(&company)
	}
	department := models.Department{Name: emp.Department}
	models.DB.Create(&department)
	if department.ID == 0 {
		models.DB.Where("name = ?", emp.Department).Find(&department)
	}
	position := models.Position{Name: emp.Position}
	models.DB.Create(&position)
	if position.ID == 0 {
		models.DB.Where("name = ?", emp.Position).Find(&position)
	}
	employee := models.Employee{
		LdapID:          emp.ObjectGUID,
		Name:            emp.Name,
		Surname:         emp.Surname,
		Birthday:        emp.Birthday,
		PositionID:      position.ID,
		DepartmentID:    department.ID,
		CompanyID:       company.ID,
		Mail:            emp.Mail,
		TelephoneNumber: emp.TelephoneNumber,
		Office:          emp.Office,
		OfficePosition:  emp.OfficePosition}
	models.DB.Create(&employee)
	if employee.ID == 0 {
		newEmployee := employee
		models.DB.Where("ldap_id = ?", emp.ObjectGUID).Find(&employee)
		models.DB.Model(&employee).Update(newEmployee)
	}
}
