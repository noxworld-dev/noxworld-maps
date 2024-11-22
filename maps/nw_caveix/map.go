package con01a

import (
	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	playerStart       ns.ObjectID
	mapSwitchTriggers ns.ObjectGroupID
)

func MapInitialize() {
	playerStart = ns.Object("PlayerStart")
	mapSwitchTriggers = ns.ObjectGroup("MapSwitchTriggers")

	tunnelEntranceTriggers = ns.ObjectGroup("TunnelEntranceTriggers")
	tunnelTriggers = ns.ObjectGroup("TunnelTriggers")
	npcsGroup = ns.ObjectGroup("NPCs")
	theStaff = ns.Object("Staff")
	tunnelAudioWP = ns.Waypoint("TunnelAudioOrigin")

	initExit()
	initSecrets()
	initElevators()
	initBarrels()
	initWaypoints()
	initShaker()
	initDialogs()
	ns.StartupScreen(1)
}

func MapEntry() {
	ns.ObjectGroupOn(mapSwitchTriggers)
}

var mapDidSwitch = false

func MapSwitch() {
	if mapDidSwitch || !ns.IsCaller(ns.GetHost()) {
		return
	}
	ns.Frozen(ns.GetHost(), true)
	mapDidSwitch = true
	ns.ObjectGroupOff(mapSwitchTriggers)
	ns.Blind()
}

func EnableTeleporter() {
	ns.ObjectOn(playerStart)
}

func PlayerDeath() {
	ns.DeathScreen(1)
}
