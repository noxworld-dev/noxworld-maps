package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var secretAudioWP ns.WaypointID

func initSecrets() {
	secretAudioWP = ns.Waypoint("SecretAudioWP")
}

func secretEffect() {
	ns.MoveWaypoint(secretAudioWP, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, secretAudioWP)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}

func Secret01Found() {
	ns.ObjectGroupOff(ns.ObjectGroup("Secret01Triggers"))
	secretEffect()
}

func Secret02Found() {
	ns.ObjectGroupOff(ns.ObjectGroup("Secret02Triggers"))
	secretEffect()
}

func SecretElevatorFound() {
	ns.ObjectOff(ns.GetTrigger())
	secretEffect()
}

func Secret03Found() {
	ns.ObjectOff(ns.GetTrigger())
	secretEffect()
}
