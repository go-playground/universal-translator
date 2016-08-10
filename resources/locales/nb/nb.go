package nb

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nb struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'nb' locale
func New() locales.Translator {
	return &nb{
		locale:   "nb",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0xe2, 0x88, 0x92},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{0xd8, 0x89},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *nb) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nb'
func (t *nb) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nb'
func (t *nb) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
