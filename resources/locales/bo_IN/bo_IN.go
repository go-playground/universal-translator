package bo_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bo_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bo_IN' locale
func New() locales.Translator {
	return &bo_IN{
		locale:  "bo_IN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *bo_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bo_IN'
func (t *bo_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bo_IN'
func (t *bo_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
