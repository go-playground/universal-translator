package vun_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vun_TZ struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'vun_TZ' locale
func New() locales.Translator {
	return &vun_TZ{
		locale:   "vun_TZ",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *vun_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vun_TZ'
func (t *vun_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vun_TZ'
func (t *vun_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
