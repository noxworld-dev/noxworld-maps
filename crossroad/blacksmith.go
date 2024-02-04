package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
)

// Blacksmith variables.
var blacksmith ns.Obj
var blacksmithSpawn ns.Pointf
var blacksmithStrikeCount int
var blacksmithHeatCount int
var quenchingWaterSpawn ns.Pointf

// Animations.
var heatingColdSword bool
var heatingWarmSword bool
var heatingHotSword bool
var swordCooledDown bool
var forging bool
var quenching bool
var idle bool

// Ideas:
// Add conditionals to implement script without the need of map editing
// - check wether brazier, water and or the forge is nearby.
// - if not only activate the forge/hammering.

// Flowchart:
// - walk to brazier and heat
// - forge at anvil
// - quench
// - bring item to merchant
// - idle and reset

// Start blacksmith script.
func initBlacksmith() {
	blacksmith = ns.Object("Blacksmith")
	blacksmithSpawn = blacksmith.Pos()
	brazier := ns.FindClosestObject(blacksmith, ns.HasTypeName{"Brazier"})
	brazier.Freeze(true)
	water := ns.FindClosestObject(blacksmith, ns.HasTypeName{"WaterBarrel"})
	quenchingWaterSpawn = water.Pos()
	// Animation
	heatingColdSword = true
	heatUpColdWeapon()
}

func heatUpColdWeapon() {
	if heatingColdSword {
		brazier := ns.Object("Brazier")
		if blacksmith.Pos().Sub(blacksmithSpawn).Len() > 10 {
			blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpColdWeapon()
			})
		} else {
			blacksmith.Equip(ns.Object("BlacksmithSword"))
			blacksmith.LookAtObject(brazier)
			blacksmith.Pause(ns.Seconds(1))
			blacksmithHeatCount = blacksmithHeatCount + 1
			if blacksmithHeatCount >= 5 {
				heatingColdSword = false
				heatingWarmSword = true
				heatUpWarmWeapon()
			}
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpColdWeapon()
			})
		}
	}
}

func heatUpWarmWeapon() {
	if heatingWarmSword {
		brazier := ns.Object("Brazier")
		if blacksmith.Pos().Sub(blacksmithSpawn).Len() > 10 {
			blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpWarmWeapon()
			})
		} else {
			blacksmith.Equip(ns.Object("BlacksmithWarmSword"))
			blacksmith.LookAtObject(brazier)
			blacksmith.Pause(ns.Seconds(1))
			blacksmithHeatCount = blacksmithHeatCount + 1
			if blacksmithHeatCount >= 10 {
				heatingWarmSword = false
				heatingHotSword = true
				heatUpHotWeapon()
			}
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpWarmWeapon()
			})
		}
	}
}

func heatUpHotWeapon() {
	if heatingHotSword {
		brazier := ns.Object("Brazier")
		if blacksmith.Pos().Sub(blacksmithSpawn).Len() > 10 {
			blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpHotWeapon()
			})
		} else {
			blacksmith.Equip(ns.Object("BlacksmithHotSword"))
			blacksmith.LookAtObject(brazier)
			blacksmith.Pause(ns.Seconds(1))
			blacksmithHeatCount = blacksmithHeatCount + 1
			if blacksmithHeatCount >= 15 {
				heatingHotSword = false
				forging = true
				blacksmithHeatCount = 0
				hitAnvil()
			}
			ns.NewTimer(ns.Seconds(1), func() {
				heatUpHotWeapon()
			})
		}
	}
}

func hitAnvil() {
	if forging {
		anvil := ns.Object("Anvil")
		blacksmith.Equip(ns.Object("BlacksmithHammer"))
		if blacksmith.Pos().Sub(blacksmithSpawn).Len() > 10 {
			blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
			ns.NewTimer(ns.Seconds(1), func() {
				hitAnvil()
			})
		} else {
			blacksmith.LookAtObject(anvil)
			blacksmith.HitMelee(blacksmith.Pos())
			blacksmithStrikeCount = blacksmithStrikeCount + 1
			if blacksmithStrikeCount <= 9 {
				ns.NewTimer(ns.Frames(7), func() {
					ns.Effect(effect.DAMAGE_POOF, anvil.Pos(), anvil.Pos())
				})
			} else {
				blacksmithStrikeCount = 0
				forging = false
				quenching = true
				quenchWeapon()
			}
			ns.NewTimer(ns.Seconds(1), func() {
				if blacksmithStrikeCount == 5 {
					swordCooledDown = true
					forging = false
					blacksmithHeatCount = 5
					heatingWarmSword = true
					heatUpWarmWeapon()
				}
				if !swordCooledDown {
					hitAnvil()
				}
				swordCooledDown = false
			})
		}
	}
}

func quenchWeapon() {
	if quenching {
		blacksmith.Equip(ns.Object("BlacksmithWarmSword"))
		water := ns.FindClosestObject(blacksmith, ns.HasTypeName{"WaterBarrel"}, ns.InCirclef{Center: blacksmith, R: 200})
		if water == nil {
			// Call next --> sell to vendor.
			rnd := ns.Random(1, 5)
			if rnd == 1 {
				ns.AudioEvent(audio.TauntShakeFist, blacksmith)
				blacksmith.ChatStr("Where's my quenching water?")
			}
			if rnd == 2 {
				ns.CreateObject("WaterBarrel", quenchingWaterSpawn)
			}
			quenching = false
			idle = true
			resetAnimation()
		} else {
			water.Freeze(true)
			blacksmith.WalkTo(water.Pos())
			blacksmith.OnEvent(ns.EventCollision, func() {
				if ns.GetCaller() == water {
					if quenching {
						quenching = false
						ns.Effect(effect.SMOKE_BLAST, water.Pos().Add(ns.Pointf{X: 0, Y: -50}), water.Pos())
						blacksmith.Unequip(ns.Object("BlacksmithWarmSword"))
						blacksmith.Pause(ns.Seconds(2))
						ns.AudioEvent(audio.RunOnWater, water)
						ns.NewTimer(ns.Seconds(2), func() {
							ns.AudioEvent(audio.MetalWeaponPickup, blacksmith)
							blacksmith.Equip(ns.Object("BlacksmithSword"))
							blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
							// Call next --> sell to vendor.
							water.Freeze(false)
							idle = true
							resetAnimation()
						})
					}
				}
			})
		}
	}
}

func resetAnimation() {
	if idle {
		blacksmith.Guard(blacksmithSpawn, blacksmithSpawn, 300)
		ns.NewTimer(ns.Seconds(60), func() {
			idle = false
			heatingColdSword = true
			blacksmithHeatCount = 0
			heatUpColdWeapon()
		})
	}
}

//func sellEquipment() {
//}
