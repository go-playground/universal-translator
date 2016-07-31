package ksh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ksh struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ksh' locale
func New() locales.Translator {
	return &ksh{
		locale:  "ksh",
		plurals: []locales.PluralRule{1, 2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ksh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksh'
func (t *ksh) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ksh'
func (t *ksh) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
