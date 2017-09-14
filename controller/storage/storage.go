package storage

import (
	"weasel/app"
	"weasel/app/crypto"
	"weasel/lib/storage"
	"weasel/lib/auth"
	"fmt"
)

func file(c *app.Context) {

	user := c.Get("user").(auth.User)

	id, _ := crypto.DecryptUrl(c.Params.ByName("fileId"))

	fmt.Println(user, id)
}

func fileForm(c *app.Context) {

	user := c.Get("user").(auth.User)

	storage.Write(user.OrganizationId)

}
