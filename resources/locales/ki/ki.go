package ki

import "github.com/go-playground/universal-translator/resources/locales"

type ki struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ki' locale
func New() locales.Translator {
	return &ki{
		locale:   "ki",
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
func (t *ki) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ki'
func (t *ki) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ki'
func (t *ki) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
