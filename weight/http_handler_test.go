package weight_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ijalalfrz/sirclo-weight-test/entity"
	"github.com/ijalalfrz/sirclo-weight-test/exception"
	"github.com/ijalalfrz/sirclo-weight-test/response"
	"github.com/ijalalfrz/sirclo-weight-test/weight"
	"github.com/ijalalfrz/sirclo-weight-test/weight/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	vld *validator.Validate
)

func TestMain(m *testing.M) {
	vld = validator.New()
	m.Run()
}
func TestNewWeightHTTPHandler(t *testing.T) {
	logger := logrus.New()
	validate := &validator.Validate{}
	router := &mux.Router{}
	usecase := &mocks.Usecase{}

	weight.NewWeightHTTPHandler(logger, validate, router, usecase)
}

func TestHttpHandler_Index_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := []entity.Weight{
		{
			Date: 1,
			Max:  2,
			Min:  1,
			Diff: 1,
		},
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindMany", mock.Anything).Return(successResponse)

	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.Index)
	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
	usecase.AssertExpectations(t)
}

func TestHttpHandler_AddForm_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.GetWeightForm)
	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_UpdateForm_Error_NoPathVariable(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.GetUpdateWeightForm)
	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusSeeOther)
}

func TestHttpHandler_UpdateForm_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)

	vars := map[string]string{
		"date": "1",
	}

	r = mux.SetURLVars(r, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.GetUpdateWeightForm)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_Detail_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)

	vars := map[string]string{
		"date": "1",
	}

	r = mux.SetURLVars(r, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.Detail)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_Detail_Error_NoPathVariable(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.Detail)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusSeeOther)
}

func TestHttpHandler_AddWeight_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	successResponse := response.NewSuccessResponse(nil, response.StatOK, "success")
	usecase.On("InsertOne", mock.Anything, mock.Anything).Return(successResponse)
	var bodyStr = []byte(`date=1&max=3&min=1`)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(bodyStr))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.AddWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusSeeOther)
}

func TestHttpHandler_AddWeight_Error_Validation(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	successResponse := response.NewSuccessResponse(nil, response.StatOK, "success")
	usecase.On("InsertOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.AddWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_AddWeight_Error_Validation_GreaterThan(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	successResponse := response.NewSuccessResponse(nil, response.StatOK, "success")
	usecase.On("InsertOne", mock.Anything, mock.Anything).Return(successResponse)
	var bodyStr = []byte(`date=1&max=1&min=4`)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(bodyStr))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.AddWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_AddWeight_Error_Unexpected(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	errorResponse := response.NewErrorResponse(exception.ErrInternalServer, http.StatusInternalServerError, nil, response.StatUnexpectedError, "fail")
	usecase.On("InsertOne", mock.Anything, mock.Anything).Return(errorResponse)
	var bodyStr = []byte(`date=1&max=3&min=1`)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(bodyStr))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.AddWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_UpdateWeight_Success(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	successResponse := response.NewSuccessResponse(nil, response.StatOK, "success")
	usecase.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(successResponse)
	var bodyStr = []byte(`max=3&min=1`)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(bodyStr))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	vars := map[string]string{
		"date": "1",
	}
	r = mux.SetURLVars(r, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.UpdateWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusSeeOther)
}

func TestHttpHandler_UpdateWeight_ErrorValidation(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	vars := map[string]string{
		"date": "1",
	}
	r = mux.SetURLVars(r, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.UpdateWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestHttpHandler_UpdateWeight_Error_NoPathVariable(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}
	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}
	successResponse := response.NewSuccessResponse(data, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.UpdateWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusSeeOther)
}

func TestHttpHandler_UpdateWeight_Error_Unexpected(t *testing.T) {
	usecase := new(mocks.Usecase)

	hh := weight.HTTPHandler{
		Logger:       logrus.New(),
		Validate:     vld,
		Usecase:      usecase,
		TemplatePath: "./template/",
	}

	errorResponse := response.NewErrorResponse(exception.ErrInternalServer, http.StatusInternalServerError, nil, response.StatUnexpectedError, "fail")
	usecase.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(errorResponse)
	successResponse := response.NewSuccessResponse(nil, response.StatOK, "success")
	usecase.On("FindOne", mock.Anything, mock.Anything).Return(successResponse)
	var bodyStr = []byte(`max=3&min=1`)
	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(bodyStr))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	vars := map[string]string{
		"date": "1",
	}
	r = mux.SetURLVars(r, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hh.UpdateWeight)

	handler.ServeHTTP(recorder, r)
	assert.Equal(t, recorder.Code, http.StatusOK)
}
