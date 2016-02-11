package chr

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "ᏆᏏᎵᎢ ᎠᏕᎳ", Symbol: ""},
		{Currency: "CAD", DisplayName: "ᎧᎾᏓ ᎠᏕᎳ", Symbol: ""},
		{Currency: "CNY", DisplayName: "ᏓᎶᏂᎨ ᎠᏕᎳ", Symbol: ""},
		{Currency: "EUR", DisplayName: "ᏳᎳᏛ", Symbol: ""},
		{Currency: "GBP", DisplayName: "ᎩᎵᏏᏲ ᎠᏕᎳ", Symbol: ""},
		{Currency: "INR", DisplayName: "ᎢᏅᏗᎾ ᎠᏕᎳ", Symbol: ""},
		{Currency: "JPY", DisplayName: "ᏣᏩᏂᏏ ᎠᏕᎳ", Symbol: ""},
		{Currency: "MXN", DisplayName: "ᏍᏆᏂ ᎠᏕᎳ", Symbol: ""},
		{Currency: "RUB", DisplayName: "ᏲᏂᎢ ᎠᏕᎳ", Symbol: ""},
		{Currency: "USD", DisplayName: "ᎤᏃᏍᏗ", Symbol: "$"},
	}
}
