package phone_util

import (
	"regexp"
	"strings"
)

var Regex, _ = regexp.Compile("(03|05|07|08|09|01[2689])+([0-9]{8})")

/**
8488888888
+8488888888
0988888888

Will return

088888888
088888888
0988888888
*/
func RemovePrefix(phone string) string {
	if strings.HasPrefix(phone, "+84") {
		return "0" + phone[3:]
	} else if strings.HasPrefix(phone, "84") {
		return "0" + phone[2:]
	}

	return phone
}

//ex:
//input :0328839667
//output:+84328839667
func AddPrefix84Plus(phone string) string {
	if strings.HasPrefix(phone, "+84") {
		return phone
	} else if strings.HasPrefix(phone, "84") {
		return "+" + phone
	} else if strings.HasPrefix(phone, "0") {
		return "+84" + phone[1:]
	}

	return phone
}

//ex:
//input :0328839667
//output:84328839667
func AddPrefix84(phone string) string {
	if strings.HasPrefix(phone, "84") {
		return phone
	} else if strings.HasPrefix(phone, "+84") {
		return phone[1:]
	} else if strings.HasPrefix(phone, "0") {
		return "84" + phone[1:]
	}

	return phone
}
