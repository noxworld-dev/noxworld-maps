package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

func dialogExampleCaptain() {
	// General
	//
	//
	//} else {
	//	ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2cStart") // The path leads to Dün Mir, the great subterranean citadel of Warriors!
	//}
	// ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorB") // Gather supplies.
	// ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2eStart") // I like my new coat. Do you?
	// ns.TellStory(audio.HumanMaleEatFood, "Wiz08a:CaptainGreet") // I've never been inside the Temple of Ix and can only pass this on this warning as it was spoken to me, years ago; beware of the Golem.

	//ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2aStart") // You need to press on now, lad.
	//ns.TellStory(audio.HumanMaleEatFood, "Con01a:CaptainProd")        // Go! Don't just stand there. Get a move on.
	//ns.TellStory(audio.HumanMaleEatFood, "War05B.scr:CaptainIdle")    // Time is running out. Now be on your way!
	//ns.TellStory(audio.HumanMaleEatFood, "War09a:CaptainImpatient")   //  go hurry up
	// Quest specific
	// Find Horvath.
	// Wiz05A.scr:CaptainGreeting : ogres attacked brin, horvath is taken.
	//ns.TellStory(audio.HumanMaleEatFood, "Con05A.scr:CaptainWaiting") // Hurry and find Horvath! He is waiting to talk to you. Be off with you! Hecubah is on the move and we need to stop her!
	//ns.TellStory(audio.HumanMaleEatFood, "Wiz05B.scr:CaptainWaiting") // Hurry and find Horvath, lad!
	// Wiz05B.scr:CaptainReturned : thanks for resquing horvath.
	// Wiz05B.scr:CaptainIdle : more thanks for resquing horvaht.
	// Ogres and brin resque ladys
	// War05B.scr:CaptainWaiting : hury up and resque the ladys from the ogre.
	// War05B.scr:CaptainSuccess3 : good work on saving the ladys
	// War05B.scr:CaptainSuccess2 : good, but mmore work saving left
	// War05B.scr:CaptainSuccess1 : good but more ladys

	// The swamp
	// War09a:CaptainGreet : arrival at the swamp
	// Swamp en mention Mordwyn.
	// War09a:CaptainLeave : weird place i know ordwyn loves it.
}

var captain ns.Obj
var captainSpawn ns.Pointf

func initCaptain() {
	if ns.Object("Captain") != nil {
		captain = ns.Object("Captain")
	} else {
		// Fix spawn
		captain = ns.CreateObject("AirshipCaptain", ns.GetHost())
	}
	captainSpawn = captain.Pos()
	ns.StoryPic(captain, "AirshipCaptainPic")
	ns.SetDialog(captain, ns.DialogNormal, captainDialogueStart, captainDialogueEnd)
}

