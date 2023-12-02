package main

import third_party "patterns/decorator/third-party"

func main() {
	basicSoldier := third_party.BasicSoldier{}
	third_party.DisplaySoldier(basicSoldier)

	soldierWithSword := SoldierWithSword{soldier: basicSoldier}
	third_party.DisplaySoldier(soldierWithSword)

	soldierWithShield := SoldierWithShield{soldier: basicSoldier}
	third_party.DisplaySoldier(soldierWithShield)

	soldierWithSwordAndShield := SoldierWithShield{
		soldier: SoldierWithSword{
			soldier: basicSoldier,
		},
	}
	third_party.DisplaySoldier(soldierWithSwordAndShield)
}

// Decorator 1: Soldier with a sword

type SoldierWithSword struct {
	soldier third_party.Soldier
}

func (s SoldierWithSword) Attack() int {
	return s.soldier.Attack() + 10
}

func (s SoldierWithSword) Defence() int {
	return s.soldier.Defence()
}

// Decorator 2: Soldier with a shield

type SoldierWithShield struct {
	soldier third_party.Soldier
}

func (s SoldierWithShield) Attack() int {
	return s.soldier.Attack() - 6
}

func (s SoldierWithShield) Defence() int {
	return s.soldier.Defence() + 20
}
