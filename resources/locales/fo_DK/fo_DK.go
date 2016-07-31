package fo_DK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fo_DK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fo_DK' locale
func New() locales.Translator {
	return &fo_DK{
		locale:  "fo_DK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fo_DK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fo_DK'
func (t *fo_DK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fo_DK'
func (t *fo_DK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
