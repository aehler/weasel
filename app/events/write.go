package business_events

import (
	"github.com/akdcode/srm-lib/discovery"
	procClient "github.com/akdcode/srm-services/procedure/client"
	gpzClient "github.com/akdcode/srm-services/gpz/client"
	contractClient "github.com/akdcode/srm-services/contract"
	"github.com/akdcode/srm-lib/rpc/client"
	"github.com/akdcode/srm-lib/crypto/url"
	"github.com/akdcode/srm-lib/business_events/event"
	"github.com/akdcode/srm-lib/business_events/worker"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"errors"
	"fmt"
)

type Write struct {
	db 			func() *sqlx.DB
	l 			*discovery.Locator
	c 			map[string]client.ObjectClient
	objLinks 	map[string]string
}

func NewWriter(db func() *sqlx.DB, l *discovery.Locator) worker.Actioner {

	return &Write{
		db : db,
		l : l,
		c : map[string]client.ObjectClient{
			BEEntityGPZ : gpzClient.NewClient(*l),
			BEEntityProcedure : procClient.NewClient(*l),
			BEEntityContract : contractClient.NewClient(*l),
		},
		objLinks : map[string]string{
			BEEntityGPZ : "/gpz/edit_item/",
			BEEntityProcedure : "/procedure/view/",
			BEEntityContract : "/contract/view/",
		},
	}

}

func (w *Write) Action(e event.Event) error {

	dbr := 0

	if entities[e.Object] == "" {
		return errors.New("BE entity not found")
	}

	obj, err := w.c[e.Object].GetObject(e.ObjectId, e.User.UserID)

	if err != nil {

		fmt.Println("Couldn't fetch object")

		return err
	}

	e.EventData["user"] = e.User

	e.EventData["link"] = fmt.Sprintf("%s%s", w.objLinks[e.Object], url.EncryptURL(fmt.Sprintf("%d", e.ObjectId)))

	e.EventData["number"] = obj.Entity_number

	if e.Object == "gpz" {

		e.EventData["number"] = fmt.Sprintf("%s/%s", obj.GpzNumber, obj.Entity_number)

		e.EventData["tender_object"] = obj.Meta_info["form_values"].(map[string]interface {})["tender_object"].(map[string]interface {})["v"]

	} else {

		e.EventData["number"] = obj.Entity_number

		e.EventData["tender_object"] = obj.Meta_info["tender_object"].(map[string]interface {})["v"]
	}

	ed, err := json.Marshal(e.EventData)
	if err != nil {

		fmt.Println("Couldn't pack data")

		return err
	}

	if dberr := w.db().Get(&dbr, "select * from personal_tasks.create_new_alert($1, $2, $3, $4)",
		string(e.BusinessEvent),
		e.Object,
		e.ObjectId,
		string(ed),
	); dberr != nil {

		fmt.Println("Couldn't write business event", e, dberr)

		return dberr

	}

//	fmt.Println("alert might have been set", string(ed))

	return nil

}
