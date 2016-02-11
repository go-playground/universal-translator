package zh_Hans_SG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ANG", DisplayName: "荷兰安的列斯盾", Symbol: ""},
		{Currency: "AWG", DisplayName: "阿鲁巴弗罗林", Symbol: ""},
		{Currency: "CNY", DisplayName: "", Symbol: "CN¥"},
		{Currency: "NIO", DisplayName: "尼加拉瓜科多巴", Symbol: ""},
		{Currency: "SGD", DisplayName: "", Symbol: "$"},
		{Currency: "XAG", DisplayName: "白银", Symbol: ""},
	}
}
