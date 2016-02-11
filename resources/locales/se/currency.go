package se

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DKK", DisplayName: "", Symbol: "Dkr"},
		{Currency: "EUR", DisplayName: "euro", Symbol: "€"},
		{Currency: "FIM", DisplayName: "suoma márkki", Symbol: ""},
		{Currency: "HKD", DisplayName: "", Symbol: "HK$"},
		{Currency: "INR", DisplayName: "", Symbol: "₹"},
		{Currency: "JPY", DisplayName: "", Symbol: "JP¥"},
		{Currency: "MXN", DisplayName: "", Symbol: "MX$"},
		{Currency: "NOK", DisplayName: "norgga kruvdno", Symbol: "kr"},
		{Currency: "SEK", DisplayName: "ruoŧŧa kruvdno", Symbol: "Skr"},
		{Currency: "THB", DisplayName: "", Symbol: "฿"},
		{Currency: "XAG", DisplayName: "uns silba", Symbol: ""},
		{Currency: "XAU", DisplayName: "uns golli", Symbol: ""},
	}
}
