/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     types/team.go
 * +===============================================
 */

package imex

// Team of domjudge
type Team struct {
	Number          int
	EId             int
	GId             int
	Name            string
	Institution     string
	InstitutionCode string
	CountryCode     string
	PrimaryEmail    string
}
