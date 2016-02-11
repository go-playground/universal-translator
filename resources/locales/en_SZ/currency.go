package en_SZ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SZL", DisplayName: "", Symbol: "E"},
	}
}
