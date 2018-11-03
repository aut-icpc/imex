# imex :left_right_arrow:
[![Travis branch](https://img.shields.io/travis/com/AUT-CEIT-SSC/ICPC-imex/master.svg?style=flat-square)](https://travis-ci.com/AUT-CEIT-SSC/ICPC-imex)

## Introduction
Teams register in ICPC website and after that we export them into json file. This stanalone program read that json file
and create accounts and teams file in TSV.

## Step by Step
1. Export data from AUT-ICPC
2. Run :runner:
3. Import teams
4. Import accounts

## File Formats
### Accounts.tsv

A text file consisting of a version line and one line for each non-team account in the contest. Each line has tab separated fields as defined below.

The first line has the following format

| Field | Description |	Example | Type |
|:----- |:----------- |:------- |:---- |
| 1     | Label | accounts | fixed string (always same value) |
| 2	| Version number | 1 | integer |

Then follow several lines with the following format (one per account).

| Field | Description |	Example | Type |
|:----- |:----------- |:------- |:---- |
| 1     | Account Type | judge | string |
| 2     | Full Name | Per Austrin | string |
| 3     | Username | austrin | string |
| 4     | Password | B!5MWJiy | string |

Account Types are: team, judge, admin, analyst.

For accounts of type team username is on the form "team-nnn" where "nnn" is the zero padded team number as given in teams.tsv.


### Teams.tsv

A text file consisting of a version line and one line for each team in the contest. Each line has tab separated fields as defined below.

The first line has the following format

| Field | Description | Example | Type |
|:----- |:----------- |:------- |:---- |
| 1     | Label	teams | fixed | string (always same value) |
| 2     | Version number | 1 | integer |

Then follow several lines with the following format (one per team).

| Field | Description | Example | Type |
|:----- |:----------- |:------- |:---- |
| 1     | Team Number | 22 | integer |
| 2     | External ID | 24314 | integer |
| 3     | Group ID | 4 | integer |
| 4     | Team name | Hoos | string |
| 5     | Institution name | University of Virginia | string |
| 6     | Institution short name | U Virginia | string |
| 7     | Country Code | USA | string [ISO 3166-1 alpha-3](http://www.nationsonline.org/oneworld/country_code_list.htm) |
