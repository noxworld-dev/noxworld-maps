package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
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
	// General Quests
	template_FollowUpQuestDialogue()
}

func template_FollowUpQuestDialogue() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.FollowUpQuestDialogue {
	case 0: // Keep empthy: 0 = inactive/not accepted/not completed, 10 = completed
	case 1:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		ns.TellStory(audio.HumanMaleEatFood, "FILL IN BLANKS") //
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.General.FollowUpQuestDialogue = 2
		})
	case 2:
	case 3:
		// Always close dialogue with end and reset if used non normal dialogue.
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
	ns.SetDialog(template, ns.DialogNormal, templateDialogueStart, templateDialogueEnd)
	return
}
