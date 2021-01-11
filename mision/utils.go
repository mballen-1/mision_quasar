package mision

// CheckSameNumberOfWordsForAllMessages ...
func CheckSameNumberOfWordsForAllMessages(messages ...[]string) bool {
	numberOfWords := len(messages[0])
	for i := range messages {
		if len(messages[i]) != numberOfWords {
			return false
		}
	}
	return true
}
