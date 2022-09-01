package cookie

import "github.com/valyala/fasthttp"

func CreateCookie(exp int, value string) fasthttp.Cookie {
	cookie := fasthttp.Cookie{}
	cookie.SetKey("cookie")
	cookie.SetValue(value)
	cookie.SetMaxAge(exp)
	//cookie.SetSecure(true)
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")
	cookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	return cookie
}
