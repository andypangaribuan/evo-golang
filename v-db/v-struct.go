package v_db

import (
	"github.com/jmoiron/sqlx"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
var Db *VS

type VS struct {}

func init() {
	Db = &VS{}
}


type DbVar struct {
	Host         string
	Port         int
	DatabaseName string
	Username     string
	Password     string
	Schema       *string
}


type IDb struct {
	Master SqlMaster
	Slave  SqlSlave
}

type Postgres struct {
	Master SqlMaster
	Slave  SqlSlave
}

type PostgresSlave struct {
	Slave SqlSlave
}



type SqlMaster struct {
	conn      string
	driveName string
	instance  *sqlx.DB
}

type SqlSlave struct {
	conn      string
	driveName string
	instance  *sqlx.DB
}

type SqlTransaction struct {
	sql        *SqlMaster
	SqlQueries []string
	SqlPars    [][]interface{}
}
