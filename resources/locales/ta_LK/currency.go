package ta_LK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "LKR", DisplayName: "", Symbol: "Rs."},
	}
}
