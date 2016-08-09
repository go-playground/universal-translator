package mgo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mgo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mgo' locale
func New() locales.Translator {
	return &mgo{
		locale:  "mgo",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mgo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mgo'
func (t *mgo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mgo'
func (t *mgo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
