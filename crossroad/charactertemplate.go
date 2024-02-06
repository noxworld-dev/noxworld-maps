package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
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
	switch data.Character.Class {
	case player.Warrior:
		// Warrior dialogue.
	case player.Conjurer:
		// Conjurer dialogue.
	case player.Wizard:
		// Wizard dialogue.
	}
}

func templateDialogueManaMinesQuest() {
	template.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if !data.Quest.FollowUpQuestDialogue && !data.Quest.FollowUpQuestDialogueCompleted {
		resetMillardDialogue()
	} else {
		//ns.SetDialog(template, ns.DialogNormal, templateDialogueManaMinesQuest, templatedDialogueEnd)
		//ns.TellStory(audio.HumanMaleEatFood, "DIALOG ADD HERE")
	}
}

func templateDialogueEnd() {
	switch ns.GetAnswer(template) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetTemplateDialogue() {
	return
}
