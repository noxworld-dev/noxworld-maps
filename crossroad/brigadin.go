package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
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
	ns.AudioEvent(audio.FireKnight1Talkable, brigadin)
	brigadin.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	switch data.Character.Class {
	case player.Warrior:
		ns.TellStory(audio.SwordsmanHurt, "War03a:DunMirGuard3")
	case player.Conjurer:
		// Conjurer dialogue.
	case player.Wizard:
		// Wizard dialogue.
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
