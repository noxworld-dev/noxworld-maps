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

type NoxWorldData struct {
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
		// ix
		TalkToTheMayor_Quest01                    QuestStatus
		BecomeConjurerer_Quest01                  QuestStatus
		MayorsScepter_Quest01                     QuestStatus
		AldwynGiveScrollWar_MayorsScepterQuest_01 QuestStatus
		TroubleAtTheManaMines_Quest01             QuestStatus
		SpidersAtTheMayor_Quest01                 QuestStatus
		MorganConManSellBow_Quest01               QuestStatus
		TestQuestStatus_Quest01                   QuestStatus
		TalkToAldwyn_Quest01                      QuestStatus
		ArcheryContest_Quests01                   QuestStatus
		ArcheryContestScore_Quest01               QuestStatus
		HenrickCongrats_Quest01                   QuestStatus
	}
}

func loadMyNoxWorldData(pl ns.Player) NoxWorldData {
	var data NoxWorldData
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Get("accountdata", &data)
	if err != nil {
		fmt.Println("cannot read data:", err)
	}
	return data
}

func saveMyNoxWorldData(pl ns.Player, data NoxWorldData) {
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Set("accountdata", &data)
	if err != nil {
		fmt.Println("cannot save data:", err)
	}
}

func updateNoxWorldData(pl ns.Player, fnc func(data *NoxWorldData)) {
	data := loadMyNoxWorldData(pl)
	fnc(&data)
	saveMyNoxWorldData(pl, data)
}
