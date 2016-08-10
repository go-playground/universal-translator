package naq_NA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type naq_NA struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'naq_NA' locale
func New() locales.Translator {
	return &naq_NA{
		locale:   "naq_NA",
		plurals:  []locales.PluralRule{2, 3, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *naq_NA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'naq_NA'
func (t *naq_NA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'naq_NA'
func (t *naq_NA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
