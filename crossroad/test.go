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
	ns.TellStory(audio.OgreBruteDie, "Do you want to help the mines?")
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			return
		}
	}
}

func testDialogueEnd() {
	data := loadMyQuestData(ns.GetCaller().Player())
	if ns.GetAnswer(test) == 0 { // Goodbye
	}
	if ns.GetAnswer(test) == 1 { // Yes
		if !data.TroubleAtTheManaMines && !data.TroubleAtTheManaMinesCompleted {
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyQuestData) {
				data.TroubleAtTheManaMines = true
			})
		}
	}
	if ns.GetAnswer(test) == 2 { // No
	}
}

func resetTestDialogue() {
	return
}
