package services

import (
	"fmt"
	"ldap_client/models"
	"log"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
)

// ReadEmployeeFromLdap ...
func ReadEmployeeFromLdap() []models.RootEmployee {
	// Search for the given username
	// Query for other attributes related to the user
	searchRequest := ldap.NewSearchRequest(
		// The base domain name to search
		"dc=office,dc=vedomosti,dc=ru",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// The filter to apply
		//extensionAttribute3=*: не пустой атрибут
		"(&(objectClass=person)(!(objectclass=computer))(!(objectclass=contact))(extensionAttribute3=*))",
		// List of attributes to retrieve
		[]string{
			"extensionAttribute3", //name
			"extensionAttribute4", //surname
			"extensionAttribute2", //birthdate
			"company",             //company
			"department",          //департамент
			"mail",
			"info",
			"telephoneNumber",
			"title", //office/officePosition
			"objectGUID"},
		nil,
	)
	searchResult, err := models.LdapConnection.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	var employees []models.RootEmployee
	// Prints the result with 2 spaces for indentation
	for i, emp := range searchResult.Entries {
		employees = append(employees, models.RootEmployee{})
		for _, attribute := range emp.Attributes {
			val := attribute.Values[0]
			switch attribute.Name {
			case "extensionAttribute3":
				employees[i].Name = val
			case "extensionAttribute4":
				employees[i].Surname = val
			case "extensionAttribute2":
				employees[i].Birthday, err = time.Parse("01-02-2006", val)
				if err != nil {
					fmt.Printf("error: не удалось запарсить birthday")
				}
			case "company":
				employees[i].Company = val
			case "department":
				employees[i].Department = val
			case "mail":
				employees[i].Mail = val
			case "info":
				info := strings.Split(val, ":")
				employees[i].Position = info[len(info)-1]
			case "telephoneNumber":
				employees[i].TelephoneNumber = val
			case "title":
				title := strings.Split(val, "/")
				employees[i].Office = title[0]
				if len(title) > 1 {
					employees[i].OfficePosition = title[1]
				}
			case "objectGUID":
				hash := []byte(val)
				val = fmt.Sprintf("%x", hash)
				employees[i].ObjectGUID = val
			}
		}
	}
	return employees
}
