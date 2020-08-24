package filterT2

import (
	"TestGOWeb/function/filter"
	"TestGOWeb/function/filter2/proxy"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func TestAccount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a := filterT.New()
	a.TestAccount(w, r, p)

}

func init() {
	filterT.New = func() filterT.FilterT {
		a := &filterT.FilterTlmpl{}
		p := &proxyFilter.ProxyFilter{a}
		return p
	}
}
