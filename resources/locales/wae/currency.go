package wae

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "Brasilianiši Real", Symbol: ""},
		{Currency: "CNY", DisplayName: "Chinesiši Yuan", Symbol: ""},
		{Currency: "EUR", DisplayName: "Euro", Symbol: ""},
		{Currency: "GBP", DisplayName: "Pfund", Symbol: ""},
		{Currency: "INR", DisplayName: "Indiši Rupie", Symbol: ""},
		{Currency: "JPY", DisplayName: "Yen", Symbol: "¥"},
		{Currency: "RUB", DisplayName: "Rubel", Symbol: ""},
		{Currency: "USD", DisplayName: "Dollar", Symbol: "$"},
		{Currency: "XXX", DisplayName: "Unbekannti Wãrig", Symbol: ""},
	}
}
