package nw_ix

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
	"github.com/noxworld-dev/opennox-lib/object"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

var MayorsFear ns.Obj
var MayorsFearIsOwned bool
var NecroEventActive bool
var NecroEntrance bool
var ArcheryContestActive bool

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		ns.StoryPic(ns.Object("Henrick"), "HenrickPic")
		henrickInit()
		ns.StoryPic(ns.Object("Lydia"), "SnobbyGirlPic")
		lydiaInit()
		ns.StoryPic(ns.Object("Joyce"), "MaidenPic")
		joyceInit()
		ns.StoryPic(ns.Object("Jacob"), "WardenPic")
		jailerInit()
		ns.StoryPic(ns.Object("Tanya"), "MaidenPic3")
		tanyaInit()
		ns.StoryPic(ns.Object("Bryan"), "MalePic4")
		bryanInit()
		ns.StoryPic(ns.Object("Clyde"), "MalePic11")
		clydeInit()
		ns.StoryPic(ns.Object("Gretchen"), "MaidenPic4")
		gretchenInit()
		ns.StoryPic(ns.Object("Geoff"), "Townsman4Pic")
		GeoffInit()
		ns.StoryPic(ns.Object("Contest_Guard"), "Townsman2Pic")
		Contest_GuardInit()
		ns.StoryPic(ns.Object("Morgan"), "MorganPic")
		MorganInit()
		ns.StoryPic(ns.Object("Aldwyn"), "AldwynPic")
		AldwynInit()
		ns.StoryPic(ns.Object("Contest_Official"), "MalePic10")
		Contest_OfficialInit()
		ns.StoryPic(ns.Object("Seth"), "Townsman1Pic")
		SethInit()
		ns.StoryPic(ns.Object("Floyd"), "Townsman3Pic")
		FloydInit()
		ns.StoryPic(ns.Object("BridgeGuard"), "MalePic9")
		BridgeGuardInit()
		ns.StoryPic(ns.Object("F6Townswoman4"), "MaidenPic6")
		F6Townswoman4Init()
		ns.StoryPic(ns.Object("Mayor's_Guard"), "Warrior3Pic")
		Mayors_GuardInit()
		ns.StoryPic(ns.Object("Tommy"), "MalePic12")
		TommyInit()
		ns.StoryPic(ns.Object("Julie"), "MaidenPic2")
		JulieInit()
		ns.StoryPic(ns.Object("Mayor_Theogrin"), "TheogrinPic")
		Mayor_TheogrinInit()
		ns.StoryPic(ns.Object("Drunk"), "DrunkPic")
		DrunkInit()
		ns.StoryPic(ns.Object("Ed"), "IxGuard1Pic")
		EdInit()
		checkQuests()
		ns.StoryPic(ns.Object("Ryan"), "IxGuard2Pic")
		RyanInit()
		ns.StoryPic(ns.Object("Mystic"), "Mystic")
		MysticInit()
	})
}

func IxToCrossRoads() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapCrossroads, nw.GoToMapOptions{
		Exit: "IxToCrossRoads",
	})
}

func IxToIxCem() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapIxCemetery, nw.GoToMapOptions{
		Exit: "IxToIxCem",
	})
}

func IxToTemple01() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapIxTemple1, nw.GoToMapOptions{
		Exit: "IxToTemple01",
	})
}

// Mystic: War08a
func MysticInit() {
	//ns.SetDialog(ns.Object("Mystic"), ns.DialogNormal, MysticDialogueStart, MysticDialogueEnd)
}
func MysticDialogueStart() {
	ns.Object("Mystic").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	//data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0:
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "War08a:Mystic")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "War08a:Mystic2")
		}
	case 1:
		//Wiz08a:Mystic	Ahhhhhhhh... I see you carry the Halberd of Horrendous, young Wizard!
		ns.TellStory(audio.SwordsmanHurt, "Wiz08a:Mystic2")
	case 2:
		// Con02a:Mystic
		//Con08a:Mystic	So, now you have both the Halberd of Horrendous and our Heart of Nox. Well done, young Conjurer!
		//Con08a:Mystic2	Perhaps you could use a scroll or potion?
	}

}
func MysticDialogueEnd() {
}

// Julie: War08a
func JulieInit() {
	ns.SetDialog(ns.Object("Julie"), ns.DialogNormal, JulieDialogueStart, JulieDialogueEnd)

}
func JulieDialogueStart() {
	ns.Object("Julie").LookAtObject(ns.GetCaller())
	ns.Object("Julie").Idle()
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0:
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02A:Maiden6Talk03")
		case 10:
			rnd := ns.Random(1, 2)
			switch rnd {
			case 1:
				ns.TellStory(audio.SwordsmanHurt, "War03b:Maiden1")
			case 2:
				ns.TellStory(audio.SwordsmanHurt, "War06a:WomenSpeak2")
			}
		}
	case 1:
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02A:Maiden6Talk03")
		case 10:
			ns.TellStory(audio.SwordsmanHurt, "War06a:WomenSpeak2")
		}
	case 2:
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02A:Maiden6Talk03")
		case 10:
			ns.TellStory(audio.SwordsmanHurt, "War06a:WomenSpeak2")
		}
	}
}
func JulieDialogueEnd() {
	ns.Object("Julie").Wander()
}

// Ryan: War08a
func RyanInit() {
	ns.SetDialog(ns.Object("Ryan"), ns.DialogNormal, RyanDialogueStart, RyanDialogueEnd)
}
func RyanDialogueStart() {
	ns.Object("Ryan").LookAtObject(ns.GetCaller())
	ns.TellStory("SwordsmanHurt", "Wiz08a:Guard01Greet")
}
func RyanDialogueEnd() {
}

