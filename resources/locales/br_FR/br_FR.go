package br_FR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type br_FR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'br_FR' locale
func New() locales.Translator {
	return &br_FR{
		locale:  "br_FR",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *br_FR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'br_FR'
func (t *br_FR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'br_FR'
func (t *br_FR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && (n%100 != 11 && n%100 != 71 && n%100 != 91) {
		return locales.PluralRuleOne, nil
	} else if n%10 == 2 && (n%100 != 12 && n%100 != 72 && n%100 != 92) {
		return locales.PluralRuleTwo, nil
	} else if n%10 >= 3 && n%10 <= 4 && (n%10 == 9) && (n%100 < 10 && n%100 > 19) || (n%100 < 70 && n%100 > 79) || (n%100 < 90 && n%100 > 99) {
		return locales.PluralRuleFew, nil
	} else if n != 0 && n%1000000 == 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
