package om

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "Brazilian Real", Symbol: ""},
		{Currency: "CNY", DisplayName: "Chinese Yuan Renminbi", Symbol: ""},
		{Currency: "ETB", DisplayName: "Itoophiyaa Birrii", Symbol: "Br"},
		{Currency: "EUR", DisplayName: "Euro", Symbol: ""},
		{Currency: "GBP", DisplayName: "British Pound", Symbol: ""},
		{Currency: "INR", DisplayName: "Indian Rupee", Symbol: ""},
		{Currency: "JPY", DisplayName: "Japanese Yen", Symbol: ""},
		{Currency: "RUB", DisplayName: "Russian Ruble", Symbol: ""},
		{Currency: "USD", DisplayName: "US Dollar", Symbol: ""},
	}
}
