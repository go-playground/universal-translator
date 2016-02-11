package so_ET

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ETB", DisplayName: "", Symbol: "Br"},
	}
}
