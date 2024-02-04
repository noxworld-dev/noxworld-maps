package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var millard ns.Obj
var millardSpawn ns.Pointf

func initMillard() {
	if ns.Object("Millard") != nil {
		millard = ns.Object("Millard")
	} else {
		// Fix spawn
		millard = ns.CreateObject("NPC", ns.GetHost())
	}
	millardSpawn = millard.Pos()
	ns.StoryPic(millard, "MalePic2")
	ns.SetDialog(millard, ns.DialogNormal, millardDialogueStart, millardDialogueEnd)
	ns.Object("MineDoor1").Lock(true)
	ns.Object("MineDoor2").Lock(true)
}

func millardDialogueStart() {
	millard.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if data.TroubleAtTheManaMines && !data.TroubleAtTheManaMinesCompleted {
		ns.SetDialog(millard, ns.DialogNext, millardDialogueManaMinesQuest, millardDialogueManaMinesQuest)
		ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:MineGuardA") // You have arrived! Lives are at stake! Please hurry to the mines, just west, down the path. The Foreman is waiting for you outside the Miner's Lodge. He'll direct you from there.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyQuestData) {
			data.TroubleAtTheManaMines = true
		})
		return
	}
	if data.TroubleAtTheManaMinesCompleted {
		ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:MineGuardC") // Your valor knows no peer. We all thank you for saving our men.
		return
	}
	ns.TellStory(audio.SwordsmanHurt, "War03a:MineGuard") // Sorry, the Mana Mines are closed.
	return
}

func millardDialogueManaMinesQuest() {
	millard.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if !data.TroubleAtTheManaMines && !data.TroubleAtTheManaMinesCompleted {
		resetMillardDialogue()
	} else {
		ns.SetDialog(millard, ns.DialogNormal, millardDialogueManaMinesQuest, millardDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:MineGuardB") // The foreman will tell you the details of our crisis. He waits for you by the entrance of the mines.
	}
}

func millardDialogueEnd() {
	if ns.GetAnswer(millard) == 0 { // Goodbye
	}
	if ns.GetAnswer(millard) == 1 { // Yes
	}
	if ns.GetAnswer(millard) == 2 { // No
	}
}

func resetMillardDialogue() {
	ns.SetDialog(millard, ns.DialogNormal, millardDialogueStart, millardDialogueEnd)
	return
}
