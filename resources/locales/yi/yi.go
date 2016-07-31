package yi

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yi' locale
func New() locales.Translator {
	return &yi{
		locale:  "yi",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *yi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yi'
func (t *yi) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'yi'
func (t *yi) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
