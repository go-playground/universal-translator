package lo

import "github.com/go-playground/universal-translator/resources/locales"

type lo struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lo' locale
func New() locales.Translator {
	return &lo{
		locale:   "lo",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *lo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lo'
func (t *lo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lo'
func (t *lo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
