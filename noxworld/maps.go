package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

const (
	MapIx         = "nw_ix"
	MapIxCemetery = "nw_ixcem"
	MapIxTemple1  = "nw_temple01"
	MapCrossroads = "nw_xroad"
)

type GoToMapOptions struct {
	Exit string
}

func GoToMap(p ns.Player, mapname string, opts GoToMapOptions) {
	UpdatePlayer(p, func(data *PlayerData) {
		data.Character.LastExitUsed = opts.Exit
		ns.LoadMap(mapname, &ns.LoadMapOptions{HideTitleScreen: true})
	})
}
