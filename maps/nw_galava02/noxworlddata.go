package noxworld

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
)

type NoxWorldData struct {
	Character struct {
		// General
		Registered    bool
		Name          string
		Health        int
		Mana          int
		Level         int
		Class         player.Class
		LastExitUsed  string
		WolfCompanion int
	}
}

func loadMyNoxWorldData(pl ns.Player) NoxWorldData {
	var data NoxWorldData
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Get("accountdata", &data)
	if err != nil {
		fmt.Println("cannot read data:", err)
	}
	return data
}

func saveMyNoxWorldData(pl ns.Player, data NoxWorldData) {
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Set("accountdata", &data)
	if err != nil {
		fmt.Println("cannot save data:", err)
	}
}

func updateNoxWorldData(pl ns.Player, fnc func(data *NoxWorldData)) {
	data := loadMyNoxWorldData(pl)
	fnc(&data)
	saveMyNoxWorldData(pl, data)
}
