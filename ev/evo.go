package ev

import (
	"github.com/andypangaribuan/evo-golang/v-db"
	"github.com/andypangaribuan/evo-golang/v-ext"
	"github.com/andypangaribuan/evo-golang/v-json"
	"github.com/andypangaribuan/evo-golang/v-log"
	"github.com/andypangaribuan/evo-golang/v-utils"
	"github.com/andypangaribuan/evo-golang/v-var"
	"github.com/andypangaribuan/evo-golang/vi"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */

//noinspection GoUnusedGlobalVariable
var Conf *v_ext.VSConf
//noinspection GoUnusedGlobalVariable
var Store *v_ext.VSStore

//noinspection GoUnusedGlobalVariable
var Db *v_db.VS
var Json *v_json.VS
var Log *v_log.VS
var Utils *v_utils.VS
//noinspection GoUnusedGlobalVariable
var Var *v_var.VS


func init() {
	Conf = v_ext.Conf
	Store = v_ext.Store

	Db = v_db.Db
	Json = v_json.Json
	Log = v_log.Log
	Utils = v_utils.Utils
	Var = v_var.Var

	//comment
	vi.Json = Json
	vi.Log = Log
	vi.Utils = Utils
}
