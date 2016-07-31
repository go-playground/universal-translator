package uz_Cyrl_UZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Cyrl_UZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'uz_Cyrl_UZ' locale
func New() locales.Translator {
	return &uz_Cyrl_UZ{
		locale:  "uz_Cyrl_UZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *uz_Cyrl_UZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz_Cyrl_UZ'
func (t *uz_Cyrl_UZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'uz_Cyrl_UZ'
func (t *uz_Cyrl_UZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
