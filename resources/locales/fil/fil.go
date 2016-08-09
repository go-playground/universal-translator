package fil

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fil struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fil' locale
func New() locales.Translator {
	return &fil{
		locale:  "fil",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fil) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fil'
func (t *fil) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fil'
func (t *fil) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (i%10 != 4 && i%10 != 6 && i%10 != 9)) || (v != 0 && (f%10 != 4 && f%10 != 6 && f%10 != 9)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
