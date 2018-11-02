/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     account.go
 * +===============================================
 */

package domjudge

import (
	"fmt"
	"math/rand"
)

// Account represents domjudge account.
// each account assigns to a team.
type Account struct {
	Type     string
	FullName string
	Username string
	Password string
}

// NewOnsiteAccount generates an account for on-site team.
// This function generates the same password between different program runs.
func NewOnsiteAccount(t Team) Account {
	return Account{
		Type:     "team",
		FullName: t.Name,
		Username: fmt.Sprintf("U%03d", t.Number),                             // team username
		Password: fmt.Sprintf("p%03d%03d", rand.Intn(1000), rand.Intn(1000)), // random password
	}

}

// NewOnlineAccount generates an account for on-line team.
// This function generates the same password between different program runs.
func NewOnlineAccount(t Team) Account {
	return Account{
		Type:     "team",
		FullName: t.Name,
		Username: fmt.Sprintf("U%03d", t.Number),                             // team username
		Password: fmt.Sprintf("P%03d%03d", rand.Intn(1000), rand.Intn(1000)), // random password
	}

}
