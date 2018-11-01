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
	"crypto/tls"
	"html/template"

	gomail "github.com/go-mail/mail"
)

type mailInfo struct {
	Name string
	User string
	Pass string
}

var tmpl *template.Template

// parse template in initation phase
func init() {
	tmpl = template.Must(template.ParseFiles("./mail.tmpl"))
}

// SendMail sends mail into given mail address with template that is defined in mail.tmpl.
// This mail contains username (u), password (p) of team (n) that has given email address (e).
func SendMail(n string, u string, p string, e string) error {
	buf := bytes.NewBufferString("")
	if err := tmpl.Execute(buf, mailInfo{
		Name: n,
		User: u,
		Pass: p,
	}); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "ceit.ssc94@gmail.com")
	m.SetHeader("To", e)
	m.SetAddressHeader("Bcc", "parham.alvani@gmail.com", "1995parham")
	m.SetHeader("Subject", "18th AUT ACM ICPC")
	m.SetBody("text/html", buf.String())

	d := gomail.NewDialer("smtp.gmail.com", 465, "ceit.ssc94@gmail.com", "anjomananjoman13")
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
