package ti_ER

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ti_ER struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ti_ER' locale
func New() locales.Translator {
	return &ti_ER{
		locale:   "ti_ER",
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
func (t *ti_ER) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ti_ER'
func (t *ti_ER) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ti_ER'
func (t *ti_ER) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
