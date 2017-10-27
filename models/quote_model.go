package models

import (

)

type Quote struct {
	Kind 			string		`json:"kind"`
	Id 				string		`json:"id"`
	Created			string		`json:"created"`
	Expires			string		`json:"expires"`
	Fee 			int			`json:"fee"`
	Currency		string		`json:"currency"`
	Dropoff_eta 	string		`json:"dropoff_eta"`
	Duration  		int			`json:"duration"`
}