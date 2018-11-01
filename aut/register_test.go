/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2018
 * |
 * | File Name:     register_test.go
 * +===============================================
 */

package aut

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	onsite, online, err := Import("../data/test.json")

	assert.NoError(t, err)

	assert.Equal(t, len(onsite), 1)
	// on-site team information
	assert.Equal(t, onsite[0].Name, "703")
	assert.Equal(t, onsite[0].Site, "Iran")
	assert.Equal(t, onsite[0].Institute, "Amirkabir University of Technology")
	// on-site first member
	assert.Equal(t, onsite[0].Members.First.FirstName, "Parham")
	assert.Equal(t, onsite[0].Members.First.LastName, "Alvani")
	assert.Equal(t, onsite[0].Members.First.Email, "parham.alvani@gmail.com")
	// on-site second member
	assert.Equal(t, onsite[0].Members.Second.FirstName, "S")
	assert.Equal(t, onsite[0].Members.Second.LastName, "Z")
	assert.Equal(t, onsite[0].Members.Second.Email, "parham.alvani@gmail.com")
	// on-site third member
	assert.Equal(t, onsite[0].Members.Third.FirstName, "E")
	assert.Equal(t, onsite[0].Members.Third.LastName, "J")
	assert.Equal(t, onsite[0].Members.Third.Email, "parham.alvani@gmail.com")

	assert.Equal(t, len(online), 1)
	// on-line team information
	assert.Equal(t, online[0].Name, "703")
	assert.Equal(t, online[0].Site, "Iran")
	assert.Equal(t, online[0].Institute, "Amirkabir University of Technology")
	// on-line first member
	assert.Equal(t, online[0].Members.First.FirstName, "Parham")
	assert.Equal(t, online[0].Members.First.LastName, "Alvani")
	assert.Equal(t, online[0].Members.First.Email, "parham.alvani@gmail.com")
}