// Ed: War08a
func EdInit() {
	ns.SetDialog(ns.Object("Ed"), ns.DialogNormal, EdDialogueStart, EdDialogueEnd)
}
func EdDialogueStart() {
	ns.Object("Ed").LookAtObject(ns.GetCaller())

	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0:
		ns.TellStory("SwordsmanHurt", "War08a:Guard02Greet")
	case 1:
		ns.TellStory("SwordsmanHurt", "War08a:Guard02Greet")
	case 2:
		switch data.Quest.ArcheryContest_Quests01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02a:Gatekeeper2Greet")
		case 10:
			ns.TellStory("SwordsmanHurt", "War08a:Guard02Greet")
		}

	}

}
func EdDialogueEnd() {

}

// Drunk: war08a
func DrunkInit() {
	ns.SetDialog(ns.Object("Drunk"), ns.DialogNormal, DrunkDialogueStart, DrunkDialogueEnd)
}
func DrunkDialogueStart() {
	ns.Object("Drunk").LookAtObject(ns.GetCaller())
	ns.TellStory("SwordsmanHurt", "War05A.scr:DrunkGreeting")
}
func DrunkDialogueEnd() {
}

func necroEvent() {
	if NecroEventActive {
		return
	} else {
		ns.NewTimer(ns.Seconds(120), func() {
			NecroEventActive = true
			Necro := ns.CreateObject("Necromancer", ns.Waypoint("NecSpawn"))
			ns.NewTimer(ns.Frames(5), func() {
				ns.AudioEvent(audio.InvisibilityOff, Necro)
			})
			Necro.SetMonsterStatus(object.MonStatusAlwaysRun)
			Necro.SetMonsterStatus(object.MonStatusHoldYourGround)
			Necro.Enchant(enchant.BLINDED, ns.Infinite())
			Necro.SetMaxHealth(1000)
			Necro.SetBaseSpeed(4)
			Necro.WalkTo(ns.Waypoint("NecWalkTo").Pos())
			NecroEntrance = true
			Necro.ChatStr("Greetings, citizens of Ix! I come bearing gifts for the future subjects of my Queen -- Hecubah!")
			ns.CastSpell(spell.SUMMON_SPIDER, Necro, Necro.Pos().Add(ns.Pointf{X: -50, Y: -60}))
			Necro.OnEvent(ns.EventDeath, func() {
				NecroEntrance = false
				NecroEventActive = false
			})
			ns.NewTimer(ns.Seconds(7), func() {
				ns.Object("MayorsBedroomDoor").Lock(true)
				spidey := ns.FindClosestObject(Necro, ns.HasTypeName{"Spider"})
				if spidey.HasOwner(Necro) {
					MayorsFear = spidey
					MayorsFear.OnEvent(ns.EventEnemyHeard, func() {
						ns.CastSpell(spell.CHARM, ns.Object("Aldwyn"), MayorsFear)
					})
					spdr := ns.CreateObject("SmallSpider", spidey)
					spdr.SetOwner(Necro)
					spdr = ns.CreateObject("SmallSpider", spidey)
					spdr.SetOwner(Necro)
					spdr = ns.CreateObject("SmallSpider", spidey)
					spdr.SetOwner(Necro)
					spidey.SetBaseSpeed(4)
					spidey.SetMaxHealth(1000)
					spidey.AggressionLevel(0.16)
					spidey.Move(ns.Waypoint("BigSpiderWP"))
					spidey.OnEvent(ns.EventDeath, func() {
						FearStays()
					})
					MayorsFear.OnEvent(ns.EventCollision, func() {
						if ns.Object("Aldwyn").CanSee(MayorsFear) {
							ns.CastSpell(spell.CHARM, ns.Object("Aldwyn"), MayorsFear)
						}
					})
					spidey.OnEvent(ns.EventEndOfWaypoint, func() {
						spidey.SetMaxHealth(100)
					})
				}
				Necro.Move(ns.Waypoint("Nexit"))
				Necro.OnEvent(ns.EventIsHit, func() {
					Necro.Move(ns.Waypoint("Nexit"))
					Necro.ChatStr("Get away from me! Filthy peasants!")
				})
				Necro.OnEvent(ns.EventEndOfWaypoint, func() {
					smoke := ns.CreateObject("Smoke", Necro)
					smoke.DeleteAfter(ns.Seconds(3))
					Necro.Delete()

				})
			})
		})
	}
}

func checkQuests() {
	ns.Object("MayorsBedroomDoor").Lock(true)
	MayorsFear = ns.CreateObject("Spider", ns.Waypoint("BigSpiderWP"))
	MayorsFear.Guard(ns.Waypoint("BigSpiderWP").Pos(), ns.Waypoint("BigSpiderWP").Pos(), 150)
	MayorsFear.OnEvent(ns.EventDeath, func() {
		ply := ns.Players()
		for _, i := range ply {
			if MayorsFear.HasOwner(i.Unit()) {
				MayorsFearIsOwned = true
				necroEvent()
			}
		}
		if MayorsFearIsOwned {
		} else {
			FearStays()
		}
	})
	MayorsFear.OnEvent(ns.EventCollision, func() {
		if ns.Object("Aldwyn").CanSee(MayorsFear) {
			ns.CastSpell(spell.CHARM, ns.Object("Aldwyn"), MayorsFear)
		}
	})
}

