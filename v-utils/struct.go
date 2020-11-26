package v_utils

import (
	"net/http"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
var Utils *VS

type filesStruct struct { }

type FileScanResult struct {
	FileName string
	DirPath  string
	FilePath string
}

type VS struct {
	Files *filesStruct
}


type evoHttpClient struct {
	client http.Client
}


func init() {
	Utils = &VS{
		Files: &filesStruct{},
	}
}
