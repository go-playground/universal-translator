package so

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DJF", DisplayName: "Faran Jabbuuti", Symbol: ""},
		{Currency: "ETB", DisplayName: "Birta Itoobbiya", Symbol: ""},
		{Currency: "EUR", DisplayName: "Yuuroo", Symbol: ""},
		{Currency: "SAR", DisplayName: "Riyaalka Sacuudiga", Symbol: ""},
		{Currency: "SOS", DisplayName: "Shilin soomaali", Symbol: "S"},
		{Currency: "TZS", DisplayName: "Shilin Tansaani", Symbol: ""},
		{Currency: "USD", DisplayName: "Doollar maraykan", Symbol: ""},
		{Currency: "XXX", DisplayName: "Lacag aan la qoon ama aan saxnayn", Symbol: ""},
	}
}
