package ksb_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ksb_TZ struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ksb_TZ' locale
func New() locales.Translator {
	return &ksb_TZ{
		locale:   "ksb_TZ",
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
func (t *ksb_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksb_TZ'
func (t *ksb_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ksb_TZ'
func (t *ksb_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
