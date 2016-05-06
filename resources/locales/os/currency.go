package os

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Бритайнаг Фунт", Symbol: "£"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Лар", Symbol: "₾"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Сом", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "АИШ-ы Доллар", Symbol: "$"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Нӕзонгӕ валютӕ", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилиаг реал", Symbol: "R$"}}
