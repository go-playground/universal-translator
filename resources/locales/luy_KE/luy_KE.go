package luy_KE

import "github.com/go-playground/universal-translator/resources/locales"

type luy_KE struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'luy_KE' locale
func New() locales.Translator {
	return &luy_KE{
		locale:   "luy_KE",
		plurals:  nil,
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *luy_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'luy_KE'
func (t *luy_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'luy_KE'
func (t *luy_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
