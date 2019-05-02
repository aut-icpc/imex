package evand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	registers, err := Import("../data/apl.xlsx")

	assert.NoError(t, err)
	assert.Equal(t, len(registers), 69)
	assert.Equal(t, registers[0].Name, ";—null")
	assert.Equal(t, registers[0].Members.First.FirstName, "Sara")
}
