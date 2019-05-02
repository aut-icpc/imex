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

	"github.com/aut-icpc/imex/domjudge"
	gomail "github.com/go-mail/mail"
)

// mailInfo is a collection of information provided for mail template
type mailInfo struct {
	Name    string
	Team    string
	Account domjudge.Account
}

// Mailer sends email using gmail mail server.
type Mailer struct {
	Addr    string
	Pass    string
	Subject string

	tmpl *template.Template
}

// New creates a new mailer
func New(addr string, pass string, subject string) (Mailer, error) {

	tmpl, err := template.New("mail").Parse(tmplString)
	if err != nil {
		return Mailer{}, err
	}

	return Mailer{
		Addr:    addr,
		Pass:    pass,
		Subject: subject,

		tmpl: tmpl,
	}, nil
}

// SendMail sends mail into given mail address with template that is defined in mail.tmpl.
// This mail contains team and account information and sends per member.
func (mailer Mailer) SendMail(t domjudge.Team, a domjudge.Account) error {
	for _, m := range t.Members {

		if m.Email == "" { // do not send an email to who has no email.
			continue
		}

		buf := bytes.NewBufferString("")
		if err := mailer.tmpl.Execute(buf, mailInfo{
			Name:    m.FirstName,
			Team:    t.Name,
			Account: a,
		}); err != nil {
			return err
		}

		ms := gomail.NewMessage()
		ms.SetHeader("From", mailer.Addr)
		ms.SetHeader("To", m.Email)
		ms.SetAddressHeader("Bcc", "parham.alvani@gmail.com", "Parham Alvani")
		ms.SetHeader("Subject", mailer.Subject)
		ms.SetBody("text/html", buf.String())

		d := gomail.NewDialer("smtp.gmail.com", 465, mailer.Addr, mailer.Pass)
		d.StartTLSPolicy = gomail.MandatoryStartTLS

		if err := d.DialAndSend(ms); err != nil {
			return err
		}
	}

	return nil
}