func FearStays() {
	ns.Object("Mayor_Theogrin").ChatStr("Stop making it worse and get Aldwyn to banish those bloody arachnids!")
	HuntingSpider01 := ns.CreateObject("SmallSpider", ns.Waypoint("BigSpiderWP"))
	HuntingSpider01.Hunt()
	HuntingSpider02 := ns.CreateObject("SmallSpider", ns.Waypoint("BigSpiderWP"))
	HuntingSpider02.Hunt()
	HuntingSpider03 := ns.CreateObject("SmallSpider", ns.Waypoint("BigSpiderWP"))
	HuntingSpider03.Hunt()
	MayorsFear = ns.CreateObject("Spider", ns.Waypoint("BigSpiderWP"))
	MayorsFear.Guard(ns.Waypoint("BigSpiderWP").Pos(), ns.Waypoint("BigSpiderWP").Pos(), 150)
	MayorsFear.OnEvent(ns.EventDeath, func() {
		ply := ns.Players()
		for _, i := range ply {
			if MayorsFear.HasOwner(i.Unit()) {
				MayorsFearIsOwned = true
				necroEvent()
			}
		}
		if MayorsFearIsOwned {
		} else {
			FearStays()
		}
	})
	MayorsFear.OnEvent(ns.EventCollision, func() {
		if ns.Object("Aldwyn").CanSee(MayorsFear) {
			ns.CastSpell(spell.CHARM, ns.Object("Aldwyn"), MayorsFear)
		}
	})
}

// Journal:War03aIxQuest	Go to Mayor Theogrin in the Village of Ix.
// Journal:War03bGetScepter	Retrieve the Mayor's scepter from the Urchins.

func Mayor_TheogrinInit() {
	ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
}
func Mayor_TheogrinDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0:
			ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
			ns.PrintStr("Find Aldwyn to help banish the spiders in the Mayor's home.")
			ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorGreeting")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 1
			})
		case 1:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Find Aldwyn to help banish the spiders in the Mayor's home.")
				ns.Object("Mayor_Theogrin").ChatStr("There's still a spider out there! I just know it!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		case 3:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				ns.Object("Mayor_Theogrin").ChatStr("I can hear it crawling! Remove those blasted spiders!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		case 4:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				ns.Object("Mayor_Theogrin").ChatStr("I can hear it crawling! Remove those blasted spiders!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		case 10:
			switch data.Quest.MayorsScepter_Quest01 {
			case 0:
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorIntro")
			case 1:
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorPre")
			case 9:
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorPost")
			case 10:
			}
		}
	case 1: // wiz
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0:
			ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
			ns.PrintStr("Find Aldwyn to help banish the spiders in the Mayor's home.")
			ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorGreeting")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 1
			})
		case 1:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Find Aldwyn to help banish the spiders in the Mayor's home.")
				ns.Object("Mayor_Theogrin").ChatStr("There's still a spider out there! I just know it!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		case 3:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				ns.Object("Mayor_Theogrin").ChatStr("Finally! A Conjurer has come to my aid!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		case 4:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				ns.Object("Mayor_Theogrin").ChatStr("Finally! A Conjurer has come to my aid!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 5
				})
			}
		}
	case 2: // con
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0:
			ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
			ns.PrintStr("Charm and banish the spiders in the Mayor's home.")
			ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorGreeting")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 1
			})
		case 1:
			if ns.FindAllObjects(ns.HasTypeName{"Spider"}) != nil {
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorProd")
				ns.PrintStr("Charm and banish the spiders in the Mayor's home.")
				ns.Object("Mayor_Theogrin").ChatStr("There's still a spider out there! I just know it!")
			} else {
				ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNext, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("You have completed an objective!")
				ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorFree")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.SpidersAtTheMayor_Quest01 = 2
				})
			}
		}
	}

}
func Mayor_TheogrinDialogueEnd() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 5:
			ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
			ns.Object("Mayor_Theogrin").ChatStr("Finally! Thank you Warrior for finding Aldwyn!")
			ns.GetCaller().ChangeGold(+200)
			ns.PrintStr("You gained 200 gold!")
			ns.AudioEvent(audio.TreasurePickup, ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 10
			})
			ns.Object("MayorsBedroomDoor").Lock(false)
			ns.NewTimer(ns.Seconds(2), func() {
				posal := ns.Object("Aldwyn").Pos()
				ns.Object("Aldwyn").SetPos(ns.Waypoint("HideAldwyn").Pos())
				ns.Object("Aldwyn").Idle()
				ImpMorph := ns.CreateObject("Imp", posal)
				ns.Effect(effect.SMOKE_BLAST, ImpMorph, ImpMorph)
				ns.Object("Aldwyn").DestroyChat()
				ns.AudioEvent(audio.ImpRecognize, ImpMorph)
				ImpMorph.ChatStr("I need to get back to my study. Good day, Warrior.")
				ImpMorph.SetOwner(ns.Object("Aldwyn"))
				ns.NewTimer(ns.Seconds(2), func() {
					ns.Effect(effect.BLUE_SPARKS, ImpMorph.Pos(), ImpMorph.Pos())
					ns.AudioEvent(audio.BlinkCast, ImpMorph.Pos())
					ImpMorph.Delete()
					ns.Object("Aldwyn").SetPos(ns.Waypoint("Study").Pos())
					ns.Effect(effect.SMOKE_BLAST, ns.Object("Aldwyn").Pos(), ns.Object("Aldwyn").Pos())
					ns.Object("Aldwyn").Guard(ns.Waypoint("Study").Pos(), ns.Waypoint("Study").Pos(), 150)

				})
			})

		case 10:
		}
	case 1: // wiz
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 5:
			ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
			ns.Object("Mayor_Theogrin").ChatStr("Finally! Thank you Wizard for finding Aldwyn!")
			ns.GetCaller().ChangeGold(+200)
			ns.PrintStr("You gained 200 gold!")
			ns.AudioEvent(audio.TreasurePickup, ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 10
			})
			ns.Object("MayorsBedroomDoor").Lock(false)
			ns.NewTimer(ns.Seconds(2), func() {
				posal := ns.Object("Aldwyn").Pos()
				ns.Object("Aldwyn").SetPos(ns.Waypoint("HideAldwyn").Pos())
				ns.Object("Aldwyn").Idle()
				ImpMorph := ns.CreateObject("Imp", posal)
				ns.Effect(effect.SMOKE_BLAST, ImpMorph, ImpMorph)
				ns.AudioEvent(audio.ImpRecognize, ImpMorph)
				ns.Object("Aldwyn").DestroyChat()
				ImpMorph.ChatStr("I need to get back to my study. Good day, Wizard.")
				ImpMorph.SetOwner(ns.Object("Aldwyn"))
				ns.NewTimer(ns.Seconds(2), func() {
					ns.Effect(effect.BLUE_SPARKS, ImpMorph.Pos(), ImpMorph.Pos())
					ns.AudioEvent(audio.BlinkCast, ImpMorph.Pos())
					ImpMorph.Delete()
					ns.Object("Aldwyn").SetPos(ns.Waypoint("Study").Pos())
					ns.Effect(effect.SMOKE_BLAST, ns.Object("Aldwyn").Pos(), ns.Object("Aldwyn").Pos())
					ns.Object("Aldwyn").Guard(ns.Waypoint("Study").Pos(), ns.Waypoint("Study").Pos(), 150)
				})
			})
		case 10:
		}
	case 2: // con
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 2:
			ns.Object("MayorsBedroomDoor").Lock(false)
			ns.SetDialog(ns.Object("Mayor_Theogrin"), ns.DialogNormal, Mayor_TheogrinDialogueStart, Mayor_TheogrinDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorsReward")
			ns.GetCaller().ChangeGold(+200)
			ns.PrintStr("You gained 200 gold!")
			ns.AudioEvent(audio.TreasurePickup, ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 10
			})
		case 10:
		}
	}
}

