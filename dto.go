package ghasedak

type ResponseDto struct {
	IsSuccess  bool   `json:"isSuccess"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
