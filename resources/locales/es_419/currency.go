package es_419

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AMD", DisplayName: "dram armenio", Symbol: ""},
		{Currency: "BGN", DisplayName: "lev búlgaro", Symbol: ""},
		{Currency: "CAD", DisplayName: "", Symbol: "CAD"},
		{Currency: "EGP", DisplayName: "", Symbol: "E£"},
		{Currency: "ERN", DisplayName: "nafka", Symbol: ""},
		{Currency: "EUR", DisplayName: "", Symbol: "EUR"},
		{Currency: "THB", DisplayName: "", Symbol: "THB"},
		{Currency: "USD", DisplayName: "", Symbol: "USD"},
		{Currency: "VEF", DisplayName: "", Symbol: "BsF"},
		{Currency: "VND", DisplayName: "", Symbol: "VND"},
		{Currency: "XXX", DisplayName: "(unidad de moneda desconocida)", Symbol: ""},
	}
}
