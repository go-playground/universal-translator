package tzm_MA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type tzm_MA struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'tzm_MA' locale
func New() locales.Translator {
	return &tzm_MA{
		locale:   "tzm_MA",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x63, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x30, 0x7d},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *tzm_MA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'tzm_MA'
func (t *tzm_MA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'tzm_MA'
func (t *tzm_MA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if (n >= 0 && n <= 1) || (n >= 11 && n <= 99) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
