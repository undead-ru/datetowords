package datetowords

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ReDate регулярное выражение для разбора входящей строки содержащей дату
var ReDate = regexp.MustCompile(`^(\d{1,2})[,|\.|\\|\-|/]{1}(\d{1,2})[,|\.|\\|\-|/]{1}(\d{2}|\d{4})$`)

// DateToString переводит дату в цифровом формате в текстовый
func DateToString(d string) (res string, err error) {

	var day, month, year string

	if ReDate.MatchString(d) == false {
		return "", fmt.Errorf("date pattern didn't match: %v", res)
	}

	s := ReDate.FindStringSubmatch(d)

	// день
	if i, err := strconv.Atoi(s[1]); err != nil {
		return "", err
	} else if i > 0 && i <= 31 {
		if len(s[1]) < 2 {
			day = dates["0"+s[1]].d
		} else {
			day = dates[s[1]].d
		}
	} else {
		return "", fmt.Errorf("day didn't match %d", i)
	}

	// месяц
	if i, err := strconv.Atoi(s[2]); err != nil {
		return "", err
	} else if i > 0 && i <= 12 {
		if len(s[2]) == 1 {
			month = dates["m0"+s[2]].d
		} else {
			month = dates["m"+s[2]].d

		}
	} else {
		return "", fmt.Errorf("month didn't match %d", i)
	}

	// год
	i, err := strconv.Atoi(s[3])
	if err == nil {
		if len(s[3]) == 2 && i > 0 && i <= 99 {
			if i <= 99 && i >= 40 {
				year = "одна тысяча девятьсот " + dates[s[3]].y + " года"
			} else if i >= 0 && i < 40 {
				year = "две тысячи " + dates[s[3]].y + " года"
			}
		} else if i >= 1000 && i <= 2999 {
			y := strings.Split(s[3], "")

			if i%1000 == 0 {
				year = dates[y[0]].d
			} else if i%100 == 0 {
				year = dates[y[0]].y + " "
				year += dates[(y[1] + "00")].d
			} else {
				year = dates[y[0]].y + " "
				year += dates[(y[1]+"00")].y + " "
				year += dates[y[2]+y[3]].y
			}

			year += " года"
		} else {
			year = ""
		}
	} else {
		year = ""
	}

	return day + " " + month + " " + year, nil
}

type dateTypes struct {
	d, y string
}

