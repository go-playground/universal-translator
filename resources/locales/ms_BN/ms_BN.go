package ms_BN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_BN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ms_BN' locale
func New() locales.Translator {
	return &ms_BN{
		locale:  "ms_BN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ms_BN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms_BN'
func (t *ms_BN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ms_BN'
func (t *ms_BN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
