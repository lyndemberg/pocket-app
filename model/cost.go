package model

import "time"

// Cost representa cada gasto presente em um mês
type Cost struct {
	name    string
	month   Month
	year    int
	created time.Time
	updated time.Time
	user    User
}
