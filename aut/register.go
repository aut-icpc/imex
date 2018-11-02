/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     types/register.go
 * +===============================================
 */

package aut

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/AUT-CEIT-SSC/ICPC-imex/common"
)

// Register contains registration data of AUT-ICPC website available
// on http://icpc.aut.ac.ir/
type Register struct {
	Name      string `json:"team_name"`
	Institute string `json:"institute_name"`
	Site      string `json:"site"`
	Members   struct {
		First  common.Member `json:"first"`
		Second common.Member `json:"second"`
		Third  common.Member `json:"third"`
	} `json:"members"`
	Status struct {
		Status string `json:"status"`
	} `json:"status"`
}

// Import imports data from a json file that is given by its path.
// Please note this function expects AUT-ICPC format (refer to data/test.json for more information).
func Import(path string) (onsite map[int]Register, online map[int]Register, err error) {
	var rs map[string]map[string]Register

	f, err := os.Open(path)
	if err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&rs)
	if err != nil {
		return
	}

	onsite = make(map[int]Register)
	for s, r := range rs["onsite"] {
		var i int
		i, err = strconv.Atoi(s)
		if err != nil {
			return
		}

		if r.Status.Status != "Finalized" {
			continue
		}
		onsite[i] = r
	}

	online = make(map[int]Register)
	for s, r := range rs["online"] {
		var i int
		i, err = strconv.Atoi(s)
		if err != nil {
			return
		}

		if r.Status.Status != "Finalized" {
			continue
		}
		online[i] = r
	}

	return
}
