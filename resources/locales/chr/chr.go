package chr

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type chr struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'chr' locale
func New() locales.Translator {
	return &chr{
		locale:  "chr",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *chr) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'chr'
func (t *chr) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'chr'
func (t *chr) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
