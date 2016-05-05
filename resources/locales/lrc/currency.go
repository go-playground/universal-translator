package lrc

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CNY": ut.Currency{Currency: "CNY", DisplayName: "یوان چین", Symbol: "CN¥"}, "INR": ut.Currency{Currency: "INR", DisplayName: "روٙپیه هئن", Symbol: "₹"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "روٙبل روٙسیه", Symbol: "RUB"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "پیل نادیار", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "رئال بئرئزیل", Symbol: "R$"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "یورو", Symbol: "€"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "پوند بئریتانیا", Symbol: "£"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "یئن جاپوٙن", Symbol: "JP¥"}, "USD": ut.Currency{Currency: "USD", DisplayName: "USD", Symbol: "US$"}}
