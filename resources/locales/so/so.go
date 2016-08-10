package so

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type so struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'so' locale
func New() locales.Translator {
	return &so{
		locale:   "so",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *so) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'so'
func (t *so) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'so'
func (t *so) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
