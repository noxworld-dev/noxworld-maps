package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var horvathApprentice ns.Obj
var horvathApprenticeSpawn ns.Pointf

func initWizardApprentice() {
	if ns.Object("HorvathApprentice") != nil {
		horvathApprentice = ns.Object("HorvathApprentice")
	} else {
		// Fix spawn
		horvathApprentice = ns.CreateObject("NPC", ns.GetHost())
	}
	horvathApprenticeSpawn = horvathApprentice.Pos()
	ns.StoryPic(horvathApprentice, "WoundedApprenticePic")
	ns.SetDialog(horvathApprentice, ns.DialogNormal, horvathApprenticeDialogueStart, horvathApprenticeDialogueEnd)
}

func horvathApprenticeDialogueStart() {
	ns.AudioEvent(audio.WoundedNPCTalkable, horvathApprentice)
	horvathApprentice.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		// Warrior dialogue.
	case player.Conjurer:
		// Conjurer dialogue.
	case player.Wizard:
		horvathApprentice_FindApprentice()
	}
	// General Quests

}

func horvathApprentice_FindApprentice() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.Wizard.FindApprentice {
	case QuestInactive:
	case QuestAccepted:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01:ApprenticeTalk01") // Thank you for trying to rescue me, but I am already too far gone. The urchins... they weren't acting alone! It was a nec... nec... necromancer...
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.Wizard.FindApprentice = QuestComplete
			data.Quest.Wizard.TellHorvathYouFoundTheApprentice = QuestAccepted
			data.Quest.Wizard.TravelToTheApprenticeHouse = QuestComplete
		})
	}
}

func horvathApprenticeDialogueEnd() {
	switch ns.GetAnswer(horvathApprentice) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetWizardApprenticeDialogue() {
	return
}