func TommyInit() {
	ns.SetDialog(ns.Object("Tommy"), ns.DialogNormal, TommyDialogueStart, TommyDialogueEnd)
}
func TommyDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0:
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk01")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk02")
		}
	case 1:
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk01")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk02")
		}
	case 2:
		if data.Quest.ArcheryContest_Quests01 == 10 {
			ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk03")
		} else {
			rnd := ns.Random(1, 2)
			switch rnd {
			case 1:
				ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk01")
			case 2:
				ns.TellStory(audio.SwordsmanHurt, "Con02A:Townsman9Talk02")
			}
		}
	}
}
func TommyDialogueEnd() {

}

func Mayors_GuardInit() {
	ns.SetDialog(ns.Object("Mayor's_Guard"), ns.DialogNormal, Mayors_GuardDialogueStart, Mayors_GuardDialogueEnd)
}
func Mayors_GuardDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch data.Quest.SpidersAtTheMayor_Quest01 {
	case 0:
		if data.Quest.TalkToTheMayor_Quest01 == 0 && data.Quest.SpidersAtTheMayor_Quest01 == 0 {
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			ns.PrintStr("Talk to the Mayor.")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.TalkToTheMayor_Quest01 = 10
			})
		}
		ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorsGuard")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		ns.TellStory(audio.SwordsmanHurt, "Con02a:MeetTheMayor")
	case 10:
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			switch p.Unit().GetClass() {
			case 0: // war
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorsGuardPost2")
			case 1: // wiz
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorsGuardIntro")
			case 2: // con
				ns.TellStory(audio.SwordsmanHurt, "War03b:MayorsGuardIntro")
			}
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "Con02a:MayorsGuardIdle")
		}
	}
}
func Mayors_GuardDialogueEnd() {

}

func BridgeGuardInit() {
	ns.SetDialog(ns.Object("BridgeGuard"), ns.DialogNormal, BridgeGuardDialogueStart, BridgeGuardDialogueEnd)
}
func BridgeGuardDialogueStart() {
	ns.TellStory(audio.SwordsmanHurt, "War03b:BridgeGuardPre")
	// Qesut mayor
	//War03b:BridgeGuardPost	You have returned! You should go see Mayor Theogrin -- his house is located in the center of town.
}
func BridgeGuardDialogueEnd() {}

func FloydInit() {
	ns.SetDialog(ns.Object("Floyd"), ns.DialogNormal, FloydDialogueStart, FloydDialogueEnd)
}
func FloydDialogueStart() {
	ns.Object("Floyd").ChatStr("The Urchins stole my best boots when I was bathing! If you bring 'em back, mate, I'll teach you a damn fine spell.")
	// War03b:T4Pre	Excuse me! Hello!? This is my home. Can I help you?
	// War03b:T4Post	Hmmm. I thought I locked the door.
	// if completed: Con02a:BridgeGuardReward	My boots! I never thought I'd see them again! I'm eternally grateful, mate. Here's the Spell, as promised.
}
func FloydDialogueEnd() {}

func F6Townswoman4Init() {
	ns.SetDialog(ns.Object("F6Townswoman4"), ns.DialogNormal, F6Townswoman4DialogueStart, F6Townswoman4DialogueEnd)
}
func F6Townswoman4DialogueStart() {
	ns.TellStory(audio.SwordsmanHurt, "War03b:Maiden2")
}
func F6Townswoman4DialogueEnd() {}

func SethInit() {
	ns.SetDialog(ns.Object("Seth"), ns.DialogNormal, SethDialogueStart, SethDialogueEnd)
}
func SethDialogueStart() {
	p := ns.GetCaller().Player()
	switch p.Unit().GetClass() {
	case 0:
		ns.TellStory(audio.SwordsmanHurt, "War03b:T1Pre")
	case 1:
		ns.TellStory(audio.SwordsmanHurt, "War03b:T1Post")
	case 2:
		ns.TellStory(audio.SwordsmanHurt, "War03b:T1Post")
	}
}
func SethDialogueEnd() {}

