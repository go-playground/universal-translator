package kde

import "github.com/go-playground/universal-translator/resources/locales"

type kde struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'kde' locale
func New() locales.Translator {
	return &kde{
		locale:   "kde",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *kde) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kde'
func (t *kde) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kde'
func (t *kde) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
