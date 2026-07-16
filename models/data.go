package models

// KFSData represents the dynamic fields present in the Key Facts Statement (KFS) document.
// The field names match the placeholders used in the HTML template exactly.
type KFSData struct {
	LoadAccountNumber    string // Matches {{.Data.LoadAccountNumber}}
	SanctionedLoanAmount string // Matches {{.Data.SanctionedLoanAmount}}
	LoanTerm             string // Matches {{.Data.LoanTerm}}
	BenchmarkRate        string // Matches {{.Data.BenchmarkRate}}
	Spread               string // Matches {{.Data.Spread}}
	FinatRate            string // Matches {{.Data.FinatRate}}
	APR                  string // Matches {{.Data.APR}}
	RateOfInterest       string // Matches {{.Data.RateOfInterest}}
}
