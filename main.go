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
	"strconv"
	"strings"

	"github.com/aut-icpc/imex/lib"
	"github.com/aut-icpc/imex/mail"
)

func main() {
	var rs map[string]map[string]imex.Register
	var fn string

	fmt.Printf("AUT-ICPC JSON Filename: ")
	fmt.Scanf("%s", &fn)

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(f).Decode(&rs); err != nil {
		panic(err)
	}
	fmt.Println(rs)

	// Onsite Teams
	// generates password and account number for each team.
	// final results are written to files.
	fts, err := os.Create("teams-onsite.tsv")
	if err != nil {
		panic(err)
	}
	fas, err := os.Create("accounts-onsite.tsv")
	if err != nil {
		panic(err)
	}
	fus, err := os.Create("userpass-onsite.csv")
	if err != nil {
		panic(err)
	}
	fts.WriteString("teams\t1\n")
	fas.WriteString("accounts\t1\n")
	for strID, r := range rs["onsite"] {
		id, _ := strconv.Atoi(strID)
		t := imex.Team{
			Number:       id + 1,
			EId:          id + 100,
			GId:          3,
			Name:         r.Name,
			Institution:  r.Institute,
			CountryCode:  imex.CountryCodes[r.Site],
			PrimaryEmail: r.Members["first"].Email,
		}
		fmt.Printf("%+v\n", t)
		fts.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.EId, t.GId, t.Name, t.Institution, t.InstitutionCode, t.CountryCode))
		a := imex.Account{
			Type:     "team",
			FullName: t.Name,
			Username: fmt.Sprintf("%03d", t.Number),
			Password: fmt.Sprintf("p%d", rand.Intn(100)),
		}
		fmt.Printf("%+v\n", a)
		fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password))
		fus.WriteString(fmt.Sprintf("%s, %s, %s\n", a.FullName, a.Username, a.Password))
	}
	return

	// Online Teams
	// generates password and account number for each team.
	// final results are written to files and sent to their emails.
	fto, err := os.Create("teams-online.tsv")
	if err != nil {
		panic(err)
	}
	fao, err := os.Create("accounts-online.tsv")
	if err != nil {
		panic(err)
	}
	fto.WriteString("teams\t1\n")
	fao.WriteString("accounts\t1\n")
	for strID, r := range rs["online"] {
		id, _ := strconv.Atoi(strID)
		t := imex.Team{
			Number:       id + 1,
			EId:          id + 100,
			GId:          3,
			Name:         r.Name,
			Institution:  r.Institute,
			CountryCode:  imex.CountryCodes[strings.Title(r.Site)],
			PrimaryEmail: r.Members["first"].Email,
		}
		fmt.Printf("%+v\n", t)
		fto.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.EId, t.GId, t.Name, t.Institution, t.InstitutionCode, t.CountryCode))
		a := imex.Account{
			Type:     "team",
			FullName: t.Name,
			Username: fmt.Sprintf("%03d", t.Number),
			Password: fmt.Sprintf("P%03d%03d", rand.Intn(1000), rand.Intn(1000)),
		}
		fmt.Printf("%+v\n", a)
		fao.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password))
		mail.SendMail(a.FullName, a.Username, a.Password, t.PrimaryEmail)
	}
}
