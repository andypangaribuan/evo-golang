package v_db

import (
	"fmt"
	"github.com/andypangaribuan/evo-golang/v-ext"
	"github.com/andypangaribuan/evo-golang/vi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */

func (slf *SqlMaster) connect() (*sqlx.DB, error) {
	if slf.instance != nil {
		return slf.instance, nil
	}

	instance, err := sqlxConnect(slf.driveName, slf.conn)
	if err == nil {
		slf.instance = instance
	}

	return instance, err
}





func (slf *SqlMaster) Exec(sqlQuery string, sqlPars ...interface{}) (err error) {
	instance, _err := slf.connect()
	if _err != nil {
		err = _err
		return
	}

	pars := sqlPars
	if len(sqlPars) == 1 {
		switch data := sqlPars[0].(type) {
		case []interface{}: pars = data
		}
	}

	_, err = sqlxExec(instance, sqlQuery, pars...)
	return
}

func (slf *SqlMaster) Transaction() *SqlTransaction {
	return &SqlTransaction{
		sql:        slf,
		SqlQueries: make([]string, 0),
		SqlPars:    make([][]interface{}, 0),
	}
}


func (slf *SqlTransaction) Exec(sqlQuery string, sqlPars ...interface{}) {
	pars := sqlPars
	if len(sqlPars) == 1 {
		switch data := sqlPars[0].(type) {
		case []interface{}: pars = data
		}
	}

	slf.SqlQueries = append(slf.SqlQueries, sqlQuery)
	slf.SqlPars = append(slf.SqlPars, pars)
}


func (slf *SqlTransaction) Commit() *v_ext.DbTxError {
	if len(slf.SqlQueries) == 0 {
		return nil
	}

	instance, err := slf.sql.connect()
	if err != nil {
		return &v_ext.DbTxError{
			Err: err,
		}
	}

	tx, err := instance.Begin()
	if err != nil {
		return &v_ext.DbTxError{
			Err: errors.WithStack(err),
		}
	}

	msgError := ""
	for i, query := range slf.SqlQueries {
		_, err = tx.Exec(query, slf.SqlPars[i]...)
		if err != nil {
			pars := slf.SqlPars[i]
			msg := ":: START ERROR ON SQL TRANSACTION"
			msg += "\n# error at query no. %v"
			msg += "\n# query:\n%v"
			if len(pars) == 0 {
				msg = fmt.Sprintf(msg, i+1, query)
			} else {
				ms := make(map[int]interface{}, 0)
				for i, v := range pars {
					ms[i+1] = v
				}

				if encode, err := vi.Json.JsonEncode(ms); err == nil {
					msg = fmt.Sprintf(msg + "\n# pars:\n%v", i+1, query, encode)
				} else {
					msg = fmt.Sprintf(msg + "\n# pars:\n%+v", i+1, query, ms)
				}
			}

			msg += "\n\n:: ALL QUERY"
			for _i, _q := range slf.SqlQueries {
				p := slf.SqlPars[_i]
				m := "\n# query no. %v:"
				m += "\n$ %v"
				if len(p) == 0 {
					m = fmt.Sprintf(m, _i+1, _q)
				} else {
					if p1, _e := vi.Json.JsonEncode(p); _e == nil {
						m = fmt.Sprintf(m + "\n$ %v", _i+1, _q, p1)
					} else {
						m = fmt.Sprintf(m + "\n$ sql pars encoding error: %v", _i+1, _q, _e)
					}
				}
				m += "\n"
			}
			msg += ":: END ERROR ON SQL TRANSACTION"
			msgError = msg

			err = errors.WithStack(err)
			_ = tx.Rollback()
			break
		}
	}

	if err == nil {
		err = tx.Commit()
	}

	if err == nil {
		return nil
	}

	return &v_ext.DbTxError{
		Err: err,
		Msg: msgError,
	}
}
