package helperModel

type BaseResponseModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

type ErrRespValidationModel struct {
	Message         string      `json:"message"`
	Data            interface{} `json:"data"`
	Error           string      `json:"error"`
	ErrorValidation interface{} `json:"error_validation"`
}
