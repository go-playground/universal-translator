package uz_Arab_AF

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Arab_AF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'uz_Arab_AF' locale
func New() locales.Translator {
	return &uz_Arab_AF{
		locale:  "uz_Arab_AF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *uz_Arab_AF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz_Arab_AF'
func (t *uz_Arab_AF) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'uz_Arab_AF'
func (t *uz_Arab_AF) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
