package usecases_test

import (
	"reflect"
	"testing"

	uc "github.com/devpablocristo/demo/coupon/usecases"
)

func TestMaximizeTotal(t *testing.T) {

	itemsList1 := map[string]int{
		"MLA1": 100,
		"MLA2": 210,
		"MLA3": 260,
		"MLA4": 80,
		"MLA5": 90,
	}

	wantList1 := []string{"MLA1", "MLA2", "MLA4", "MLA5"}
	gotList1 := uc.MaximizeTotal(500, itemsList1)

	if !reflect.DeepEqual(wantList1, gotList1) {
		t.Fatalf("\nWant: %q.\nGot: %q", wantList1, gotList1)
	}

}
