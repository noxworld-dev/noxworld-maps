package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
)

// Start blacksmith script.
func init() {
	OnLateInit(func() {
		NewBlacksmith("Blacksmith")
	})
}

type BlacksmithState int

const (
	BlacksmithHeatingColdSword = BlacksmithState(iota)
	BlacksmithHeatingWarmSword
	BlacksmithHeatingHotSword
	BlacksmithForging
	BlacksmithQuenchLookForWater
	BlacksmithQuenching
	BlacksmithIdle
)

// Blacksmith variables.
type Blacksmith struct {
	blacksmith          ns.Obj
	brazier             ns.Obj
	anvil               ns.Obj
	sword               ns.Obj
	swordWarm           ns.Obj
	swordHot            ns.Obj
	hammer              ns.Obj
	spawn               ns.Pointf
	quenchingWaterSpawn ns.Pointf
	state               BlacksmithState
	strikeCount         int
	heatCount           int
	water               ns.Obj
}

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

func NewBlacksmith(name string) *Blacksmith {
	b := &Blacksmith{
		blacksmith: ns.Object(name),
	}
	b.spawn = b.blacksmith.Pos()
	b.brazier = ns.FindClosestObject(b.blacksmith, ns.HasTypeName{"Brazier"})
	b.brazier.Freeze(true)
	b.anvil = ns.Object("Anvil")
	b.sword = ns.Object("BlacksmithSword")
	b.swordWarm = ns.Object("BlacksmithWarmSword")
	b.swordHot = ns.Object("BlacksmithHotSword")
	b.hammer = ns.Object("BlacksmithHammer")
	water := ns.FindClosestObject(b.blacksmith, ns.HasTypeName{"WaterBarrel"})
	b.quenchingWaterSpawn = water.Pos()

	b.SwitchState(BlacksmithHeatingColdSword)
	return b
}

func (b *Blacksmith) Update() {
	switch b.state {
	case BlacksmithHeatingColdSword:
		b.heatUpColdWeapon()
	case BlacksmithHeatingWarmSword:
		b.heatUpWarmWeapon()
	case BlacksmithHeatingHotSword:
		b.heatUpHotWeapon()
	case BlacksmithForging:
		b.hitAnvil()
	case BlacksmithQuenchLookForWater:
		b.quenchLookForWater()
	case BlacksmithQuenching:
		b.quenchingWeapon()
	case BlacksmithIdle:
		b.resetAnimation()
	}
}

func (b *Blacksmith) SwitchState(st BlacksmithState) {
	b.state = st
	b.Update()
}

func (b *Blacksmith) updateLater(sec float64) {
	ns.NewTimer(ns.Seconds(sec), func() {
		b.Update()
	})
}

func (b *Blacksmith) heatUpColdWeapon() {
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		b.updateLater(1)
	} else {
		b.blacksmith.Equip(b.sword)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 5 {
			b.SwitchState(BlacksmithHeatingWarmSword)
		} else {
			b.updateLater(1)
		}
	}
}

func (b *Blacksmith) heatUpWarmWeapon() {
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		b.updateLater(1)
	} else {
		b.blacksmith.Equip(b.swordWarm)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 10 {
			b.SwitchState(BlacksmithHeatingHotSword)
		} else {
			b.updateLater(1)
		}
	}
}

func (b *Blacksmith) heatUpHotWeapon() {
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		b.updateLater(1)
	} else {
		b.blacksmith.Equip(b.swordHot)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 15 {
			b.heatCount = 0
			b.SwitchState(BlacksmithForging)
		} else {
			b.updateLater(1)
		}
	}
}

func (b *Blacksmith) hitAnvil() {
	if b.state != BlacksmithForging {
		return
	}
	b.blacksmith.Equip(b.hammer)
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		b.updateLater(1)
	} else {
		b.blacksmith.LookAtObject(b.anvil)
		b.blacksmith.HitMelee(b.blacksmith.Pos())
		b.strikeCount++
		if b.strikeCount <= 9 {
			ns.NewTimer(ns.Frames(7), func() {
				ns.Effect(effect.DAMAGE_POOF, b.anvil.Pos(), b.anvil.Pos())
			})
		} else {
			b.strikeCount = 0
			b.SwitchState(BlacksmithQuenchLookForWater)
		}
		ns.NewTimer(ns.Seconds(1), func() {
			swordCooledDown := b.strikeCount == 5
			if swordCooledDown {
				b.heatCount = 5
				b.SwitchState(BlacksmithHeatingWarmSword)
			} else {
				b.hitAnvil()
			}
		})
	}
}

func (b *Blacksmith) quenchLookForWater() {
	if b.state != BlacksmithQuenchLookForWater {
		return
	}
	b.blacksmith.Equip(b.swordWarm)
	b.water = ns.FindClosestObject(b.blacksmith, ns.HasTypeName{"WaterBarrel"}, ns.InCirclef{Center: b.blacksmith, R: 200})
	if b.water == nil {
		// Call next --> sell to vendor.
		rnd := 2
		if !isTesting {
			rnd = ns.Random(1, 5)
		}
		switch rnd {
		case 1:
			ns.AudioEvent(audio.TauntShakeFist, b.blacksmith)
			b.blacksmith.ChatStr("Where's my quenching water?")
		case 2:
			b.water = ns.CreateObject("WaterBarrel", b.quenchingWaterSpawn)
		}
		b.SwitchState(BlacksmithIdle)
	} else {
		b.water.Freeze(true)
		b.blacksmith.WalkTo(b.water.Pos())
		b.blacksmith.OnEvent(ns.EventCollision, func() {
			if ns.GetCaller() != b.water {
				return
			}
			if b.state != BlacksmithQuenchLookForWater {
				return
			}
			b.SwitchState(BlacksmithQuenching)
		})
	}
}

func (b *Blacksmith) quenchingWeapon() {
	ns.Effect(effect.SMOKE_BLAST, b.water.Pos().Add(ns.Pointf{X: 0, Y: -50}), b.water.Pos())
	b.blacksmith.Unequip(b.swordWarm)
	b.blacksmith.Pause(ns.Seconds(2))
	ns.AudioEvent(audio.RunOnWater, b.water)
	ns.NewTimer(ns.Seconds(2), func() {
		ns.AudioEvent(audio.MetalWeaponPickup, b.blacksmith)
		b.blacksmith.Equip(b.sword)
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		// Call next --> sell to vendor.
		b.water.Freeze(false)
		b.state = BlacksmithIdle
		b.Update()
	})
}

func (b *Blacksmith) resetAnimation() {
	b.blacksmith.Guard(b.spawn, b.spawn, 300)
	ns.NewTimer(ns.Seconds(60), func() {
		b.heatCount = 0
		b.state = BlacksmithHeatingColdSword
		b.Update()
	})
}
