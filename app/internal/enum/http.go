package enum

const (
	FOR_SUCCESS     int = 1000
	FOR_ERROR       int = 500
	FOR_AUTH_ERROR  int = 401
	FOR_BAD_REQUEST int = 400
)

const (
	SUCCESS_STR     string = "Success"
	ERROR_STR       string = "Internal Error"
	AUTH_ERROR_STR  string = "Unauthorized"
	BAD_REQUEST_STR string = "Bad Request"
	//
	IS_ACTIVE string = "IS_ACTIVE"
)

const (
	NO_DATA_PROVIDED_CRE_STR string = "No Data Create Provided"
	NO_DATA_PROVIDED_UPD_STR string = "No Data Create Provided"
)
