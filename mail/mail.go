/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 03-11-2017
 * |
 * | File Name:     main.go
 * +===============================================
 */

package mail

import (
	"bytes"
	"html/template"

	"github.com/AUT-CEIT-SSC/ICPC-imex/domjudge"
	gomail "github.com/go-mail/mail"
)

type mailInfo struct {
	Name    string
	Team    string
	Account domjudge.Account
}

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("mail").Parse(tmplString)
	if err != nil {
		panic(err)
	}
}

// SendMail sends mail into given mail address with template that is defined in mail.tmpl.
// This mail contains team and account information and sends per member.
func SendMail(t domjudge.Team, a domjudge.Account) error {
	for _, m := range t.Members {

		if m.Email == "" { // do not send an email to who has no email.
			continue
		}

		buf := bytes.NewBufferString("")
		if err := tmpl.Execute(buf, mailInfo{
			Name:    m.FirstName,
			Team:    t.Name,
			Account: a,
		}); err != nil {
			return err
		}

		ms := gomail.NewMessage()
		ms.SetHeader("From", "ceit.ssc94@gmail.com")
		ms.SetHeader("To", m.Email)
		ms.SetAddressHeader("Bcc", "parham.alvani@gmail.com", "1995parham")
		ms.SetHeader("Subject", "18th AUT ACM ICPC")
		ms.SetBody("text/html", buf.String())

		d := gomail.NewDialer("smtp.gmail.com", 465, "ceit.ssc94@gmail.com", "anjomananjoman13")
		d.StartTLSPolicy = gomail.MandatoryStartTLS

		if err := d.DialAndSend(ms); err != nil {
			return err
		}
	}

	return nil
}
