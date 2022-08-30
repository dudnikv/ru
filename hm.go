//

package ru

import (
	"fmt"
	"strconv"
	"strings"
//	"unicode"
//	"unicode/utf8"
)

type HmValue interface {
	HmString() string
	HmSet(string) error
}

var rusWeek = [7]string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}

func HmWeekList(list ...HmPos) string {
	return HmList(1, func(i HmPos) string { return rusWeek[i] }, list...)
}

func HmDayList(list ...HmPos) string {
	return HmList(3, func(i HmPos) string { return strconv.Itoa(int(i) + 1) }, list...)
}

func HmList(lim int, f func(p HmPos) string, list ...HmPos) string {

	rn := make([]string, 0)

	for i, j := 0, 0; j < len(list); i++ {

		if i > 0 && (i == len(list) || list[i] != list[i-1]+1) {
			if lim > 0 && i-j > lim {
				rn = append(rn, fmt.Sprintf("%s-%s", f(list[j]), f(list[i-1])))
			} else {
				for k := j; k < i; k++ {
					rn = append(rn, f(list[k]))
				}
			}
			j = i
		}
	}
	return strings.Join(rn, ",")
}

func HmHours(i HmInterval) string {
	h1 := i.from % CapDay / CapHour
	m1 := i.from % CapHour
	h2 := i.to % CapDay / CapHour
	m2 := i.to % CapHour
	if m1 == 0 && m2 == 0 {
		return fmt.Sprintf("%d-%d", h1, h2)
	}
	return fmt.Sprintf("%d:%02d-%d:%02d", h1, m1, h2, m2)

}

func TreatWeekDay(name string) (HmPos, error) {

	for i, s := range rusWeek {
		if strings.EqualFold(s, name) {
			return HmPos(i), nil
		}
	}
	return 0, fmt.Errorf("Недопустимое значение дня недели '%s'", name)
}

func TreatH(s string) (HmPos, error) {
	s = strings.TrimSpace(s)
	h, err := strconv.Atoi(s)
	if err != nil || h < 0 || h >= 24 {
		return 0, fmt.Errorf("Недопустимое значение часов '%s'", s)
	}
	return HmPos(h * CapHour), nil
}

func TreatM(s string) (HmPos, error) {
	s = strings.TrimSpace(s)
	h, err := strconv.Atoi(s)
	if err != nil || h < 0 || h >= 60 {
		return 0, fmt.Errorf("Недопустимое значение минут '%s'", s)
	}
	return HmPos(h), nil
}

func TreatHM(s string) (HmPos, error) {
	s = strings.TrimSpace(s)
	pos := strings.IndexAny(s, ":.-=")
	if pos == -1 {
		return TreatH(s)
	}
	h, eh := TreatH(s[:pos])
	if eh != nil {
		return h, eh
	}
	m, em := TreatM(s[pos+1:])
	if em != nil {
		return m, em
	}
	return h + m, nil

	return 0, fmt.Errorf("Недопустимое значение времени '%s'", s)
}

func TreatHI(s string) (*HmInterval, error) {
	s = strings.TrimSpace(s)

	pos := strings.IndexAny(s, "-")
	if pos == -1 {
		return nil, fmt.Errorf("Недопустимое значение диапазона времени '%s'", s)
	}
	t1, e1 := TreatHM(s[:pos])
	if e1 != nil {
		return nil, e1
	}
	t2, e2 := TreatHM(s[pos+1:])
	if e2 != nil {
		return nil, e2
	}
	return &HmInterval{t1, t2}, nil
}

func TreatWeekService(s string) (*HmService, error) {
	slist := strings.Split(s, ",")
	var i int = 0
	ilist := make([]HmInterval,0,7)
	for _, w := range slist {
		
		w = strings.TrimSpace(w)
		w1, w2, w3 := "", "", ""
		var wd1, wd2 HmPos
		var p int
		var err error
		
		p = strings.IndexByte(w, ' ')
		if p >= 0 {
			w1 = w[:p]
			w3 = w[p+1:]
		}
		p = strings.IndexByte(w1, '-')
		if p >= 0 {
			w2 = w1[p+1:]
			w1 = w1[:p]
		}
		
//		fmt.Println(w,w1,w2,w3)
		wd1, err = TreatWeekDay(w1)
		if err != nil {
			return nil, err
		}
		wd2 = wd1
		if w2 != "" {
			wd2, err = TreatWeekDay(w2)
			if err != nil {
				return nil, err
			}
		}

		for wd := wd1; wd <= wd2; wd++ {
			ilist = append(ilist, HmInterval{wd * CapDay, 0})
		}

		if w3 != "" {
			hi, err := TreatHI(w3)
			if err != nil {
				return nil, err
			}
			for j := i; j < len(ilist); j++ {
				ilist[j].to = ilist[j].from + hi.to
				ilist[j].from = ilist[j].from + hi.from
			}
			i = len(ilist)
		}
	}
	return &HmService{RLWeek, &ilist}, nil
}

