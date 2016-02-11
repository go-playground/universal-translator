package eo

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AUD", DisplayName: "Aŭstralia dolaro", Symbol: "AU$"},
		{Currency: "BRL", DisplayName: "Brazila realo", Symbol: "R$"},
		{Currency: "CAD", DisplayName: "Kanada dolaro", Symbol: "CA$"},
		{Currency: "CHF", DisplayName: "Svisa franko", Symbol: "CHF"},
		{Currency: "CNY", DisplayName: "Ĉina juano", Symbol: "CN¥"},
		{Currency: "DKK", DisplayName: "Dana krono", Symbol: "DKK"},
		{Currency: "EUR", DisplayName: "Eŭro", Symbol: "€"},
		{Currency: "GBP", DisplayName: "Brita pundo", Symbol: "£"},
		{Currency: "HKD", DisplayName: "Honkonga dolaro", Symbol: "HK$"},
		{Currency: "IDR", DisplayName: "Indonezia rupio", Symbol: "IDR"},
		{Currency: "INR", DisplayName: "Barata rupio", Symbol: "₹"},
		{Currency: "JPY", DisplayName: "Japana eno", Symbol: "JP¥"},
		{Currency: "KRW", DisplayName: "Sud-korea ŭono", Symbol: "₩"},
		{Currency: "MXN", DisplayName: "Meksika peso", Symbol: "MX$"},
		{Currency: "NOK", DisplayName: "Norvega krono", Symbol: "NOK"},
		{Currency: "PLN", DisplayName: "Pola zloto", Symbol: "PLN"},
		{Currency: "RUB", DisplayName: "Rusa rublo", Symbol: "RUB"},
		{Currency: "SAR", DisplayName: "Sauda rialo", Symbol: "SAR"},
		{Currency: "SEK", DisplayName: "Sveda krono", Symbol: "SEK"},
		{Currency: "THB", DisplayName: "Taja bahto", Symbol: "฿"},
		{Currency: "TRY", DisplayName: "Turka liro", Symbol: "₺"},
		{Currency: "TWD", DisplayName: "Nova tajvana dolaro", Symbol: "NT$"},
		{Currency: "USD", DisplayName: "Usona dolaro", Symbol: "US$"},
		{Currency: "XAG", DisplayName: "arĝento", Symbol: ""},
		{Currency: "XAU", DisplayName: "oro", Symbol: ""},
		{Currency: "XBB", DisplayName: "eŭropa monunuo", Symbol: ""},
		{Currency: "XFO", DisplayName: "franca ora franko", Symbol: ""},
		{Currency: "XPD", DisplayName: "paladio", Symbol: ""},
		{Currency: "XPT", DisplayName: "plateno", Symbol: ""},
		{Currency: "XXX", DisplayName: "Nekonata valuto", Symbol: ""},
		{Currency: "ZAR", DisplayName: "Sud-afrika rando", Symbol: "ZAR"},
	}
}
