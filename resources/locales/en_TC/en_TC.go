package en_TC

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_TC struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_TC' locale
func New() locales.Translator {
	return &en_TC{
		locale:  "en_TC",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_TC) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_TC'
func (t *en_TC) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_TC'
func (t *en_TC) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
