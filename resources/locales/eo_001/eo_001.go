package eo_001

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type eo_001 struct {
	locale string
}

// New returns a new instance of translator for the 'eo_001' locale
func New() locales.Translator {
	return &eo_001{
		locale: "eo_001",
	}
}

// Locale returns the current translators string locale
func (l *eo_001) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *eo_001) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
