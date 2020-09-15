package v_db


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (slf *IDb) SetDB(obj interface{}) {
	switch data := obj.(type) {
	case Postgres:
		slf.Master = data.Master
		slf.Slave = data.Slave
	}
}


func (slf *IDb) Transaction() *SqlTransaction {
	return slf.Master.Transaction()
}