func AldwynInit() {
	ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
}
func AldwynDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0:
		case 1:
			if data.Quest.TalkToAldwyn_Quest01 == 0 {
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.TalkToAldwyn_Quest01 = 10
				})
			}
			ns.Object("Aldwyn").ChatStr("Lead the way.")
			ns.TellStory(audio.SwordsmanHurt, "War03b:AldwynIntro")
			ns.Object("Aldwyn").Follow(ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 3
			})
		case 3:
			if MayorsFear != nil {
				if MayorsFear.HasOwner(ns.Object("Aldwyn")) {
					ns.Object("Aldwyn").ChatStr("It is done. The beast is banished.")
					ns.CastSpell(spell.SUMMON_SPIDER, MayorsFear, MayorsFear)
					MayorsFear.Delete()
					necroEvent()
					ns.Effect(effect.SUMMON_CANCEL, MayorsFear, MayorsFear)
					ns.AudioEvent(audio.SummonAbort, ns.Object("Aldwyn"))
					nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
						data.Quest.SpidersAtTheMayor_Quest01 = 4
					})
				} else {
					ns.Object("Aldwyn").Follow(ns.GetCaller())
					ns.Object("Aldwyn").ChatStr("Bring me to the spider and I'll charm the beast.")
				}
			}
		case 10:
			switch data.Quest.MayorsScepter_Quest01 {
			case 1:
				switch data.Quest.AldwynGiveScrollWar_MayorsScepterQuest_01 {
				case 0:
					ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
					ns.TellStory(audio.SwordsmanHurt, "War03b:AldwynIntro")
					nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
						data.Quest.AldwynGiveScrollWar_MayorsScepterQuest_01 = 10
					})
				case 10:
					ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
					ns.TellStory(audio.SwordsmanHurt, "War03b:AldwynPre")
				}
			case 10:
				ns.TellStory(audio.SwordsmanHurt, "War03b:AldwynPost")
			}
		}
	case 1: // wiz
		switch data.Quest.SpidersAtTheMayor_Quest01 {
		case 0:
		case 1:
			if data.Quest.TalkToAldwyn_Quest01 == 0 {
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("Bring Aldwyn to the spider to banish it.")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.TalkToAldwyn_Quest01 = 10
				})
			}
			ns.Object("Aldwyn").ChatStr("Theogrin is trying to kill the spiders? Allow me to assist instead.")
			ns.Object("Aldwyn").Follow(ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.SpidersAtTheMayor_Quest01 = 3
			})
		case 3:
			if MayorsFear != nil {
				if MayorsFear.HasOwner(ns.Object("Aldwyn")) {
					ns.Object("Aldwyn").ChatStr("It is done. The beast is banished.")
					ns.CastSpell(spell.SUMMON_SPIDER, MayorsFear, MayorsFear)
					MayorsFear.Delete()
					necroEvent()
					ns.Effect(effect.SUMMON_CANCEL, MayorsFear, MayorsFear)
					ns.AudioEvent(audio.SummonAbort, ns.Object("Aldwyn"))
					nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
						data.Quest.SpidersAtTheMayor_Quest01 = 4
					})
				} else {
					ns.Object("Aldwyn").Follow(ns.GetCaller())
					ns.Object("Aldwyn").ChatStr("Bring me to the spider and I'll charm the beast.")
				}
			}
		}
	case 2: // con
		switch data.Quest.BecomeConjurerer_Quest01 {
		case 0:
			ns.SetDialog(ns.Object("Aldwyn"), ns.DialogYesNo, AldwynDialogueStart, AldwynDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Con02a:AldwinGreeting")
		case 10:
			switch data.Quest.TroubleAtTheManaMines_Quest01 {
			case 0:
				ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:AldwinImp")
				ns.AudioEvent(audio.JournalEntryAdd, p.Unit().Pos())
				ns.PrintStr("Go to the mines and locate the mine foreman.")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.TroubleAtTheManaMines_Quest01 = 1
				})
			case 1:
				ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:AldwinProd")
				ns.PrintStr("Go to the mines and locate the mine foreman.")
			case 10:
			}
		}
	}
}

func AldwynDialogueEnd() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch data.Quest.BecomeConjurerer_Quest01 {
	case 0:
		// find aldwyn update quest
		// 	You have completed an objective; your journal has been updated!
		switch ns.GetAnswer(ns.Object("Aldwyn")) {
		case ns.AnswerGoodbye:
			ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
		case ns.AnswerNo:
			ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
		case ns.AnswerYes:
			if ns.GetCaller().GetGold() < 30 {
				ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNormal, AldwynDialogueStart, AldwynDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:AldwinPoor")
			} else {
				SpiderFieldGUide := ns.FindClosestObject(ns.GetCaller(), ns.HasTypeName{"FieldGuide"})
				if SpiderFieldGUide != nil {
					ns.PrintStr("A new Field Guide has been added to your inventory!")
					ns.GetCaller().Equip(SpiderFieldGUide)
				}
				ns.SetDialog(ns.Object("Aldwyn"), ns.DialogNext, AldwynDialogueStart, AldwynDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:BecomeConjurer")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.BecomeConjurerer_Quest01 = 1
				})
			}
		}
	case 1:
		ns.TellStory("SwordsmanHurt", "Con02a:BecomeConjurer2")
		nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
			data.Quest.BecomeConjurerer_Quest01 = 2
		})
	case 2:
		ns.TellStory("SwordsmanHurt", "Con02a:BecomeConjurer3")
		nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
			data.Quest.BecomeConjurerer_Quest01 = 10
		})
	case 3:
	case 10:
	}
}

