package teawaf

import (
	"github.com/TeaWeb/code/teawaf/requests"
	"net/http"
)

type LogAction struct {
}

func (this *LogAction) Perform(waf *WAF, request *requests.Request, writer http.ResponseWriter) (allow bool) {
	return true
}
