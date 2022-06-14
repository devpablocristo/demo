package usecases

import (
	"fmt"
	"sort"
)

func MaximizeTotal(amout float32, items map[string]int) []string {
	//var itemsSorted []string

	itemsSorted := sortMapByValue(items)

	fmt.Println(items)
	fmt.Println(itemsSorted)

	//Sort your prices from smallest to largest before running your loop.
	//For your example it adds 2, then 3, then 5 and sees that is larger
	//than 7 so it returns 2. If it was in order it would add 1, 2, and 3 before getting to 5.

	itemsMaxComplete := mapToSlice(itemsSorted)

	checkAmouth := float32(0)
	itemsMax := make([]string, 0, len(itemsMaxComplete))
	for i := 0; i < len(itemsMaxComplete); i++ {
		checkAmouth += amout
		if checkAmouth > amout {
			break
		}

		itemsMax = append(itemsMax, itemsMaxComplete[i])

	}

	return itemsMax
}

func mapToSlice(m map[string]int) []string {
	v := make([]string, 0, len(m))

	for s := range m {
		v = append(v, s)
	}

	return v
}

func sortMapByValue(m map[string]int) map[string]int {
	// basket := map[string]int{"orange": 5, "apple": 7,
	// "mango": 3, "strawberry": 9}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	fmt.Println(m)

	r := make(map[string]int)
	for _, k := range keys {
		fmt.Println(k, m[k])
		r[k] = m[k]
	}

	return r
}
