package sv_AX

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sv_AX struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'sv_AX' locale
func New() locales.Translator {
	return &sv_AX{
		locale:   "sv_AX",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x66, 0x2c, 0x20, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x38, 0x2c, 0x20, 0x30, 0x78, 0x39, 0x32, 0x7d},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *sv_AX) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv_AX'
func (t *sv_AX) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sv_AX'
func (t *sv_AX) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
