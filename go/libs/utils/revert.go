package utils

func ParseRevertMessage(untrimmedBytes []byte) string {
	if len(untrimmedBytes) > 68 {
		return BytesToString(untrimmedBytes)[68:]
	}
	return BytesToString(untrimmedBytes)
}
