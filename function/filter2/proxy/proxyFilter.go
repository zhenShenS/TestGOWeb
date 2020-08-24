package proxyFilter

import (
	"TestGOWeb/function/filter"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProxyFilter struct {
	ProxyFilter filterT.FilterT
}

func (pf *ProxyFilter) TestAccount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("被过滤方法 前")
	pf.ProxyFilter.TestAccount(w, r, p)
	fmt.Println("被过滤方法后")
}
