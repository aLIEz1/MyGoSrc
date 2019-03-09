/**
 * Created by super on 2019/3/8.
 *awesomeProject.
 */
package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
type ErroResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{HttpSC: 400,
		Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser = ErroResponse{HttpSC: 401,
		Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
	ErrorDBEroor = ErroResponse{HttpSC: 500,
		Error: Err{"DB ops failed", "003"}}
	ErrorInternalFaults = ErroResponse{HttpSC: 500,
		Error: Err{"Internal service error", "004"}}
)
