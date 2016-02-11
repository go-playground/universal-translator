package pa_Arab

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "EUR", DisplayName: "يورو", Symbol: ""},
		{Currency: "INR", DisplayName: "روپئیہ [INR]", Symbol: ""},
		{Currency: "PKR", DisplayName: "روپئیہ", Symbol: "ر"},
	}
}
