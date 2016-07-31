package zh_Hans_CN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_CN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans_CN' locale
func New() locales.Translator {
	return &zh_Hans_CN{
		locale:  "zh_Hans_CN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans_CN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans_CN'
func (t *zh_Hans_CN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh_Hans_CN'
func (t *zh_Hans_CN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
