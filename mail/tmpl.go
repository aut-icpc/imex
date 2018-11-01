/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2018
 * |
 * | File Name:     tmpl.go
 * +===============================================
 */

package mail

var tmplString = `
<html>
        <body>
                <h1>ACM ICPC AUT 18th</h1>
                <p>Dear {{.Name}} of {{.Team}}</p>
                <p>Please write down your login information:</p>
                <p><strong>Team {{.Team}}</strong></p>
                <p><strong>Your username: {{.Account.Username}}</strong></p>
                <p><strong>Your Password: {{.Account.Password}}</strong></p>

                <p>
                        Wishing ever more success
                        Team of the 18th International AUT ACM ICPC
                </p>
        </body>
</html>
`
