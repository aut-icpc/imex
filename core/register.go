/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     types/register.go
 * +===============================================
 */

package core

// Register contains registration data of AUT-ICPC website available
// on http://icpc.aut.ac.ir/
type Register struct {
	Name      string `json:"team_name"`
	Institute string `json:"institute_name"`
	Site      string `json:"site"`
	Members   map[string]struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"members"`
}