func MorganInit() {
	ns.SetDialog(ns.Object("Morgan"), ns.DialogYesNo, MorganDialogueStart, MorganDialogueEnd)
}
func MorganDialogueStart() {
	ns.SetDialog(ns.Object("Morgan"), ns.DialogYesNo, MorganDialogueStart, MorganDialogueEnd)
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		switch data.Quest.MorganConManSellBow_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02a:ConManSalesPitch3")
		case 10:
			ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
			ns.TellStory("SwordsmanHurt", "Con02a:ConManIdle2")
		}
	case 1: // wiz
		switch data.Quest.MorganConManSellBow_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02a:ConManSalesPitch3")
		case 10:
			ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
			ns.TellStory("SwordsmanHurt", "Con02a:ConManIdle2")
		}
	case 2: // con
		switch data.Quest.MorganConManSellBow_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			rnd := ns.Random(1, 2)
			switch rnd {
			case 1:
				ns.TellStory("SwordsmanHurt", "Con02a:ConManSalesPitch")
			case 2:
				ns.TellStory("SwordsmanHurt", "Con02a:ConManSalesPitch3")
			}
		case 10:
			ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
			ns.TellStory("SwordsmanHurt", "Con02a:ConManIdle2")
		}

	}
}
func MorganDialogueEnd() {
	switch ns.GetAnswer(ns.Object("Morgan")) {
	case ns.AnswerGoodbye:
	case ns.AnswerYes:
		if ns.GetCaller().GetGold() < 100 {
			ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
			ns.TellStory("SwordsmanHurt", "Con02a:ConManNotEnoughGold")
		} else {
			ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
			ns.TellStory("SwordsmanHurt", "Con02a:ConManSale")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.MorganConManSellBow_Quest01 = 10
			})
			bow := ns.CreateObject("Bow", ns.GetCaller())
			bow.FlagsEnable(object.FlagMarked)
			if bow != nil {
				ns.PrintStr("A bow has been added to your inventory!")
				ns.GetCaller().Equip(bow)
			}
		}
	case ns.AnswerNo:
		ns.SetDialog(ns.Object("Morgan"), ns.DialogNormal, MorganDialogueStart, MorganDialogueEnd)
		ns.TellStory("SwordsmanHurt", "Con02a:ConManSaleFailed")
	}
}

func Contest_GuardInit() {
	ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogNormal, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
}
func Contest_GuardDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		switch data.Quest.MayorsScepter_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory(audio.SwordsmanHurt, "War03b:T2Pre")
		case 10:
			ns.TellStory(audio.SwordsmanHurt, "War03b:T2Post") // if quest completed
		}
	case 1: // wiz
		switch data.Quest.MayorsScepter_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.AudioEvent(audio.HumanMaleEatApple, ns.GetCaller())
			ns.Object("Contest_Guard").ChatStr("We don't do Inversion lessons anymore. Go to Galava instead. I'm sure they'll be able to help you out.")
		case 10:
			ns.TellStory(audio.SwordsmanHurt, "War03b:T2Post") // if quest completed
		}
	case 2: // con
		switch data.Quest.ArcheryContest_Quests01 {
		case 0:
			if ns.GetCaller().InItems().FindObjects(nil, ns.HasTypeName{"Bow", "CrossBow"}) != 0 {
				// conjurere quest archery contest
				ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogYesNo, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:ContestGuard")
				// if yes but no bow: Con02a:NoBow
			} else {
				ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogNormal, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:NoBow")
				// You need your own bow for the contest. Byzanti's shop has them for a fair price.
			}
		case 1:
		case 10:
			ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogNormal, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
			switch data.Quest.MayorsScepter_Quest01 {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			case 10:
				ns.TellStory(audio.SwordsmanHurt, "War03b:T2Post") // if quest completed
			}
		}
	}
}
func Contest_GuardDialogueEnd() {
	p := ns.GetCaller().Player()
	switch p.Unit().GetClass() {
	case 0: // war
	case 1: // wiz
	case 2: // con
		switch ns.GetAnswer(ns.Object("Contest_Guard")) {
		case ns.AnswerGoodbye:
			Contest_GuardInit()
		case ns.AnswerYes:
			if ns.GetCaller().GetGold() < 20 {
				ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogNormal, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
				ns.TellStory(audio.HumanMaleEatFood, "Con02a:NotEnoughGold")
				Contest_GuardInit()
			} else {
				ns.SetDialog(ns.Object("Contest_Guard"), ns.DialogNormal, Contest_GuardDialogueStart, Contest_GuardDialogueEnd)
				ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
				ns.PrintStr("Talk to the contest official to begin the archery contest.")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.ArcheryContest_Quests01 = 1
				})
				ns.TellStory(audio.HumanMaleEatFood, "Con02a:EnterContest")
				ns.GetCaller().ChangeGold(-20)
				quiver := ns.CreateObject("Quiver", ns.GetCaller())
				if quiver != nil {
					ns.PrintStr("A quiver has been added to your inventory!")
					p.Unit().Equip(quiver)
				}
				Contest_GuardInit()
			}
		case ns.AnswerNo:
			Contest_GuardInit()
		}
	}
}

