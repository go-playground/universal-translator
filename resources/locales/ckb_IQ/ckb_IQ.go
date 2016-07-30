package ckb_IQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ckb_IQ struct {
	locale string
}

// New returns a new instance of translator for the 'ckb_IQ' locale
func New() locales.Translator {
	return &ckb_IQ{
		locale: "ckb_IQ",
	}
}

// Locale returns the current translators string locale
func (l *ckb_IQ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ckb_IQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
