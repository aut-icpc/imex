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
	"github.com/AUT-CEIT-SSC/ICPC-imex/evand"
	"github.com/AUT-CEIT-SSC/ICPC-imex/mail"
	log "github.com/sirupsen/logrus"
)

func main() {
	var path string
	var isSendMail bool
	var tp string

	flag.StringVar(&path, "path", "./data/test.json, ./data/apl.xslx", "path to the team export of AUT-ICPC website or Evand excel")
	flag.BoolVar(&isSendMail, "sendmail", false, "send mail to online contestant")
	flag.StringVar(&tp, "type", "aut", "aut, evand")

	flag.Parse()

	switch tp {
	case "aut":
		AUT(path, isSendMail)
	case "evand":
		Evand(path)
	default:
		fmt.Println()
	}
}

// Evand parse Evand excel report and creates the required file.
func Evand(path string) {
	onsite, err := evand.Import(path)
	if err != nil {
		log.Fatal(err)
	}

	// Onsite Teams
	// generates password and account number for each team.
	// final results are written to files.

	// domjudge team descriptions
	fts, err := os.Create("teams-onsite.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// domjudge accounts
	fas, err := os.Create("accounts-onsite.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// teams username and passwords for printing.
	// these are printed for each team and putted in their boxes.
	fus, err := os.Create("userpass-onsite.csv")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := fts.WriteString("teams\t1\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := fas.WriteString("accounts\t1\n"); err != nil {
		log.Fatal(err)
	}
	for i, r := range onsite {
		t := convert.Convert1(i+400, r)
		log.Infof("On-Site Team: %+v\n", t)

		if _, err := fts.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.ExternalID, t.GroupID, t.Name, t.Institution, t.InstitutionCode, t.CountryCode)); err != nil {
			log.Fatal(err)
		}

		a := domjudge.NewOnsiteAccount(t)

		log.Infof("On-Site Account: %+v\n", a)

		if _, err := fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password)); err != nil {
			log.Fatal(err)
		}
		if _, err := fus.WriteString(fmt.Sprintf("%s, %s, %s\n", a.FullName, a.Username, a.Password)); err != nil {
			log.Fatal(err)
		}
	}
}

// AUT parses AUT-ICPC website JSON report and creates required files.
func AUT(path string, isSendMail bool) {
	onsite, online, err := aut.Import(path)
	if err != nil {
		log.Fatal(err)
	}

	// Onsite Teams
	// generates password and account number for each team.
	// final results are written to files.

	// domjudge team descriptions
	fts, err := os.Create("teams-onsite.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// domjudge accounts
	fas, err := os.Create("accounts-onsite.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// teams username and passwords for printing.
	// these are printed for each team and putted in their boxes.
	fus, err := os.Create("userpass-onsite.csv")
	if err != nil {
		log.Fatal(err)
	}
	// teams description for reception guys. These guys check these information when teams arrive.
	frs, err := os.Create("reception-onsite.csv")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := fts.WriteString("teams\t1\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := fas.WriteString("accounts\t1\n"); err != nil {
		log.Fatal(err)
	}
	for i, r := range onsite {
		t := convert.Convert(i+300, r)
		log.Infof("On-Site Team: %+v\n", t)

		if _, err := fts.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.ExternalID, t.GroupID, t.Name, t.Institution, t.InstitutionCode, t.CountryCode)); err != nil {
			log.Fatal(err)
		}

		for _, m := range t.Members {
			if _, err := frs.WriteString(fmt.Sprintf("%s, %s, %s, %s, %s\n", t.Name, m.FirstName, m.LastName, m.TShirt, m.StudentID)); err != nil {
				log.Fatal(err)
			}
		}

		a := domjudge.NewOnsiteAccount(t)

		log.Infof("On-Site Account: %+v\n", a)

		if _, err := fas.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password)); err != nil {
			log.Fatal(err)
		}
		if _, err := fus.WriteString(fmt.Sprintf("%s, %s, %s\n", a.FullName, a.Username, a.Password)); err != nil {
			log.Fatal(err)
		}
	}
	// Online Teams
	// generates password and account number for each team.
	// final results are written to files and sent to their emails.

	// domjudge team descriptions
	fto, err := os.Create("teams-online.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// domjudge accounts
	fao, err := os.Create("accounts-online.tsv")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := fto.WriteString("teams\t1\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := fao.WriteString("accounts\t1\n"); err != nil {
		log.Fatal(err)
	}
	for i, r := range online {
		t := convert.Convert(i+300, r)
		log.Infof("On-Line Team: %+v\n", t)

		if _, err := fto.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", t.Number, t.ExternalID, t.GroupID, t.Name, t.Institution, t.InstitutionCode, t.CountryCode)); err != nil {
			log.Fatal(err)
		}

		a := domjudge.NewOnlineAccount(t)

		log.Infof("On-Line Account: %+v\n", a)

		if _, err := fao.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", a.Type, a.FullName, a.Username, a.Password)); err != nil {
			log.Fatal(err)
		}

		if isSendMail {
			if err := mail.SendMail(t, a); err != nil {
				log.Errorf("Error on sending an email to team %s -- %s", t.Name, err)
			}
			log.Infof("Successfully send an email to team %s", t.Name)
		}
	}

	log.Infof("On-Site teams: %d", len(onsite))
	log.Infof("On-Line teams: %d", len(online))
}
