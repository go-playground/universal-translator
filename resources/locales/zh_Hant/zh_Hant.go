package zh_Hant

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hant' locale
func New() locales.Translator {
	return &zh_Hant{
		locale:  "zh_Hant",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hant) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hant'
func (t *zh_Hant) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hant'
func (t *zh_Hant) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
