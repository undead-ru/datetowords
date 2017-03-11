package datetowords

import (
	"testing"
)

func TestDateToString(t *testing.T) {
	var tests = []struct {
		input string
		want string
		wantErr error
	}{
		{"01.02.03", "первое февраля две тысячи третьего года", nil},
		{"10/10/1899", "десятое октября одна тысяча восемьсот девяносто девятого года", nil},
		{"10-10-1000", "десятое октября однотысячного года", nil},
		{"1/1/01", "первое января две тысячи первого года", nil},

	}

	for _, test := range tests {
		if got, got2 := DateToString(test.input); got != test.want || got2 != nil {
			t.Errorf("DateToString(%q) = %v, %v", test.input, got, got2)
		}
	}
}
