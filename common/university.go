package common

// InstitutionCodes contains a map between 3 character code of universities in the contest and their name.
// AUT-ICPC website works with institution name so we must have a way to convert them.
func InstitutionCodes(institution string) string {
	return map[string]string{
		"Amirkabir University of Technology": "AUT",
		"Sharif University of Technology":    "SUT",
	}[institution]
}
