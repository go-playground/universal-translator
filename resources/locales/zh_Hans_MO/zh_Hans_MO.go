package zh_Hans_MO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_MO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans_MO' locale
func New() locales.Translator {
	return &zh_Hans_MO{
		locale:  "zh_Hans_MO",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans_MO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans_MO'
func (t *zh_Hans_MO) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hans_MO'
func (t *zh_Hans_MO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
