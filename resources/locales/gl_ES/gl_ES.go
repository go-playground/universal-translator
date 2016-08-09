package gl_ES

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gl_ES struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gl_ES' locale
func New() locales.Translator {
	return &gl_ES{
		locale:  "gl_ES",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gl_ES) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gl_ES'
func (t *gl_ES) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gl_ES'
func (t *gl_ES) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
