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
		must.EqOp(t, BlacksmithHeatingColdSword, b.state)
		must.EqOp(t, 1, b.heatCount)
		must.True(t, npc.HasEquipment(cold))

		// Heating it progress
		r.UpdateFor(time.Second)
		must.EqOp(t, BlacksmithHeatingColdSword, b.state)
		must.EqOp(t, 2, b.heatCount)
		must.True(t, npc.HasEquipment(cold))
		r.UpdateFor(3 * time.Second)

		// Sword is getting warm
		must.EqOp(t, BlacksmithHeatingWarmSword, b.state)
		must.EqOp(t, 6, b.heatCount)
		must.True(t, npc.HasEquipment(warm))
		r.UpdateFor(4 * time.Second)

		// Sword is getting hot
		must.EqOp(t, BlacksmithHeatingHotSword, b.state)
		must.EqOp(t, 11, b.heatCount)
		must.True(t, npc.HasEquipment(hot))
		r.UpdateFor(4 * time.Second)

		// Starts forging
		must.EqOp(t, BlacksmithForging, b.state)
		must.EqOp(t, 0, b.heatCount)
		must.EqOp(t, 1, b.strikeCount)
		must.True(t, npc.HasEquipment(hm))
		r.UpdateFor(5 * time.Second)

		// Sword cooled down
		must.EqOp(t, BlacksmithHeatingWarmSword, b.state)
		must.EqOp(t, 5, b.strikeCount)
		must.EqOp(t, 6, b.heatCount)
		must.True(t, npc.HasEquipment(warm))
		r.UpdateFor(4 * time.Second)

		// Sword is getting hot again
		must.EqOp(t, BlacksmithHeatingHotSword, b.state)
		must.EqOp(t, 11, b.heatCount)
		must.True(t, npc.HasEquipment(hot))
		r.UpdateFor(4 * time.Second)

		// Continue forging
		must.EqOp(t, BlacksmithForging, b.state)
		must.EqOp(t, 0, b.heatCount)
		must.EqOp(t, 6, b.strikeCount)
		must.True(t, npc.HasEquipment(hm))
		r.UpdateFor(4 * time.Second)

		// Starts quenching (walk to water barrel)
		const walkFrames = 11
		if hasWater {
			must.EqOp(t, BlacksmithQuenchLookForWater, b.state)
			must.EqOp(t, 0, b.heatCount)
			must.EqOp(t, 0, b.strikeCount)
			must.True(t, npc.HasEquipment(warm))
			r.UpdateN(walkFrames - 1)
			must.EqOp(t, BlacksmithQuenchLookForWater, b.state)
			r.Update()
			// Reached the barrel
			must.EqOp(t, BlacksmithQuenching, b.state)
			must.False(t, npc.HasEquipment(warm))
			r.UpdateFor(2 * time.Second)
		}

		// Ending the sequence
		must.EqOp(t, BlacksmithIdle, b.state)
		if hasWater {
			must.True(t, npc.HasEquipment(cold))
		}
		r.UpdateFor(60 * time.Second)

		// Start again
		must.EqOp(t, BlacksmithHeatingColdSword, b.state)
		must.EqOp(t, 1, b.heatCount)
		must.True(t, npc.HasEquipment(cold))
	}
}
