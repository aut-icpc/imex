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

// before contest email. this mail contains just general information about online contest.
/*
var tmplString = `
<html>
        <body>
                <h1>ACM ICPC AUT 18th</h1>
                <p>Dear {{.Name}} of the team {{.Team}}</p>
                <p>
		Today's contest is scheduled to start at 2:30 p.m. Your team's username, password,
		and any further information will be sent to you 30 minutes before the contest.
		</p>
		<p>
		The contest is available <a href="http://daavar.ceit-ssc.ir/domjudge/login">here</a>.
		</p>
        </body>
</html>
`
*/

// on contest mail. this mail contains the login information.
const tmplString = `
<html>
        <body>
		<h1>ACM ICPC AUT 18th</h1>
                <p>Dear {{.Name}} of the team {{.Team}}</p>
		<p>
		The contest is available <a href="http://daavar.ceit-ssc.ir/domjudge/login">here</a>.
		</p>
		<p>
		Username: {{.Account.Username}}<br>
		Password: {{.Account.Password}}
		</p>
        </body>
</html>
`
