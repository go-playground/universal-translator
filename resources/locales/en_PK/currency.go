package en_PK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PKR", DisplayName: "", Symbol: "Rs"},
	}
}
