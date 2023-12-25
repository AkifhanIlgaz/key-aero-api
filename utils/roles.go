package utils

import "strings"

const (
	AdminRole = "admin"
)

func ParseRoles(roles string) []string {
	return strings.Split(roles, ",")
}

func GenerateRolesString(roles []string) string {
	return strings.Join(roles, ",")
}
