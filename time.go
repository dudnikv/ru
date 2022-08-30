//

package ru

import (
	"fmt"
	"time"
)

// Date(t) = DateFormat("d-m-Y",t)
func Date(t time.Time) string { return fmt.Sprintf("%02d-%02d-%d", t.Day(), int(t.Month()), t.Year()) }

// DateTime(t) = DateFormat("d-m-Y H:i",t)
func DateTime(t time.Time) string {
	return fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), int(t.Month()), t.Year(), t.Hour(), t.Minute())
}

var longDayNames = []string{
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
	"Воскресенье",
}

var shortDayNames = []string{
	"Пн",
	"Вт",
	"Ср",
	"Чт",
	"Пт",
	"Сб",
	"Вс",
}

var shortMonthNames = []string{
	"ЯНВ",
	"ФЕВ",
	"МАР",
	"АПР",
	"МАЙ",
	"ИЮН",
	"ИЮЛ",
	"АВГ",
	"СЕН",
	"ОКТ",
	"НОЯ",
	"ДЕК",
}

var longMonthNames = []string{
	"Январь",
	"Февраль",
	"Март",
	"Апрель",
	"Май",
	"Июнь",
	"Июль",
	"Август",
	"Сентябрь",
	"Октябрь",
	"Ноябрь",
	"Декабрь",
}

var genCaseMonthNames = []string{
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

var amPM = []string{
	"am",
	"pm",
	"AM",
	"PM",
}

func DayName(d int, f int) string {
	if d < 0 || d > 6 {
		return ""
	}
	switch f {
	case AB:
		return shortDayNames[d]
	case Nm:
		return longDayNames[d]
	}
	return ""
}

func MonthName(d int, f int) string {
	if d < 1 || d > 12 {
		return ""
	}
	switch f {
	case AB:
		return shortMonthNames[d-1]
	case Nm:
		return longMonthNames[d-1]
	case Gn:
		return genCaseMonthNames[d-1]
	}
	return ""
}

func TimeFormat(format string, t time.Time) string {

	var fmts []byte
	var args []interface{}
	var slash bool = false

	for _, b := range []byte(format) {
		if slash {
			fmts = append(fmts, b)
			slash = false
			continue
		}
		switch b {
		case '\\':
			slash = true
		case '%':
			fmts = append(fmts, "%%"...)
		case 'd':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Day())
		case 'D':
			fmts = append(fmts, "%s"...)
			args = append(args, shortDayNames[(int(t.Weekday())+6)%7])
		case 'j':
			fmts = append(fmts, "%d"...)
			args = append(args, t.Day())
		case 'l':
			fmts = append(fmts, "%s"...)
			args = append(args, longDayNames[(int(t.Weekday())+6)%7])
		case 'N':
			fmts = append(fmts, "%d"...)
			args = append(args, (int(t.Weekday())+6)%7+1)
		case 'z':
			fmts = append(fmts, "%d"...)
			args = append(args, t.YearDay()-1)
		case 'W':
			_, week := t.ISOWeek()
			fmts = append(fmts, "%d"...)
			args = append(args, week)
		case 'F':
			fmts = append(fmts, "%s"...)
			args = append(args, longMonthNames[int(t.Month())-1])
		case 'f':
			fmts = append(fmts, "%s"...)
			args = append(args, genCaseMonthNames[int(t.Month())-1])
		case 'M':
			fmts = append(fmts, "%s"...)
			args = append(args, shortMonthNames[int(t.Month())-1])
		case 'm':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Month())
		case 'n':
			fmts = append(fmts, "%d"...)
			args = append(args, t.Month())
		case 'Y':
			fmts = append(fmts, "%d"...)
			args = append(args, t.Year())
		case 'y':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Year()%100)
		case 'a':
			h := t.Hour()
			fmts = append(fmts, "%s"...)
			args = append(args, amPM[(h-h%12)/12])
		case 'A':
			h := t.Hour()
			fmts = append(fmts, "%s"...)
			args = append(args, amPM[(h-h%12)/12+2])
		case 'G':
			fmts = append(fmts, "%d"...)
			args = append(args, t.Hour())
		case 'H':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Hour())
		case 'g':
			fmts = append(fmts, "%d"...)
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			args = append(args, h)
		case 'h':
			fmts = append(fmts, "%02d"...)
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			args = append(args, h)
		case 'i':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Minute())
		case 's':
			fmts = append(fmts, "%02d"...)
			args = append(args, t.Second())
		case 'u':
			nano := t.UnixNano() % 1000000000
			fmts = append(fmts, "%06d"...)
			args = append(args, nano/1000)
		case 'v':
			nano := t.UnixNano() % 1000000000
			fmts = append(fmts, "%03d"...)
			args = append(args, nano/1000000)
		case 'T':
			name, _ := t.Zone()
			fmts = append(fmts, "%s"...)
			args = append(args, name)
		case 'Z':
			_, offset := t.Zone()
			fmts = append(fmts, "%d"...)
			args = append(args, offset)
		default:
			fmts = append(fmts, b)
		}
	}
	return fmt.Sprintf(string(fmts), args...)
}

