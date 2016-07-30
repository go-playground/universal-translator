package pt_LU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_LU struct {
	locale string
}

// New returns a new instance of translator for the 'pt_LU' locale
func New() locales.Translator {
	return &pt_LU{
		locale: "pt_LU",
	}
}

// Locale returns the current translators string locale
func (l *pt_LU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *pt_LU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
