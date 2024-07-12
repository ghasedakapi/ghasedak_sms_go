package main

type ResponseDto struct {
	IsSuccess  bool   `json:"isSuccess"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

//type AccountInformationResponse struct {
//	IsSuccess bool   `json:"isSuccess"`
//	StatusCode int  `json:"statusCode"`
//	Message   string `json:"message"`
//	Data      struct {
//		Credit     int    `json:"credit"`
//		ExpireDate string `json:"expireDate"`
//		Plan       string `json:"plan"`
//	} `json:"data"`
//}
