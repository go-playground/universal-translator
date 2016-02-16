package lrc

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BRL": ut.Currency{Currency: "BRL", DisplayName: "رئال بئرئزیل", Symbol: "R$"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "یوان چین", Symbol: "CN¥"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "یئن جاپوٙن", Symbol: "JP¥"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "روٙبل روٙسیه", Symbol: "RUB"}, "USD": ut.Currency{Currency: "USD", DisplayName: "USD", Symbol: "US$"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "پیل نادیار", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "یورو", Symbol: "€"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "پوند بئریتانیا", Symbol: "£"}, "INR": ut.Currency{Currency: "INR", DisplayName: "روٙپیه هئن", Symbol: "₹"}}
