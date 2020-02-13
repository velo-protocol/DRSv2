package utils

func ParseRevertMessage(untrimmedBytes []byte) string {
	if len(untrimmedBytes) > 68 {
		return string(untrimmedBytes[68:])
	}
	return string(untrimmedBytes)
}
