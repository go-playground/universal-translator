package ro_RO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ro_RO struct {
	locale string
}

// New returns a new instance of translator for the 'ro_RO' locale
func New() locales.Translator {
	return &ro_RO{
		locale: "ro_RO",
	}
}

// Locale returns the current translators string locale
func (l *ro_RO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ro_RO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if (v != 0) || (n == 0) || (n != 1 && n%100 >= 1 && n%100 <= 19) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
