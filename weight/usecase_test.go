package weight_test

import (
	"context"
	"testing"

	"github.com/ijalalfrz/sirclo-weight-test/entity"
	"github.com/ijalalfrz/sirclo-weight-test/exception"
	"github.com/ijalalfrz/sirclo-weight-test/model"
	"github.com/ijalalfrz/sirclo-weight-test/weight"
	"github.com/ijalalfrz/sirclo-weight-test/weight/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecaseInsertOne_Success(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	resultRepo := entity.Weight{
		Date: 0,
		Max:  0,
		Min:  0,
		Diff: 0,
	}

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(resultRepo, nil)
	repoMock.On("InsertOne", mock.Anything, mock.Anything).Return(nil)

	result := usecase.InsertOne(context.TODO(), model.WeightPayload{})

	assert.Nil(t, result.Error(), "should be no error")
	repoMock.AssertExpectations(t)

}
func TestUsecaseInsertOne_Unexpected(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	resultRepo := entity.Weight{
		Date: 0,
		Max:  0,
		Min:  0,
		Diff: 0,
	}

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(resultRepo, nil)
	repoMock.On("InsertOne", mock.Anything, mock.Anything).Return(exception.ErrInternalServer)

	result := usecase.InsertOne(context.TODO(), model.WeightPayload{})

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrInternalServer, "should be internal server error")
	repoMock.AssertExpectations(t)

}

func TestUsecaseInsertOne_Error_Unexpected_When_FindWeight(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(entity.Weight{}, exception.ErrInternalServer)
	result := usecase.InsertOne(context.TODO(), model.WeightPayload{})

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrInternalServer, "should be internal server error")
	repoMock.AssertExpectations(t)

}

func TestUsecaseInsertOne_Error_Already_Exist_Weight(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(entity.Weight{Date: 1}, nil)
	result := usecase.InsertOne(context.TODO(), model.WeightPayload{})

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrConflict, "should be conflict error")
	repoMock.AssertExpectations(t)

}

func TestUsecaseUpdateOne_Success(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	payload := model.WeightPayload{
		Date: 1,
		Max:  1,
		Min:  1,
	}

	repoMock.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	result := usecase.UpdateOne(context.TODO(), payload.Date, payload)

	assert.Nil(t, result.Error(), "should be no error")
	repoMock.AssertExpectations(t)

}

func TestUsecaseUpdateOne_Error_Unexpected(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	payload := model.WeightPayload{
		Date: 1,
		Max:  1,
		Min:  1,
	}

	repoMock.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(exception.ErrInternalServer)

	result := usecase.UpdateOne(context.TODO(), payload.Date, payload)

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrInternalServer, "should be internal server error")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindMany_Error_Unexpected(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindMany", mock.Anything, mock.Anything, mock.Anything).Return([]entity.Weight{}, exception.ErrInternalServer)

	result := usecase.FindMany(context.TODO())

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrInternalServer, "should be internal server error")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindMany_Error_NotFound(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindMany", mock.Anything, mock.Anything, mock.Anything).Return([]entity.Weight{}, exception.ErrNotFound)

	result := usecase.FindMany(context.TODO())

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrNotFound, "should be not found error")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindMany_Success(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})
	data := []entity.Weight{
		{
			Date: 1,
			Max:  2,
			Min:  1,
			Diff: 1,
		},
	}
	repoMock.On("FindMany", mock.Anything, mock.Anything, mock.Anything).Return(data, nil)

	result := usecase.FindMany(context.TODO())

	assert.Nil(t, result.Error(), "should be no error")
	resultData := result.Data().(model.WeightResponse)
	assert.Equal(t, len(resultData.List), 1, "should be equal one")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindOne_Error_Unexpected(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(entity.Weight{}, exception.ErrInternalServer)

	result := usecase.FindOne(context.TODO(), 1)

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrInternalServer, "should be internal server error")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindOne_Error_NotFound(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(entity.Weight{}, exception.ErrNotFound)

	result := usecase.FindOne(context.TODO(), 1)

	assert.Error(t, result.Error(), "should be error")
	assert.Equal(t, result.Error(), exception.ErrNotFound, "should be not found error")
	repoMock.AssertExpectations(t)
}

func TestUsecaseFindOne_Success(t *testing.T) {
	repoMock := new(mocks.Repository)
	usecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: "test-service",
		Logger:      logrus.New(),
		Repository:  repoMock,
	})

	data := entity.Weight{
		Date: 1,
		Max:  2,
		Min:  1,
		Diff: 1,
	}

	repoMock.On("FindOne", mock.Anything, mock.Anything).Return(data, nil)

	result := usecase.FindOne(context.TODO(), 1)

	assert.Nil(t, result.Error(), "should be no error")
	resultData := result.Data().(model.WeighDetailResponse)
	assert.Equal(t, resultData.Date, int64(1), "should be equal one")
	repoMock.AssertExpectations(t)
	repoMock.AssertExpectations(t)
}
