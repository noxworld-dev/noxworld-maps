package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	barrelAudioWP ns.WaypointID
)

func initBarrels() {
	barrelAudioWP = ns.Waypoint("BarrelAudioOrigin")
}

func BarrelCaveIn() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	g := ns.ObjectGroup("BarrelExplosionTriggers")
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.WallDestroyed, barrelAudioWP)
	ns.AudioEvent(ns.FireballExplode, barrelAudioWP)
	ns.ObjectGroupOff(g)
}

func BarrelsExplode() {
	u := ns.Object("PowderBarrel")
	ns.Damage(u, 0, 30, 1)
}

func EndExplodingBarrelSetPiece() {
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
}

var barrelsStarted = false

func BeginExplodingBarrelSetPiece() {
	if barrelsStarted || !ns.IsCaller(ns.GetHost()) {
		return
	}
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
	ns.AudioEvent(ns.EarthRumbleMajor, barrelAudioWP)
	barrelsStarted = true
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
}
