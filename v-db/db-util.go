package v_db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strconv"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (*VS) BuildPostgresConnection(model DbVar) (connectionStr string) {
	if model.Schema == nil {
		connectionStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", model.Host, model.Port, model.Username, model.Password, model.DatabaseName)
	} else {
		connectionStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", model.Host, model.Port, model.Username, model.Password, model.DatabaseName, *model.Schema)
	}

	return
}


func (*VS) BuildPostgresInstance(masterConnection string, slaveConnection string) Postgres {
	return Postgres{
		Master: SqlMaster{conn: masterConnection, driveName: "postgres"},
		Slave:  SqlSlave{conn: slaveConnection, driveName: "postgres"},
	}
}

func (*VS) BuildPostgresInstanceSlave(slaveConnection string) Postgres {
	return Postgres{
		Slave: SqlSlave{conn: slaveConnection, driveName: "postgres"},
	}
}


func (*VS) SymbolParsPostgres(sqlQuery *string, length int) string {
	symbol := ""
	for i:=0; i<length; i++ {
		if i>0 {
			symbol += ", "
		}
		symbol += "$" + strconv.Itoa(i+1)
	}
	symbol = " ( " + symbol + " )"

	query := *sqlQuery + symbol
	if query[:1] == "\n" {
		query = query[1:]
	}
	*sqlQuery = query

	return query
}


func (*VS) SqlIn(sqlQuery string, sqlPars ...interface{}) (string, []interface{}, error) {
	_sqlQuery, _sqlPars, err := sqlx.In(sqlQuery, sqlPars...)
	if err != nil {
		return sqlQuery, sqlPars, err
	}
	return _sqlQuery, _sqlPars, err
}


func (*VS) Rebind(sqlQuery string) string {
	return sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
}


func sqlxConnect(driverName string, conn string) (*sqlx.DB, error) {
	instance, err := sqlx.Connect(driverName, conn)
	if err == nil {
		instance.SetConnMaxLifetime(time.Hour)
		instance.SetMaxIdleConns(10)
		err = instance.Ping()
	}

	if err != nil {
		err = errors.WithStack(err)
	}

	return instance, err
}


func sqlxExec(instance *sqlx.DB, sqlQuery string, pars ...interface{}) (sql.Result, error) {
	stmt, err := instance.Prepare(sqlQuery)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(pars...)
	if err != nil {
		err = errors.WithStack(err)
	}

	return res, err
}
