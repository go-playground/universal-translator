package eo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type eo struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'eo' locale
func New() locales.Translator {
	return &eo{
		locale:   "eo",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0xe2, 0x88, 0x92},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *eo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'eo'
func (t *eo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'eo'
func (t *eo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
