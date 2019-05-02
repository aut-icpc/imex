package evand

import (
	"strconv"

	"github.com/AUT-CEIT-SSC/ICPC-imex/common"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

// Register contains registration data of AUT-ICPC website available
// on http://icpc.aut.ac.ir/
type Register struct {
	Name      string
	Institute string
	Members   struct {
		First  common.Member
		Second common.Member
	}
	Status struct {
		Status string
	}
}

// Import imports data from a excel file that is given by its path.
// Please note this function expects Evand (https://evand.com/)
// format (refer to data/apl.xlsx for more information).
func Import(path string) (registers []Register, err error) {
	f, err := xlsx.OpenFile(path)
	if err != nil {
		return
	}

	// use the first sheet
	sheet := f.Sheets[0]

	// remove the header row
	rows := sheet.Rows[1:]

	for n, row := range rows {
		if n%2 == 0 {
			// first team member
			if _, err := strconv.ParseFloat(row.Cells[0].Value, 64); err != nil {
				logrus.Info(err)
				break
			}
			register := Register{
				Name:      row.Cells[1].Value,
				Institute: row.Cells[10].Value,
			}
			register.Members.First = common.Member{
				FirstName: row.Cells[4].Value,
				LastName:  row.Cells[6].Value,
				TShirt:    row.Cells[7].Value,
				StudentID: row.Cells[9].Value,
			}
			registers = append(registers, register)
		} else {
			// the second team member
			registers[len(registers)-1].Members.Second = common.Member{
				FirstName: row.Cells[4].Value,
				LastName:  row.Cells[6].Value,
				TShirt:    row.Cells[7].Value,
				StudentID: row.Cells[9].Value,
			}
		}
	}

	return
}
