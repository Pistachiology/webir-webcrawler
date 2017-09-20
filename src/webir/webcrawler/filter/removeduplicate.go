package filter

// RemoveDuplicate url
func RemoveDuplicate(list []string) []string {
	length := len(list) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if list[i] == list[j] {
				list[j] = list[length]
				list = list[:length]
				length--
				j--
			}
		}
	}
	return list
}
