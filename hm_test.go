// ЭТО ТЕСТЫ

package ru

import (
	"fmt"
	"testing"
)

func TestHM(t *testing.T) {

	fmt.Println("Test HM:", RLHour, RLDay, RLWeek, RLYear, CapWeek, CapYear)
/*
	s := "   :   iii "
	fmt.Println(SkipWSchr(&s, ",.:"), "["+s+"]")
	fmt.Println("[" + ScanToken(&s) + "]")

	fmt.Println(TreatHI("9.01- 17= 01"))
*/
	r := NewWeekService()
	r.addWeekService(9, 0, 17, 0, 1, 2, 3, 4)
//	r.addWeekService(9, 0, 17, 0, 1)
	r.addWeekService(9, 0, 16, 0, 5)
	
	fmt.Println(r, *r.list)
	rHm := r.HmString()
	fmt.Println(rHm)
	r2,err := TreatWeekService(rHm)
        fmt.Println(err)
        fmt.Println(r2,*r2.list)
//	fmt.Println(r, *r.list, r.posOnService(1979), r.posOnService(2000), r.posOnService(2461), r.posOnService(6780), r.OnServiceNow())
	fmt.Println(HmDayList(0, 1, 2, 4, 6))
	fmt.Println(HmWeekList())
	fmt.Println(HmWeekList(0))
	fmt.Println(HmHours((*r.list)[0]))
	fmt.Println(TreatWeekDay("пп"))
}
