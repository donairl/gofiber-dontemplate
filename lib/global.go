package lib

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func SetStore(store *session.Store) {
	Store = store
}

// gets the session from the store
func GetSession(ctx *fiber.Ctx) (*session.Session, error) {
	return Store.Get(ctx)
}
