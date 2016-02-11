package mgo

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "XAF", DisplayName: "shirè", Symbol: "FCFA"},
		{Currency: "XXX", DisplayName: "iku ikap mɔʼɔ", Symbol: ""},
	}
}
