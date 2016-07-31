package ckb_IQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ckb_IQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ckb_IQ' locale
func New() locales.Translator {
	return &ckb_IQ{
		locale:  "ckb_IQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ckb_IQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ckb_IQ'
func (t *ckb_IQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ckb_IQ'
func (t *ckb_IQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
