package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

func Galava02ToGalava01() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "Galava02ToGalava01"
		ns.LoadMap("nw_galva01", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}

func Galava02ToToi01() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "Galava02ToToi01"
		ns.LoadMap("nw_toi01", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}
