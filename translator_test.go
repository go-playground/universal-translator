package ut

import (
	"fmt"
	"testing"

	"github.com/go-playground/locales"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestBasicTranslation(t *testing.T) {

	uni, err := New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	en := uni.GetTranslator("en") // or fallback if fails to find 'en'
	en.Add("test_trans", "Welcome {0} to the {1}.")
	en.Add("test_trans2", "{0} to the {1}.")
	en.Add("test_trans3", "Welcome {0} to the {1}")
	en.Add("test_trans4", "{0}{1}")

	tests := []struct {
		key           interface{}
		params        []string
		expected      string
		expectedError bool
	}{
		{
			key:      "test_trans",
			params:   []string{"Joeybloggs", "The Test"},
			expected: "Welcome Joeybloggs to the The Test.",
		},
		{
			key:      "test_trans2",
			params:   []string{"Joeybloggs", "The Test"},
			expected: "Joeybloggs to the The Test.",
		},
		{
			key:      "test_trans3",
			params:   []string{"Joeybloggs", "The Test"},
			expected: "Welcome Joeybloggs to the The Test",
		},
		{
			key:      "test_trans4",
			params:   []string{"Joeybloggs", "The Test"},
			expected: "JoeybloggsThe Test",
		},
		// bad translation
		{
			key:           "non-existant-key",
			params:        []string{"Joeybloggs", "The Test"},
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		s := en.T(tt.key, tt.params...)
		if s != tt.expected {
			if !tt.expectedError && err != ErrUnknowTranslation {
				t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
			}
		}
	}
}

func TestCardinalTranslation(t *testing.T) {

	uni, err := New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	en := uni.GetTranslator("en")

	translations := []struct {
		key           interface{}
		trans         string
		rule          locales.PluralRule
		expected      error
		expectedError bool
	}{
		// bad translation
		{
			key:           "cardinal_test",
			trans:         "You have a day left.",
			rule:          locales.PluralRuleOne,
			expected:      &ErrCardinalTranslation{text: fmt.Sprintf("error: parameter '%s' not found, may want to use 'Add' instead of 'AddCardinal'", paramZero)},
			expectedError: true,
		},
		{
			key:      "cardinal_test",
			trans:    "You have {0} day left.",
			rule:     locales.PluralRuleOne,
			expected: nil,
		},
		{
			key:      "cardinal_test",
			trans:    "You have {0} days left.",
			rule:     locales.PluralRuleOther,
			expected: nil,
		},
		{
			key:           "cardinal_test",
			trans:         "You have {0} days left.",
			rule:          locales.PluralRuleOther,
			expected:      &ErrConflictingTranslation{key: "cardinal_test", rule: locales.PluralRuleOther, text: "You have {0} days left."},
			expectedError: true,
		},
	}

	for _, tt := range translations {

		err := en.AddCardinal(tt.key, tt.trans, tt.rule)
		if err != tt.expected {
			if !tt.expectedError || err.Error() != tt.expected.Error() {
				t.Errorf("Expected '<nil>' Got '%s'", err)
			}
		}
	}

	tests := []struct {
		key           interface{}
		num           float64
		digits        uint64
		param         string
		expected      string
		expectedError bool
	}{
		{
			key:      "cardinal_test",
			num:      1,
			digits:   0,
			param:    string(en.FmtNumber(1, 0)),
			expected: "You have 1 day left.",
		},
		// bad translation key
		{
			key:           "non-existant",
			num:           1,
			digits:        0,
			param:         string(en.FmtNumber(1, 0)),
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {

		s, err := en.C(tt.key, tt.num, tt.digits, tt.param)
		if err != nil {
			if !tt.expectedError && err != ErrUnknowTranslation {
				t.Errorf("Expected '<nil>' Got '%s'", err)
			}
		}

		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestOrdinalTranslation(t *testing.T) {

	uni, err := New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	en := uni.GetTranslator("en")

	translations := []struct {
		key           interface{}
		trans         string
		rule          locales.PluralRule
		expected      error
		expectedError bool
	}{
		// bad translation
		{
			key:           "day",
			trans:         "st",
			rule:          locales.PluralRuleOne,
			expected:      &ErrOrdinalTranslation{text: fmt.Sprintf("error: parameter '%s' not found, may want to use 'Add' instead of 'AddOrdinal'", paramZero)},
			expectedError: true,
		},
		{
			key:      "day",
			trans:    "{0}st",
			rule:     locales.PluralRuleOne,
			expected: nil,
		},
		{
			key:      "day",
			trans:    "{0}nd",
			rule:     locales.PluralRuleTwo,
			expected: nil,
		},
		{
			key:      "day",
			trans:    "{0}rd",
			rule:     locales.PluralRuleFew,
			expected: nil,
		},
		{
			key:      "day",
			trans:    "{0}th",
			rule:     locales.PluralRuleOther,
			expected: nil,
		},
		// bad translation
		{
			key:           "day",
			trans:         "{0}th",
			rule:          locales.PluralRuleOther,
			expected:      &ErrConflictingTranslation{key: "day", rule: locales.PluralRuleOther, text: "{0}th"},
			expectedError: true,
		},
	}

	for _, tt := range translations {

		err := en.AddOrdinal(tt.key, tt.trans, tt.rule)
		if err != tt.expected {
			if !tt.expectedError || err.Error() != tt.expected.Error() {
				t.Errorf("Expected '<nil>' Got '%s'", err)
			}
		}
	}

	tests := []struct {
		key           interface{}
		num           float64
		digits        uint64
		param         string
		expected      string
		expectedError bool
	}{
		{
			key:      "day",
			num:      1,
			digits:   0,
			param:    string(en.FmtNumber(1, 0)),
			expected: "1st",
		},
		{
			key:      "day",
			num:      2,
			digits:   0,
			param:    string(en.FmtNumber(2, 0)),
			expected: "2nd",
		},
		{
			key:      "day",
			num:      3,
			digits:   0,
			param:    string(en.FmtNumber(3, 0)),
			expected: "3rd",
		},
		{
			key:      "day",
			num:      4,
			digits:   0,
			param:    string(en.FmtNumber(4, 0)),
			expected: "4th",
		},
		{
			key:      "day",
			num:      10258.43,
			digits:   0,
			param:    string(en.FmtNumber(10258.43, 0)),
			expected: "10,258th",
		},
		// bad translation
		{
			key:           "d-day",
			num:           10258.43,
			digits:        0,
			param:         string(en.FmtNumber(10258.43, 0)),
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {

		s, err := en.O(tt.key, tt.num, tt.digits, tt.param)
		if err != nil {
			if !tt.expectedError && err != ErrUnknowTranslation {
				t.Errorf("Expected '<nil>' Got '%s'", err)
			}
		}

		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestRangeTranslation(t *testing.T) {

	uni, err := New("nl", "nl")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// dutch
	nl := uni.GetTranslator("nl")

	translations := []struct {
		key           interface{}
		trans         string
		rule          locales.PluralRule
		expected      error
		expectedError bool
	}{
		// bad translation
		{
			key:           "day",
			trans:         "er -{1} dag vertrokken",
			rule:          locales.PluralRuleOne,
			expected:      &ErrRangeTranslation{text: fmt.Sprintf("error: parameter '%s' not found, are you sure you're adding a Range Translation?", paramZero)},
			expectedError: true,
		},
		// bad translation
		{
			key:           "day",
			trans:         "er {0}- dag vertrokken",
			rule:          locales.PluralRuleOne,
			expected:      &ErrRangeTranslation{text: fmt.Sprintf("error: parameter '%s' not found, a Range Translation requires two parameters", paramOne)},
			expectedError: true,
		},
		{
			key:      "day",
			trans:    "er {0}-{1} dag vertrokken",
			rule:     locales.PluralRuleOne,
			expected: nil,
		},
		{
			key:      "day",
			trans:    "er zijn {0}-{1} dagen over",
			rule:     locales.PluralRuleOther,
			expected: nil,
		},
		// bad translation
		{
			key:           "day",
			trans:         "er zijn {0}-{1} dagen over",
			rule:          locales.PluralRuleOther,
			expected:      &ErrConflictingTranslation{key: "day", rule: locales.PluralRuleOther, text: "er zijn {0}-{1} dagen over"},
			expectedError: true,
		},
	}

	for _, tt := range translations {

		err := nl.AddRange(tt.key, tt.trans, tt.rule)
		if err != tt.expected {
			if !tt.expectedError || err.Error() != tt.expected.Error() {
				t.Errorf("Expected '%#v' Got '%s'", tt.expected, err)
			}
		}
	}

	tests := []struct {
		key           interface{}
		num1          float64
		digits1       uint64
		num2          float64
		digits2       uint64
		param1        string
		param2        string
		expected      string
		expectedError bool
	}{
		{
			key:      "day",
			num1:     1,
			digits1:  0,
			num2:     2,
			digits2:  0,
			param1:   string(nl.FmtNumber(1, 0)),
			param2:   string(nl.FmtNumber(2, 0)),
			expected: "er zijn 1-2 dagen over",
		},
		{
			key:      "day",
			num1:     0,
			digits1:  0,
			num2:     1,
			digits2:  0,
			param1:   string(nl.FmtNumber(0, 0)),
			param2:   string(nl.FmtNumber(1, 0)),
			expected: "er 0-1 dag vertrokken",
		},
		{
			key:      "day",
			num1:     0,
			digits1:  0,
			num2:     2,
			digits2:  0,
			param1:   string(nl.FmtNumber(0, 0)),
			param2:   string(nl.FmtNumber(2, 0)),
			expected: "er zijn 0-2 dagen over",
		},
		// bad translations from here
		{
			key:           "d-day",
			num1:          0,
			digits1:       0,
			num2:          2,
			digits2:       0,
			param1:        string(nl.FmtNumber(0, 0)),
			param2:        string(nl.FmtNumber(2, 0)),
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {

		s, err := nl.R(tt.key, tt.num1, tt.digits1, tt.num2, tt.digits2, tt.param1, tt.param2)
		if err != nil {
			if !tt.expectedError && err != ErrUnknowTranslation {
				t.Errorf("Expected '<nil>' Got '%s'", err)
			}
		}

		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFallbackTranslator(t *testing.T) {

	uni, err := New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	en := uni.GetTranslator("en")
	if en.Locale() != "en" {
		t.Errorf("Expected '%s' Got '%s'", "en", en.Locale())
	}

	fallback := uni.GetTranslator("nl")
	if fallback.Locale() != "en" {
		t.Errorf("Expected '%s' Got '%s'", "en", fallback.Locale())
	}

	en = uni.FindTranslator("nl", "en")
	if en.Locale() != "en" {
		t.Errorf("Expected '%s' Got '%s'", "en", en.Locale())
	}

	fallback = uni.FindTranslator("nl")
	if fallback.Locale() != "en" {
		t.Errorf("Expected '%s' Got '%s'", "en", fallback.Locale())
	}
}

func TestNew(t *testing.T) {

	_, err := New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	expected := &ErrLocaleNotFound{text: fmt.Sprintf("locale '%s' could not be found, perhaps it is not supported.", "definitly doesn't exist!")}

	_, err = New("en", "definitly doesn't exist!")
	if err == nil || err.Error() != expected.Error() {
		t.Fatalf("Expected '%s' Got '%s'", expected, err)
	}

	expected = &ErrLocaleNotFound{text: fmt.Sprintf("fallback locale '%s' could not be found", "fallback definitly doesn't exist!")}

	_, err = New("fallback definitly doesn't exist!", "en")
	if err == nil || err.Error() != expected.Error() {
		t.Fatalf("Expected '%s' Got '%s'", expected, err)
	}
}

func TestVerifyTranslations(t *testing.T) {

	// dutch
	uni, err := New("nl", "nl")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	loc := uni.GetTranslator("nl")
	if loc.Locale() != "nl" {
		t.Errorf("Expected '%s' Got '%s'", "nl", loc.Locale())
	}

	// cardinal checks

	err = loc.AddCardinal("day", "je {0} dag hebben verlaten", locales.PluralRuleOne)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// fail cardinal rules
	expected := &ErrMissingPluralTranslation{translationType: "plural", rule: locales.PluralRuleOther, key: "day"}
	err = loc.VerifyTranslations()
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected '%s' Got '%s'", expected, err)
	}

	// success cardinal
	err = loc.AddCardinal("day", "je {0} dagen hebben verlaten", locales.PluralRuleOther)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	err = loc.VerifyTranslations()
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// range checks
	err = loc.AddRange("day", "je {0}-{1} dagen hebben verlaten", locales.PluralRuleOther)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// fail range rules
	expected = &ErrMissingPluralTranslation{translationType: "range", rule: locales.PluralRuleOne, key: "day"}
	err = loc.VerifyTranslations()
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected '%s' Got '%s'", expected, err)
	}

	// success range
	err = loc.AddRange("day", "je {0}-{1} dag hebben verlaten", locales.PluralRuleOne)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	err = loc.VerifyTranslations()
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// ok so 'nl' aka dutch, ony has one plural rule for ordinals, so going to switch to english from here which has 4

	uni, err = New("en", "en")
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	loc = uni.GetTranslator("en")
	if loc.Locale() != "en" {
		t.Errorf("Expected '%s' Got '%s'", "en", loc.Locale())
	}

	// ordinal checks

	err = loc.AddOrdinal("day", "{0}st", locales.PluralRuleOne)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	err = loc.AddOrdinal("day", "{0}rd", locales.PluralRuleFew)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	err = loc.AddOrdinal("day", "{0}th", locales.PluralRuleOther)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	// fail ordinal rules
	expected = &ErrMissingPluralTranslation{translationType: "ordinal", rule: locales.PluralRuleTwo, key: "day"}
	err = loc.VerifyTranslations()
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected '%s' Got '%s'", expected, err)
	}

	// success ordinal

	err = loc.AddOrdinal("day", "{0}nd", locales.PluralRuleTwo)
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}

	err = loc.VerifyTranslations()
	if err != nil {
		t.Fatalf("Expected '<nil>' Got '%s'", err)
	}
}
