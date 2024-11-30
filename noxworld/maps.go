package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

const (
	MapConStartCaveIx    = "nw_caveix"
	MapIx                = "nw_ix"
	MapIxCemetery        = "nw_ixcem"
	MapTempleOfIx01      = "nw_tox01"
	MapTempleOfIx02      = "nw_tox02"
	MapCrossroads        = "nw_xroad"
	MapDunMir01          = "nw_dnmr01"
	MapBeachHouse        = "nw_glv01"
	MapGalava            = "nw_glv02"
	MapTowerOfIllusion01 = "nw_toi01"
	MapTowerOfIllusion02 = "nw_toi02"
	MapManaMines         = "nw_manam"
	MapG_Castle          = "nw_castl"
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
