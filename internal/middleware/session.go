package middleware

import (
	"context"

	"github.com/brendonmatos/golive"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ContextKey int

const sessionID ContextKey = iota

func Session(next golive.HTTPHandlerCtx) golive.HTTPHandlerCtx {
	store := session.New()

	return func(c *fiber.Ctx, ctx context.Context) {
		// get session from storage
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		// save session
		defer sess.Save()

		ctx = context.WithValue(ctx, sessionID, sess.ID())
		next(c, ctx)
	}
}

func GetSessionID(ctx context.Context) string {
	val, _ := ctx.Value(sessionID).(string)
	return val
}
