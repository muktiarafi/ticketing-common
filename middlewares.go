package common

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CurrentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sub := c.Request().Header.Get("X-User")
		email := c.Request().Header.Get("X-Email")

		id, err := strconv.Atoi(sub)
		if err != nil {
			return &Error{
				Op:      "CurrentUser",
				Code:    EINVALID,
				Message: "Malformed payload",
				Err:     err,
			}
		}
		payload := &UserPayload{
			ID:    id,
			Email: email,
		}

		c.Set("userPayload", payload)

		return next(c)
	}
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		const op = "RequireAuth"
		if err != nil {
			return &Error{
				Op:      op,
				Code:    EINVALID,
				Message: "Missing Cookie",
				Err:     err,
			}
		}

		token, payload, err := ParseToken(cookie.Value)
		if err != nil {
			return &Error{Op: op, Err: err}
		}

		if !token.Valid {
			return &Error{
				Op:      op,
				Code:    EINVALID,
				Message: "Invalid token",
				Err:     errors.New("invalid token"),
			}
		}

		c.Set("userPayload", payload)

		return next(c)
	}
}
