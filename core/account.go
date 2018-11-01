/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     account.go
 * +===============================================
 */

package core

// Account represents domjudge account.
// each account assigns to a team.
type Account struct {
	Type     string
	FullName string
	Username string
	Password string
}
