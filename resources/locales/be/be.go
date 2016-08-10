package be

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type be struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'be' locale
func New() locales.Translator {
	return &be{
		locale:   "be",
		plurals:  []locales.PluralRule{2, 4, 5, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *be) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'be'
func (t *be) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'be'
func (t *be) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n%10 == 1 && n%100 != 11 {
		return locales.PluralRuleOne
	} else if n%10 >= 2 && n%10 <= 4 && n%100 < 12 && n%100 > 14 {
		return locales.PluralRuleFew
	} else if (n%10 == 0) || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
