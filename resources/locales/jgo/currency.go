package jgo

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CAD", DisplayName: "Ndɔ́la-Kanandâ", Symbol: ""},
		{Currency: "EUR", DisplayName: "Ʉ́lɔ", Symbol: ""},
		{Currency: "USD", DisplayName: "Ndɔ́la-Amɛlîk", Symbol: ""},
		{Currency: "XAF", DisplayName: "Fɛlâŋ", Symbol: "FCFA"},
		{Currency: "XXX", DisplayName: "ntɛ-ŋkáp yi pɛ́ ká kɛ́ jínɛ", Symbol: ""},
	}
}
