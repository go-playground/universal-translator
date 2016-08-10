package tzm

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type tzm struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'tzm' locale
func New() locales.Translator {
	return &tzm{
		locale:   "tzm",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *tzm) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'tzm'
func (t *tzm) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'tzm'
func (t *tzm) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if (n >= 0 && n <= 1) || (n >= 11 && n <= 99) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
