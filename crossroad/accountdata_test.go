package noxworld

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/noxworld-dev/noxscript/ns/v4"
	_ "github.com/noxworld-dev/noxscript/ns/v4/test"
)

func TestAccountData(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		data, err := json.Marshal(MyAccountData{})
		if err != nil {
			t.Fatal(err)
		}
		var v MyAccountData
		if err = json.Unmarshal(data, &v); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("unexported", func(t *testing.T) {
		checkUnexported(t, reflect.TypeOf(MyAccountData{}), "")
	})
	t.Run("save-and-load", func(t *testing.T) {
		pl := ns.HostPlayer()

		data := loadMyQuestData(pl)
		data.Character.Name = "Test"
		data.Quest.Warrior.JoinTheFireKnights = QuestAccepted
		saveMyQuestData(pl, data)

		data2 := loadMyQuestData(pl)
		if data != data2 {
			t.Fatal("account save failed")
		}
	})
}

func checkUnexported(t testing.TB, rt reflect.Type, pref string) {
	switch rt.Kind() {
	case reflect.Pointer:
		checkUnexported(t, rt.Elem(), pref)
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if !f.IsExported() {
				name := rt.Name()
				if name == "" {
					name = pref
				}
				t.Errorf("unexported field: %s.%s", name, f.Name)
			}
			checkUnexported(t, f.Type, f.Name)
		}
	}
}
