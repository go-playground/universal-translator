package nnh_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nnh_CM struct {
	locale string
}

// New returns a new instance of translator for the 'nnh_CM' locale
func New() locales.Translator {
	return &nnh_CM{
		locale: "nnh_CM",
	}
}

// Locale returns the current translators string locale
func (l *nnh_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nnh_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
