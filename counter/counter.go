package counter

import (
	"fmt"
	"sort"
)

// Counter сокращеет строку aaabbbccccc -> a3b3c5
func Counter(s string) string {
	cntMap := make(map[rune]int)
	cntKeys := make([]rune, 0)

	for _, c := range s {
		if cnt, ok := cntMap[c]; ok {
			// не первый раз встретили символ - инкремент
			cntMap[c] = cnt + 1
		} else {
			// первый раз встретили символ
			cntKeys = append(cntKeys, c)
			cntMap[c] = 1
		}
	}

	// сортируем cntKeys
	sort.Slice(cntKeys, func(i, j int) bool {
		return cntKeys[i] < cntKeys[j]
	})

	// собираем result по отсортированному cntKeys и берём счетчики из cntMap
	result := ""
	for _, key := range cntKeys {
		result += string(key) + fmt.Sprint(cntMap[key])
	}
	return result
}
