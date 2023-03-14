package main

// import (
// 	"fmt"
// )

// type Constructor struct {
// 	Damage int
// 	HP     int
// }

// func findMinDamage(mapStruct map[int]Constructor) int {
// 	min := mapStruct[0].Damage
// 	for i := 0; mapStruct[i].HP > 0; i++ {
// 		if min > mapStruct[i].Damage {
// 			min = mapStruct[i].Damage
// 			fmt.Println("min => ", min)
// 		}
// 	}
// 	return min
// }

// func monster(power, number int, hp, damage []int) string {
// 	mapStruct := make(map[int]Constructor)
// 	for i := 0; i < number; i++ {
// 		mapStruct[i] = Constructor{
// 			Damage: damage[i],
// 			HP:     hp[i],
// 		}
// 	}
// 	for j := 0; power > 0; j++ {
// 		for i := 0; i < len(mapStruct); i++ {
// 			mapStruct[i] = Constructor{
// 				HP:     mapStruct[i].HP - power,
// 				Damage: mapStruct[i].Damage,
// 			}
// 		}
// 		power -= findMinDamage(mapStruct)
// 		if power <= 0 {
// 			return "NO"
// 		}
// 	}
// 	return "YES"
// }

// func main() {
// 	// 3
// 	// 6 7
// 	// 18 5 13 9 10 1
// 	// 2 7 2 1 2 6
// 	// 3 4
// 	// 5 5 5
// 	// 4 4 4
// 	// 3 2
// 	// 2 1 3
// 	// 1 1 1
// 	power := 7
// 	number := 6
// 	arr := []int{18, 5, 13, 9, 10, 1}
// 	damage := []int{2, 7, 2, 1, 2, 6}
// 	fmt.Println(monster(power, number, damage, arr))
// }
