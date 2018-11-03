/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2018
 * |
 * | File Name:     convert.go
 * +===============================================
 */

package convert

import (
	"github.com/AUT-CEIT-SSC/ICPC-imex/aut"
	"github.com/AUT-CEIT-SSC/ICPC-imex/common"
	"github.com/AUT-CEIT-SSC/ICPC-imex/domjudge"
)

// Convert converts team registration information from AUT-ICPC to domjudge team.
// please note that domjudge need an unique identifier for each team. If this identifier
// is not unique then old record in domjudge will override.
// ExternalID is not important and you can even do not pass it to Domjudge.
func Convert(id int, r aut.Register) domjudge.Team {
	return domjudge.Team{
		Number:          id,
		ExternalID:      id,
		GroupID:         3,
		Name:            r.Name,
		Institution:     r.Institute,
		InstitutionCode: common.InstitutionCodes[r.Institute],
		CountryCode:     common.CountryCodes[r.Site],
		Members: [3]common.Member{
			r.Members.First,
			r.Members.Second,
			r.Members.Third,
		},
	}
}
