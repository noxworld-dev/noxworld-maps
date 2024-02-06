package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var brigadin ns.Obj
var brigadinSpawn ns.Pointf

func initBrigadin() {
	if ns.Object("Brigadin") != nil {
		brigadin = ns.Object("Brigadin")
	} else {
		// Fix spawn
		brigadin = ns.CreateObject("NPC", ns.GetHost())
	}
	brigadinSpawn = brigadin.Pos()
	ns.StoryPic(brigadin, "Warrior4Pic")
	ns.SetDialog(brigadin, ns.DialogNormal, brigadinDialogueStart, brigadinDialogueEnd)
}

func brigadinDialogueStart() {
	brigadin.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	if data.Character.Warrior {
		ns.TellStory(audio.SwordsmanHurt, "War03a:DunMirGuard3")
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

func brigadinDialogueEnd() {
	switch ns.GetAnswer(brigadin) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetBrigadinDialogue() {
	ns.SetDialog(brigadin, ns.DialogNormal, brigadinDialogueStart, brigadinDialogueEnd)
}
