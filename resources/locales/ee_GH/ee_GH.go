package ee_GH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ee_GH struct {
	locale string
}

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale: "ee_GH",
	}
}

// Locale returns the current translators string locale
func (l *ee_GH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ee_GH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
