package lrc

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "رئال بئرئزیل", Symbol: "R$"},
		{Currency: "CNY", DisplayName: "یوان چین", Symbol: "CN¥"},
		{Currency: "EUR", DisplayName: "یورو", Symbol: "€"},
		{Currency: "GBP", DisplayName: "پوند بئریتانیا", Symbol: "£"},
		{Currency: "INR", DisplayName: "روٙپیه هئن", Symbol: "₹"},
		{Currency: "JPY", DisplayName: "یئن جاپوٙن", Symbol: "JP¥"},
		{Currency: "RUB", DisplayName: "روٙبل روٙسیه", Symbol: "RUB"},
		{Currency: "USD", DisplayName: "USD", Symbol: "US$"},
		{Currency: "XXX", DisplayName: "پیل نادیار", Symbol: ""},
	}
}
