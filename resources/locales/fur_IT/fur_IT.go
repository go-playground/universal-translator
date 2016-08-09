package fur_IT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fur_IT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fur_IT' locale
func New() locales.Translator {
	return &fur_IT{
		locale:  "fur_IT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fur_IT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fur_IT'
func (t *fur_IT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fur_IT'
func (t *fur_IT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
