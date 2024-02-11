package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var horst ns.Obj
var horstSpawn ns.Pointf

func dialogExampleHorst() {
}

func initHorst() {
	if ns.Object("Horst") != nil {
		horst = ns.Object("Horst")
	} else {
		// Fix spawn
		horst = ns.CreateObject("NPC", ns.GetHost())
	}
	horstSpawn = horst.Pos()
	ns.StoryPic(horst, "IxGuard2Pic")
	ns.SetDialog(horst, ns.DialogNormal, horstDialogueStart, horstDialogueEnd)
}

func horstDialogueStart() {
	ns.AudioEvent(audio.MaleNPC2Talkable, horst)
	horst.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
	case player.Conjurer:
	case player.Wizard:
	}
	horst_TroubleAtTheManaMines()
	horst_MayorTheogrinNeedsHelp()
}

func horst_MayorTheogrinNeedsHelp() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.FollowUpQuestDialogue {
	case QuestInactive: // Keep empty
	case QuestAccepted:
		switch data.Character.Class {
		case player.Warrior:
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2Intro") //	Good day, Warrior! Welcome to the Village of Ix! The Mayor is expecting you.
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.General.MayorTheogrinNeedsHelp = 2
			})
		case player.Conjurer:
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.General.MayorTheogrinNeedsHelp = 2
			})
		case player.Wizard:
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.General.MayorTheogrinNeedsHelp = 2
			})
		}
	case 2:
	case 3:
		// Always close dialogue with end and reset if used non normal dialogue.
	}
}

func horst_TroubleAtTheManaMines() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.TroubleAtTheManaMines {
	case QuestInactive, QuestComplete:
	default:
		ns.TellStory(audio.HumanMaleEatApple, "Con03A.scr:IxGuard2") // Good luck in the mines, boy. Watch out for bandits on the way!
		// Always close dialogue with end and reset if used non normal dialogue.
	}
}

func horstDialogueEnd() {
	switch ns.GetAnswer(horst) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetHorstDialogue() {
	ns.SetDialog(horst, ns.DialogNormal, horstDialogueStart, horstDialogueEnd)
}
