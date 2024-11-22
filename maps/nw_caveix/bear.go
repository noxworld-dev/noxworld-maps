package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var caveBoulder ns.ObjectID

func getBear() ns.ObjectID {
	return ns.Object("BigSpider")
}

func KillBear() {
	ns.AudioEvent(ns.BearHurt, tunnelAudioWP)
	ns.Damage(getBear(), 0, 200, 0)
}

func DropCrushingRock() {
	caveBoulder = ns.CreateObject("CaveBoulders", ns.Waypoint("SpiderWP"))
	ns.Raise(caveBoulder, 100)
}

func BearDie() {
	ns.AudioEvent(ns.FistHit, tunnelAudioWP)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}

func SquashSpider() {
	x := ns.GetObjectX(caveBoulder)
	y := ns.GetObjectY(caveBoulder)
	ns.MoveObject(caveBoulder, x, y+15)
	ns.WallGroupOpen(ns.WallGroup("BearCageWalls"))
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.NoWallSound(false)
}

func ReleaseStun() {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	_ = ns.Waypoint("BoulderWP2") // unused
	_ = ns.Waypoint("BoulderWP3") // unused
	ns.AudioEvent(ns.FistHit, tunnelAudioWP)
	ns.Effect(ns.JIGGLE, x, y, 10, 0)
}

func GoBigSpider() {
	u := getBear()
	wp := ns.Waypoint("SpiderWP")
	ns.AudioEvent(ns.BearRecognize, tunnelAudioWP)
	ns.SetCallback(u, 11, func() {
		ns.SetCallback(u, 11, func() {})
		ns.LookAtObject(u, ns.Object("BearFaceLocation"))
		ns.HitLocation(u, ns.GetObjectX(u)-1, ns.GetObjectY(u)+1)
		ns.AudioEvent(ns.BearAttackInit, tunnelAudioWP)
		ns.WallGroupClose(ns.WallGroup("BearCageWalls"))
	})
	ns.Move(u, wp)
}
