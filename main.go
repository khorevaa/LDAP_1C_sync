package main

import (
	"ldap_client/configs"
	"ldap_client/models"
	"ldap_client/services"
)

func main() {
	// Database config
	configs.ReadConfig("config.json")
	models.ConnectDB(configs.DBconfig())
	// Try to connect to LDAP server
	models.LdapConnect(configs.LDAPconfig())
	defer models.CloseLdapConnection()

	employees := services.ReadEmployeeFromLdap()
	for _, emp := range employees {
		//controllers.PrintEmployee(emp)
		services.AddToDB(emp)
	}

}
