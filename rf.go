// Russian Federation specific

package ru

const (
	Nm = 1 << iota // Именительный падеж
	Gn             // Родительный падеж
	Dt             // Дательныый падеж
	Ac             // Винительный падеж
	In             // Творительный падеж
	Lc             // Предложный падеж
	Sn             // Единственное число
	Pl             // Множественное число
	Ms             // Мужской род
	Fm             // Женский род
	Nt             // Средний род
	AB             // Абревиатура или сокращенное название
)

const (
	Gender = Ms + Fm + Nt
	Plural = Sn + Pl
	Nmcase = Nm + Gn + Dt + Ac + In + Lc
)
