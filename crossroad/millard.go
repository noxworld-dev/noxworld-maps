package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var millard ns.Obj
var millardSpawn ns.Pointf

func init() {
	OnLateInit(func() {
		if ns.Object("Millard") != nil {
			millard = ns.Object("Millard")
		} else {
			// Fix spawn
			millard = ns.CreateObject("NPC", ns.GetHost())
		}
		millardSpawn = millard.Pos()
		ns.StoryPic(millard, "MalePic2")
		ns.SetDialog(millard, ns.DialogNormal, millard_TroubleAtTheManaMines, millardDialogueEnd)
		ns.Object("MineDoor1").Lock(true)
		ns.Object("MineDoor2").Lock(true)
	})
}

func millard_TroubleAtTheManaMines() {
	ns.AudioEvent(audio.NPCTalkable, millard)
	millard.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.TroubleAtTheManaMines {
	case QuestInactive:
		ns.TellStory(audio.SwordsmanHurt, "War03a:MineGuard") // Sorry, the Mana Mines are closed.
	case QuestAccepted:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		ns.SetDialog(millard, ns.DialogNext, millard_TroubleAtTheManaMines, millard_TroubleAtTheManaMines)
		ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:MineGuardA") // You have arrived! Lives are at stake! Please hurry to the mines, just west, down the path. The Foreman is waiting for you outside the Miner's Lodge. He'll direct you from there.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.General.TroubleAtTheManaMines = 2
		})
	case 2:
		ns.SetDialog(millard, ns.DialogNormal, millard_TroubleAtTheManaMines, millardDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:MineGuardB") // The foreman will tell you the details of our crisis. He waits for you by the entrance of the mines.
	case 3:
		ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:MineGuardC") // Your valor knows no peer. We all thank you for saving our men.
	}
}

func millardDialogueEnd() {
	switch ns.GetAnswer(millard) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetMillardDialogue() {
	ns.SetDialog(millard, ns.DialogNormal, millard_TroubleAtTheManaMines, millardDialogueEnd)
}
