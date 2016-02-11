package en_FK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "FKP", DisplayName: "", Symbol: "£"},
		{Currency: "GBP", DisplayName: "", Symbol: "GB£"},
	}
}
