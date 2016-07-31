package be

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type be struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'be' locale
func New() locales.Translator {
	return &be{
		locale:  "be",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *be) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'be'
func (t *be) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'be'
func (t *be) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && n%100 != 11 {
		return locales.PluralRuleOne, nil
	} else if n%10 >= 2 && n%10 <= 4 && n%100 < 12 && n%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (n%10 == 0) || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
