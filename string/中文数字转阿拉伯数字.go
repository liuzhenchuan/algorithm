package string

func wordTonum(s string) (ans uint8) {
	m := make(map[string]int)
	m["一"] = 1
	m["二"] = 2
	m["三"] = 3
	m["四"] = 4
	m["五"] = 5
	m["六"] = 6
	m["七"] = 7
	m["八"] = 8
	m["九"] = 9
	m["十"] = 10
	m["百"] = 100
	m["千"] = 1000
	m["零"] = 0

	len := len(s)
	if len == 0 {
		return 0
	}
	for i := len - 1; i >= 0; i-- {
		if string(s[i]) == "零" {
			continue
		}
		tmp := s[i]
		if i == 0 {
			ans = ans + tmp
			break
		}
		tmp1 := s[i-1]
		tmp2 := m[string(tmp1)]
		ans = ans + tmp*uint8(tmp2)
		i--
	}
	return ans
}
