package ckb

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ckb struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ckb' locale
func New() locales.Translator {
	return &ckb{
		locale:   "ckb",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{0xe2, 0x80, 0x8e, 0x2d},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ckb) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ckb'
func (t *ckb) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ckb'
func (t *ckb) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
