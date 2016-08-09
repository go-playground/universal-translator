package ksh_DE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ksh_DE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ksh_DE' locale
func New() locales.Translator {
	return &ksh_DE{
		locale:  "ksh_DE",
		plurals: []locales.PluralRule{1, 2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ksh_DE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksh_DE'
func (t *ksh_DE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ksh_DE'
func (t *ksh_DE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
