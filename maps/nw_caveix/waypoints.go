package con01a

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	tunnelPitWPs     [16]ns.WaypointID
	exitPitWPs       [6]ns.WaypointID
	entranceShakeWPs [4]ns.WaypointID
	mineShake1WPs    [8]ns.WaypointID
	mineShake2WPs    [4]ns.WaypointID
	mineShake3WPs    [4]ns.WaypointID
	mineShake4WPs    [4]ns.WaypointID
	mineShake5WPs    [4]ns.WaypointID
	cavePitWPs       [4]ns.WaypointID
	flameWPs         [20]ns.WaypointID
)

func initWaypoints() {
	for i := range tunnelPitWPs {
		tunnelPitWPs[i] = ns.Waypoint(fmt.Sprintf("TunnelPitWP%d", i))
	}
	for i := range exitPitWPs {
		exitPitWPs[i] = ns.Waypoint(fmt.Sprintf("ExitPitWP%d", i))
	}
	for i := range entranceShakeWPs {
		entranceShakeWPs[i] = ns.Waypoint(fmt.Sprintf("EntranceShakePit%d", i))
	}
	for i := range mineShake1WPs {
		mineShake1WPs[i] = ns.Waypoint(fmt.Sprintf("MineShake1Pit%d", i+1))
	}
	for i := range mineShake2WPs {
		mineShake2WPs[i] = ns.Waypoint(fmt.Sprintf("MineShake2Pit%d", i+1))
	}
	for i := range mineShake3WPs {
		mineShake3WPs[i] = ns.Waypoint(fmt.Sprintf("MineShake3Pit%d", i+1))
	}
	for i := range mineShake4WPs {
		mineShake4WPs[i] = ns.Waypoint(fmt.Sprintf("MineShake4Pit%d", i+1))
	}
	for i := range mineShake5WPs {
		mineShake5WPs[i] = ns.Waypoint(fmt.Sprintf("MineShake5Pit%d", i+1))
	}
	for i := range cavePitWPs {
		cavePitWPs[i] = ns.Waypoint(fmt.Sprintf("CavePit%d", i+1))
	}
	for i := range flameWPs {
		flameWPs[i] = ns.Waypoint(fmt.Sprintf("FlameWP%d", i+1))
	}
}
