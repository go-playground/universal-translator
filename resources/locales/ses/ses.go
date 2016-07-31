package ses

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ses struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ses' locale
func New() locales.Translator {
	return &ses{
		locale:  "ses",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ses) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ses'
func (t *ses) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ses'
func (t *ses) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
