package con01a

import (
	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	theStaff               ns.ObjectID
	tunnelEntranceTriggers ns.ObjectGroupID
	tunnelTriggers         ns.ObjectGroupID
	npcsGroup              ns.ObjectGroupID
	tunnelAudioWP          ns.WaypointID
)

func MovePlayerToExit() {
	exit := ns.Waypoint("Exit")
	x := ns.GetWaypointX(exit)
	y := ns.GetWaypointY(exit)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), x, y)
}

func NoMonsters() {
	if !ns.IsAttackedBy(ns.GetCaller(), ns.GetHost()) {
		return
	}
	ns.GoBackHome(ns.GetCaller())
}

func TunnelCollapse() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.GolemMove, tunnelAudioWP)
}

func FallingRocksSEG1() {
	_ = ns.GetObjectX(ns.GetHost())
	_ = ns.GetObjectY(ns.GetHost())
	ns.AudioEvent(ns.EarthRumbleMajor, tunnelAudioWP)
}

func SetpieceShake1() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.EarthRumbleMinor, tunnelAudioWP)
}

func FallingRocksSEG5() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	treasure := ns.Object("Treasure")
	walls := ns.WallGroup("TunnelWalls")
	dwalls1 := ns.WallGroup("TunnelDestructableWalls2")
	dwalls2 := ns.WallGroup("TunnelDestructableWalls3")
	_ = ns.Waypoint("BoulderWP1") // unused
	_ = ns.Waypoint("BoulderWP4") // unused
	_ = ns.Waypoint("BoulderWP5") // unused
	ns.WallGroupOpen(dwalls1)
	ns.WallGroupBreak(dwalls2)
	ns.NoWallSound(false)
	ns.AudioEvent(ns.EarthRumbleMajor, tunnelAudioWP)
	ns.NoWallSound(true)
	ns.WallGroupClose(walls)
	ns.Delete(treasure)
}

func FallingRocksSEG4() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	g := ns.WallGroup("TunnelDestructableWalls1")
	ns.Effect(ns.JIGGLE, x, y, 15, 0)
	ns.NoWallSound(true)
	ns.WallGroupOpen(g)
	ns.AudioEvent(ns.FistHit, tunnelAudioWP)
}

func FallingRocksSEG3() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.EarthRumbleMajor, tunnelAudioWP)
}

func FallingRocksSEG2() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 15, 0)
	ns.AudioEvent(ns.GolemMove, tunnelAudioWP)
}

func LittleSpiderRecognize() {
	u := ns.Object("LittleSpider")
	ns.CreatureIdle(u)
	ns.AggressionLevel(u, 0.83)
	ns.SetCallback(u, 3, func() {})
}
