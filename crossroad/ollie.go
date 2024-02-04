package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

var ollie ns.Obj
var ollieSpawn ns.Pointf

var followPlayer bool

func initOllie() {
	if ns.Object("Ollie") != nil {
		ollie = ns.Object("Ollie")
	} else {
		// Fix spawn
		ollie = ns.CreateObject("Wolf", ns.GetHost())
		ollie.SetTeam(ns.Teams()[1])
	}
	ollieSpawn = ollie.Pos()
	place()
}

func heel() {
}

func place() {
	ollie.Guard(ollieSpawn, ollieSpawn, 300)
	ns.NewTimer(ns.Seconds(20), func() {
		explore()
	})
}

func explore() {
	ollie.Wander()
	ns.NewTimer(ns.Seconds(5), func() {
		beg()
	})
}

func beg() {
	target := ns.FindClosestObject(ollie, ns.HasTypeName{"NPC", "NewPlayer"})
	ollie.Follow(target)
	if ns.Random(1, 4) == 1 {
		ollie.Chat("War05A.scr:HoundGreeting")
	}
	ns.NewTimer(ns.Seconds(8), func() {
		place()
	})
}
