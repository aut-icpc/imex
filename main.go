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
	"flag"
	"fmt"
	"os"

	"github.com/AUT-CEIT-SSC/ICPC-imex/aut"
	"github.com/AUT-CEIT-SSC/ICPC-imex/convert"
	"github.com/AUT-CEIT-SSC/ICPC-imex/domjudge"
	"github.com/AUT-CEIT-SSC/ICPC-imex/mail"
	log "github.com/sirupsen/logrus"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "./data/test.json", "path to the team export of AUT-ICPC website")
	flag.Parse()

	onsite, online, err := aut.Import(path)
	if err != nil {
		panic(err)
	}

	// Onsite Teams
	// generates password and account number for each team.
	// final results are written to files.

	// domjudge team descriptions
	fts, err := os.Create("teams-onsite.tsv")
	if err != nil {
		panic(err)
	}
	// domjudge accounts
	fas, err := os.Create("accounts-onsite.tsv")
	if err != nil {
		panic(err)
	}
	// teams username and passwords for printing.
	// these are printed for each team and putted in their boxes.
	fus, err := os.Create("userpass-onsite.csv")
	if err != nil {
		panic(err)
	}

	if _, err := fts.WriteString("teams\t1\n"); err != nil {
		panic(err)
	}
	if _, err := fas.WriteString("accounts\t1\n"); err != nil {
		panic(err)
	}
	for i, r := range onsite {
		t := convert.Convert(i+1, i+300, r)
		log.Infof("On-Site Team: %+v\n", t)

		if _, err := fts.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.EId, t.GId, t.Name, t.Institution, t.InstitutionCode, t.CountryCode)); err != nil {
			panic(err)
		}

		a := domjudge.NewOnsiteAccount(t)

		log.Infof("On-Site Account: %+v\n", a)

		if _, err := fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password)); err != nil {
			panic(err)
		}
		if _, err := fus.WriteString(fmt.Sprintf("%s, %s, %s\n", a.FullName, a.Username, a.Password)); err != nil {
			panic(err)
		}
	}

	// Online Teams
	// generates password and account number for each team.
	// final results are written to files and sent to their emails.

	// domjudge team descriptions
	fto, err := os.Create("teams-online.tsv")
	if err != nil {
		panic(err)
	}
	// domjudge accounts
	fao, err := os.Create("accounts-online.tsv")
	if err != nil {
		panic(err)
	}
	fto.WriteString("teams\t1\n")
	fao.WriteString("accounts\t1\n")
	for i, r := range online {
		t := convert.Convert(i+1, i+300, r)
		log.Infof("On-Line Team: %+v\n", t)

		if _, err := fts.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.EId, t.GId, t.Name, t.Institution, t.InstitutionCode, t.CountryCode)); err != nil {
			panic(err)
		}

		a := domjudge.NewOnlineAccount(t)

		log.Infof("On-Line Account: %+v\n", a)

		if _, err := fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password)); err != nil {
			panic(err)
		}
		if _, err := fus.WriteString(fmt.Sprintf("%s, %s, %s\n", a.FullName, a.Username, a.Password)); err != nil {
			panic(err)
		}

		if err := mail.SendMail(t, a); err != nil {
			panic(err)
		}
		log.Infof("Successfully send an email to team %s", t.Name)
	}
}
