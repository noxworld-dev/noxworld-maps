package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	entranceLiftAudioWP1 ns.WaypointID
	entranceLiftAudioWP2 ns.WaypointID
	entranceLiftAndGears ns.ObjectGroupID

	exitLiftAudioWP1 ns.WaypointID
	exitLiftAudioWP2 ns.WaypointID
	exitLiftAudioWP3 ns.WaypointID
	exitLiftAndGears ns.ObjectGroupID

	secretElevGroup   ns.ObjectGroupID
	secretElevWalls   ns.WallGroupID
	secretElevAudioWP ns.WaypointID
)

func initElevators() {
	entranceLiftAudioWP1 = ns.Waypoint("EntranceLiftAudioOrigin1")
	entranceLiftAudioWP2 = ns.Waypoint("EntranceLiftAudioOrigin2")
	entranceLiftAndGears = ns.ObjectGroup("EntranceLiftAndGears")

	exitLiftAudioWP1 = ns.Waypoint("ExitLiftAudioOrigin1")
	exitLiftAudioWP2 = ns.Waypoint("ExitLiftAudioOrigin2")
	exitLiftAudioWP3 = ns.Waypoint("ExitLiftAudioOrigin3")
	exitLiftAndGears = ns.ObjectGroup("ExitLiftAndGears")

	secretElevGroup = ns.ObjectGroup("SecretElevatorGroup")
	secretElevWalls = ns.WallGroup("SecretElevatorWalls")
	secretElevAudioWP = ns.Waypoint("SecretElevatorAudioOrigin")
}

func PlayEntranceLiftSounds() {
	ns.AudioEvent(ns.ChangeSpellbar, entranceLiftAudioWP1)
	ns.AudioEvent(ns.ChangeSpellbar, entranceLiftAudioWP2)
}

func PlayExitLiftSounds() {
	ns.AudioEvent(ns.ChangeSpellbar, exitLiftAudioWP1)
	ns.AudioEvent(ns.ChangeSpellbar, exitLiftAudioWP2)
	ns.AudioEvent(ns.ChangeSpellbar, exitLiftAudioWP3)
}

func ActivateEntranceElevator() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOff(ns.Object("EntranceSwitch"))
	ns.AudioEvent(ns.CreatureCageAppears, entranceLiftAudioWP1)
	ns.AudioEvent(ns.ChangeSpellbar, entranceLiftAudioWP1)
	ns.ObjectGroupOn(entranceLiftAndGears)
}

func ActivateExitElevator() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectOff(ns.Object("ExitSwitch"))
	ns.AudioEvent(ns.CreatureCageAppears, exitLiftAudioWP1)
	ns.AudioEvent(ns.ChangeSpellbar, exitLiftAudioWP1)
	ns.ObjectGroupOn(exitLiftAndGears)
	MineShake4()
}

func ActivateSecretElevator() {
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.CreatureCageAppears, secretElevAudioWP)
	ns.AudioEvent(ns.ChangeSpellbar, secretElevAudioWP)
	ns.WallGroupOpen(secretElevWalls)
	ns.ObjectGroupOn(secretElevGroup)
}

func DisableEntranceLift() {
	ns.ObjectGroupOff(entranceLiftAndGears)
}
