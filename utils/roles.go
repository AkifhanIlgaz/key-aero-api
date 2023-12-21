package utils

import "strings"

// roles => <role1>,<role2>...
// TODO: Role enum || constants
func ParseRoles(roles string) []string {
	return strings.Split(roles, ",")
}
