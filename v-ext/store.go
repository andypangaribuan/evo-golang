package v_ext


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
var Store *VSStore

type VSStore struct {
	ServiceName    string
	ServiceVersion string
	CLogBaseUrl    string
}


func init() {
	Store = &VSStore{
		ServiceName:    "vo:default",
		ServiceVersion: "0.0.0",
	}
}
