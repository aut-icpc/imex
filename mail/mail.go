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

	gomail "gopkg.in/gomail.v2"
)

type mailInfo struct {
	Name string
	User string
	Pass string
}

// SendMail sends mail into given mail
func SendMail(n string, u string, p string, e string) {
	buf := bytes.NewBufferString("")
	tmpl := template.Must(template.ParseFiles("./mail/mail.tmpl"))
	tmpl.Execute(buf, mailInfo{
		Name: n,
		User: u,
		Pass: p,
	})

	m := gomail.NewMessage()
	m.SetHeader("From", "ceit.ssc94@gmail.com")
	m.SetHeader("To", e)
	m.SetAddressHeader("Bcc", "parham.alvani@gmail.com", "1995parham")
	m.SetHeader("Subject", "Hi")
	m.SetBody("text/html", buf.String())

	d := gomail.NewDialer("smtp.gmail.com", 465, "ceit.ssc94@gmail.com", "anjomananjoman13")
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
