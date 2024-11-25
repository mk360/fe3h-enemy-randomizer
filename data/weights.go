package data

import (
	"fe3h-randomizer/common"
	"fmt"
	"math/rand"
)

func getWeights(difficultyLevel uint) [6]int {
	if difficultyLevel == 1 {
		return [6]int{75, 25, 0, 0, 0, 0}
	} else if difficultyLevel == 2 {
		return [6]int{20, 55, 20, 5, 0, 0}
	} else if difficultyLevel == 3 {
		return [6]int{10, 20, 55, 15, 0, 0}
	} else if difficultyLevel == 4 {
		return [6]int{1, 10, 20, 50, 19, 0}
	} else if difficultyLevel == 5 {
		return [6]int{0, 0, 20, 30, 40, 10}
	} else {
		return [6]int{0, 0, 10, 10, 35, 45}
	}
}

func GenerateWeightedCharacter(difficulty uint) error {
	var weights = getWeights(difficulty)
	var classRoll = rand.Intn(101)
	var classRank = getRankFromWeight(weights, classRoll)
	var class = common.PickNItems(SkillLevels[classRank].Classes, 1)[0]
	fmt.Println("Class:", class)

	var weapons = make([]string, getItemSlots(difficulty))
	for i := range weapons {
		var slotRoll = getRankFromWeight(weights, rand.Intn(101))
		var weapon = common.PickNItems(SkillLevels[slotRoll].Weapons, 1)
		weapons[i] = weapon[0]
	}

	fmt.Println("Loadout:", weapons)

	var shouldHaveSpells = rand.Intn(51) > 1

	if shouldHaveSpells {
		var spellCount = rand.Intn(3)
		var spells = make([]string, spellCount)
		for i := range spells {
			var diceRoll = rand.Intn(101)
			var spellRoll = getRankFromWeight(weights, diceRoll)
			var spell = common.PickNItems(SkillLevels[spellRoll].Spells, 1)
			spells[i] = spell[0]
		}

		fmt.Println("Spells:", spells)
	}

	var combatArt = common.PickNItems(COMBAT_ARTS, 1)
	fmt.Println("Combat Art:", combatArt[0])

	var maxItemCount = rand.Intn(6 - len(weapons))
	var items []string = common.PickNItems(ITEMS, maxItemCount)

	fmt.Println("Items:", items)

	var shouldDropItem = rand.Intn(101) <= int(SkillLevels[difficulty].ItemDropOdds)
	var shouldHaveBattalion = rand.Intn(101) <= int(SkillLevels[difficulty].BattalionOdds)

	if shouldDropItem {
		var inventory = append(weapons, items...)
		fmt.Println(inventory)
		var droppedItemIndex = rand.Intn(len(inventory))
		var droppedItem = inventory[droppedItemIndex]
		fmt.Println("Dropped Item:", droppedItem)
	}

	if shouldHaveBattalion {
		var level = rand.Intn(int(difficulty)) + 1
		var battalion = common.PickNItems(BATTALIONS, 1)[0]
		fmt.Println("Battalion:", battalion, "Level:", level)
	}

	var skillSlots = make([]string, getSkillSlots(difficulty))

	for i := range skillSlots {
		var slotRoll = getRankFromWeight(weights, rand.Intn(101))
		var skill = common.PickNItems(SkillLevels[slotRoll].Skills, 1)
		skillSlots[i] = skill[0]
	}

	fmt.Println("Skills:", skillSlots)

	return nil
}

func getItemSlots(difficulty uint) int {
	if difficulty <= 2 {
		return 2
	}

	if difficulty <= 5 {
		return 3
	}

	return 4
}

func getSkillSlots(difficulty uint) int {
	if difficulty <= 3 {
		return 3
	}

	return int(difficulty)
}

func getRankFromWeight(weighted [6]int, roll int) uint {
	var lowerBound = 0
	var higherBound = 0
	for i, weight := range weighted {
		higherBound += weight
		if i > 0 {
			lowerBound += weighted[i-1]
		}
		if lowerBound <= roll && roll <= higherBound {
			return uint(i + 1)
		}
	}

	return uint(1)
}
