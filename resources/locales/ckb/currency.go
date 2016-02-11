package ckb

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AFN", DisplayName: "ئەفغانیی ئەفغانستان", Symbol: ""},
		{Currency: "BHD", DisplayName: "دیناری بەحرەینی", Symbol: ""},
		{Currency: "BZD", DisplayName: "دۆلاری بەلیزی", Symbol: ""},
		{Currency: "DZD", DisplayName: "دیناری جەزائیری", Symbol: ""},
		{Currency: "EUR", DisplayName: "یورۆ", Symbol: ""},
		{Currency: "IQD", DisplayName: "دیناری عێراقی", Symbol: ""},
		{Currency: "IRR", DisplayName: "ڕیاڵی ئێرانی", Symbol: ""},
		{Currency: "JOD", DisplayName: "دیناری ئوردنی", Symbol: ""},
		{Currency: "KWD", DisplayName: "دیناری کووەیتی", Symbol: ""},
		{Currency: "OMR", DisplayName: "ڕیاڵی عومانی", Symbol: ""},
		{Currency: "QAR", DisplayName: "ڕیاڵی قەتەری", Symbol: ""},
		{Currency: "SAR", DisplayName: "ڕیاڵی سەعوودی", Symbol: ""},
		{Currency: "TND", DisplayName: "دیناری توونس", Symbol: ""},
		{Currency: "TRY", DisplayName: "لیرەی تورکیا", Symbol: ""},
		{Currency: "TTD", DisplayName: "دۆلاری ترینیداد و تۆباگۆ", Symbol: ""},
		{Currency: "XAU", DisplayName: "زێڕ", Symbol: ""},
	}
}
