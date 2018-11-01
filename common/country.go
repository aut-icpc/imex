/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     country/country.go
 * +===============================================
 */

package common

// CountryCodes contains a map between 3 character code of countries around the world and their name.
// AUT-ICPC website works with country name so we must have a way to convert them.
var CountryCodes = map[string]string{
	"Iran":                     "IRN",
	"Germany":                  "DEU",
	"Sweden":                   "SWE",
	"United States of America": "USA",
}
