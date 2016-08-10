package mt

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mt struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'mt' locale
func New() locales.Translator {
	return &mt{
		locale:   "mt",
		plurals:  []locales.PluralRule{2, 4, 5, 6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *mt) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mt'
func (t *mt) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mt'
func (t *mt) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if (n == 0) || (n%100 >= 2 && n%100 <= 10) {
		return locales.PluralRuleFew
	} else if n%100 >= 11 && n%100 <= 19 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
