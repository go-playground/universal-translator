package sv

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sv struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'sv' locale
func New() locales.Translator {
	return &sv{
		locale:   "sv",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{0xe2, 0x80, 0x8f, 0xe2, 0x88, 0x92},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *sv) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv'
func (t *sv) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sv'
func (t *sv) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
