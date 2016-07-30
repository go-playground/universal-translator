package ar_YE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_YE struct {
	locale string
}

// New returns a new instance of translator for the 'ar_YE' locale
func New() locales.Translator {
	return &ar_YE{
		locale: "ar_YE",
	}
}

// Locale returns the current translators string locale
func (l *ar_YE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ar_YE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew, nil
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
