package weight

import (
	"context"
	"net/http"
	"time"

	"github.com/ijalalfrz/sirclo-weight-test/entity"
	"github.com/ijalalfrz/sirclo-weight-test/exception"
	"github.com/ijalalfrz/sirclo-weight-test/model"
	"github.com/ijalalfrz/sirclo-weight-test/response"
	"github.com/sirupsen/logrus"
)

// collection of message
const (
	insertOneUnexpectedErrMessage = "Unexpected error while inserting weight"
	insertOneSuccessMessage       = "Weight has been successfully inserted"
	updateOneUnexpectedErrMessage = "Unexpected error while updating weight"
	updateOneSuccessMessage       = "Weight has been successfully updated"
	weightNotFoundErrMessage      = "Weight not found"
	weightSuccessMessage          = "List of weight"
	weightUnexpectedErrMessage    = "Unexpected error while geting weight data"
	weightAllreadyExistErrMessage = "Weight is already exist"
)

// Usecase is collection of behaviour usecase
type Usecase interface {
	InsertOne(ctx context.Context, payload model.WeightPayload) (resp response.Response)
	UpdateOne(ctx context.Context, key int64, payload model.WeightPayload) (resp response.Response)
	FindMany(ctx context.Context) (resp response.Response)
	FindOne(ctx context.Context, key int64) (resp response.Response)
}

type weightUsecase struct {
	serviceName string
	logger      *logrus.Logger
	repository  Repository
}

// NewWeightUsecase is constructor
func NewWeightUsecase(property UsecaseProperty) Usecase {
	return &weightUsecase{
		serviceName: property.ServiceName,
		logger:      property.Logger,
		repository:  property.Repository,
	}
}

func (u weightUsecase) InsertOne(ctx context.Context, payload model.WeightPayload) (resp response.Response) {

	findWeight, err := u.repository.FindOne(ctx, payload.Date)
	if err != nil {
		u.logger.Error(err)
		if err != exception.ErrNotFound {
			return response.NewErrorResponse(err, http.StatusInternalServerError, nil, response.StatUnexpectedError, weightUnexpectedErrMessage)
		}
	}

	if findWeight.Date != 0 {
		err := exception.ErrConflict
		return response.NewErrorResponse(err, http.StatusConflict, nil, response.StatAlreadyExist, weightAllreadyExistErrMessage)
	}

	weight := entity.Weight{
		Date: payload.Date,
		Max:  payload.Max,
		Min:  payload.Min,
		Diff: payload.Max - payload.Min,
	}
	err = u.repository.InsertOne(ctx, weight)
	if err != nil {
		u.logger.Error(err)
		return response.NewErrorResponse(err, http.StatusInternalServerError, nil, response.StatUnexpectedError, insertOneUnexpectedErrMessage)

	}
	return response.NewSuccessResponse(nil, response.StatCreated, insertOneSuccessMessage)
}

func (u weightUsecase) UpdateOne(ctx context.Context, key int64, payload model.WeightPayload) (resp response.Response) {
	weight := entity.Weight{
		Date: payload.Date,
		Max:  payload.Max,
		Min:  payload.Min,
		Diff: payload.Max - payload.Min,
	}
	err := u.repository.UpdateOne(ctx, payload.Date, weight)
	if err != nil {
		u.logger.Error(err)
		return response.NewErrorResponse(err, http.StatusInternalServerError, nil, response.StatUnexpectedError, updateOneUnexpectedErrMessage)

	}
	return response.NewSuccessResponse(nil, response.StatOK, updateOneSuccessMessage)
}
func (u weightUsecase) FindMany(ctx context.Context) (resp response.Response) {

	weight, err := u.repository.FindMany(ctx, "date", -1)
	if err != nil {
		u.logger.Error(err)
		if err != exception.ErrNotFound {
			return response.NewErrorResponse(err, http.StatusInternalServerError, nil, response.StatUnexpectedError, weightUnexpectedErrMessage)
		}

		return response.NewErrorResponse(exception.ErrNotFound, http.StatusNotFound, nil, response.StatNotFound, weightNotFoundErrMessage)

	}
	sumMax := 0
	sumMin := 0
	sumDiff := 0
	var weightDetail []model.WeighDetailResponse
	for _, w := range weight {
		sumMax += w.Max
		sumMin += w.Min
		sumDiff += w.Diff

		wd := model.WeighDetailResponse{
			DateString: u.unixToDateString(w.Date),
			Date:       w.Date,
			Max:        w.Max,
			Min:        w.Min,
			Diff:       w.Diff,
		}
		weightDetail = append(weightDetail, wd)
	}
	totalData := len(weight)
	weightResponse := model.WeightResponse{
		List:        weightDetail,
		AverageMax:  float32(sumMax) / float32(totalData),
		AverageMin:  float32(sumMin) / float32(totalData),
		AverageDiff: float32(sumDiff) / float32(totalData),
	}
	return response.NewSuccessResponse(weightResponse, response.StatOK, weightSuccessMessage)
}
func (u weightUsecase) FindOne(ctx context.Context, key int64) (resp response.Response) {
	weight, err := u.repository.FindOne(ctx, key)
	if err != nil {
		u.logger.Error(err)
		if err != exception.ErrNotFound {
			return response.NewErrorResponse(err, http.StatusInternalServerError, nil, response.StatUnexpectedError, weightUnexpectedErrMessage)
		}

		return response.NewErrorResponse(exception.ErrNotFound, http.StatusNotFound, nil, response.StatNotFound, weightNotFoundErrMessage)

	}
	weightDetail := model.WeighDetailResponse{
		Date:       weight.Date,
		DateString: u.unixToDateString(weight.Date),
		Max:        weight.Max,
		Min:        weight.Min,
		Diff:       weight.Diff,
	}
	return response.NewSuccessResponse(weightDetail, response.StatOK, weightSuccessMessage)
}

func (u weightUsecase) unixToDateString(timestamp int64) string {
	tUnix := timestamp / int64(time.Second)
	tUnixNanoRemainder := (timestamp % int64(time.Second))
	date := time.Unix(tUnix, tUnixNanoRemainder)
	return date.Format("2006-01-02")
}
