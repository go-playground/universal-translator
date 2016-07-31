package zh_Hant_TW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant_TW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hant_TW' locale
func New() locales.Translator {
	return &zh_Hant_TW{
		locale:  "zh_Hant_TW",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hant_TW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hant_TW'
func (t *zh_Hant_TW) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hant_TW'
func (t *zh_Hant_TW) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
