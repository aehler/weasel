package organizations

import (
	"time"
	"encoding/json"
	"errors"
	"weasel/app/grid"
)

type Organization struct {
	ID uint `weaselform:"hidden" formLabel:""`
	Fullname string `weaselform:"textarea" formLabel:"Полное наименование"`
	Shortname string `weaselform:"text" formLabel:"Сокращенное наименование"`
	IsIB bool `weaselform:"checkbox" formLabel:"Индивидуальный предприниматель"`	//В случае проставления галочки ИП появляютяс дополнительный поля
	LicenceSeries string `weaselform:"text" formLabel:"Серия свидетельства"`
	LicenceNumber string `weaselform:"text" formLabel:"Номер свидетельства"`
	LicenceIssuer string `weaselform:"textarea" formLabel:"Кем выдано"`
	LicenceIssued time.Time `weaselform:"date" formLabel:"Дата выдачи"`
	Passport string `weaselform:"textarea" formLabel:"Пасспортные данные"`
	INN string `weaselform:"inn" formLabel:"ИНН"`
	KPP string `weaselform:"kpp" formLabel:"КПП"`
	JurAddress string `weaselform:"textarea" formLabel:"Адрес юридический (почтовый)"`
	PhysAddress string `weaselform:"textarea" formLabel:"Адрес местонахождения"`
	Phones string `weaselform:"phone" formLabel:"Телефон"`
	Emails string `weaselform:"email" formLabel:"Электронная почта"`
	WebSite string `weaselform:"text" formLabel:"Сайт"`
	OGRN string `weaselform:"text" formLabel:"ОГРН"`
	OKPO string `weaselform:"text" formLabel:"ОКПО"`
	OKVED []string `weaselform:"okved" formLabel:"ОКВЭД"`
	Rs string `weaselform:"number" formLabel:"Р/С"`
	BIK string `weaselform:"bic" formLabel:"БИК"`
	Bank string `weaselform:"text" formLabel:"Банк"`
	Ks string `weaselform:"number" formLabel:"К/С"`
	Chairmen string `weaselform:"textarea" formLabel:"Руководители организации"`
	Accountant string `weaselform:"textarea" formLabel:"Главный бухгалтер"`
	TaxSystem string `weaselform:"text" formLabel:"Система налогообложения"`
	IsMSP bool `weaselform:"checkbox" formLabel:"Принадлежность к МСП"`

}

type StoredOrganization struct {
	ID uint `db:"id"`
	INN string `db:"inn"`
	KPP string `db:"kpp"`
	OrganizationID uint `db:"organization_id"`
	UserID uint `db:"user_id"`
	Data Organization `db:"meta_info"`
}

type OrganizationRow struct {
	Items []*StoredOrganization
	grid.Grider
}

func (o *Organization) Scan(src interface {}) error {

	var source []byte

	switch src.(type) {

	case string:

		source = []byte(src.(string))

	case []byte:

		source = src.([]byte)

	default:

		return errors.New("Incompatible type for organization")
	}

	if err := json.Unmarshal(source, &o); err != nil {

		return err
	}

	return nil
}
