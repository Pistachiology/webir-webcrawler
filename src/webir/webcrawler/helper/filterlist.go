package helper

// FilterList convert func filter url to urlList
func FilterList(filter func(string) bool) func([]string) []string {
	return func(list []string) []string {
		var filtered []string
		for _, item := range list {
			if filter(item) {
				filtered = append(filtered, item)
			}
		}
		return filtered
	}
}
