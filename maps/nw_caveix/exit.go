package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	exitAudioWP ns.WaypointID
)

func initExit() {
	exitAudioWP = ns.Waypoint("ExitAudioOrigin")
}

func AwardExitExperience() {
	ns.MoveWaypoint(exitAudioWP, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagDrop, exitAudioWP)
	ns.GiveXp(ns.GetHost(), 200)
}

func CaveExitCollapseSEG2() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	_ = ns.Waypoint("ExitBoulderWP") // unused
	ns.NoWallSound(false)
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.FistHit, exitAudioWP)
}

var caveExitCollapsed = false

func CaveExitCollapse() {
	wg := ns.WallGroup("CaveExitWalls")
	trig := ns.ObjectGroup("CaveExitTriggers")
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.ObjectGroupOff(trig)
	caveExitCollapsed = true
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.EarthRumbleMajor, exitAudioWP)
	ns.NoWallSound(true)
	ns.WallGroupClose(wg)
	shake.Enabled = false
}

func EnterExitPolygon() {
	if caveExitCollapsed || !ns.IsCaller(ns.GetHost()) {
		return
	}
	CaveExitCollapse()
	PlayOutdoorMusic()
}

var collapseBegan = false

func BeginCollapseSetPiece() {
	if collapseBegan || !ns.IsCaller(ns.GetHost()) {
		return
	}
	PlayUndergroundMusic()
	collapseBegan = true
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
	ns.AudioEvent(ns.EarthRumbleMajor, tunnelAudioWP)
}

func PlayOutdoorMusic() {
	ns.Music(21, 100)
}

func PlayUndergroundMusic() {
	ns.Music(18, 100)
}
