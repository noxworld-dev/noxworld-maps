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
	strikeCount         int
	heatCount           int
	quenchingWaterSpawn ns.Pointf

	// Animations.
	heatingColdSword bool
	heatingWarmSword bool
	heatingHotSword  bool
	swordCooledDown  bool
	forging          bool
	quenching        bool
	idle             bool
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
	// Animation
	b.heatingColdSword = true
	b.heatUpColdWeapon()
	return b
}

func (b *Blacksmith) heatUpColdWeapon() {
	if !b.heatingColdSword {
		return
	}
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpColdWeapon()
		})
	} else {
		b.blacksmith.Equip(b.sword)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 5 {
			b.heatingColdSword = false
			b.heatingWarmSword = true
			b.heatUpWarmWeapon()
		}
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpColdWeapon()
		})
	}
}

func (b *Blacksmith) heatUpWarmWeapon() {
	if !b.heatingWarmSword {
		return
	}
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpWarmWeapon()
		})
	} else {
		b.blacksmith.Equip(b.swordWarm)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 10 {
			b.heatingWarmSword = false
			b.heatingHotSword = true
			b.heatUpHotWeapon()
		}
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpWarmWeapon()
		})
	}
}

func (b *Blacksmith) heatUpHotWeapon() {
	if !b.heatingHotSword {
		return
	}
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpHotWeapon()
		})
	} else {
		b.blacksmith.Equip(b.swordHot)
		b.blacksmith.LookAtObject(b.brazier)
		b.blacksmith.Pause(ns.Seconds(1))
		b.heatCount++
		if b.heatCount >= 15 {
			b.heatingHotSword = false
			b.forging = true
			b.heatCount = 0
			b.hitAnvil()
		}
		ns.NewTimer(ns.Seconds(1), func() {
			b.heatUpHotWeapon()
		})
	}
}

func (b *Blacksmith) hitAnvil() {
	if !b.forging {
		return
	}
	b.blacksmith.Equip(b.hammer)
	if b.blacksmith.Pos().Sub(b.spawn).Len() > 10 {
		b.blacksmith.Guard(b.spawn, b.spawn, 300)
		ns.NewTimer(ns.Seconds(1), func() {
			b.hitAnvil()
		})
	} else {
		b.blacksmith.LookAtObject(b.anvil)
		b.blacksmith.HitMelee(b.blacksmith.Pos())
		b.strikeCount = b.strikeCount + 1
		if b.strikeCount <= 9 {
			ns.NewTimer(ns.Frames(7), func() {
				ns.Effect(effect.DAMAGE_POOF, b.anvil.Pos(), b.anvil.Pos())
			})
		} else {
			b.strikeCount = 0
			b.forging = false
			b.quenching = true
			b.quenchWeapon()
		}
		ns.NewTimer(ns.Seconds(1), func() {
			if b.strikeCount == 5 {
				b.swordCooledDown = true
				b.forging = false
				b.heatCount = 5
				b.heatingWarmSword = true
				b.heatUpWarmWeapon()
			}
			if !b.swordCooledDown {
				b.hitAnvil()
			}
			b.swordCooledDown = false
		})
	}
}

func (b *Blacksmith) quenchWeapon() {
	if !b.quenching {
		return
	}
	b.blacksmith.Equip(b.swordWarm)
	water := ns.FindClosestObject(b.blacksmith, ns.HasTypeName{"WaterBarrel"}, ns.InCirclef{Center: b.blacksmith, R: 200})
	if water == nil {
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
			ns.CreateObject("WaterBarrel", b.quenchingWaterSpawn)
		}
		b.quenching = false
		b.idle = true
		b.resetAnimation()
	} else {
		water.Freeze(true)
		b.blacksmith.WalkTo(water.Pos())
		b.blacksmith.OnEvent(ns.EventCollision, func() {
			if caller := ns.GetCaller(); caller != water {
				return
			}
			if !b.quenching {
				return
			}
			b.quenching = false
			ns.Effect(effect.SMOKE_BLAST, water.Pos().Add(ns.Pointf{X: 0, Y: -50}), water.Pos())
			b.blacksmith.Unequip(b.swordWarm)
			b.blacksmith.Pause(ns.Seconds(2))
			ns.AudioEvent(audio.RunOnWater, water)
			ns.NewTimer(ns.Seconds(2), func() {
				ns.AudioEvent(audio.MetalWeaponPickup, b.blacksmith)
				b.blacksmith.Equip(b.sword)
				b.blacksmith.Guard(b.spawn, b.spawn, 300)
				// Call next --> sell to vendor.
				water.Freeze(false)
				b.idle = true
				b.resetAnimation()
			})
		})
	}
}

func (b *Blacksmith) resetAnimation() {
	if !b.idle {
		return
	}
	b.blacksmith.Guard(b.spawn, b.spawn, 300)
	ns.NewTimer(ns.Seconds(60), func() {
		b.idle = false
		b.heatingColdSword = true
		b.heatCount = 0
		b.heatUpColdWeapon()
	})
}
