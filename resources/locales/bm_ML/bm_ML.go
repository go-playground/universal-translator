package bm_ML

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bm_ML struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bm_ML' locale
func New() locales.Translator {
	return &bm_ML{
		locale:  "bm_ML",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *bm_ML) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bm_ML'
func (t *bm_ML) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bm_ML'
func (t *bm_ML) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
