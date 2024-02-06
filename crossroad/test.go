package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var test ns.Obj
var testSpawn ns.Pointf

func initTest() {
	if ns.Object("Test") != nil {
		test = ns.Object("Test")
	} else {
		// Fix spawn
		test = ns.CreateObject("NPC", ns.GetHost())
	}
	testSpawn = test.Pos()
	ns.SetDialog(test, ns.DialogYesNo, testDialogueStart, testDialogueEnd)
}

func testDialogueStart() {
	test.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	ns.TellStory(audio.OgreBruteDie, "Do you want to help the mines?")
	// Warrior dialogue.
	if data.Character.Warrior {
		return
	}
	// Conjurer dialogue.
	if data.Character.Conjurer {
		return
	}
	// Wizard dialogue.
	if data.Character.Wizard {
		return
	}
}

func testDialogueEnd() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch ns.GetAnswer(test) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		if !data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.TroubleAtTheManaMines = true
			})
		}
	case ns.AnswerNo:
		// No
	}
}

func resetTestDialogue() {
	return
}
