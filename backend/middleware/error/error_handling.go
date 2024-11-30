package middleware

import (
	"encoding/json"
	"fmt"
	"backend/shared"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

type ReturnResponseError struct {
	RespCode string `json:"responseCode"`
	RespMsg  string `json:"responseMessage"`
}

type ResponseList struct {
	ReturnResponseError
	HttpStatusCode string `json:"httpStatusCode"`
}

type ResponseFromDTO struct {
	shared.BaseResponse
	httpStatusCode ResponseList
}

type ListResponseFromJsonFile struct {
	ListResponseFromJsonFile map[string]ResponseList
}

var errorDataList map[string]ResponseList

func LoadErrorListFromJsonFile(pathfilename string) error {
	var file []byte
	var err error

	file, err = os.ReadFile(pathfilename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &errorDataList)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	fmt.Println("ERORR DI HADNLING : ", err)
	resp := ResponseError(c, err.Error())

	convertHTTPStatusCode, _ := strconv.Atoi(resp.HttpStatusCode)

	return c.Status(convertHTTPStatusCode).JSON(ReturnResponseError{
		RespCode: resp.RespCode,
		RespMsg:  resp.RespMsg,
	})
}

func ResponseError(ctx *fiber.Ctx, respCode string) ResponseList {

	var loadResponse = SearchResponseValueFromJsonFile(respCode)

	return ResponseList{
		ReturnResponseError: ReturnResponseError{
			RespCode: respCode,
			RespMsg:  loadResponse.RespMsg,
		},
		HttpStatusCode: loadResponse.HttpStatusCode,
	}
}

func ResponseSuccess(ctx *fiber.Ctx, respCode string, data interface{}) error {

	var loadResponse = SearchResponseValueFromJsonFile(respCode)

	convertRespCode, _ := strconv.Atoi(respCode)
	convertHTTPStatusCode, _ := strconv.Atoi(loadResponse.HttpStatusCode)

	return ctx.Status(convertHTTPStatusCode).JSON(ResponseFromDTO{
		BaseResponse: shared.BaseResponse{
			ResponseCode:    convertRespCode,
			ResponseMessage: loadResponse.RespMsg,
			Data:            data,
		},
	})
}

func SearchResponseValueFromJsonFile(resCode string) ResponseList {
	var loadListResponse = errorDataList
	resCodeValue, errResCodeValue := loadListResponse[resCode]
	if errResCodeValue {
		return resCodeValue
	} else {
		return ResponseList{
			ReturnResponseError: ReturnResponseError{
				RespCode: resCode,
				RespMsg:  "Error message is not defined!",
			},
			HttpStatusCode: strconv.Itoa(fiber.StatusInternalServerError),
		}
	}

}
