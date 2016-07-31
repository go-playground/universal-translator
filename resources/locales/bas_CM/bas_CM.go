package bas_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bas_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bas_CM' locale
func New() locales.Translator {
	return &bas_CM{
		locale:  "bas_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *bas_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bas_CM'
func (t *bas_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bas_CM'
func (t *bas_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
