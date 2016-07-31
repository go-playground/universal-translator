package brx_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type brx_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'brx_IN' locale
func New() locales.Translator {
	return &brx_IN{
		locale:  "brx_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *brx_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'brx_IN'
func (t *brx_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'brx_IN'
func (t *brx_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
