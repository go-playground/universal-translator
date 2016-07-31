package ee_GH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ee_GH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale:  "ee_GH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ee_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ee_GH'
func (t *ee_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ee_GH'
func (t *ee_GH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
