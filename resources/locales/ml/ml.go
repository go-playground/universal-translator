package ml

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ml struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ml' locale
func New() locales.Translator {
	return &ml{
		locale:  "ml",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ml) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ml'
func (t *ml) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ml'
func (t *ml) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
