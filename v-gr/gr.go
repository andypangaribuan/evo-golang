package v_gr


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (slf *GR) AppVersionNoLongerSupported() {
	slf.Code = 1
	slf.Message = "this app/endpoint version is no longer supported or has been retired"
}

func (slf *GR) UnderMaintenance() {
	slf.Code = 2
	slf.Message = "under maintenance"
}

func (slf *GR) InvalidAuthToken() {
	slf.Code = 3
	slf.Message = "invalid auth-token"
}

func (slf *GR) UserNotFound() {
	slf.Code = 4
	slf.Message = "user not found"
}

func (slf *GR) AccountSuspended() {
	slf.Code = 5
	slf.Message = "your account is suspended and is not permitted to access this feature"
}

func (slf *GR) AccountLocked() {
	slf.Code = 6
	slf.Message = "your account is locked and is not permitted to access this feature"
}

func (slf *GR) NotAuthorizedToAccessThisApi() {
	slf.Code = 7
	slf.Message = "you are not authorized to access this api"
}



func (slf *GR) InternalError() {
	slf.Code = 10
	slf.Message = "internal error"
}

func (slf *GR) ParameterConversionFailed() {
	slf.Code = 11
	slf.Message = "failed to convert the body to the model"
}

func (slf *GR) EmptyParameter() {
	slf.Code = 12
	slf.Message = "parameter cannot be empty"
}

func (slf *GR) InvalidParameter() {
	slf.Code = 13
	slf.Message = "invalid parameter value"
}

func (slf *GR) DataExists() {
	slf.Code = 14
	slf.Message = "data already exists"
}

func (slf *GR) DataNotExist() {
	slf.Code = 15
	slf.Message = "data does not exist"
}

func (slf *GR) DataDuplicate() {
	slf.Code = 16
	slf.Message = "found duplicate data"
}

func (slf *GR) DataInvalid() {
	slf.Code = 17
	slf.Message = "invalid data"
}

func (slf *GR) TokenNotExist() {
	slf.Code = 18
	slf.Message = "token does not exist"
}

func (slf *GR) NoDataFound() {
	slf.Code = 20
	slf.Message = "there are no data found"
}
