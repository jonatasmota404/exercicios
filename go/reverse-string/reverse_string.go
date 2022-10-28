package reverse

import "strings"


func Reverse(input string) string {
	str := strings.Split(input, "")

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	
	join := strings.Join(str, "")

	return join
}
