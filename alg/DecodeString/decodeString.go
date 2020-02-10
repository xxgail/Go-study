package DecodeString

/**
 * @Time: 2020/2/10 17:30
 * @DESC: 394. 字符串解码 中等
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 * 示例:
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * @param $s
 * @return string
 */
func DecodeString(s string) string {
	lens := len(s)
	var num int
	var str string
	var numbers []int
	var strings []string
	for i := 0; i < lens; i++ {
		// s[i] 的类型是uint8。rune()相当于转换为ASCII码，对比PHP的ord()
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 { // 数字
			num = num*10 + int(s[i]) - 48
		} else if (rune(s[i]) >= 97 && rune(s[i]) <= 122) || (rune(s[i]) >= 65 && rune(s[i]) <= 90) {
			// 字母
			str = str + string(s[i])
		} else if rune(s[i]) == 91 { // "["
			numbers = append(numbers, num) // array_push
			num = 0
			strings = append(strings, str)
			str = ""
		} else { // "]"
			var times int
			if len(numbers) != 0 {
				times = numbers[len(numbers)-1]    // 取数组的最后一位
				numbers = numbers[:len(numbers)-1] // 返回去掉最后一位的数组
				// 以上两步相当于array_pop
			}
			newStr := ""
			for j := 0; j < times; j++ {
				newStr += str
			}
			if len(strings) != 0 {
				str = strings[len(strings)-1] + newStr
				strings = strings[:len(strings)-1]
				// 以上两步相当于array_pop
			} else {
				str = newStr
			}
		}
		//fmt.Println(str,num,strings,numbers)
	}
	return str
}
