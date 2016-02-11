package os

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "Бразилиаг реал", Symbol: "R$"},
		{Currency: "EUR", DisplayName: "Евро", Symbol: "€"},
		{Currency: "GBP", DisplayName: "Бритайнаг Фунт", Symbol: "£"},
		{Currency: "GEL", DisplayName: "Лар", Symbol: ""},
		{Currency: "RUB", DisplayName: "Сом", Symbol: ""},
		{Currency: "USD", DisplayName: "АИШ-ы Доллар", Symbol: "$"},
		{Currency: "XXX", DisplayName: "Нӕзонгӕ валютӕ", Symbol: ""},
	}
}