func captainDialogueStart() {
	captain.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Wizard:
		// Wizards starting quest.
		if !data.Quest.BecomeTheWizardApprentice && !data.Quest.BecomeTheWizardApprenticeCompleted {
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.BecomeTheWizardApprentice = true
			})
			ns.SetDialog(captain, ns.DialogNext, captainDialogueWizardStartingQuest, captainDialogueWizardStartingQuest)
			ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:JandorTalk01") // Hard, cold facts, lad. You must become a wizard's apprentice before you can gain entry to the Tower of Illusion. The one controlling that is the Arch-Wizard Horvath, an old friend of mine.I know Horvath is in need of a worthy apprentice. But are you worthy, lad? heh, heh, heh...Follow the beach road. Find the home of his new apprentice. And beware of urchins. They can be the most nettlesome of pests, at best. I guarantee, you don't want to learn their worst, lad.
			return
		}
		if data.Quest.BecomeTheWizardApprentice && !data.Quest.BecomeTheWizardApprenticeCompleted {
			captainDialogueWizardStartingQuest()
		}
		// Tomb of Valor starting quest.
		if data.Quest.TroubleAtTheManaMinesCompleted && !data.Quest.ExploreTheTombsOfValor && !data.Quest.ExploreTheTombsOfValorCompleted {
			ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorA") // Good work with the miners, lad! With the Mana supply reestablished, the wizards of Galava will be busy tonight. But if I know Hecubah, I fear rumors of her practicing the forbidden Black Arts may be true. And it is up to us to journey to the Field of Valor -- where you must find out for sure if she is communing with the Undead through her ancestors' occult art of necromancy.
			// Wiz03a:AirshipCaptainSendoff : well done, compliments from horvath, go to the fields of valor.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainIdle")         // Enter the Tomb of Valor, lad! We must know if reports of Hecubah's acts are true. Hurry! Go into the crypts to see if Hecubah has already been here.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainGreet")        // The entrance to the tombs is in the building at the end of the path. See if Hecubah's been here. You'll know. And careful. Many of the crypts have traps against grave robbers.
			return
		}
		// if tomb completed {}
		// //ns.TellStory(audio.HumanMaleEatFood, "Con05A.scr:CaptainGreeting") // The trials in the Tomb of Valor have tested your mettle well, lad. But no time to rest on laurels. You must now recover the Amulet of Teleportation, which the Ogres took from Horvath when they raided the unlucky Hamlet of Brin! With the Ogres on the move and with this as their prize, Hecubah is most certainly behind it all. Horvath is waiting in the village. Now, off with you!
		// }
		// ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2dStart") // But find the smuggler's tunnel in the cavern and you are in!
		// ns.TellStory(audio.HumanMaleEatFood, "Wiz06a:Captain2") // Invisibility will be invaluable for sneaking past the guards. The Lock spell will help keep anyone from following you on your way out. Good luck and good speed.
		// ns.TellStory(audio.HumanMaleEatFood, "Wiz05C.scr:CaptainIdle") // Thanks again, Apprentice. You have performed a noble service, one which I can't easily repay... Thank you. I suppose you've outgrown being an apprentice.
	case player.Warrior:
		if !data.Quest.JoinTheFireKnights && !data.Quest.JoinTheFireKnightsCompleted {
			updateMyQuestData(ns.GetCaller().Player(), func(data *MyAccountData) {
				data.Quest.JoinTheFireKnights = true
			})
			ns.SetDialog(captain, ns.DialogNext, captainDialogueWizardStartingQuest, captainDialogueWizardStartingQuest)
			ns.TellStory(audio.HumanMaleEatFood, "War01A.scr:CaptainTalkStart") // Gain entry to Dün Mir, find the Academy and then be ready for the test of your life -- The Gauntlet. It weeds out the weak recruits.return
			return
		}
		if data.Quest.JoinTheFireKnights && !data.Quest.JoinTheFireKnightsCompleted {
			captainDialogueWarriorStartingQuest()
		}
		if data.Quest.TroubleAtTheManaMinesCompleted && !data.Quest.ExploreTheTombsOfValor && !data.Quest.ExploreTheTombsOfValorCompleted {
			ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorA") // Good work with the miners, lad! With the Mana supply reestablished, the wizards of Galava will be busy tonight. But if I know Hecubah, I fear rumors of her practicing the forbidden Black Arts may be true. And it is up to us to journey to the Field of Valor -- where you must find out for sure if she is communing with the Undead through her ancestors' occult art of necromancy.
			return
		}
		// Tomb of Valor starting quest.
		if data.Quest.TroubleAtTheManaMinesCompleted && !data.Quest.ExploreTheTombsOfValor && !data.Quest.ExploreTheTombsOfValorCompleted {
			ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorA") // Good work with the miners, lad! With the Mana supply reestablished, the wizards of Galava will be busy tonight. But if I know Hecubah, I fear rumors of her practicing the forbidden Black Arts may be true. And it is up to us to journey to the Field of Valor -- where you must find out for sure if she is communing with the Undead through her ancestors' occult art of necromancy.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainIdle")         // Enter the Tomb of Valor, lad! We must know if reports of Hecubah's acts are true. Hurry! Go into the crypts to see if Hecubah has already been here.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainGreet")        // The entrance to the tombs is in the building at the end of the path. See if Hecubah's been here. You'll know. And careful. Many of the crypts have traps against grave robbers.
			return
		}
		captain.ChatStr("works")
		// War03b:AirshipCaptainIxSpeech : // What, ho! You've returned, lad, or should I say Warrior! Heh, heh, heh...
	case player.Conjurer:
		if data.Quest.TroubleAtTheManaMinesCompleted && !data.Quest.ExploreTheTombsOfValor && !data.Quest.ExploreTheTombsOfValorCompleted {
			ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorA") // Good work with the miners, lad! With the Mana supply reestablished, the wizards of Galava will be busy tonight. But if I know Hecubah, I fear rumors of her practicing the forbidden Black Arts may be true. And it is up to us to journey to the Field of Valor -- where you must find out for sure if she is communing with the Undead through her ancestors' occult art of necromancy.
			return
		}
		// Tomb of Valor starting quest.
		if data.Quest.TroubleAtTheManaMinesCompleted && !data.Quest.ExploreTheTombsOfValor && !data.Quest.ExploreTheTombsOfValorCompleted {
			ns.TellStory(audio.HumanMaleEatFood, "Con03A.scr:JandorA") // Good work with the miners, lad! With the Mana supply reestablished, the wizards of Galava will be busy tonight. But if I know Hecubah, I fear rumors of her practicing the forbidden Black Arts may be true. And it is up to us to journey to the Field of Valor -- where you must find out for sure if she is communing with the Undead through her ancestors' occult art of necromancy.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainIdle")         // Enter the Tomb of Valor, lad! We must know if reports of Hecubah's acts are true. Hurry! Go into the crypts to see if Hecubah has already been here.
			//ns.TellStory(audio.HumanMaleEatFood, "Con04a:CaptainGreet")        // The entrance to the tombs is in the building at the end of the path. See if Hecubah's been here. You'll know. And careful. Many of the crypts have traps against grave robbers.
			return
		}
		captain.ChatStr("works")
	}
	// ns.TellStory(audio.HumanMaleEatFood, "Con01a:CaptainGreet") // I'll bring you no further, lad. Too many prying eyes the closer we get by air to the Village of Ix.Follow the tunnel which leads to Ix.Find my old friend Aldwyn. He'll help you if he can.Take this staff. Forest beasts will heed its bite should they threaten. Careful as you go, lad.

	// General
	// ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2cStart") // The path leads to Dün Mir, the great subterranean citadel of Warriors!
}

func captainDialogueWarriorStartingQuest() {
	captain.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if !data.Quest.JoinTheFireKnights && !data.Quest.JoinTheFireKnightsCompleted {
		resetCaptainDialogue()
	} else {
		ns.SetDialog(captain, ns.DialogNormal, captainDialogueWarriorStartingQuest, captainDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "War01a:CaptainTalk2bStart") // This sword is the best blade I can offer you. Wear it with pride, lad!
	}
}

func captainDialogueWizardStartingQuest() {
	captain.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if !data.Quest.BecomeTheWizardApprentice && !data.Quest.BecomeTheWizardApprenticeCompleted {
		resetCaptainDialogue()
	} else {
		ns.SetDialog(captain, ns.DialogNormal, captainDialogueWizardStartingQuest, captainDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "Wiz01A.scr:JandorTalk02") // This is as far as I take you, lad. Find Horvath's apprentice at his house -- just north of the beach. He'll guide you to Galava from there.
	}
}

func captainDialogueEnd() {
	switch ns.GetAnswer(captain) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetCaptainDialogue() {
	ns.SetDialog(captain, ns.DialogNormal, captainDialogueStart, captainDialogueEnd)
	return
}
