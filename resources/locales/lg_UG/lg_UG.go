package lg_UG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lg_UG struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lg_UG' locale
func New() locales.Translator {
	return &lg_UG{
		locale:   "lg_UG",
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
func (t *lg_UG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lg_UG'
func (t *lg_UG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lg_UG'
func (t *lg_UG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
