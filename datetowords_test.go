package datetowords

import (
	"testing"
)

func TestDateToString(t *testing.T) {
	var tests = []struct {
		input string
		want string
	}{
		{"01.02.03", "первое февраля две тысячи третьего года"},
		{"10/10/1899", "десятое октября одна тысяча восемьсот девяносто девятого года"},
		{"10-10-1000", "десятое октября однотысячного года"},

	}

	for _, test := range tests {
		if got := DateToString(test.input); got != test.want {
			t.Errorf("DateToString(%q) = %v", test.input, got)
		}
	}


}