var dates = map[string]dateTypes{
	"01": dateTypes{d: "первое", y: "первого"},
	"02": dateTypes{d: "второе", y: "второго"},
	"03": dateTypes{d: "третье", y: "третьего"},
	"04": dateTypes{d: "четвертое", y: "четвертого"},
	"05": dateTypes{d: "пятое", y: "пятого"},
	"06": dateTypes{d: "шестое", y: "шестого"},
	"07": dateTypes{d: "седьмое", y: "седьмого"},
	"08": dateTypes{d: "восьмое", y: "восьмого"},
	"09": dateTypes{d: "девятое", y: "девятого"},
	"10": dateTypes{d: "десятое", y: "десятого"},
	"11": dateTypes{d: "одиннадцатое", y: "одиннадцатого"},
	"12": dateTypes{d: "двенадцатое", y: "двенадцатого"},
	"13": dateTypes{d: "тринадцатое", y: "тринадцатого"},
	"14": dateTypes{d: "четырнадцатое", y: "четырнадцатого"},
	"15": dateTypes{d: "пятнадцатое", y: "пятнадцатого"},
	"16": dateTypes{d: "шестнадцатое", y: "шестнадцатого"},
	"17": dateTypes{d: "семнадцатое", y: "семнадцатого"},
	"18": dateTypes{d: "восемнадцатое", y: "восемнадцатого"},
	"19": dateTypes{d: "девятнадцатое", y: "девятнадцатого"},
	"20": dateTypes{d: "двадцатое", y: "двадцатого"},
	"21": dateTypes{d: "двадцать первое", y: "двадцать первого"},
	"22": dateTypes{d: "двадцать второе", y: "двадцать второго"},
	"23": dateTypes{d: "двадцать третье", y: "двадцать третьего"},
	"24": dateTypes{d: "двадцать четвертое", y: "двадцать четвертого"},
	"25": dateTypes{d: "двадцать пятое", y: "двадцать пятого"},
	"26": dateTypes{d: "двадцать шестое", y: "двадцать шестого"},
	"27": dateTypes{d: "двадцать седьмое", y: "двадцать седьмого"},
	"28": dateTypes{d: "двадцать восьмое", y: "двадцать восьмого"},
	"29": dateTypes{d: "двадцать девятое", y: "двадцать девятого"},
	"30": dateTypes{d: "тридцатое", y: "тридцатого"},
	"31": dateTypes{d: "тридцать первое", y: "тридцать первого"},
	"32": dateTypes{y: "тридцать второго"},
	"33": dateTypes{y: "тридцать третьего"},
	"34": dateTypes{y: "тридцать четвертого"},
	"35": dateTypes{y: "тридцать пятого"},
	"36": dateTypes{y: "тридцать шестого"},
	"37": dateTypes{y: "тридцать седьмого"},
	"38": dateTypes{y: "тридцать восьмого"},
	"39": dateTypes{y: "тридцать девятого"},
	"40": dateTypes{y: "сорокового"},
	"41": dateTypes{y: "сорок первого"},
	"42": dateTypes{y: "сорок второго"},
	"43": dateTypes{y: "сорок третьего"},
	"44": dateTypes{y: "сорок четвертого"},
	"45": dateTypes{y: "сорок пятого"},
	"46": dateTypes{y: "сорок шестого"},
	"47": dateTypes{y: "сорок седьмого"},
	"48": dateTypes{y: "сорок восьмого"},
	"49": dateTypes{y: "сорок девятого"},
	"50": dateTypes{y: "пятидесятого"},
	"51": dateTypes{y: "пятьдесят первого"},
	"52": dateTypes{y: "пятьдесят второго"},
	"53": dateTypes{y: "пятьдесят третьего"},
	"54": dateTypes{y: "пятьдесят четвертого"},
	"55": dateTypes{y: "пятьдесят пятого"},
	"56": dateTypes{y: "пятьдесят шестого"},
	"57": dateTypes{y: "пятьдесят седьмого"},
	"58": dateTypes{y: "пятьдесят восьмого"},
	"59": dateTypes{y: "пятьдесят девятого"},
	"60": dateTypes{y: "шестидесятого"},
	"61": dateTypes{y: "шестьдесят первого"},
	"62": dateTypes{y: "шестьдесят второго"},
	"63": dateTypes{y: "шестьдесят третьего"},
	"64": dateTypes{y: "шестьдесят четвертого"},
	"65": dateTypes{y: "шестьдесят пятого"},
	"66": dateTypes{y: "шестьдесят шестого"},
	"67": dateTypes{y: "шестьдесят седьмого"},
	"68": dateTypes{y: "шестьдесят восьмого"},
	"69": dateTypes{y: "шестьдесят девятого"},
	"70": dateTypes{y: "семидесятого"},
	"71": dateTypes{y: "семьдесят первого"},
	"72": dateTypes{y: "семьдесят второго"},
	"73": dateTypes{y: "семьдесят третьего"},
	"74": dateTypes{y: "семьдесят четвертого"},
	"75": dateTypes{y: "семьдесят пятого"},
	"76": dateTypes{y: "семьдесят шестого"},
	"77": dateTypes{y: "семьдесят седьмого"},
	"78": dateTypes{y: "семьдесят восьмого"},
	"79": dateTypes{y: "семьдесят девятого"},
	"80": dateTypes{y: "восьмидесятого"},
	"81": dateTypes{y: "восемьдесят первого"},
	"82": dateTypes{y: "восемьдесят второго"},
	"83": dateTypes{y: "восемьдесят третьего"},
	"84": dateTypes{y: "восемьдесят четвертого"},
	"85": dateTypes{y: "восемьдесят пятого"},
	"86": dateTypes{y: "восемьдесят шестого"},
	"87": dateTypes{y: "восемьдесят седьмого"},
	"88": dateTypes{y: "восемьдесят восьмого"},
	"89": dateTypes{y: "восемьдесят девятого"},
	"90": dateTypes{y: "девяностого"},
	"91": dateTypes{y: "девяносто первого"},
	"92": dateTypes{y: "девяносто второго"},
	"93": dateTypes{y: "девяносто третьего"},
	"94": dateTypes{y: "девяносто четвертого"},
	"95": dateTypes{y: "девяносто пятого"},
	"96": dateTypes{y: "девяносто шестого"},
	"97": dateTypes{y: "девяносто седьмого"},
	"98": dateTypes{y: "девяносто восьмого"},
	"99": dateTypes{y: "девяносто девятого"},

	"1":   dateTypes{d: "однотысячного", y: "одна тысяча"},
	"2":   dateTypes{d: "двухтысячного", y: "две тысячи"},
	"100": dateTypes{d: "сотого", y: "сто"},
	"200": dateTypes{d: "двухсотого", y: "двести"},
	"300": dateTypes{d: "трехсотого", y: "триста"},
	"400": dateTypes{d: "четырехсотого", y: "четыреста"},
	"500": dateTypes{d: "пятисотого", y: "пятьсот"},
	"600": dateTypes{d: "шестисотого", y: "шестьсот"},
	"700": dateTypes{d: "семисотого", y: "семьсот"},
	"800": dateTypes{d: "восьмисотого", y: "восемьсот"},
	"900": dateTypes{d: "девитясотого", y: "девятьсот"},

	"m01": dateTypes{d: "января"},
	"m02": dateTypes{d: "февраля"},
	"m03": dateTypes{d: "марта"},
	"m04": dateTypes{d: "апреля"},
	"m05": dateTypes{d: "мая"},
	"m06": dateTypes{d: "июня"},
	"m07": dateTypes{d: "июля"},
	"m08": dateTypes{d: "августа"},
	"m09": dateTypes{d: "сентября"},
	"m10": dateTypes{d: "октября"},
	"m11": dateTypes{d: "ноября"},
	"m12": dateTypes{d: "декабря"},
}
