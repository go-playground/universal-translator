package agq_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type agq_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'agq_CM' locale
func New() locales.Translator {
	return &agq_CM{
		locale:  "agq_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *agq_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'agq_CM'
func (t *agq_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'agq_CM'
func (t *agq_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
