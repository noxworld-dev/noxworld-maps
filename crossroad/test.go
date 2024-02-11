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
	//data := loadMyQuestData(ns.GetCaller().Player())
	ns.TellStory(audio.OgreBruteDie, "Do you want to help the mines?")

}

func testDialogueEnd() {
	//data := loadMyQuestData(ns.GetCaller().Player())
	switch ns.GetAnswer(test) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.General.TroubleAtTheManaMines = 1
		})
	case ns.AnswerNo:
		// No
	}
}

func resetTestDialogue() {
	ns.SetDialog(test, ns.DialogYesNo, testDialogueStart, testDialogueEnd)
}
