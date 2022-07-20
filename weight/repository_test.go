package weight_test

import (
	"context"
	"testing"

	"github.com/ijalalfrz/sirclo-weight-test/entity"
	"github.com/ijalalfrz/sirclo-weight-test/exception"
	"github.com/ijalalfrz/sirclo-weight-test/mongodb/mocks"
	"github.com/ijalalfrz/sirclo-weight-test/weight"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertOne_Success(t *testing.T) {
	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	err := repo.InsertOne(context.TODO(), entity.Weight{})
	assert.NoError(t, err, "should be no error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestInsertOne_Error_Unexpected(t *testing.T) {
	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("InsertOne", mock.Anything, mock.Anything).Return(nil, exception.ErrInternalServer)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	err := repo.InsertOne(context.TODO(), entity.Weight{})
	assert.Error(t, err, "should be error")
	assert.Equal(t, err, exception.ErrInternalServer, "should be error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)

}

func TestUpdateOne_Success(t *testing.T) {
	updateResult := &mongo.UpdateResult{
		MatchedCount: 1,
	}
	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, options.Update().SetUpsert(true)).Return(updateResult, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	err := repo.UpdateOne(context.TODO(), 1, entity.Weight{})
	assert.NoError(t, err, "should be no error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestUpdateOne_Error_Unexpected(t *testing.T) {

	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, options.Update().SetUpsert(true)).Return(nil, exception.ErrInternalServer)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	err := repo.UpdateOne(context.TODO(), 1, entity.Weight{})
	assert.Error(t, err, "should be error")
	assert.Equal(t, err, exception.ErrInternalServer, "should be error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestUpdateOne_Error_NotFound(t *testing.T) {

	updateResult := &mongo.UpdateResult{
		MatchedCount: 0,
	}
	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, options.Update().SetUpsert(true)).Return(updateResult, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	err := repo.UpdateOne(context.TODO(), 1, entity.Weight{})
	assert.Error(t, err, "should be error")
	assert.Equal(t, err, exception.ErrNotFound, "should be not found error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}
func TestFindMany_Success(t *testing.T) {
	cursorMock := new(mocks.Cursor)

	col := new(mocks.Collection)
	db := new(mocks.Database)
	cursorMock.On("Next", mock.Anything).Return(true).Once()
	cursorMock.On("Next", mock.Anything).Return(false).Once()
	cursorMock.On("Decode", mock.AnythingOfType("*entity.Weight")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Weight)
		arg.Date = 1656633600000000000
	})
	col.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cursorMock, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindMany(context.TODO(), "date", -1)
	assert.NoError(t, err, "should be no error")
	assert.Equal(t, result[0].Date, int64(1656633600000000000), "should be the same")
	cursorMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindMany_Error_Unexpected_When_Decode(t *testing.T) {
	cursorMock := new(mocks.Cursor)

	col := new(mocks.Collection)
	db := new(mocks.Database)
	cursorMock.On("Next", mock.Anything).Return(true).Once()
	cursorMock.On("Decode", mock.AnythingOfType("*entity.Weight")).Return(mongo.ErrNoDocuments)
	col.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cursorMock, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindMany(context.TODO(), "date", -1)
	assert.Nil(t, result, "should  be null")
	assert.Error(t, err, "should be error")
	assert.Equal(t, err, exception.ErrInternalServer, "should be not internal server error")
	cursorMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindMany_Error_Unexpected(t *testing.T) {
	col := new(mocks.Collection)
	db := new(mocks.Database)

	col.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(nil, mongo.ErrClientDisconnected)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindMany(context.TODO(), "date", -1)
	assert.Nil(t, result, "should  be null")
	assert.Error(t, err, "should be error")
	assert.Equal(t, err, exception.ErrInternalServer, "should be not internal server error")
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindMany_Error_NotFound(t *testing.T) {
	cursorMock := new(mocks.Cursor)

	col := new(mocks.Collection)
	db := new(mocks.Database)
	cursorMock.On("Next", mock.Anything).Return(false).Once()
	col.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cursorMock, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindMany(context.TODO(), "date", -1)
	assert.Nil(t, result)
	assert.Equal(t, err, exception.ErrNotFound, "should be not found error")
	assert.Error(t, err, "should be error")
	cursorMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindOne_Success(t *testing.T) {
	singleResultMock := new(mocks.SingleResult)

	date := int64(1656633600000000000)
	col := new(mocks.Collection)
	db := new(mocks.Database)
	singleResultMock.On("Decode", mock.AnythingOfType("*entity.Weight")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Weight)
		arg.Date = date
	})
	col.On("FindOne", mock.Anything, mock.Anything).Return(singleResultMock, nil)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindOne(context.TODO(), date)
	assert.NoError(t, err, "should be no error")
	assert.Equal(t, result.Date, date, "should be the same")
	singleResultMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindOne_Error_NotFound(t *testing.T) {
	singleResultMock := new(mocks.SingleResult)

	date := int64(1656633600000000000)
	col := new(mocks.Collection)
	db := new(mocks.Database)

	singleResultMock.On("Decode", mock.AnythingOfType("*entity.Weight")).Return(mongo.ErrNoDocuments)
	col.On("FindOne", mock.Anything, mock.Anything).Return(singleResultMock)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindOne(context.TODO(), date)
	assert.Error(t, err, "should be error")
	assert.Equal(t, exception.ErrNotFound, err)
	assert.Equal(t, int64(0), result.Date, "should be 0")
	singleResultMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestFindOne_Error_Unexpected(t *testing.T) {
	singleResultMock := new(mocks.SingleResult)

	date := int64(1656633600000000000)
	col := new(mocks.Collection)
	db := new(mocks.Database)

	singleResultMock.On("Decode", mock.AnythingOfType("*entity.Weight")).Return(mongo.ErrClientDisconnected)
	col.On("FindOne", mock.Anything, mock.Anything).Return(singleResultMock)
	db.On("Collection", mock.AnythingOfType("string")).Return(col)

	repo := weight.NewWeightRepository(logrus.New(), db)

	result, err := repo.FindOne(context.TODO(), date)
	assert.Error(t, err, "should be error")
	assert.Equal(t, exception.ErrInternalServer, err)
	assert.Equal(t, int64(0), result.Date, "should be 0")
	singleResultMock.AssertExpectations(t)
	col.AssertExpectations(t)
	db.AssertExpectations(t)
}
