package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var horvath ns.Obj
var horvathSpawn ns.Pointf

func initHorvath() {
	if ns.Object("Horvath") != nil {
		horvath = ns.Object("Horvath")
	} else {
		// Fix spawn
		horvath = ns.CreateObject("NPC", ns.GetHost())
	}
	horvathSpawn = horvath.Pos()
	ns.StoryPic(horvath, "HorvathPic")
	ns.SetDialog(horvath, ns.DialogNormal, horvathDialogueStart, horvathDialogueEnd)
}

func horvathDialogueStart() {
	horvath.LookAtObject(ns.GetCaller())
	ns.AudioEvent(audio.HorvathTalkable, horvath)
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		// Warrior dialogue.
	case player.Conjurer:
		// Conjurer dialogue.
	case player.Wizard:
		horvath_FindApprentice()
		horvath_TellHorvathYouFoundTheApprentice()
	}
	// General Quests

}

func horvath_FindApprentice() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.Wizard.TravelToTheApprenticeHouse {
	case 0:
	case 1:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		ns.SetDialog(horvath, ns.DialogNext, horvath_FindApprentice, horvath_FindApprentice)
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk01") //	Thieving little creatures! The Urchins just moved into our area and have already made absolute pests of themselves! I am Arch-Wizard Horvath. I am looking for my apprentice. I fear he has been lost in a nearby Urchin Den.If you could find him and bring him back to me, I would be willing to bring you along with us to Galava. Come, I will show you the way to the wretched Urchin Den outside my apprentice's home.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.Wizard.BecomeTheWizardApprentice = 10
			data.Quest.Wizard.TravelToTheApprenticeHouse = 2
			data.Quest.Wizard.FindApprentice = 1
		})
	case 2:
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.Wizard.TravelToTheApprenticeHouse = 3
			ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk02") //	Here's a spellbook and some other articles you might need.
		})
	case 3:
		ns.SetDialog(horvath, ns.DialogNormal, horvathDialogueStart, horvathDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk02a") // I've unlocked the gates. Just follow the road and I expect you will have no trouble finding these detestable interloping urchins. Good luck. I hope to see you again.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.Wizard.TravelToTheApprenticeHouse = 4
		})
	case 4:
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk03") //	Have you found my apprentice yet? Oh dear, I fear the worst.
	}
}

func horvath_TellHorvathYouFoundTheApprentice() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.Wizard.TellHorvathYouFoundTheApprentice {
	case 0:
	case 1:
		ns.PrintStr("You have gained a new Quest.")
		ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk04") //	My apprentice's robe! Dark fortunes must have befallen him.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.Wizard.GoToHorvathHisOffice = 1
			data.Quest.Wizard.TellHorvathYouFoundTheApprentice = 2
		})
	case 2:
		ns.TellStory(audio.HumanMaleDrinkJug, "Wiz01A.scr:HorvathTalk05") // When you are ready, just step onto the teleporter and we will be on our way.
	}
}

func horvathDialogueEnd() {
	switch ns.GetAnswer(horvath) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetHorvathDialogue() {
	ns.SetDialog(horvath, ns.DialogNormal, horvathDialogueStart, horvathDialogueEnd)
	return
}
