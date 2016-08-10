package mg_MG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mg_MG struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'mg_MG' locale
func New() locales.Translator {
	return &mg_MG{
		locale:   "mg_MG",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *mg_MG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mg_MG'
func (t *mg_MG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mg_MG'
func (t *mg_MG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
