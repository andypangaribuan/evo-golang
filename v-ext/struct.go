package v_ext


/* ============================================
	Created by andy pangaribuan on 2020/06/17
	Copyright BoltIdea. All rights reserved.
   ============================================ */
type DbTxError struct {
	Err error
	Msg string
}


type DbUnsafeSelectError struct {
	LogType    string
	SqlQuery   string
	SqlPars    []interface{}
	LogMessage *string
	LogTrace   *string
}
