//

package ru

import (
	//	"time"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Интервалы времени выраженные в минутах

const ( // Базовые интнрвалы относительно начала котороых измеряются относительные интервалы в минутах
	RLHour = iota + 1
	RLDay
	RLWeek
	RLMonth
	RLYear
	RLCenture
)

const ( // Максимальная емкость базовых интервалов в минутах
	CapHour  = 60
	CapDay   = CapHour * 24
	CapWeek  = CapDay * 7
	CapMonth = CapDay * 31
	CapYear  = CapDay * 366
)

type HmBase int                          // Тип базового интнрвала
type HmPos uint32                        // Смещение относительно начала базового интервала в минутах
type HmInterval struct{ from, to HmPos } // Относительный интервал
type HmService struct {                  // Режим обслуживания в базовом интервале, как список относительных интервалов
	base HmBase
	list *[]HmInterval
}

func errMinute(v int) (err error) {
	if v < 0 || v > 59 {
		err = fmt.Errorf("Недопустимое значение минут %d", v)
	}
	return
}

func errHour(v int) (err error) {
	if v < 0 || v > 23 {
		err = fmt.Errorf("Недопустимое значение часов %d", v)
	}
	return
}

func errWeekDay(v int) (err error) {
	if v < 0 || v > 6 {
		err = fmt.Errorf("Недопустимое значение номера дня недели %d", v)
	}
	return
}

func NewWeekService() *HmService {
	list := make([]HmInterval, 0, 2)
	return &HmService{RLWeek, &list}
}

func (h *HmService) addWeekService(h1, m1, h2, m2 int, wdl ...int) {
	for _, wd := range wdl {
		from := HmPos(wd*CapDay + h1*CapHour + m1)
		to := HmPos(wd*CapDay + h2*CapHour + m2)
		*h.list = append(*h.list, HmInterval{from, to})
	}
}

func (h *HmService) OnServiceNow() bool { return h.OnService(time.Now()) }

func (h *HmService) OnService(t time.Time) bool {
	var pos int = -1
	switch h.base {
	case RLHour:
		pos = t.Minute()
	case RLDay:
		pos = t.Hour()*CapHour + t.Minute()
	case RLWeek:
		pos = (int(t.Weekday())+6)%7*CapDay + t.Hour()*CapHour + t.Minute()
	case RLMonth:
		pos = (t.Day()-1)*CapDay + t.Hour()*CapHour + t.Minute()
	case RLYear:
		pos = (t.YearDay()-1)*CapDay + t.Hour()*CapHour + t.Minute()
	}
	//	fmt.Println(pos, t.Day(), t.YearDay())
	return h.posOnService(HmPos(pos))

}

func (h *HmService) posOnService(pos HmPos) bool {
	//        fmt.Println(pos,"...")
	n := len(*h.list)
	c := sort.Search(n, func(i int) bool { return (*h.list)[i].from > pos })
	if c == 0 {
		return false
	}
	cp := (*h.list)[c-1]
	//	fmt.Println(c-1,cp)
	return cp.from <= pos && cp.to >= pos
}

// Emit

func (h HmInterval) DownToDay() (HmInterval, HmPos) {
	var d HmPos = CapDay
	return HmInterval{h.from % d, h.to % d}, h.from / d
}

type HmExt struct {
	HmInterval
	list *[]HmPos
}

type HmExtList []HmExt

func (h *HmExtList) add(v HmInterval, p HmPos) {

	for i, c := range *h {
		if c.from == v.from && c.to == v.to {
			*c.list = append(*c.list, p)
			(*h)[i] = c
			return
		}
	}
	l := make([]HmPos, 1)
	l[0] = p
	c := HmExt{v, &l}
	*h = append(*h, c)
	return
}

func (h *HmService) HmString() string {
	s := make(HmExtList, 0, 1)
	for _, c := range *h.list {
		(&s).add(c.DownToDay())
	}
	ss := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ss[i] = HmWeekList(*s[i].list...) + " " + HmHours(s[i].HmInterval)
	}
	return strings.Join(ss, ", ")
}

// Parse

func (h *HmService) HmSet(s string) error {
	return nil
}
