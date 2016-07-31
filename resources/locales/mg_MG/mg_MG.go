package mg_MG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mg_MG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mg_MG' locale
func New() locales.Translator {
	return &mg_MG{
		locale:  "mg_MG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mg_MG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mg_MG'
func (t *mg_MG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mg_MG'
func (t *mg_MG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