func Contest_OfficialInit() {
	ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
}
func Contest_OfficialDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
	case 1: // wiz
	case 2: // con
		switch data.Quest.ArcheryContest_Quests01 {
		case 1:
			ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNext, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Con02a:ContestGreeting")
			// I'm today's contest judge. Today's best score is 8 out of 10 hits on the target. You need at least 9 to win.
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.ArcheryContest_Quests01 = 2
			})
		case 2:
			if ArcheryContestActive {
				ns.Object("Contest_Official").ChatStr("Hold on! We've to finish this first.")
			} else {
				ArcheryContestActive = true
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.ArcheryContest_Quests01 = 3
				})
				ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:ContestOfficialWaiting")
				ns.NewTimer(ns.Seconds(4), func() {
					ns.Object("Contest_Official").ChatStr("Ready...")
					ns.NewTimer(ns.Seconds(1), func() {
						ns.Object("Contest_Official").ChatStr("Set...")
						ns.NewTimer(ns.Seconds(1), func() {
							ns.Object("Contest_Official").ChatStr("BEGIN!")
							archeryContest(p)
						})
					})
				})
			}
		case 3:
			if !ArcheryContestActive {
				if data.Quest.ArcheryContestScore_Quest01 < 9 {
					ns.SetDialog(ns.Object("Contest_Official"), ns.DialogYesNo, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
					ns.TellStory(audio.SwordsmanHurt, "Con02a:RetryContestChoice")
				} else {
					ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
					ns.TellStory(audio.SwordsmanHurt, "Con02a:AwardContestPrize")
					ns.GetCaller().ChangeGold(+50)
					ns.PrintStr("You gained 50 gold!")
					ns.AudioEvent(audio.TreasurePickup, ns.GetCaller())
					nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
						data.Quest.ArcheryContest_Quests01 = 10
					})
					// 	Congratulations to our new Champion Archer! Your prize is 50 gold pieces.
				}
			}
		case 10:
			ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
			// You've already competed here. Sorry the rules disallow reentrants.
			ns.TellStory(audio.SwordsmanHurt, "Con02a:NoMoreContestTries")
		}
	}

}
func Contest_OfficialInitDialogueEnd() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
	case 1: // wiz
	case 2: // con
		switch data.Quest.ArcheryContest_Quests01 {
		case 2:
			if ArcheryContestActive {
				ns.Object("Contest_Official").ChatStr("Hold on! We've to finish this first.")
			} else {
				ArcheryContestActive = true
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.ArcheryContest_Quests01 = 3
				})
				ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con02a:ContestOfficialWaiting")
				ns.NewTimer(ns.Seconds(4), func() {
					ns.Object("Contest_Official").ChatStr("Ready...")
					ns.NewTimer(ns.Seconds(1), func() {
						ns.Object("Contest_Official").ChatStr("Set...")
						ns.NewTimer(ns.Seconds(1), func() {
							ns.Object("Contest_Official").ChatStr("BEGIN!")
							archeryContest(p)
						})
					})
				})
			}
		case 3:
			switch ns.GetAnswer(ns.Object("Contest_Official")) {
			case ns.AnswerGoodbye:
				Contest_OfficialInit()
			case ns.AnswerYes:
				if ns.GetCaller().GetGold() < 20 {
					ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNormal, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
					ns.Object("Contest_Official").ChatStr("You don't have enough gold.")
					Contest_OfficialInit()
				} else {
					if ArcheryContestActive {
						Contest_OfficialInit()
						ns.Object("Contest_Official").ChatStr("Hold on! We've to finish this first.")
						nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
							data.Quest.ArcheryContest_Quests01 = 2
						})
					} else {
						nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
							data.Quest.ArcheryContest_Quests01 = 2
						})
						ns.SetDialog(ns.Object("Contest_Official"), ns.DialogNext, Contest_OfficialDialogueStart, Contest_OfficialInitDialogueEnd)
						ns.TellStory(audio.HumanMaleEatFood, "Con02a:RetryContest")
						ns.GetCaller().ChangeGold(-20)
						quiver := ns.CreateObject("Quiver", ns.GetCaller())
						if quiver != nil {
							ns.PrintStr("A quiver has been added to your inventory!")
							p.Unit().Equip(quiver)
						}
					}
				}
			case ns.AnswerNo:
				Contest_OfficialInit()
			}
		}
	}
}

func ContestJudge() {

}

func GeoffInit() {
	ns.SetDialog(ns.Object("Geoff"), ns.DialogNormal, GeoffDialogueStart, GeoffDialogueEnd)
	// gatekeep
}
func GeoffDialogueStart() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // war
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory("SwordsmanHurt", "War03b:GatekeeperPost")
		case 2:
			ns.TellStory("SwordsmanHurt", "War03b:GatekeeperPre")
		}
	case 1: // wiz
		ns.TellStory("SwordsmanHurt", "Con02a:Gatekeeper3Greet")
	case 2: // con
		switch data.Quest.BecomeConjurerer_Quest01 {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			ns.TellStory("SwordsmanHurt", "Con02a:Gatekeeper3Greet")
		case 10:
			if data.Quest.TroubleAtTheManaMines_Quest01 == 0 {
				ns.TellStory("SwordsmanHurt", "Con08a:GuardGreet")
			} else {
				ns.TellStory("SwordsmanHurt", "Con02a:Gatekeeper3Greet")
			}
		}
	}
}
func GeoffDialogueEnd() {
}

func gretchenInit() {
	ns.SetDialog(ns.Object("Gretchen"), ns.DialogNormal, gretchenDialogueStart, gretchenDialogueEnd)
}

func gretchenDialogueEnd() {
	ns.Object("Gretchen").Wander()
}

func gretchenDialogueStart() {
	ns.Object("Gretchen").Idle()
	ns.Object("Gretchen").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	switch p.Unit().GetClass() {
	case 0: // war
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman3")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman1")
		}
	case 1: // wiz
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman3")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman1")
		}
	case 2: // con
		rnd := ns.Random(1, 2)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman3")
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "War02a:NewTownswoman1")
		}
	}
}

func clydeInit() {
	ns.SetDialog(ns.Object("Clyde"), ns.DialogNormal, clydeDialogueStart, clydeDialogueEnd)
}

