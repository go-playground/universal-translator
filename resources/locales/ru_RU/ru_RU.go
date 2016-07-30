package ru_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ru_RU struct {
	locale string
}

// New returns a new instance of translator for the 'ru_RU' locale
func New() locales.Translator {
	return &ru_RU{
		locale: "ru_RU",
	}
}

// Locale returns the current translators string locale
func (l *ru_RU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ru_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%10 == 1 && i%100 != 11 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (v == 0 && i%10 == 0) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 11 && i%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
