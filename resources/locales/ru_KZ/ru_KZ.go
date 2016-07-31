package ru_KZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ru_KZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ru_KZ' locale
func New() locales.Translator {
	return &ru_KZ{
		locale:  "ru_KZ",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ru_KZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ru_KZ'
func (t *ru_KZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ru_KZ'
func (t *ru_KZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if v == 0 && i%10 == 1 && i%100 != 11 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (v == 0 && i%10 == 0) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 11 && i%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