/*  Взято из PHP (https://www.php.net/manual/ru/datetime.format.php)
    День 	--- 	---
    d 	День месяца, 2 цифры с ведущим нулём 	от 01 до 31
    D 	Текстовое представление дня недели, 2 символа 	от Пн до Вс	// !Изменено (3->2)
    j 	День месяца без ведущего нуля 	от 1 до 31
    l   Полное наименование дня недели 	от (понедельник) до (воскресенье)
    N 	Порядковый номер дня недели в соответствии со стандартом ISO-8601 от 1 (понедельник) до 7 (воскресенье)
    w 	Порядковый номер дня недели 	от 0 (воскресенье) до 6 (суббота)
    z 	Порядковый номер дня в году (начиная с 0) 	От 0 до 365
    Неделя 	--- 	---
    W 	Порядковый номер недели года в соответствии со стандартом ISO-8601; недели начинаются с понедельника 	Например: 42 (42-я неделя года)
    Месяц 	--- 	---
    F 	Полное наименование месяца, например, Январь
    f   Наименование месяца в р.п, например 1 сентября 			// !Добавлено (не из PHP)
    m 	Порядковый номер месяца с ведущим нулём от 01 до 12
    M 	Сокращенное наименование месяца, 3 символа 	от ЯНВ до ДЕК
    n 	Порядковый номер месяца без ведущего нуля 	от 1 до 12
//  t 	Количество дней в указанном месяце 	от 28 до 31
    Год 	--- 	---
//  L 	Признак високосного года 	1, если год високосный, иначе 0.
    Y 	Порядковый номер года, 4 цифры 	Примеры: 1999, 2003
    y 	Номер года, 2 цифры 	Примеры: 99, 03
    Время 	--- 	---
    a 	Ante meridiem (лат. "до полудня") или Post meridiem (лат. "после полудня") в нижнем регистре 	am или pm
    A 	Ante meridiem или Post meridiem в верхнем регистре 	AM или PM
    g 	Часы в 12-часовом формате без ведущего нуля 	от 1 до 12
    G 	Часы в 24-часовом формате без ведущего нуля 	от 0 до 23
    h 	Часы в 12-часовом формате с ведущим нулём 	от 01 до 12
    H 	Часы в 24-часовом формате с ведущим нулём 	от 00 до 23
    i 	Минуты с ведущим нулём 	от 00 до 59
    s 	Секунды с ведущим нулём 	от 00 до 59
    u 	Микросекунды
    v 	Миллисекунды
    -----------------------------
    Временная зона 	--- 	---
//  e 	Идентификатор временной зоны (добавлено в PHP 5.1.0) 	Примеры: UTC, GMT, Atlantic/Azores
//  I (заглавная i) 	Признак летнего времени 	1, если дата соответствует летнему времени, 0 в противном случае.
//  O 	Разница с временем по Гринвичу без двоеточия между часами и минутами 	Например: +0200
//  P 	Разница с временем по Гринвичу с двоеточием между часами и минутами (добавлено в PHP 5.1.3) 	Например: +02:00
    T 	Аббревиатура временной зоны 	Примеры: EST, MDT ...
    Z 	Смещение временной зоны в секундах. Для временных зон, расположенных западнее UTC возвращаются отрицательные числа, а расположенных восточнее UTC - положительные. 	от -43200 до 50400
    Полная дата/время 	--- 	---
//  c 	Дата в формате стандарта ISO 8601 (добавлено в PHP 5) 	2004-02-12T15:19:21+00:00
//  r 	Дата в формате » RFC 2822 	Например: Thu, 21 Dec 2000 16:01:07 +0200
//  U 	Количество секунд, прошедших с начала Эпохи Unix (1 января 1970 00:00:00 GMT) 	Смотрите также time()

*/
