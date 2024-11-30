package nw_glv01

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.Music(21, 100)
	ns.NewTimer(ns.Frames(1), func() {
		ns.StoryPic(ns.Object("Lance"), "WizardGuard1Pic")
		LanceInit()
		ns.StoryPic(ns.Object("HorvathApprentice"), "WoundedApprenticePic")
		HorvathApprenticeInit()
	})
}

func HorvathApprenticeInit() {
	ns.SetDialog(ns.Object("HorvathApprentice"), ns.DialogNormal, HorvathApprenticeDialogueStart, HorvathApprenticeDialogueEnd)
}
func HorvathApprenticeDialogueStart() {
	ns.Object("HorvathApprentice").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // Warrior
	case 1: // Wizard
		switch data.Quest.JandorWizStart_Quest01 {
		case 4:
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			ns.TellStory(audio.HumanMaleEatFood, "Wiz01:ApprenticeTalk01")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 5
			})
		case 0, 1, 2, 3, 5, 6, 7, 8, 9, 10:
			ns.AudioEvent(audio.HumanMaleHurtMedium, ns.GetCaller())
			ns.Object("HorvathApprentice").ChatStr("I'm too far gone. Tell Horvath it was a nec... nec... necromancer...")
		}
	case 2: // Conjurer
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01:ApprenticeTalk01")
	}

}
func HorvathApprenticeDialogueEnd() {
	switch ns.GetAnswer(ns.Object("HorvathApprentice")) {
	case ns.AnswerGoodbye:
	case ns.AnswerNo:
	case ns.AnswerYes:
	}
}

func LanceInit() {
	ns.SetDialog(ns.Object("Lance"), ns.DialogNormal, LanceDialogueStart, LanceDialogueEnd)
}
func LanceDialogueStart() {
	ns.Object("Lance").LookAtObject(ns.GetCaller())
	ns.Object("Lance").ChatStr("Travel safe!")
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	// Warrior
	case 0:
		// Wizard
		//TellStory("SwordsmanHurt", "")
	case 1:
		// Conjurer
	case 2:
	}

}
func LanceDialogueEnd() {
}

func BeachHouseToGalava() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapGalava, nw.GoToMapOptions{
		Exit: "BeachHouseToGalava",
	})
}

func BeachHouseToCrossRoads() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapCrossroads, nw.GoToMapOptions{
		Exit: "BeachHouseToCrossRoads",
	})
}
