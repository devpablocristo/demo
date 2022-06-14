package usecases_test

import (
	"reflect"
	"testing"

	uc "github.com/devpablocristo/demo/coupon/usecases"
)

func TestMaximizeTotal(t *testing.T) {

	t.Run("given items list", func(t *testing.T) {

		itemsList := map[string]float32{
			"MLA1": 100,
			"MLA2": 210,
			"MLA3": 260,
			"MLA4": 80,
			"MLA5": 90,
		}
		amount := float32(500)
		wantList := []string{"MLA1", "MLA2", "MLA4", "MLA5"}
		gotList := uc.MaximizeTotal(amount, itemsList)

		if !reflect.DeepEqual(wantList, gotList) {
			t.Fatalf("\nWant: %q.\nGot: %q", wantList, gotList)
		}

		t.Logf("Want: %q", wantList)
		t.Logf("Got: %q", gotList)
	})

	t.Run("new items list ", func(t *testing.T) {

		itemsList := map[string]float32{
			"MLA122": 100,
			"MLA132": 200,
			"MLA43":  23,
			"MLA42":  84,
			"MLA5":   91.23,
			"MLA9":   92.20,
			"MLA15":  191.12,
			"MLA225": 911.23,
			"MLA67":  123.90,
			"MLA512": 19,
		}
		amount := float32(466)
		wantList := []string{"MLA43", "MLA67", "MLA122", "MLA132", "MLA512"}
		gotList := uc.MaximizeTotal(amount, itemsList)

		if !reflect.DeepEqual(wantList, gotList) {
			t.Fatalf("\nWant: %q.\nGot: %q", wantList, gotList)
		}

		t.Logf("Want: %q", wantList)
		t.Logf("Got: %q", gotList)
	})

}

func BenchmMaximizeTotal(b *testing.B) {

	b.Run("test given example ", func(b *testing.B) {
		itemsList1 := map[string]float32{
			"MLA1": 100,
			"MLA2": 210,
			"MLA3": 260,
			"MLA4": 80,
			"MLA5": 90,
		}

		for i := 0; i < b.N; i++ {
			uc.MaximizeTotal(500, itemsList1)
		}
	})
}
