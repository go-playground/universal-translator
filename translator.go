package ut

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/locales"
)

const (
	paramZero          = "{0}"
	paramOne           = "{1}"
	unknownTranslation = ""
)

var (
	// ErrUnknowTranslation indicates the translation could not be found
	ErrUnknowTranslation = errors.New("Unknown Translation")
)

// Translator is universal translators
// translator instance which is a thin wrapper
// around locales.Translator instance providing
// some extra functionality
type Translator interface {
	locales.Translator

	// adds a normal translation for a particular language/locale
	// {#} is the only replacement type accepted and are add infintium
	// eg. one: '{0} day left' other: '{0} days left'
	Add(key interface{}, text string)

	// adds a cardinal plural translation for a particular language/locale
	// {0} is the only replacement type accepted and only one variable is accepted as
	// multiple cannot be used for a plural rule determination, unless it is a range;
	// see AddRange below.
	// eg. in locale 'en' one: '{0} day left' other: '{0} days left'
	AddCardinal(key interface{}, text string, rule locales.PluralRule) error

	// adds an ordinal plural translation for a particular language/locale
	// {0} is the only replacement type accepted and only one variable is accepted as
	// multiple cannot be used for a plural rule determination, unless it is a range;
	// see AddRange below.
	// eg. in locale 'en' one: '{0}st day of spring' other: '{0}nd day of spring'
	// - 1st, 2nd, 3rd...
	AddOrdinal(key interface{}, text string, rule locales.PluralRule) error

	// adds a range plural translation for a particular language/locale
	// {0} and {1} are the only replacement types accepted and only these are accepted.
	// eg. in locale 'nl' one: '{0}-{1} day left' other: '{0}-{1} days left'
	AddRange(key interface{}, text string, rule locales.PluralRule) error

	// creates the translation for the locale given the 'key' and params passed in
	T(key interface{}, params ...string) string

	// creates the cardinal translation for the locale given the 'key', 'num' and 'digit' arguments
	//  and param passed in
	C(key interface{}, num float64, digits uint64, param string) (string, error)

	// creates the ordinal translation for the locale given the 'key', 'num' and 'digit' arguments
	// and param passed in
	O(key interface{}, num float64, digits uint64, param string) (string, error)

	//  creates the range translation for the locale given the 'key', 'num1', 'digit1', 'num2' and
	//  'digit2' arguments and 'param1' and 'param2' passed in
	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) (string, error)
}

var _ Translator = new(translator)
var _ locales.Translator = new(translator)

type translator struct {
	locales.Translator
	translations        map[interface{}]*transText
	cardinalTanslations map[interface{}][]*transText // array index is mapped to locales.PluralRule index + the locales.PluralRuleUnknown
	ordinalTanslations  map[interface{}][]*transText
	rangeTanslations    map[interface{}][]*transText
}

type transText struct {
	text    string
	indexes []int
}

func newTranslator(trans locales.Translator) Translator {
	return &translator{
		Translator:          trans,
		translations:        make(map[interface{}]*transText), // translation text broken up by byte index
		cardinalTanslations: make(map[interface{}][]*transText),
		ordinalTanslations:  make(map[interface{}][]*transText),
		rangeTanslations:    make(map[interface{}][]*transText),
	}
}

// Add adds a normal translation for a particular language/locale
// {#} is the only replacement type accepted and are add infintium
// eg. one: '{0} day left' other: '{0} days left'
func (t *translator) Add(key interface{}, text string) {

	trans := &transText{
		text: text,
	}

	var i int
	var idx int
	var cum int

	for {
		s := "{" + strconv.Itoa(i) + "}"
		idx = strings.Index(text[idx:], s)
		if idx == -1 {
			break
		}

		trans.indexes = append(trans.indexes, idx+cum)
		idx += len(s)
		trans.indexes = append(trans.indexes, idx+cum)
		cum += idx
		i++
	}

	t.translations[key] = trans
}

