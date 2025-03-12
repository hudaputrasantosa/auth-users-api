package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberResponseWriter struct {
	Ctx *fiber.Ctx
}

func (f *FiberResponseWriter) Header() http.Header {
	return http.Header{}
}

func (f *FiberResponseWriter) Write(b []byte) (int, error) {
	return f.Ctx.Write(b)
}

func (f *FiberResponseWriter) WriteHeader(statusCode int) {
	f.Ctx.Status(statusCode)
}
