package common

import "math/rand"

func PickNItems(list []string, items int) []string {
	if items == 0 {
		return []string{}
	}
	var picks []string = make([]string, items)
	var l = len(list)
	if l == 0 {
		return []string{}
	}
	for i := 0; i < items; i++ {
		var randomChoice = rand.Intn(l)
		picks[i] = (list)[randomChoice]
	}

	return picks
}
