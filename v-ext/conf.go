package v_ext


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
var Conf *VSConf

type VSConf struct {
	DbUnsafeCompatibility bool
	DbType                string
}


func init() {
	Conf = &VSConf{}
}



func (*VSConf) TraceIdKey() string {
	return "EVO_API_TRACE_ID"
}

func (*VSConf) RequestUidKey() string {
	return "EVO_API_REQUEST_UID"
}

func (*VSConf) RequestFromServiceName() string {
	return "EVO_API_REQUEST_FROM_SERVICE_NAME"
}

func (*VSConf) RequestFromServiceVersion() string {
	return "EVO_API_REQUEST_FROM_SERVICE_VERSION"
}
