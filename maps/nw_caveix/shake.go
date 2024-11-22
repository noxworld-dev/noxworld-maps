package con01a

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

const (
	rocksZMin      = 75
	rocksZMax      = 150
	rocksSpreadMin = 30
	rocksSpreadMax = 70
)

var (
	rocks                 [50]ns.ObjectID
	rockInd               int
	shake                 *Shaker
	entranceShakeTriggers ns.ObjectGroupID
	entranceExitTriggers  ns.ObjectGroupID
)

func initShaker() {
	shake = NewShaker()
	entranceShakeTriggers = ns.ObjectGroup("EntranceShakeTriggers")
	entranceExitTriggers = ns.ObjectGroup("EntranceExitTriggers")
}

func NewShaker() *Shaker {
	return &Shaker{
		NumMax:      5,
		NumMin:      2,
		RepeatMax:   600,
		Var55:       5,
		JiggleForce: 7.0,
		AudioWP:     ns.Waypoint("QuakeAudioOrigin"),
		CreationWP:  ns.Waypoint("CreationWP"),
		Enabled:     true,
	}
}

type Shaker struct {
	AudioWP     ns.WaypointID
	CreationWP  ns.WaypointID
	NumMax      int
	NumMin      int
	RepeatMax   int
	Var55       int
	JiggleForce float32
	Enabled     bool
	Debug       bool
}

func (s *Shaker) Start() {
	shake.Enabled = true
	shake.RepeatMax = 600
	Quake()
}

func (s *Shaker) Quake() {
	if s.Debug {
		ns.PrintToAll("Quake initiated")
	}
	if !s.Enabled {
		return
	}
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	n := ns.Random(s.NumMin, s.NumMax)
	ns.MoveWaypoint(s.AudioWP, x, y)
	ns.AudioEvent(ns.EarthRumbleMinor, s.AudioWP)
	ns.AudioEvent(ns.BoulderMove, s.AudioWP)
	ns.Effect(ns.JIGGLE, x, y, s.JiggleForce, 0)

	for i := 0; i <= n; i++ {
		if s.Debug {
			ns.PrintToAll("''For'' loop")
		}
		dx := ns.RandomFloat(rocksSpreadMin, rocksSpreadMax)
		dy := ns.RandomFloat(rocksSpreadMin, rocksSpreadMax)
		dir := ns.Random(0, 3)
		if s.Debug {
			ns.PrintToAll("Attempting switch")
		}
		switch dir {
		case 0:
			ns.MoveWaypoint(s.CreationWP, x+dx, y+dy)
		case 1:
			ns.MoveWaypoint(s.CreationWP, x+dx, y-dy)
		case 2:
			ns.MoveWaypoint(s.CreationWP, x-dx, y+dy)
		case 3:
			ns.MoveWaypoint(s.CreationWP, x-dx, y-dy)
		}
		ns.Delete(rocks[rockInd])
		rock := ns.CreateObject("CaveRocksTiny", s.CreationWP)
		rocks[rockInd] = rock
		rockInd++
		if rockInd >= len(rocks) {
			rockInd = 0
		}
		z := ns.RandomFloat(rocksZMin, rocksZMax)
		ns.Raise(rock, z)
	}
	dt := ns.Random(5, s.RepeatMax)
	ns.FrameTimer(dt, s.Quake)
}

func DebugOn() {
	shake.Debug = true
}

func DebugOff() {
	shake.Debug = false
}

func Quake() {
	shake.Quake()
}

func EntranceShake() {
	ns.ObjectGroupOff(entranceShakeTriggers)
	ns.ObjectGroupOn(entranceExitTriggers)
	shake.Start()
}

func StopShakes() {
	ns.ObjectGroupOff(entranceExitTriggers)
	shake.Enabled = false
	ns.ObjectGroupOn(entranceShakeTriggers)
}

func mineShakeSeg(i int) {
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	wp := ns.Waypoint(fmt.Sprintf("MineShake%dAudioOrigin", i))
	ns.Effect(ns.JIGGLE, x, y, 5, 0)
	ns.AudioEvent(ns.EarthRumbleMinor, wp)
}

func MineShake1SEG2() {
	mineShakeSeg(1)
	shake.RepeatMax = 500
	shake.Var55 = 200
}

func MineShake2SEG2() {
	mineShakeSeg(2)
	shake.JiggleForce = 8
	shake.RepeatMax = 120
}

func MineShake3SEG2() {
	mineShakeSeg(3)
}

func MineShake4SEG2() {
	mineShakeSeg(4)
	shake.JiggleForce = 8
	shake.RepeatMax = 120
}

func MineShake5SEG2() {
	mineShakeSeg(5)
	shake.JiggleForce = 9
	shake.NumMax = 2
	shake.NumMin = 1
	shake.RepeatMax = 20
}

func mineShake(i int) {
	wp := ns.Waypoint(fmt.Sprintf("MineShake%dAudioOrigin", 1))
	g := ns.ObjectGroup(fmt.Sprintf("MineShake%dTriggers", 1))
	ns.AudioEvent(ns.EarthRumbleMajor, wp)
	x := ns.GetObjectX(ns.GetHost())
	y := ns.GetObjectY(ns.GetHost())
	ns.ObjectGroupOff(g)
	ns.Effect(ns.JIGGLE, x, y, 5, 0)
}

var mineDidShake1 = false

func MineShake1() {
	if mineDidShake1 || !ns.IsCaller(ns.GetHost()) {
		return
	}
	mineDidShake1 = true
	mineShake(1)
}

var mineDidShake2 = false

func MineShake2() {
	if mineDidShake2 || !ns.IsCaller(ns.GetHost()) {
		return
	}
	mineDidShake2 = true
	mineShake(2)
}

var mineDidShake3 = false

func MineShake3() {
	if mineDidShake3 || !ns.IsCaller(ns.GetHost()) {
		return
	}
	mineDidShake3 = true
	mineShake(3)
}

func MineShake4() {
	mineShake(4)
}

func MineShake5() {
	mineShake(5)
}
