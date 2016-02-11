package ig

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CVE", DisplayName: "Escudo Caboverdiano", Symbol: ""},
		{Currency: "NGN", DisplayName: "Naịra", Symbol: "₦"},
	}
}