func clydeDialogueEnd() {
	ns.Object("Clyde").Wander()
}

func clydeDialogueStart() {
	ns.Object("Clyde").Idle()
	ns.Object("Clyde").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 2)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman13Talk02")
	case 2:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman13Talk03")
	}

}

func bryanInit() {
	ns.SetDialog(ns.Object("Bryan"), ns.DialogNormal, bryanDialogueStart, bryanDialogueEnd)
}

func bryanDialogueEnd() {
	ns.Object("Bryan").Wander()
}

func bryanDialogueStart() {
	ns.Object("Bryan").Idle()
	ns.Object("Bryan").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 3)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk01")
	case 2:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk02")
	case 3:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk03")
	}
}

func tanyaInit() {
	ns.SetDialog(ns.Object("Tanya"), ns.DialogNormal, tanyaDialogueStart, tanyaDialogueEnd)
}

func tanyaDialogueEnd() {
	ns.Object("Tanya").Wander()
}

// Kristine: War08a
func tanyaDialogueStart() {
	ns.Object("Tanya").Idle()
	ns.Object("Tanya").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	switch p.Unit().GetClass() {
	case 0: // war
		ns.TellStory(audio.SwordsmanHurt, "Wiz02A.scr:MaidenTalk05")
	case 1: // wiz
		ns.TellStory(audio.SwordsmanHurt, "Wiz02A.scr:MaidenTalk05")
	case 2: // con
		rnd := ns.Random(1, 3)
		switch rnd {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "Wiz02A.scr:MaidenTalk05")
		case 2:
			ns.TellStory(audio.HumanFemaleEatApple, "Con02A:Maiden3Talk01") // 	Leave the bickering to the Warriors and the Wizards. If you need something done, get a Conjurer.
		case 3:
			data := nw.LoadPlayer(p)
			if data.Quest.MayorsScepter_Quest01 == 10 {
				ns.TellStory(audio.SwordsmanHurt, "Con02A:Maiden2Talk01")
			} else {
				rnd := ns.Random(1, 2)
				switch rnd {
				case 1:
					ns.TellStory(audio.SwordsmanHurt, "Wiz02A.scr:MaidenTalk05")
				case 2:
					ns.TellStory(audio.HumanFemaleEatApple, "Con02A:Maiden3Talk01") // 	Leave the bickering to the Warriors and the Wizards. If you need something done, get a Conjurer.
				}
			}
		}
	}
}

func jailerInit() {
	ns.SetDialog(ns.Object("Jacob"), ns.DialogNormal, jailerDialogueStart, jailerDialogueEnd)
}

func jailerDialogueEnd() {
}

func jailerDialogueStart() {
	//Con02A:JailerTalk04	Don't hassle my prisoners.
	// ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk02")
	//Con02A:JailerTalk05	Last warning, don't bother the prisoners!
	ns.Object("Jacob").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 2)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk01")
	case 3:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk03")
	}

}

func joyceInit() {
	ns.SetDialog(ns.Object("Joyce"), ns.DialogNormal, joyceDialogueStart, joyceDialogueEnd)
}

func joyceDialogueEnd() {
}

func joyceDialogueStart() {
	ns.Object("Joyce").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Con02a:BarMaiden")
}

func lydiaInit() {
	ns.SetDialog(ns.Object("Lydia"), ns.DialogNormal, lydiaDialogueStart, lydiaDialogueEnd)
}

func lydiaDialogueEnd() {
}

func lydiaDialogueStart() {
	ns.Object("Lydia").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Con02a:BarMaiden2")
}

func henrickInit() {
	ns.SetDialog(ns.Object("Henrick"), ns.DialogYesNo, henrickDialogueStart, henrickDialogueEnd)
}

func henrickDialogueStart() {
	// Con02A:HenrickTalk01	Stand back! I'll handle this!
	//Con02A:HenrickTalk02	I make a good living charming wolves.
	//Con02A:HenrickTalk03	Greetings! My name is Henrick. I charm and tame the local wolves to keep the countryside safe for travelers.
	//Con02A:HenrickTalk04	I heard about the Necromancer that made a surprise visit in town. Very odd, since necromancers have not been seen for nearly 50 years now. Dark times ahead for all of us if the necromancers have returned.
	//Con02A:HenrickTalk05	The Mayor is still locked away in his home. How a man can be that frightened by a little spider is beyond me.
	//Con02A:HenrickTalk06	So you're a Conjurer now! Welcome to the brotherhood! As keepers of the lands, we have many responsibilities. I'm sure you'll honor us through your travels.
	ns.Object("Henrick").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch data.Character.WolfCompanion {
	case 0:
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSalesPitchA")
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSalesPitchB")
	case 2:
		ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickNoMoreWolves")
		henrickInit()
	}
}

func henrickDialogueEnd() {
	switch ns.GetAnswer(ns.Object("Henrick")) {
	case ns.AnswerGoodbye:
		henrickInit()
	case ns.AnswerYes:
		if ns.GetCaller().GetGold() < 200 {
			ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
			ns.TellStory(audio.HumanMaleEatFood, "war08b:HenrickNotEnoughGold")
			henrickInit()
		} else {
			ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Character.WolfCompanion++
			})
			ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSaleSuccessful")
			ns.GetCaller().ChangeGold(-200)
			wolf := ns.CreateObject("Wolf", ns.Waypoint("WolfCreateWP"))
			wolf.SetOwner(ns.GetCaller())
			wolf.Follow(ns.GetCaller())
			henrickInit()
			wolf.OnEvent(ns.EventDeath, func() {})
		}
	case ns.AnswerNo:
		henrickInit()
	}
}
