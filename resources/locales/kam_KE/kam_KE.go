package kam_KE

import "github.com/go-playground/universal-translator/resources/locales"

type kam_KE struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'kam_KE' locale
func New() locales.Translator {
	return &kam_KE{
		locale:   "kam_KE",
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
func (t *kam_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kam_KE'
func (t *kam_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kam_KE'
func (t *kam_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
