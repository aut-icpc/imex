/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/aut-icpc/imex/lib"
)

func main() {
	var rs map[string]map[int]imex.Register
	var fn string

	fmt.Printf("AUT-ICPC JSON Filename: ")
	fmt.Scanf("%s", &fn)

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(f).Decode(&rs)

	// Onsite
	fts, err := os.Create("teams-onsite.tsv")
	if err != nil {
		panic(err)
	}
	fas, err := os.Create("accounts-onsite.tsv")
	if err != nil {
		panic(err)
	}
	fus, err := os.Create("userpass-onsite.txt")
	if err != nil {
		panic(err)
	}
	fts.WriteString("onsite-teams\t1\n")
	fas.WriteString("onsite-accounts\t1\n")
	for id, r := range rs["onsite"] {
		t := imex.Team{
			Number:       id + 100,
			Name:         r.Name,
			Institution:  r.Institute,
			CountryCode:  imex.CountryCodes[r.Site],
			PrimaryEmail: r.Members["first"].Email,
		}
		fts.WriteString(fmt.Sprintf("%d\t\t3\t%s\t%s\t%s\t%s\n", t.Number, t.Name, t.Institution, t.InstitutionCode, t.CountryCode))
		a := imex.Account{
			Type:     "team",
			FullName: t.Name,
			Username: fmt.Sprintf("t-%3d", t.Number),
			Password: fmt.Sprintf("p%d", rand.Intn(100)),
		}
		fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password))
		fus.WriteString(fmt.Sprintf("- %s: %s %s\n", a.FullName, a.Username, a.Password))
	}

	// Online
	fto, err := os.Create("teams-online.tsv")
	if err != nil {
		panic(err)
	}
	fao, err := os.Create("accounts-online.tsv")
	if err != nil {
		panic(err)
	}
	fto.WriteString("online-teams\t1\n")
	fao.WriteString("online-accounts\t1\n")
	for id, r := range rs["online"] {
		t := imex.Team{
			Number:       id + 200,
			Name:         r.Name,
			Institution:  r.Institute,
			CountryCode:  imex.CountryCodes[strings.Title(r.Site)],
			PrimaryEmail: r.Members["first"].Email,
		}
		fto.WriteString(fmt.Sprintf("%d\t\t3\t%s\t%s\t%s\t%s\n", t.Number, t.Name, t.Institution, t.InstitutionCode, t.CountryCode))
		a := imex.Account{
			Type:     "team",
			FullName: t.Name,
			Username: fmt.Sprintf("t-%3d", t.Number),
			Password: fmt.Sprintf("P%d", rand.Intn(100)),
		}
		fao.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password))
	}
}
