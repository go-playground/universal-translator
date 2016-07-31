package teo_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type teo_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'teo_KE' locale
func New() locales.Translator {
	return &teo_KE{
		locale:  "teo_KE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *teo_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'teo_KE'
func (t *teo_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'teo_KE'
func (t *teo_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
