package es_UY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_UY struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_UY' locale
func New() locales.Translator {
	return &es_UY{
		locale:  "es_UY",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_UY) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_UY'
func (t *es_UY) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'es_UY'
func (t *es_UY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
