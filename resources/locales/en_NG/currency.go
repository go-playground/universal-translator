package en_NG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "NGN", DisplayName: "", Symbol: "₦"},
	}
}
