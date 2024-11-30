package nw_toi01

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

var HorvathInLab bool
var InversionLessonActive bool
var InversionTarget ns.Obj
var InversionScore int

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		ns.StoryPic(ns.Object("Horvath"), "HorvathPic")
		HorvathInit()
		ns.StoryPic(ns.Object("Teacher"), "WizardGuard1Pic")
		TeacherInit()
	})
}

func TriggerTeacher() {
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	if data.Quest.WizardLesson01_Quest01 == 1 {
		ns.StartDialog(ns.Object("Teacher"), ns.GetCaller())
	}
}

func TeacherInit() {
	ns.SetDialog(ns.Object("Teacher"), ns.DialogYesNo, TeacherDialogueStart, TeacherDialogueEnd)
}
func TeacherDialogueStart() {
	ns.Object("Teacher").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	ns.GetCaller().LookAtObject(ns.Object("Teacher"))
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // Warrior
	case 1: // Wizard
		switch data.Quest.WizardLesson01_Quest01 {
		case 1:
			ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk01")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.WizardLesson01_Quest01 = 2
				data.Quest.JandorWizStart_Quest01 = 7
			})
		case 2:
			ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk01")
		case 3:
			ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.WizardLesson01_Quest01 = 2
			})
		case 4:
			ns.SetDialog(ns.Object("Teacher"), ns.DialogYesNo, TeacherDialogueStart, TeacherDialogueEndTwo)
			ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk03")
		case 10:
			ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk05")
		}
	case 2: // Conjurer
	}
}

func TeacherDialogueEndTwo() {
	switch ns.GetAnswer(ns.Object("Teacher")) {
	case ns.AnswerGoodbye:
		TeacherInit()
		if InversionLessonActive {
			InversionLesson()
		}
	case ns.AnswerNo:
		ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
		ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk04a")
	case ns.AnswerYes:
		ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
		ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk02")
		InversionLessonActive = true
		InversionTarget = ns.GetCaller()
	}
}

func TeacherDialogueEnd() {
	switch ns.GetAnswer(ns.Object("Teacher")) {
	case ns.AnswerGoodbye:
		TeacherInit()
		if InversionLessonActive {
			InversionLesson()
		}
	case ns.AnswerNo:
	case ns.AnswerYes:
		if ns.GetCaller().GetGold() < 100 {
			ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk04")
		} else {
			if InversionLessonActive {
				ns.AudioEvent(audio.HumanMaleEatFood, ns.GetCaller())
				ns.Object("Teacher").ChatStr("Wait your turn, apprentice.")
			} else {
				InversionLessonActive = true
				ns.GetCaller().ChangeGold(-100)
				ns.GetCaller().AwardSpell(spell.INVERSION)
				ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
				ns.SetDialog(ns.Object("Teacher"), ns.DialogNormal, TeacherDialogueStart, TeacherDialogueEnd)
				ns.TellStory(audio.SwordsmanHurt, "Con08b:InversionBoyTalk02")
				nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
					data.Quest.WizardLesson01_Quest01 = 3
				})
				InversionTarget = ns.GetCaller()
			}
		}
	}
}

func InversionLesson() {
	ns.Object("InversionGate").Lock(true)
	ns.Object("ApprenticeTeacher").SetPos(ns.Waypoint("OutFence").Pos())
	ns.Object("ApprenticeTeacher").Guard(ns.Waypoint("OutFence").Pos(), ns.Waypoint("OutFence").Pos(), 10)
	InversionTarget.SetPos(ns.Waypoint("InFence").Pos())
	ns.CastSpell(spell.TELEPORT_TO_TARGET, InversionTarget, ns.Waypoint("InFence"))
	ns.CastSpell(spell.TELEPORT_TO_TARGET, ns.Object("ApprenticeTeacher"), ns.Waypoint("OutFence"))
	InversionTarget.Enchant(enchant.ANCHORED, ns.Infinite())
	ns.NewTimer(ns.Frames(60), func() {
		ApprenticeTeacherCastLight()
		ns.NewTimer(ns.Frames(60), func() {
			CheckInversionSuccesful()
			ApprenticeTeacherCastLight()
			ns.NewTimer(ns.Frames(60), func() {
				CheckInversionSuccesful()
				ApprenticeTeacherCastLight()
				ns.NewTimer(ns.Frames(60), func() {
					CheckInversionSuccesful()
					ApprenticeTeacherCastLight()
					ns.NewTimer(ns.Frames(30), func() {
						ApprenticeTeacherCastInversion()
					})
					ns.NewTimer(ns.Frames(90), func() {
						CheckInversionSuccesful()
						ApprenticeTeacherCastLight()
						ns.NewTimer(ns.Frames(30), func() {
							ApprenticeTeacherCastInversion()
						})
						ns.NewTimer(ns.Frames(90), func() {
							CheckInversionSuccesful()
							EndInversionLession()
						})
					})
				})
			})
		})
	})
}

func EndInversionLession() {
	InversionTarget.EnchantOff(enchant.ANCHORED)
	ns.Object("ApprenticeTeacher").Guard(ns.Waypoint("AppWork").Pos(), ns.Waypoint("AppWork").Pos(), 10)
	ns.Object("InversionGate").Lock(false)
	if InversionScore == 5 {
		ns.Object("Teacher").ChatStr("Well done!")
		nw.UpdatePlayer(InversionTarget.Player(), func(data *nw.PlayerData) {
			data.Quest.WizardLesson01_Quest01 = 10
		})
	} else {
		ns.Object("Teacher").ChatStr("I'm afraid that's not good enough! You want to try again?")
		nw.UpdatePlayer(InversionTarget.Player(), func(data *nw.PlayerData) {
			data.Quest.WizardLesson01_Quest01 = 4
		})
	}
	InversionLessonActive = false
	InversionTarget = nil
	InversionScore = 0
}

