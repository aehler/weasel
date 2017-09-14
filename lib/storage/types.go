package storage

import (
	"weasel/lib/auth"
)

type File struct {
	Name string
	HashName string
	Bucket bucket
	MD5 string
	Path string
	AVCheck bool
	AVMessage string
	Size uint
	ContentType string
	Meta string
	Version uint
	Entity string
	EntityId uint
	Owner *auth.User
}

type Files []File

/* required for organization
Свидетельство о регистрации компании 	Может быть не одно
2 	Свидетельство о постановки на налоговый учет 	Может быть не одно
3 	Устав
4 	Выписка из ЕГРЮЛ 	Контроль срока действия
5 	Приказ о назначении ГД
6 	Доверенности на уполномоченных лиц
7 	Бенефициары 	Документы по бенефициарам (выписки, учредительные документы)
8 	МСП
9 	Лицензии и допуски
10 	Бухгалтерская отчетность
11 	Штатное расписание (при необходимости) 	При необходимости формируется база на персонал с приложением договоров и трудовых книжек
12 	Уведомление о применении УСН
13 	Документы по опыту 	Договоры, акты, реестры
14 	Документы по МТР 	Договоры, бухгалтерские справки
	*/
