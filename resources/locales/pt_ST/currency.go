package pt_ST

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "STD", DisplayName: "", Symbol: "Db"},
	}
}
