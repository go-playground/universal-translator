package zh_Hans_SG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans_SG' locale
func New() locales.Translator {
	return &zh_Hans_SG{
		locale:  "zh_Hans_SG",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans_SG'
func (t *zh_Hans_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hans_SG'
func (t *zh_Hans_SG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
