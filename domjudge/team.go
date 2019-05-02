/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     types/team.go
 * +===============================================
 */

package domjudge

import "github.com/aut-icpc/imex/common"

// Team represents domjudge team.
type Team struct {
	Number          int
	ExternalID      int
	GroupID         int
	Name            string
	Institution     string
	InstitutionCode string
	CountryCode     string
	Members         [3]common.Member
}
