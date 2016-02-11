package kl

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DKK", DisplayName: "danmarkimut koruuni", Symbol: "kr."},
		{Currency: "EUR", DisplayName: "euro", Symbol: "€"},
		{Currency: "NOK", DisplayName: "norskit koruuni", Symbol: "Nkr"},
		{Currency: "SEK", DisplayName: "svenskit koruuni", Symbol: "Skr"},
	}
}
