package deck

import "testing"

func TestSingle(t *testing.T) {
	var tests = map[string]struct {
		expected string
		given    Rank
	}{
		"ace":   {"A", Ace},
		"two":   {"2", Two},
		"three": {"3", Three},
		"four":  {"4", Four},
		"five":  {"5", Five},
		"six":   {"6", Six},
		"seven": {"7", Seven},
		"eight": {"8", Eight},
		"nine":  {"9", Nine},
		"ten":   {"10", Ten},
		"jack":  {"J", Jack},
		"queen": {"Q", Queen},
		"king":  {"K", King},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tt.given.Single()
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}
		})
	}
}
