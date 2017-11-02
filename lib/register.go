/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     types/register.go
 * +===============================================
 */

package imex

// Register of AUT-ICPC
type Register struct {
	Name      string `json:"team_name"`
	Institute string `json:"institute_name"`
	Site      string
	Members   map[string]rMember
}

type rMember struct {
	Email string `json:"email"`
}
