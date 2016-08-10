package pa_Arab

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pa_Arab struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'pa_Arab' locale
func New() locales.Translator {
	return &pa_Arab{
		locale:   "pa_Arab",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x62, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		minus:    []byte{0xe2, 0x80, 0x8e, 0x2d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x61, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x38, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x39, 0x7d},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *pa_Arab) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pa_Arab'
func (t *pa_Arab) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pa_Arab'
func (t *pa_Arab) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
