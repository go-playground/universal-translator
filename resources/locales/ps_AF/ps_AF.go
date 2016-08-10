package ps_AF

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ps_AF struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ps_AF' locale
func New() locales.Translator {
	return &ps_AF{
		locale:   "ps_AF",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x62, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x63, 0x7d},
		minus:    []byte{},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x61, 0x7d},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ps_AF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ps_AF'
func (t *ps_AF) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ps_AF'
func (t *ps_AF) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
