package ms_MY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_MY struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ms_MY' locale
func New() locales.Translator {
	return &ms_MY{
		locale:  "ms_MY",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ms_MY) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms_MY'
func (t *ms_MY) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ms_MY'
func (t *ms_MY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
