package fil_PH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fil_PH struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'fil_PH' locale
func New() locales.Translator {
	return &fil_PH{
		locale:   "fil_PH",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x65, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x35, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x62, 0x30, 0x7d},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *fil_PH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fil_PH'
func (t *fil_PH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fil_PH'
func (t *fil_PH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (i%10 != 4 && i%10 != 6 && i%10 != 9)) || (v != 0 && (f%10 != 4 && f%10 != 6 && f%10 != 9)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
