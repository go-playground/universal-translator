package uz_Arab

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Arab struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'uz_Arab' locale
func New() locales.Translator {
	return &uz_Arab{
		locale:   "uz_Arab",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xd9, 0xac},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x38, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x39, 0x7d},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *uz_Arab) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz_Arab'
func (t *uz_Arab) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'uz_Arab'
func (t *uz_Arab) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
