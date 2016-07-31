package te

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type te struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'te' locale
func New() locales.Translator {
	return &te{
		locale:  "te",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *te) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'te'
func (t *te) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'te'
func (t *te) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
