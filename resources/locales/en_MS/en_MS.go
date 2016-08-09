package en_MS

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_MS struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_MS' locale
func New() locales.Translator {
	return &en_MS{
		locale:  "en_MS",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_MS) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_MS'
func (t *en_MS) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_MS'
func (t *en_MS) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
