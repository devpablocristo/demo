package usecases

import (
	"fmt"
	"sort"
)

func MaximizeTotal(amout float32, items map[string]float32) []string {
	itemsSorted := sortMapByValueAsc(items)

	amoutSum := float32(0)
	amoutAux := float32(0)
	itemsMax := []string{}
	for itemName, itemPrice := range itemsSorted {
		amoutSum += itemPrice
		if amoutSum < amout {
			itemsMax = append(itemsMax, itemName)
			amoutAux += itemPrice
		}
	}

	sort.Strings(itemsMax)

	fmt.Println("amountAux:", amoutAux)
	return itemsMax
}

func sortMapByValueAsc(m map[string]float32) map[string]float32 {

	type kv struct {
		Key   string
		Value float32
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	r := make(map[string]float32)
	for _, kv := range ss {
		r[kv.Key] = kv.Value
	}

	return r
}
