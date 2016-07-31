package pl_PL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type pl_PL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pl_PL' locale
func New() locales.Translator {
	return &pl_PL{
		locale:  "pl_PL",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *pl_PL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pl_PL'
func (t *pl_PL) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'pl_PL'
func (t *pl_PL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (v == 0 && i != 1 && i%10 >= 0 && i%10 <= 1) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 12 && i%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
