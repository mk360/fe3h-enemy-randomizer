package main

import (
	"bufio"
	"fe3h-randomizer/common"
	"fe3h-randomizer/data"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type ChestInfo struct {
	X       int
	Y       int
	Content string
}

// randomiser les boss
// intégrer la vérification de positions valides pour les coffres
// Améliorer le logging des résultats

func main() {
	var args = os.Args
	if args[1] == "weighted" {
		data.LoadLevels()
		var convInt, _ = strconv.Atoi(os.Args[2])
		data.GenerateWeightedCharacter(uint(convInt))
	} else if args[1] == "chest" {
		var itemPool = append(append(data.ITEMS, data.BEAST_WEAPONS...), data.WEAPONS...)
		fmt.Println(itemPool[rand.Intn(len(itemPool))])
	} else {
		pickCompletelyRandomCharacter()
		var mapX = 20
		var mapY = 19
		reader := bufio.NewReader(os.Stdin)
		var potentialChestContents = append(data.WEAPONS, data.ITEMS...)

		var hasAChest = rand.Intn(101) > 65 // 35% chance
		if hasAChest {
			var chestCount = rand.Intn(3) // max. 2
			var validChestLocations = make([]ChestInfo, chestCount)

			fmt.Println("Chests to set:", chestCount)

			for i := 0; i < chestCount; i++ {
				var chosenX = rand.Intn(mapX + 1)
				var chosenY = rand.Intn(mapY + 1)
				fmt.Println("Chest #", i+1, ": X =", chosenX, "Y =", chosenY)
				text, _, _ := reader.ReadLine()
				if strings.ToLower(string(text)) == "yes." {
					var chestContent = common.PickNItems(potentialChestContents, 1)
					validChestLocations[i] = ChestInfo{
						X:       chosenX,
						Y:       chosenY,
						Content: chestContent[0],
					}
				} else {
					fmt.Println([]byte(strings.ToLower(string(text))), []byte("yes."))
					i--
				}
			}

			for i := 0; i < chestCount; i++ {
				fmt.Println("Chest #", i+1, ": X =", validChestLocations[i].X, "Y =", validChestLocations[i].Y, "Contents", validChestLocations[i].Content)
			}

		}
	}
}

func pickCompletelyRandomCharacter() {
	var skillCount = rand.Intn(6)
	var shouldBeGiant = rand.Intn(101) > 68 // 32% chance
	var skills = common.PickNItems(data.SKILLS, skillCount)
	var shouldDropItem = rand.Intn(101) > 70 // 30% chance
	var weaponsCount = rand.Intn(6) + 1      // at least one weapon
	var weapons = common.PickNItems(data.WEAPONS, weaponsCount)
	var itemsCount = 6 - weaponsCount
	var items = common.PickNItems(data.ITEMS, itemsCount)

	if shouldBeGiant {
		fmt.Println("Giant")
		var species = common.PickNItems(data.BEASTS, 1)
		fmt.Println("Species", species)
		var weapon = common.PickNItems(data.BEAST_WEAPONS, 1)
		fmt.Println("Weapon", weapon)
		var shouldHaveWeakness = rand.Intn(101) > 50 // 50% chance
		if shouldHaveWeakness {
			var weakness = data.WEAKNESSES[rand.Intn(len(data.WEAKNESSES))]
			fmt.Println("Weakness", weakness)
		}
	} else {
		var class = common.PickNItems(data.CLASSES, 1)[0]
		fmt.Println("class", class)
		var shouldHaveBattalion = rand.Intn(101) > 75 // 25% chance
		var extraSpells = common.PickNItems(data.SPELLS, 2)
		var combatArt = common.PickNItems(data.COMBAT_ARTS, 1)

		fmt.Println("skills", skills)

		if shouldHaveBattalion {
			var battalion = common.PickNItems(data.BATTALIONS, 1)
			var level = rand.Intn(6)
			fmt.Println("battalion", battalion[0])
			fmt.Println("battalion level", level)
		}

		fmt.Println("extra spells", extraSpells)
		fmt.Println("combat art", combatArt)

		var genderRand = rand.Intn(101)
		var gender = "Male"

		if genderRand > 51 {
			gender = "Female"
		}

		var shouldHaveWeakness = rand.Intn(101) > 85 // 15% chance
		if shouldHaveWeakness {
			var weakness = data.WEAKNESSES[rand.Intn(len(data.WEAKNESSES))]
			fmt.Println("Weakness", weakness)
		}

		fmt.Println("Gender", gender)
	}

	fmt.Println("weapons", weapons)
	fmt.Println("items", items)

	if shouldDropItem {
		var dropPossibilities = append(weapons, items...)
		var droppedItem = common.PickNItems(dropPossibilities, 1)
		fmt.Println("dropped item", droppedItem)
	}

	data.LoadLevels()

	generateLevelBasedCharacter(4)
}

func generateLevelBasedCharacter(level uint) {
	_ = data.GenerateWeightedCharacter(level)
	// fmt.Println(weight)
}
