package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var osborn ns.Obj
var osbornSpawn ns.Pointf

func initOsborn() {
	if ns.Object("Osborn") != nil {
		osborn = ns.Object("Osborn")
	} else {
		// Fix spawn
		osborn = ns.CreateObject("NPC", ns.GetHost())
	}
	osbornSpawn = osborn.Pos()
	ns.StoryPic(osborn, "OsbornPic")
	ns.SetDialog(osborn, ns.DialogNormal, osbornDialogueStart, osbornDialogueEnd)
}

func osbornDialogueStart() {
	osborn.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Start LostSpectales quest.
	if !data.Quest.LostSpectacles && !data.Quest.LostSpectaclesCompleted {
		ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:HermitMeet01") // Gahhhhhh! No! Don't kill me! Oh. A young man?! I can't see well at all. But I know you're not one of those infernal bandits who stole my spectacles! I'm almost blind without them. If you could get them back, you'd save my life and I'd be eternally grateful.
		updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
			data.Quest.LostSpectacles = true
		})
		return
	}
	if data.Quest.LostSpectacles && !data.Quest.LostSpectaclesCompleted {
		ns.TellStory(audio.ArcherHurt, "Con03A.scr:HermitMeet02")
		return
	}
	// on complete with spectacles // Con03A.scr:HermitHappy	My spectacles! You brought them back! May all that is great bless you! And please, take this scroll. It contains all I have learned about bats. It would be invaluable to any conjurer.
}

func osbornDialogueEnd() {
	if ns.GetAnswer(osborn) == 0 { // Goodbye
	}
	if ns.GetAnswer(osborn) == 1 { // Yes
	}
	if ns.GetAnswer(osborn) == 2 { // No
	}
}

func resetOsbornDialogue() {
	ns.SetDialog(osborn, ns.DialogNormal, osbornDialogueStart, osbornDialogueEnd)
	return
}
