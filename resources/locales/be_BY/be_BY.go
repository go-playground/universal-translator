package be_BY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type be_BY struct {
	locale string
}

// New returns a new instance of translator for the 'be_BY' locale
func New() locales.Translator {
	return &be_BY{
		locale: "be_BY",
	}
}

// Locale returns the current translators string locale
func (l *be_BY) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *be_BY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && n%100 != 11 {
		return locales.PluralRuleOne, nil
	} else if n%10 >= 2 && n%10 <= 4 && n%100 < 12 && n%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (n%10 == 0) || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