// AddCardinal adds a cardinal plural translation for a particular language/locale
// {0} is the only replacement type accepted and only one variable is accepted as
// multiple cannot be used for a plural rule determination, unless it is a range;
// see AddRange below.
// eg. in locale 'en' one: '{0} day left' other: '{0} days left'
func (t *translator) AddCardinal(key interface{}, text string, rule locales.PluralRule) error {

	tarr, ok := t.cardinalTanslations[key]
	if ok {
		// verify not adding a conflicting record
		if len(tarr) > 0 && tarr[rule] != nil {
			return fmt.Errorf("warning: conflicting key '%#v' rule '%d' with text '%s', value being ignored", key, rule, text)
		}

	} else {
		tarr = make([]*transText, 7, 7)
		// tarr = make([]*transText, len(t.PluralsCardinal())+1, len(t.PluralsCardinal())+1)
		t.cardinalTanslations[key] = tarr
	}

	trans := &transText{
		text:    text,
		indexes: make([]int, 2, 2),
	}

	tarr[rule] = trans

	idx := strings.Index(text, paramZero)
	if idx == -1 {
		tarr[rule] = nil
		return fmt.Errorf("error: parameter '%s' not found, may want to use 'Add' instead of 'AddCardinal'", paramZero)
	}

	trans.indexes[0] = idx
	trans.indexes[1] = idx + len(paramZero)

	return nil
}

// AddOrdinal adds an ordinal plural translation for a particular language/locale
// {0} is the only replacement type accepted and only one variable is accepted as
// multiple cannot be used for a plural rule determination, unless it is a range;
// see AddRange below.
// eg. in locale 'en' one: '{0}st day of spring' other: '{0}nd day of spring' - 1st, 2nd, 3rd...
func (t *translator) AddOrdinal(key interface{}, text string, rule locales.PluralRule) error {

	tarr, ok := t.ordinalTanslations[key]
	if ok {
		// verify not adding a conflicting record
		if len(tarr) > 0 && tarr[rule] != nil {
			return fmt.Errorf("warning: conflicting key '%#v' rule '%d' with text '%s', value being ignored", key, rule, text)
		}

	} else {
		tarr = make([]*transText, 7, 7)
		t.ordinalTanslations[key] = tarr
	}

	trans := &transText{
		text:    text,
		indexes: make([]int, 2, 2),
	}

	tarr[rule] = trans

	idx := strings.Index(text, paramZero)
	if idx == -1 {
		tarr[rule] = nil
		return fmt.Errorf("error: parameter '%s' not found, may want to use 'Add' instead of 'AddOrdinal'", paramZero)
	}

	trans.indexes[0] = idx
	trans.indexes[1] = idx + len(paramZero)

	return nil
}

// AddRange adds a range plural translation for a particular language/locale
// {0} and {1} are the only replacement types accepted and only these are accepted.
// eg. in locale 'nl' one: '{0}-{1} day left' other: '{0}-{1} days left'
func (t *translator) AddRange(key interface{}, text string, rule locales.PluralRule) error {

	tarr, ok := t.rangeTanslations[key]
	if ok {
		// verify not adding a conflicting record
		if len(tarr) > 0 && tarr[rule] != nil {
			return fmt.Errorf("warning: conflicting key '%#v' rule '%s' with text '%s', value being ignored", key, rule, text)
		}

	} else {
		tarr = make([]*transText, 7, 7)
		t.rangeTanslations[key] = tarr
	}

	trans := &transText{
		text:    text,
		indexes: make([]int, 4, 4),
	}

	tarr[rule] = trans

	idx := strings.Index(text, paramZero)
	if idx == -1 {
		tarr[rule] = nil
		return fmt.Errorf("error: parameter '%s' not found, are you sure you're adding a Range Translation?", paramZero)
	}

	trans.indexes[0] = idx
	trans.indexes[1] = idx + len(paramZero)

	idx = strings.Index(text, paramOne)
	if idx == -1 {
		tarr[rule] = nil
		return fmt.Errorf("error: parameter '%s' not found, a Range Translation requires two parameters", paramOne)
	}

	trans.indexes[2] = idx
	trans.indexes[3] = idx + len(paramOne)

	return nil
}

