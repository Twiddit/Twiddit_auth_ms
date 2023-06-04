package utils

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

const (
	AdminUsername = "cn=admin,dc=twiddit,dc=unal,dc=edu,dc=co"
	AdminPassword = "admin"
	BaseDN        = "ou=users,dc=twiddit,dc=unal,dc=edu,dc=co"
	Filter        = "(objectClass=*)"
	IP            = "34.138.201.211"
)

func Connect() (*ldap.Conn, error) {
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:389", IP))
	if err != nil {
		return nil, err
	}

	return l, nil
}

func BindAndSearch(l *ldap.Conn, email string, pwd string) (*ldap.SearchResult, error) {
	dn := fmt.Sprintf("cn=%s,%s", email, BaseDN)
	l.Bind(dn, pwd)
	searchReq := ldap.NewSearchRequest(
		dn,
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

func CreateUser(l *ldap.Conn, username string, mail string, pwd string) error {
	l.Bind(AdminUsername, AdminPassword)
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", mail, BaseDN), []ldap.Control{})
	//changetype: add
	addReq.Attribute("objectClass", []string{"inetOrgPerson"})
	addReq.Attribute("cn", []string{mail})
	addReq.Attribute("givenname", []string{username})
	addReq.Attribute("sn", []string{"Cyberplace"})
	addReq.Attribute("displayname", []string{username})
	addReq.Attribute("mail", []string{mail})
	addReq.Attribute("userpassword", []string{pwd})
	err := l.Add(addReq)
	if err != nil {
		return err
	}
	return nil
}

func ResetPwd(l *ldap.Conn, username string, oldPassword string, newPassword string) error {
	dn := fmt.Sprintf("cn=%s,%s", username, BaseDN)
	l.Bind(AdminUsername, AdminPassword)
	passwdModReq := ldap.NewPasswordModifyRequest(dn, oldPassword, newPassword)
	_, err := l.PasswordModify(passwdModReq)
	if err != nil {
		return err
	}
	return nil
}