func CheckInversionSuccesful() {
	if InversionTarget.HasEnchant(enchant.LIGHT) {
		InversionTarget.EnchantOff(enchant.LIGHT)
		rnd := ns.Random(1, 5)
		switch rnd {
		case 1:
			ns.Object("Teacher").ChatStr("That's bad!")
		case 2:
			ns.Object("Teacher").ChatStr("React quicker!")
		case 3:
			ns.Object("Teacher").ChatStr("Be prepared!")
		case 4:
			ns.Object("Teacher").ChatStr("Oh boy...")
		case 5:
			ns.Object("Teacher").ChatStr("Come on!")
		}
	} else {
		rnd := ns.Random(1, 5)
		switch rnd {
		case 1:
			ns.Object("Teacher").ChatStr("Very good!")
		case 2:
			ns.Object("Teacher").ChatStr("Nice!")
		case 3:
			ns.Object("Teacher").ChatStr("Quick reflexes!")
		case 4:
			ns.Object("Teacher").ChatStr("Fine work!")
		case 5:
			ns.Object("Teacher").ChatStr("Right on!")
		}
		InversionScore++
	}
}

func ApprenticeTeacherCastLight() {
	ns.Object("ApprenticeTeacher").LookAtObject(InversionTarget)
	nw.CastPhonemes(ns.Object("ApprenticeTeacher"), []audio.Name{nw.PhLeft, nw.PhRight, nw.PhUp}, func() {
		ns.CastSpell(spell.LIGHT, ns.Object("ApprenticeTeacher"), InversionTarget)
	})
}

func ApprenticeTeacherCastInversion() {
	ns.Object("ApprenticeTeacher").LookAtObject(InversionTarget)
	nw.CastPhonemes(ns.Object("ApprenticeTeacher"), []audio.Name{nw.PhUpLeft, nw.PhUpRight}, func() {
		ns.CastSpell(spell.INVERSION, ns.Object("ApprenticeTeacher"), InversionTarget)
	})
}

func HorvathInit() {
	ns.SetDialog(ns.Object("Horvath"), ns.DialogYesNo, HorvathDialogueStart, HorvathDialogueEnd)
}
func HorvathDialogueStart() {
	ns.Object("Horvath").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := nw.LoadPlayer(p)
	switch p.Unit().GetClass() {
	case 0: // Warrior
	case 1: // Wizard
		switch data.Quest.JandorWizStart_Quest01 {
		case 0:
		case 1:
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNext, HorvathDialogueStart, HorvathDialogueStart)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk01")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 2
			})
		case 2:
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNext, HorvathDialogueStart, HorvathDialogueStart)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk02")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 3
			})
			ns.GetCaller().AwardSpell(spell.LESSER_HEAL)
			ns.GetCaller().AwardSpell(spell.MAGIC_MISSILE)
		case 3:
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNormal, HorvathDialogueStart, HorvathDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk02a")
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 4
			})
		case 4:
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNormal, HorvathDialogueStart, HorvathDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk03")
		case 5:
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNormal, HorvathDialogueStart, HorvathDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk04")
			ns.AudioEvent(audio.JournalEntryAdd, ns.GetCaller())
			nw.UpdatePlayer(ns.GetCaller().Player(), func(data *nw.PlayerData) {
				data.Quest.JandorWizStart_Quest01 = 6
				data.Quest.WizardLesson01_Quest01 = 1
				HorvathInLab = true
			})
		case 6:
			ns.SetDialog(ns.Object("Horvath"), ns.DialogNormal, HorvathDialogueStart, HorvathDialogueEnd)
			ns.TellStory(audio.SwordsmanHurt, "Wiz01A.scr:HorvathTalk05")
			HorvathInLab = false
		case 7:
			ns.AudioEvent(audio.HorvathTalkable, ns.GetCaller())
			ns.Object("Horvath").ChatStr("Complete your Inversion lesson before you come to see me again.")
		case 10:
			switch data.Quest.WizardLesson01_Quest01 {
			case 10:
				switch data.Quest.FirstTaskAsWizardApprentice_Quest01 {
				case 0:
				case 1:
				}
			}
		}
	case 2: // Conjurer
	}

}
func HorvathDialogueEnd() {
	switch ns.GetAnswer(ns.Object("Horvath")) {
	case ns.AnswerGoodbye:
		if HorvathInLab {
			ns.Object("Horvath").Guard(ns.Waypoint("HorvathLab").Pos(), ns.Waypoint("HorvathLab").Pos(), 10)
		} else {
			ns.Object("Horvath").Guard(ns.Waypoint("HorvathOffice").Pos(), ns.Waypoint("HorvathOffice").Pos(), 10)
		}
	case ns.AnswerNo:
	case ns.AnswerYes:
	}
}

func TowerOfIllusion01ToGalava() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapGalava, nw.GoToMapOptions{
		Exit: "TowerOfIllusion01ToGalava",
	})
}

func TowerOfIllusion01ToTowerOfIllusion02() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapTowerOfIllusion02, nw.GoToMapOptions{
		Exit: "TowerOfIllusion01ToTowerOfIllusion02",
	})
}
