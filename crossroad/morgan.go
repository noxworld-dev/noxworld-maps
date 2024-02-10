package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
)

var morgan ns.Obj
var morganSpawn ns.Pointf

// Warrior
//War07A.scr:MorganTalk01	Psst...over here, Fire Knight...
//War07A.scr:MorganTalk02	Hail, cell-brother! Morgan Lightfingers' the name! You seem like a man on a mission, so I thought I might offer this help. The last time I was arrested, they threw me in the very cell you're in now... And I escaped by digging a tunnel. The entrance should still be hidden under the cot in your cell. I'll keep an eye out for the warden while you search for it.
//War07A.scr:MorganTalk03	Thanks for freeing me. Let's get out of here.
//War07A.scr:MorganTalk04	Hi there. Name's Morgan Lightfingers. Could you get me out of here? If you can, I'll get you into the Tower of Illusion.
//War07A.scr:MorganTalk05	Thanks, brother. Follow me. I know a secret way out.
//War07A.scr:MorganTalk06	Thanks for my freedom. Accept this gold as token of my gratitude. Ah, yes! And here's a key which gives you access to the Tower of Illusion. Be careful when you go in. The wizards won't react well to the sight of a Fire Knight in their tower. Thanks again!

// Conjurer
//Con02a:ConManIdle	Sorry, the bow's been sold.
//Con02a:ConManIdle2	Nice bow.
//Con02a:ConManNotEnoughGold	Come back when you have 100 gold.
//Con02a:ConManSale	There ya go! That bow was made for you. You can't lose the contest with a bow like that.
//Con02a:ConManSaleFailed	You're making a mistake. Go to Byzanti's and check out the new bows. Then get back here quick -- I've got someone else interested in it too.
//Con02a:ConManSalesPitch	No doubt you're here for the archery contest. I've got a real nice bow for sale. Since I like your face, I'll let it go for 100 gold. Are you interested?
//Con02a:ConManSalesPitch2	I still have that bow. You want it?
//Con02a:ConManSalesPitch3	This bow is a beauty. The previous owner only used it on weekends.Over in Byzanti's shop a bow like this would go for 300 gold or more. But since I like your face, I'll let it go for 100. You want it?

// Wizard
// Wiz02A.scr:MorganTalk01	I surrender! Please, don't hurt me! I'll go quietly!

func initMorgan() {
	if ns.Object("Morgan") != nil {
		morgan = ns.Object("Morgan")
	} else {
		// Fix spawn
		if ns.Object("IxEntrance") != nil {
			IxEntrance = ns.Object("IxEntrance")
		}
		if ns.Object("GalavaEntrance") != nil {
			GalavaEntrance = ns.Object("GalavaEntrance")
		}
		if ns.Object("DunMirEntrance") != nil {
			GalavaEntrance = ns.Object("DunMirEntrance")
		}
		morgan = ns.CreateObject("NPC", IxEntrance)
	}
	morganSpawn = morgan.Pos()
	ns.StoryPic(morgan, "MorganPic")
	ns.SetDialog(morgan, ns.DialogNormal, morganDialogueStart, morganDialogueEnd)
}

func morganDialogueStart() {
	morgan.LookAtObject(ns.GetCaller())
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

func morganDialogueManaMinesQuest() {
	morgan.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	if 1 == 1 {
		resetMillardDialogue()
	} else {
		//ns.SetDialog(morgan, ns.DialogNormal, morganDialogueManaMinesQuest, morgandDialogueEnd)
		//ns.TellStory(audio.HumanMaleEatFood, "DIALOG ADD HERE")
	}
}

func morganDialogueEnd() {
	switch ns.GetAnswer(morgan) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetMorganDialogue() {
	return
}
