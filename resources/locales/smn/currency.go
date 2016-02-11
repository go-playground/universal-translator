package smn

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DKK", DisplayName: "Tanska ruvnâ", Symbol: ""},
		{Currency: "EEK", DisplayName: "Eesti ruvnâ", Symbol: ""},
		{Currency: "EUR", DisplayName: "euro", Symbol: ""},
		{Currency: "FIM", DisplayName: "Suomâ märkki", Symbol: ""},
		{Currency: "ISK", DisplayName: "Island ruvnâ", Symbol: ""},
		{Currency: "LVR", DisplayName: "Latvia ruble", Symbol: ""},
		{Currency: "NOK", DisplayName: "Taažâ ruvnâ", Symbol: ""},
		{Currency: "SEK", DisplayName: "Ruotâ ruvnâ", Symbol: ""},
	}
}
