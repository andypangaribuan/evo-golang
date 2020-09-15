package v_db

import (
	"evo-lib/v-ext"
	"evo-lib/vi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (slf *SqlSlave) connect() (*sqlx.DB, error) {
	if slf.instance != nil {
		return slf.instance, nil
	}

	instance, err := sqlxConnect(slf.driveName, slf.conn)
	if err == nil {
		slf.instance = instance
	}

	return instance, err
}


func (slf *SqlSlave) Count(sqlQuery string, sqlPars ...interface{}) (int, error) {
	instance, err := slf.connect()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	pars := sqlPars
	if len(sqlPars) == 1 {
		switch data := sqlPars[0].(type) {
		case []interface{}: pars = data
		}
	}

	var count int
	if err := instance.Get(&count, sqlQuery, pars...); err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}


func (slf *SqlSlave) Select(out interface{}, sqlQuery string, sqlPars ...interface{}) (*v_ext.DbUnsafeSelectError, error) {
	instance, err := slf.connect()
	if err != nil {
		return nil, err
	}

	pars := sqlPars
	if len(sqlPars) == 1 {
		switch data := sqlPars[0].(type) {
		case []interface{}: pars = data
		}
	}

	err = instance.Select(out, sqlQuery, pars...)
	if err != nil {
		err = errors.WithStack(err)

		if v_ext.Conf.DbUnsafeCompatibility {
			msg := err.Error()
			unsafe := v_ext.DbUnsafeSelectError{
				LogType:    "error",
				SqlQuery:   sqlQuery,
				SqlPars:    pars,
				LogMessage: &msg,
				LogTrace:   vi.Log.Stack(sqlQuery, pars, err),
			}

			err = instance.Unsafe().Select(out, sqlQuery, pars...)
			if err != nil {
				err = errors.WithStack(err)
			}

			return &unsafe, err
		}
	}

	return nil, err
}
