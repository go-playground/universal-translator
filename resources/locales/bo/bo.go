package bo

import "github.com/go-playground/universal-translator/resources/locales"

type bo struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'bo' locale
func New() locales.Translator {
	return &bo{
		locale:   "bo",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *bo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bo'
func (t *bo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bo'
func (t *bo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
