package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
)

var exitUsed bool

func init() {
	for i := 0; i < len(ns.Players()); i++ {
		ns.Players()[i].Unit().Enchant(enchant.BLINDED, ns.Frames(1))
	}
	ns.Runtime().OnPlayerJoin(playerJoin)
	ns.OnChat(onCommand)
	ns.NewTimer(ns.Frames(1), func() {
		moveToEntry()
		ns.NewTimer(ns.Frames(1), func() {
			checkIfRegistered(ns.GetHost().Player())
		})
	})
}

func OnFrame() {
	checkExits()
}

func onCommand(t ns.Team, p ns.Player, obj ns.Obj, msg string) string {
	if p != nil && p == ns.GetHost().Player() {
		switch msg {
		case "-test":
			ns.LoadMap("nw_xroad", &ns.LoadMapOptions{
				HideTitleScreen: true,
			})
		}
	}
	return msg
}

func checkIfRegistered(p ns.Player) {
	data := loadMyNoxWorldData(p)
	if !data.Character.Registered {
		updateNoxWorldData(p, func(data *NoxWorldData) {
			data.Character.Registered = true
			registerCharacter(p)
			p.Unit().Player().PrintStr("Welcome to the Open World Nox Quest server!")
			p.Unit().Player().PrintStr("Your player character has been registered.")
		})
	} else {
		p.Unit().Player().PrintStr("Welcome back to the Open World Nox Quest server!")
		registerCharacter(p)
	}
}

func registerCharacter(p ns.Player) {
	checkClass(p)
	checkCompanions(p)
	//resetCharacter(p)
}

func checkCompanions(p ns.Player) {
	data := loadMyNoxWorldData(p)
	if data.Character.WolfCompanion > 2 {
		updateNoxWorldData(p, func(data *NoxWorldData) {
			data.Character.WolfCompanion = 2
		})
	} else {
		if data.Character.WolfCompanion == 1 {
			wolf := ns.CreateObject("Wolf", p.Unit().Pos())
			wolf.SetOwner(p.Unit())
			wolf.Follow(p.Unit())
			wolf.OnEvent(ns.EventDeath, func() {
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Character.WolfCompanion--
				})
			})
		}
		if data.Character.WolfCompanion == 2 {
			wolf1 := ns.CreateObject("Wolf", p.Unit().Pos())
			wolf1.SetOwner(p.Unit())
			wolf1.Follow(p.Unit())
			wolf2 := ns.CreateObject("Wolf", p.Unit().Pos())
			wolf2.SetOwner(p.Unit())
			wolf2.Follow(p.Unit())
			wolf1.OnEvent(ns.EventDeath, func() {
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Character.WolfCompanion--
				})
			})
			wolf2.OnEvent(ns.EventDeath, func() {
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Character.WolfCompanion--
				})
			})
		}
	}
}

func resetCharacter(p ns.Player) {
	//wizardLevel1(p)
}

func checkClass(p ns.Player) {
	updateNoxWorldData(p, func(data *NoxWorldData) {
		data.Character.Class = p.Unit().GetClass()
	})
}

func playerJoin(p ns.Player) bool {
	checkIfRegistered(p)
	return true
}
