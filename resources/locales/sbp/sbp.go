package sbp

import "github.com/go-playground/universal-translator/resources/locales"

type sbp struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'sbp' locale
func New() locales.Translator {
	return &sbp{
		locale:   "sbp",
		plurals:  nil,
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *sbp) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sbp'
func (t *sbp) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sbp'
func (t *sbp) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
