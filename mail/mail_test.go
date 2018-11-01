/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-11-2018
 * |
 * | File Name:     mail_test.go
 * +===============================================
 */

package mail

import (
	"testing"

	"github.com/AUT-CEIT-SSC/ICPC-imex/common"
	"github.com/AUT-CEIT-SSC/ICPC-imex/domjudge"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	account := domjudge.Account{
		Username: "1995parham",
		Password: "123123",
	}
	team := domjudge.Team{
		Name: "Travis",
		Members: [3]common.Member{
			common.Member{
				FirstName: "Parham",
				Email:     "parham.alvani@gmail.com",
			},
		},
	}
	err := SendMail(team, account)
	assert.NoError(t, err)
}
