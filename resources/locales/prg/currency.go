package prg

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "Brazīlijas reals", Symbol: ""},
		{Currency: "CNY", DisplayName: "Kīnas juāns", Symbol: ""},
		{Currency: "EUR", DisplayName: "eurō", Symbol: ""},
		{Currency: "GBP", DisplayName: "punds sterlings", Symbol: ""},
		{Currency: "INR", DisplayName: "Īndijas rūpija", Symbol: ""},
		{Currency: "JPY", DisplayName: "Japānijas jāns", Symbol: ""},
		{Currency: "RUB", DisplayName: "Russis rūbels", Symbol: ""},
		{Currency: "USD", DisplayName: "APW dālars", Symbol: ""},
		{Currency: "XXX", DisplayName: "niwaistā walūta", Symbol: ""},
	}
}
