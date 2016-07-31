package bo_CN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bo_CN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bo_CN' locale
func New() locales.Translator {
	return &bo_CN{
		locale:  "bo_CN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *bo_CN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bo_CN'
func (t *bo_CN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bo_CN'
func (t *bo_CN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
