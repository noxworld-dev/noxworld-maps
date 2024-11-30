package noxworld

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
)

type QuestStatus int

const (
	QuestInactive = QuestStatus(0)
	QuestAccepted = QuestStatus(1)
	QuestComplete = QuestStatus(10)
)

type PlayerData struct {
	Character struct {
		// General
		Registered       bool
		Name             string
		Health           int
		Mana             int
		Level            int
		Class            player.Class
		LastExitUsed     string
		WolfCompanion    int
		ItemsInInventory []ns.Obj
	}
	Quest struct {
		// Conjurer Main Questline
		JandorStartQuestCon_Quest01 QuestStatus
		BecomeConjurerer_Quest01    QuestStatus
		TalkToTheMayor_Quest01      QuestStatus
		// Ix
		HenrickCongrats_Quest01                   QuestStatus
		MayorsScepter_Quest01                     QuestStatus
		AldwynGiveScrollWar_MayorsScepterQuest_01 QuestStatus
		TroubleAtTheManaMines_Quest01             QuestStatus
		SpidersAtTheMayor_Quest01                 QuestStatus
		MorganConManSellBow_Quest01               QuestStatus
		TalkToAldwyn_Quest01                      QuestStatus
		ArcheryContest_Quests01                   QuestStatus
		ArcheryContestScore_Quest01               QuestStatus
		JandorFieldsOfValor_Quest01               QuestStatus
		// Wizard Questline
		JandorWizStart_Quest01              QuestStatus
		HorvathFindApprenticeStart_Quest01  QuestStatus
		WizardLesson01_Quest01              QuestStatus
		FirstTaskAsWizardApprentice_Quest01 QuestStatus
	}
}

func LoadPlayer(pl ns.Player) PlayerData {
	var data PlayerData
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Get("accountdata", &data)
	if err != nil {
		fmt.Println("cannot read data:", err)
	}
	return data
}

func SavePlayer(pl ns.Player, data PlayerData) {
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Set("accountdata", &data)
	if err != nil {
		fmt.Println("cannot save data:", err)
	}
}

func UpdatePlayer(pl ns.Player, fnc func(data *PlayerData)) {
	data := LoadPlayer(pl)
	fnc(&data)
	SavePlayer(pl, data)
}
