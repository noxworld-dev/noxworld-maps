package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

func Galava01ToGalava02() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "Galava01ToGalava02"
		ns.LoadMap("nw_galava02", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}

func Galava01ToCrossRoads() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "Galava01ToCrossRoads"
		ns.LoadMap("nw_xroad", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}
