package zh_Hant_HK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant_HK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hant_HK' locale
func New() locales.Translator {
	return &zh_Hant_HK{
		locale:  "zh_Hant_HK",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hant_HK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hant_HK'
func (t *zh_Hant_HK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hant_HK'
func (t *zh_Hant_HK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
