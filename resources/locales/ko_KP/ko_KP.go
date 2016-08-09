package ko_KP

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ko_KP struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ko_KP' locale
func New() locales.Translator {
	return &ko_KP{
		locale:  "ko_KP",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ko_KP) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ko_KP'
func (t *ko_KP) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ko_KP'
func (t *ko_KP) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
