package weight

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ijalalfrz/sirclo-weight-test/model"
	"github.com/sirupsen/logrus"
)

const (
	basePath = "/weight"
)

// HTTPHandler is a concrete struct of weight http handler.
type HTTPHandler struct {
	Logger   *logrus.Logger
	Validate *validator.Validate
	Usecase  Usecase
}

func NewWeightHTTPHandler(logger *logrus.Logger, validate *validator.Validate, router *mux.Router, usecase Usecase) {
	handler := &HTTPHandler{
		Logger:   logger,
		Validate: validate,
		Usecase:  usecase,
	}
	router.HandleFunc(basePath+"/add", handler.GetWeightForm).Methods(http.MethodGet)
	router.HandleFunc(basePath+"/{date}/update", handler.GetUpdateWeightForm).Methods(http.MethodGet)

	router.HandleFunc(basePath, handler.Index).Methods(http.MethodGet)
	router.HandleFunc(basePath+"/{date}", handler.Detail).Methods(http.MethodGet)

	router.HandleFunc(basePath, handler.AddWeight).Methods(http.MethodPost)
	router.HandleFunc(basePath+"/{date}", handler.UpdateWeight).Methods(http.MethodPost)

}

func (handler HTTPHandler) GetWeightForm(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Error": r.Header.Get("error"),
	}
	tmpl := template.Must(template.ParseFiles("./weight/template/add.html"))

	tmpl.Execute(w, data)
	return

}

func (handler HTTPHandler) GetUpdateWeightForm(w http.ResponseWriter, r *http.Request) {

	pathVariables := mux.Vars(r)
	dateStr := pathVariables["date"]
	if dateStr == "" {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
	}

	date, err := strconv.ParseInt(dateStr, 10, 64)
	if err != nil {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
	}

	resp := handler.Usecase.FindOne(r.Context(), date)

	data := map[string]interface{}{
		"Error": r.Header.Get("error"),
		"Data":  resp.Data(),
	}

	tmpl := template.Must(template.ParseFiles("./weight/template/update.html"))
	tmpl.Execute(w, data)
	return

}

func (handler HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	resp := handler.Usecase.FindMany(r.Context())

	tmpl := template.Must(template.ParseFiles("./weight/template/index.html"))
	tmpl.Execute(w, resp.Data())
	return

}

func (handler HTTPHandler) Detail(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	dateStr := pathVariables["date"]
	if dateStr == "" {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
	}
	date, _ := strconv.ParseInt(dateStr, 10, 64)

	resp := handler.Usecase.FindOne(r.Context(), date)

	tmpl := template.Must(template.ParseFiles("./weight/template/detail.html"))
	tmpl.Execute(w, resp.Data())
	return

}

func (handler HTTPHandler) AddWeight(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dateString := r.FormValue("date")
	dateTime, _ := time.Parse("2006-01-02", dateString)
	max, _ := strconv.Atoi(r.FormValue("max"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	payload := model.WeightPayload{
		Date: dateTime.UnixNano(),
		Max:  max,
		Min:  min,
	}

	err := handler.validateRequest(payload)
	if err != nil {
		r.Header.Set("error", err.Error())
		handler.GetWeightForm(w, r)
		return
	}

	resp := handler.Usecase.InsertOne(ctx, payload)

	if resp.Error() == nil {
		http.Redirect(w, r, basePath+"/add", http.StatusSeeOther)
	} else {
		r.Header.Set("error", resp.Message())
		handler.GetWeightForm(w, r)
		return
	}
	return
}

func (handler HTTPHandler) UpdateWeight(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pathVariables := mux.Vars(r)
	dateStr := pathVariables["date"]
	if dateStr == "" {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
	}
	max, _ := strconv.Atoi(r.FormValue("max"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	date, _ := strconv.ParseInt(dateStr, 10, 64)
	payload := model.WeightPayload{
		Date: date,
		Max:  max,
		Min:  min,
	}

	err := handler.validateRequest(payload)
	if err != nil {
		r.Header.Set("error", err.Error())
		handler.GetUpdateWeightForm(w, r)
		return
	}

	resp := handler.Usecase.UpdateOne(ctx, date, payload)
	if resp.Error() == nil {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
	} else {
		r.Header.Set("error", resp.Message())
		handler.GetUpdateWeightForm(w, r)
		return
	}
	return
}

func (handler HTTPHandler) validateRequest(payload model.WeightPayload) (err error) {
	err = handler.Validate.Struct(payload)
	if err == nil {
		if payload.Max < payload.Min {
			err = fmt.Errorf("Max must be greater than min")
			return
		}
		return
	}

	errorFields := err.(validator.ValidationErrors)
	errorField := errorFields[0]
	err = fmt.Errorf("Invalid '%s' with value '%v'", errorField.Field(), errorField.Value())

	return
}
