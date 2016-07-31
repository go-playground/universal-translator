package ms_SG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ms_SG' locale
func New() locales.Translator {
	return &ms_SG{
		locale:  "ms_SG",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ms_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms_SG'
func (t *ms_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ms_SG'
func (t *ms_SG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
