package filter

import (
	"cn.anydevelop/go_recruit/common"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Admin token filter
func AdminTokenFilter(ctx *context.Context) {
	if tokenStr := ctx.Request.Header.Get("Admin-Token"); tokenStr == "" {
		beego.Debug("Admin filter is execute!")
		marshal, _ := json.Marshal(common.Fail("Without Token!"))
		ctx.Output.Body(marshal)
	}
}
