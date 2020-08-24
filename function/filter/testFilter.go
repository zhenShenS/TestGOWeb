package filter

import (
	"TestGOWeb/function/filter/account"
	"TestGOWeb/function/filter/proxy"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func TestAccount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)
	id := queryForm.Get("id")
	a := account.New(id, "ZhangSan", 100)
	a.Query(id)
	a.Update(id, 500)

}

func init() {
	account.New = func(id, name string, value int) account.Account {
		a := &account.AccountImpl{id, name, value}
		p := &proxy.Proxy{a}
		return p
	}

}
