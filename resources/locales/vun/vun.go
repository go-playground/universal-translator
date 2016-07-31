package vun

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vun struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vun' locale
func New() locales.Translator {
	return &vun{
		locale:  "vun",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *vun) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vun'
func (t *vun) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'vun'
func (t *vun) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
