package haw_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type haw_US struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'haw_US' locale
func New() locales.Translator {
	return &haw_US{
		locale:  "haw_US",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *haw_US) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'haw_US'
func (t *haw_US) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'haw_US'
func (t *haw_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
