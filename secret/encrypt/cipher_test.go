package encrypt

import "testing"

func TestEncrypt(t *testing.T) {
	type given struct {
		key       string
		plaintext string
	}
	type expected struct {
		hasValue bool
		hasErr   bool
	}
	var tests = map[string]struct {
		given    given
		expected expected
	}{
		"basic": {
			given: given{
				key:       "foobar",
				plaintext: "Hello, world",
			},
			expected: expected{
				hasValue: true,
			},
		},
		"empty key": {
			given: given{
				key:       "",
				plaintext: "Hello, world",
			},
			expected: expected{
				hasValue: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualValue, actualErr := Encrypt(tt.given.key, tt.given.plaintext)
			if tt.expected.hasErr && actualErr == nil {
				t.Error("expected an error")
			}
			if actualValue == "" && tt.expected.hasValue {
				t.Errorf("expected a value")
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type given struct {
		key       string
		cipherHex string
	}
	type expected struct {
		value  string
		hasErr bool
	}
	var tests = map[string]struct {
		given    given
		expected expected
	}{
		"basic": {
			given: given{
				key:       "foobar",
				cipherHex: "e4d945e68912867ab85a5a04cd0233f24ebf09cea2bcdb567e469121",
			},
			expected: expected{
				value: "Hello, world",
			},
		},
		"wrong cipher hex": {
			given: given{
				key:       "foobar",
				cipherHex: "thisiswrong",
			},
			expected: expected{
				hasErr: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualValue, actualErr := Decrypt(tt.given.key, tt.given.cipherHex)
			if tt.expected.hasErr && actualErr == nil {
				t.Error("expected an error")
			}
			if actualValue != tt.expected.value {
				t.Errorf("expected %v, actual %v", tt.expected.value, actualValue)
			}
		})
	}
}
