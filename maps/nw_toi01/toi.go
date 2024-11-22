package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

func Toi01ToGalava02() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "Toi01ToGalava02"
		ns.LoadMap("nw_galava02", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}
