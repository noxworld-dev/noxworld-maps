package noxworld

import (
	"testing"
	"time"

	"github.com/noxworld-dev/noxscript/ns/v4"
	nstest "github.com/noxworld-dev/noxscript/ns/v4/test"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/shoenig/test/must"
)

func TestBlacksmith(t *testing.T) {
	r := nstest.GetRuntime()
	r.Update()
	r.ResetFrame()
	pos := ns.Ptf(100, 100)

	const name = "Blacksmith1"
	npc := r.NewObject(r.Types.NPC, name, pos)

	brazierType := r.NewObjectType(object.ClassImmobile, "Brazier")
	anvilType := r.NewObjectType(object.ClassImmobile, "Brazier")
	waterType := r.NewObjectType(object.ClassSimple, "WaterBarrel")
	r.NewObject(brazierType, "", pos.Add(ns.Ptf(10, 0)))
	r.NewObject(anvilType, "Anvil", pos.Add(ns.Ptf(10, 5)))
	water := r.NewObject(waterType, "", pos.Add(ns.Ptf(10, -5)))

	swordType := r.NewObjectType(object.ClassWeapon, "Longsword")
	hammerType := r.NewObjectType(object.ClassWeapon, "WarHammer")
	cold := r.NewObject(swordType, "BlacksmithSword", ns.Ptf(0, 0))
	warm := r.NewObject(swordType, "BlacksmithWarmSword", ns.Ptf(0, 0))
	hot := r.NewObject(swordType, "BlacksmithHotSword", ns.Ptf(0, 0))
	hm := r.NewObject(hammerType, "BlacksmithHammer", ns.Ptf(0, 0))
	for _, it := range []*nstest.Object{cold, warm, hot, hm} {
		npc.Pickup(it)
	}

	b := NewBlacksmith(name)
	for try := 0; try < 3; try++ {
		hasWater := try != 1
		if !hasWater {
			water.Delete()
		}

		// Start heating immediately
		must.True(t, b.heatingColdSword)
		must.EqOp(t, 1, b.heatCount)
		must.True(t, npc.HasEquipment(cold))

		// Heating it progress
		r.UpdateFor(time.Second)
		must.True(t, b.heatingColdSword)
		must.EqOp(t, 2, b.heatCount)
		must.True(t, npc.HasEquipment(cold))
		r.UpdateFor(3 * time.Second)

		// Sword is getting warm
		must.False(t, b.heatingColdSword)
		must.True(t, b.heatingWarmSword)
		must.EqOp(t, 6, b.heatCount)
		must.True(t, npc.HasEquipment(warm))
		r.UpdateFor(4 * time.Second)

		// Sword is getting hot
		must.False(t, b.heatingWarmSword)
		must.True(t, b.heatingHotSword)
		must.EqOp(t, 11, b.heatCount)
		must.True(t, npc.HasEquipment(hot))
		r.UpdateFor(4 * time.Second)

		// Starts forging
		must.False(t, b.heatingHotSword)
		must.True(t, b.forging)
		must.EqOp(t, 0, b.heatCount)
		must.EqOp(t, 1, b.strikeCount)
		must.False(t, b.swordCooledDown)
		must.True(t, npc.HasEquipment(hm))
		r.UpdateFor(5 * time.Second)

		// Sword cooled down
		must.False(t, b.forging)
		must.True(t, b.heatingWarmSword)
		must.EqOp(t, 5, b.strikeCount)
		must.EqOp(t, 6, b.heatCount)
		must.False(t, b.swordCooledDown)
		must.True(t, npc.HasEquipment(warm))
		r.UpdateFor(4 * time.Second)

		// Sword is getting hot again
		must.False(t, b.heatingWarmSword)
		must.True(t, b.heatingHotSword)
		must.EqOp(t, 11, b.heatCount)
		must.True(t, npc.HasEquipment(hot))
		r.UpdateFor(4 * time.Second)

		// Continue forging
		must.False(t, b.heatingHotSword)
		must.True(t, b.forging)
		must.EqOp(t, 0, b.heatCount)
		must.EqOp(t, 6, b.strikeCount)
		must.False(t, b.swordCooledDown)
		must.True(t, npc.HasEquipment(hm))
		r.UpdateFor(4 * time.Second)

		// Starts quenching (walk to water barrel)
		const walkFrames = 11
		if hasWater {
			must.False(t, b.forging)
			must.True(t, b.quenching)
			must.EqOp(t, 0, b.heatCount)
			must.EqOp(t, 0, b.strikeCount)
			must.False(t, b.swordCooledDown)
			must.True(t, npc.HasEquipment(warm))
			r.UpdateN(walkFrames - 1)
			must.True(t, b.quenching)
			r.Update()
			// Reached the barrel
			must.False(t, b.quenching)
			must.False(t, npc.HasEquipment(warm))
			r.UpdateFor(2 * time.Second)
		}

		// Ending the sequence
		must.False(t, b.quenching)
		must.True(t, b.idle)
		if hasWater {
			must.True(t, npc.HasEquipment(cold))
		}
		r.UpdateFor(60 * time.Second)

		// Start again
		must.False(t, b.idle)
		must.True(t, b.heatingColdSword)
		must.EqOp(t, 1, b.heatCount)
		must.True(t, npc.HasEquipment(cold))
	}
}
