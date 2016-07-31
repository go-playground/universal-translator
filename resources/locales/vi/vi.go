package vi

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vi' locale
func New() locales.Translator {
	return &vi{
		locale:  "vi",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *vi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vi'
func (t *vi) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'vi'
func (t *vi) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
