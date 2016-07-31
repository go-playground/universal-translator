package sg_CF

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sg_CF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sg_CF' locale
func New() locales.Translator {
	return &sg_CF{
		locale:  "sg_CF",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *sg_CF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sg_CF'
func (t *sg_CF) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sg_CF'
func (t *sg_CF) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
