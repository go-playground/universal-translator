package ast_ES

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ast_ES struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ast_ES' locale
func New() locales.Translator {
	return &ast_ES{
		locale:  "ast_ES",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ast_ES) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ast_ES'
func (t *ast_ES) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ast_ES'
func (t *ast_ES) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
