package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var ollie ns.Obj
var ollieSpawn ns.Pointf

var followPlayer bool

func init() {
	OnLateInit(func() {
		if ns.Object("Ollie") != nil {
			ollie = ns.Object("Ollie")
		} else {
			// Fix spawn
			ollie = ns.CreateObject("Wolf", ns.GetHost())
			ollie.SetTeam(ns.Teams()[1])
		}
		ollieSpawn = ollie.Pos()
		place()
		ns.OnChat(commandOllie)
	})
}

func commandOllie(t ns.Team, p ns.Player, obj ns.Obj, msg string) string {
	if p != nil {
		switch msg {
		case "come", "Come", "Come!", "come!", "come.", "Come.", "Here!", "here.", "here", "Here", "here!", "Ollie", "ollie", "Ollie!", "ollie!", "Here boy!", "Here girl!", "here boy!", "here girl!", "here boy", "here girl", "Here boy", "Here girl":
			if !followPlayer {
				ns.AudioEvent(audio.WolfIdle, ollie)
				followPlayer = true
				ollie.Follow(p.Unit())
				ollie.Chat("War05A.scr:HoundGreeting")
				ns.NewTimer(ns.Seconds(8), func() {
					followPlayer = false
					place()
				})
			}
		}
	}
	return msg
}

func place() {
	if !followPlayer {
		ns.AudioEvent(audio.WolfIdle, ollie)
		ollie.Guard(ollieSpawn, ollieSpawn, 300)
		ns.NewTimer(ns.Seconds(20), func() {
			if !followPlayer {
				explore()
			}
		})
	}
}

func explore() {
	if !followPlayer {
		ns.AudioEvent(audio.WolfRecognize, ollie)
		ollie.Wander()
		ns.NewTimer(ns.Seconds(5), func() {
			if !followPlayer {
				beg()
			}
		})
	}
}

func beg() {
	if !followPlayer {
		target := ns.FindClosestObject(ollie, ns.HasTypeName{"NPC", "NewPlayer"})
		ns.AudioEvent(audio.WolfIdle, ollie)
		ollie.Follow(target)
		if ns.Random(1, 4) == 1 {
			ollie.Chat("War05A.scr:HoundGreeting")
		}
		ns.NewTimer(ns.Seconds(8), func() {
			if !followPlayer {
				place()
			}
		})
	}
}
