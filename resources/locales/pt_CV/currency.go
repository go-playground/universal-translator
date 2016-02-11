package pt_CV

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CVE", DisplayName: "", Symbol: "\u200b"},
		{Currency: "PTE", DisplayName: "", Symbol: "\u200bPTE"},
	}
}
