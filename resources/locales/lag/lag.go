package lag

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lag struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lag' locale
func New() locales.Translator {
	return &lag{
		locale:   "lag",
		plurals:  []locales.PluralRule{1, 2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *lag) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lag'
func (t *lag) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lag'
func (t *lag) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if n == 0 {
		return locales.PluralRuleZero
	} else if (i == 0 || i == 1) && n != 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
