package wae

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiši Rupie", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen", Symbol: "¥"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianiši Real", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pfund", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rubel", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dollar", Symbol: "$"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Unbekannti Wãrig", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinesiši Yuan", Symbol: ""}}
