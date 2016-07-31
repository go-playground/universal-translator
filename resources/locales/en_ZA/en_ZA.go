package en_ZA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_ZA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_ZA' locale
func New() locales.Translator {
	return &en_ZA{
		locale:  "en_ZA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_ZA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_ZA'
func (t *en_ZA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_ZA'
func (t *en_ZA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
