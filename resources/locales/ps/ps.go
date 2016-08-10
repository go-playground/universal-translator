package ps

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ps struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ps' locale
func New() locales.Translator {
	return &ps{
		locale:   "ps",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xd9, 0xac},
		minus:    []byte{},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ps) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ps'
func (t *ps) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ps'
func (t *ps) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
