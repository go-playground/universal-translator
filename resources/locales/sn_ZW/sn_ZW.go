package sn_ZW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sn_ZW struct {
	locale string
}

// New returns a new instance of translator for the 'sn_ZW' locale
func New() locales.Translator {
	return &sn_ZW{
		locale: "sn_ZW",
	}
}

// Locale returns the current translators string locale
func (l *sn_ZW) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sn_ZW) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
