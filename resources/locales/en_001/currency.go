package en_001

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BYB", DisplayName: "Belarusian New Rouble (1994–1999)", Symbol: ""},
		{Currency: "BYR", DisplayName: "Belarusian Rouble", Symbol: ""},
		{Currency: "JPY", DisplayName: "", Symbol: "JP¥"},
		{Currency: "LVR", DisplayName: "Latvian Rouble", Symbol: ""},
		{Currency: "RUB", DisplayName: "Russian Rouble", Symbol: ""},
		{Currency: "RUR", DisplayName: "Russian Rouble (1991–1998)", Symbol: ""},
		{Currency: "TJR", DisplayName: "Tajikistani Rouble", Symbol: ""},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
