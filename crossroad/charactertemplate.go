package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
)

var template ns.Obj
var templateSpawn ns.Pointf

func initTemplate() {
	if ns.Object("Template") != nil {
		template = ns.Object("Template")
	} else {
		// Fix spawn
		template = ns.CreateObject("NPC", ns.GetHost())
	}
	templateSpawn = template.Pos()
	ns.StoryPic(template, "SEARCH IN THE GAME")
	ns.SetDialog(template, ns.DialogNormal, templateDialogueStart, templateDialogueEnd)
}

func templateDialogueStart() {
	template.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
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

func templateDialogueManaMinesQuest() {
	template.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if !data.FollowUpQuestDialogue && !data.FollowUpQuestDialogueCompleted {
		resetMillardDialogue()
	} else {
		//ns.SetDialog(template, ns.DialogNormal, templateDialogueManaMinesQuest, templatedDialogueEnd)
		//ns.TellStory(audio.HumanMaleEatFood, "DIALOG ADD HERE")
	}
}

func templateDialogueEnd() {
	if ns.GetAnswer(template) == 0 { // Goodbye
	}
	if ns.GetAnswer(template) == 1 { // Yes
	}
	if ns.GetAnswer(template) == 2 { // No
	}
}

func resetTemplateDialogue() {
	return
}
