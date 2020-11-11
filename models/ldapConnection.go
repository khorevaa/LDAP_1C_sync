package models

import (
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

// LdapConnection ..
var LdapConnection *ldap.Conn

// LdapConnect ...
func LdapConnect(rootConfigs string) {
	configs := strings.Split(rootConfigs, " ")
	ldapServer := configs[0]
	ldapPort := configs[1]
	ldapUsername := configs[2]
	ldapPassword := configs[3]
	// Try to connect to LDAP server
	LdapCon, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapServer, ldapPort))
	if err != nil {
		log.Fatal(err)
	}
	// Reconnect with TLS
	err = LdapCon.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	// Bind as the user to verify their password
	// If this operation doesn't trigger an error then it's considered as a successful login
	err = LdapCon.Bind(ldapUsername, ldapPassword)
	if err != nil {
		log.Fatal(err)
	}
	LdapConnection = LdapCon
}

// CloseLdapConnection ...
func CloseLdapConnection() {
	LdapConnection.Close()
}
