package nw_ix

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		// StoryPic(Gvar18,"IxGuard1Pic") Janero
		//StoryPic(Gvar19,"IxGuard2Pic") Horst
		//ns.StoryPic(ns.Object("Janero"), "IxGuard1Pic")
		ns.StoryPic(ns.Object("Captain"), "AirshipCaptainPic")
		CaptainInit()
	})
}

func CaptainInit() {
	ns.SetDialog(ns.Object("Captain"), ns.DialogNormal, CaptainDialogueStart, CaptainDialogueEnd)
}
func CaptainDialogueStart() {
	ns.Object("Captain").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // warrior
	case 1: // wizard
		switch data.Quest.JandorWizStart_Quest01 {
		case 0:
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:JandorTalk01") // TODO: FIX DIALOGUE
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 1
			})
			// Hard, cold facts, lad. You must become a wizard's apprentice before you can gain entry to the Tower of Illusion. The one controlling that is the Arch-Wizard Horvath, an old friend of mine.
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:JandorTalk02") // TODO: FIX DIALOGUE
			// This is as far as I take you, lad. Find Horvath's apprentice at his house -- just north of the beach. He'll guide you to Galava from there.
		case 10:
		}
	case 2: // conjurer
		switch data.Quest.JandorStartQuestCon_Quest01 {
		case 0:
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			ns.TellStory(audio.SwordsmanHurt, "Con01a:CaptainGreet") // TODO: FIX DIALOGUE
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorStartQuestCon_Quest01 = 1
			})
		case 1:
			ns.PrintStr("Find Aldwyn in Ix.")
			ns.TellStory(audio.SwordsmanHurt, "Con01a:CaptainProd")
		case 10:
			switch data.Quest.TroubleAtTheManaMines_Quest01 {
			case 10:
				switch data.Quest.JandorFieldsOfValor_Quest01 {
				case 0:
					ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
					ns.TellStory(audio.SwordsmanHurt, "Con03A.scr:JandorA") // TODO: FIX DIALOGUE
				case 1:
					ns.TellStory(audio.SwordsmanHurt, "	Con03A.scr:JandorB") // TODO: FIX DIALOGUE
				case 10:
				}
			}
		}
	}
}
func CaptainDialogueEnd() {
}

func CrossRoadsToIx() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapIx, nw.GoToMapOptions{
		Exit: "CrossRoadsToIx",
	})
}

func CrossRoadsToManaMines() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapManaMines, nw.GoToMapOptions{
		Exit: "CrossRoadsToManaMines",
	})
}

func CrossRoadsToDunMir01() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapDunMir01, nw.GoToMapOptions{
		Exit: "CrossRoadsToDunMir01",
	})
}

func CrossRoadsToBeachHouse() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapBeachHouse, nw.GoToMapOptions{
		Exit: "CrossRoadsToBeachHouse",
	})
}
