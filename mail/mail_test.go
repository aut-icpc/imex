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

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	err := SendMail("Travis", "1995parham", "123123", "parham.alvani@gmail.com")
	assert.NoError(t, err)
}
