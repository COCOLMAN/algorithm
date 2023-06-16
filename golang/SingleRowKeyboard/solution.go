package SingleRowKeyboard

func diff(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func calculateTime(keyboard string, word string) int {
	keyMap := map[rune]int{}

	for idx, key := range keyboard {
		keyMap[key] = idx
	}

	position := 0
	result := 0
	for _, w := range word {
		result += diff(keyMap[w], position)
		position = keyMap[w]
	}

	return result
}
