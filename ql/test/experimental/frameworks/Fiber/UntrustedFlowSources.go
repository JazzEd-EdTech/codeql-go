// Code generated by https://github.com/gagliardetto. DO NOT EDIT.

package main

import "github.com/gofiber/fiber"

// Package github.com/gofiber/fiber@v1.14.6
func UntrustedFlowSources_GithubComGofiberFiberV1146() {
	// Untrusted flow sources from method calls.
	{
		// Untrusted flow sources from method calls on github.com/gofiber/fiber.Ctx.
		{
			// func (*Ctx).BaseURL() string
			{
				var receiverCtx273 fiber.Ctx
				result982 := receiverCtx273.BaseURL()
				sink(result982) // $untrustedFlowSource
			}
			// func (*Ctx).Body() string
			{
				var receiverCtx458 fiber.Ctx
				result506 := receiverCtx458.Body()
				sink(result506) // $untrustedFlowSource
			}
			// func (*Ctx).BodyParser(out interface{}) error
			{
				var receiverCtx213 fiber.Ctx
				var paramOut468 interface{}
				receiverCtx213.BodyParser(paramOut468)
				sink(paramOut468) // $untrustedFlowSource
			}
			// func (*Ctx).Cookies(key string, defaultValue ...string) string
			{
				var receiverCtx219 fiber.Ctx
				result265 := receiverCtx219.Cookies("", "")
				sink(result265) // $untrustedFlowSource
			}
			// func (*Ctx).FormFile(key string) (*mime/multipart.FileHeader, error)
			{
				var receiverCtx971 fiber.Ctx
				result320, _ := receiverCtx971.FormFile("")
				sink(result320) // $untrustedFlowSource
			}
			// func (*Ctx).FormValue(key string) (value string)
			{
				var receiverCtx545 fiber.Ctx
				resultValue566 := receiverCtx545.FormValue("")
				sink(resultValue566) // $untrustedFlowSource
			}
			// func (*Ctx).Get(key string, defaultValue ...string) string
			{
				var receiverCtx497 fiber.Ctx
				result274 := receiverCtx497.Get("", "")
				sink(result274) // $untrustedFlowSource
			}
			// func (*Ctx).Hostname() string
			{
				var receiverCtx783 fiber.Ctx
				result905 := receiverCtx783.Hostname()
				sink(result905) // $untrustedFlowSource
			}
			// func (*Ctx).Method(override ...string) string
			{
				var receiverCtx389 fiber.Ctx
				result198 := receiverCtx389.Method("")
				sink(result198) // $untrustedFlowSource
			}
			// func (*Ctx).MultipartForm() (*mime/multipart.Form, error)
			{
				var receiverCtx477 fiber.Ctx
				result544, _ := receiverCtx477.MultipartForm()
				sink(result544) // $untrustedFlowSource
			}
			// func (*Ctx).OriginalURL() string
			{
				var receiverCtx382 fiber.Ctx
				result715 := receiverCtx382.OriginalURL()
				sink(result715) // $untrustedFlowSource
			}
			// func (*Ctx).Params(key string, defaultValue ...string) string
			{
				var receiverCtx179 fiber.Ctx
				result366 := receiverCtx179.Params("", "")
				sink(result366) // $untrustedFlowSource
			}
			// func (*Ctx).Path(override ...string) string
			{
				var receiverCtx648 fiber.Ctx
				result544 := receiverCtx648.Path("")
				sink(result544) // $untrustedFlowSource
			}
			// func (*Ctx).Query(key string, defaultValue ...string) string
			{
				var receiverCtx754 fiber.Ctx
				result680 := receiverCtx754.Query("", "")
				sink(result680) // $untrustedFlowSource
			}
			// func (*Ctx).QueryParser(out interface{}) error
			{
				var receiverCtx722 fiber.Ctx
				var paramOut506 interface{}
				receiverCtx722.QueryParser(paramOut506)
				sink(paramOut506) // $untrustedFlowSource
			}
			// func (*Ctx).Range(size int) (rangeData Range, err error)
			{
				var receiverCtx121 fiber.Ctx
				resultRangeData293, _ := receiverCtx121.Range(0)
				sink(resultRangeData293) // $untrustedFlowSource
			}
			// func (*Ctx).Subdomains(offset ...int) []string
			{
				var receiverCtx151 fiber.Ctx
				result849 := receiverCtx151.Subdomains(0)
				sink(result849) // $untrustedFlowSource
			}
		}
	}
	// Untrusted flow sources from struct fields.
	{
		// Untrusted flow sources from github.com/gofiber/fiber.Cookie struct fields.
		{
			structCookie322 := new(fiber.Cookie)
			sink(
				structCookie322.Domain,   // $untrustedFlowSource
				structCookie322.Name,     // $untrustedFlowSource
				structCookie322.Path,     // $untrustedFlowSource
				structCookie322.SameSite, // $untrustedFlowSource
				structCookie322.Value,    // $untrustedFlowSource
			)
		}
		// Untrusted flow sources from github.com/gofiber/fiber.Error struct fields.
		{
			structError339 := new(fiber.Error)
			sink(structError339.Message) // $untrustedFlowSource
		}
	}
}
