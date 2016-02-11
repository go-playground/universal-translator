package bo_IN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CNY", DisplayName: "", Symbol: "CN¥"},
	}
}