// T creates the translation for the locale given the 'key' and params passed in
func (t *translator) T(key interface{}, params ...string) string {

	trans, ok := t.translations[key]
	if !ok {
		return unknownTranslation
	}

	// maybe pool these later?...
	b := make([]byte, 0, 64)

	var start, end, count int

	for i := 0; i < len(trans.indexes); i++ {
		end = trans.indexes[i]
		b = append(b, trans.text[start:end]...)
		b = append(b, params[count]...)
		i++
		start = trans.indexes[i]
		count++
	}

	b = append(b, trans.text[start:]...)

	return string(b)
}

// C creates the cardinal translation for the locale given the 'key', 'num' and 'digit' arguments and param passed in
func (t *translator) C(key interface{}, num float64, digits uint64, param string) (string, error) {

	tarr, ok := t.cardinalTanslations[key]
	if !ok {
		return unknownTranslation, ErrUnknowTranslation
	}

	rule := t.CardinalPluralRule(num, digits)

	trans := tarr[rule]

	// maybe pool these later?...
	b := make([]byte, 0, 64)
	b = append(b, trans.text[:trans.indexes[0]]...)
	b = append(b, param...)
	b = append(b, trans.text[trans.indexes[1]:]...)

	return string(b), nil
}

// O creates the ordinal translation for the locale given the 'key', 'num' and 'digit' arguments and param passed in
func (t *translator) O(key interface{}, num float64, digits uint64, param string) (string, error) {

	tarr, ok := t.ordinalTanslations[key]
	if !ok {
		return unknownTranslation, ErrUnknowTranslation
	}

	rule := t.OrdinalPluralRule(num, digits)

	trans := tarr[rule]

	// maybe pool these later?...
	b := make([]byte, 0, 64)
	b = append(b, trans.text[:trans.indexes[0]]...)
	b = append(b, param...)
	b = append(b, trans.text[trans.indexes[1]:]...)

	return string(b), nil
}

// R creates the range translation for the locale given the 'key', 'num1', 'digit1', 'num2' and 'digit2' arguments
// and 'param1' and 'param2' passed in
func (t *translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) (string, error) {

	tarr, ok := t.rangeTanslations[key]
	if !ok {
		return unknownTranslation, ErrUnknowTranslation
	}

	rule := t.RangePluralRule(num1, digits1, num2, digits2)

	trans := tarr[rule]

	// maybe pool these later?...
	b := make([]byte, 0, 64)
	b = append(b, trans.text[:trans.indexes[0]]...)
	b = append(b, param1...)
	b = append(b, trans.text[trans.indexes[1]:trans.indexes[2]]...)
	b = append(b, param2...)
	b = append(b, trans.text[trans.indexes[3]:]...)

	return string(b), nil
}

// TODO: Add VerifyTranslations function to check that all plurals rules are accounted for after all have been added.

// VerifyTranslations checks to ensures that no plural rules have been
// missed within the translations.
func (t *translator) VerifyTranslations() error {

	for k, v := range t.cardinalTanslations {

		for _, rule := range t.PluralsCardinal() {

			if v[rule] == nil {
				return fmt.Errorf("error: missing cardinal plural rule '%s' for translation with key '%#v", rule, k)
			}
		}
	}

	for k, v := range t.ordinalTanslations {

		for _, rule := range t.PluralsOrdinal() {

			if v[rule] == nil {
				return fmt.Errorf("error: missing ordinal plural rule '%s' for translation with key '%#v", rule, k)
			}
		}
	}

	for k, v := range t.rangeTanslations {

		for _, rule := range t.PluralsRange() {

			if v[rule] == nil {
				return fmt.Errorf("error: missing range plural rule '%s' for translation with key '%#v", rule, k)
			}
		}
	}

	return nil
}
