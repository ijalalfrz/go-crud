package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-playground/validator/v10"
	"github.com/ijalalfrz/sirclo-weight-test/mongodb"
	"github.com/ijalalfrz/sirclo-weight-test/weight"

	"github.com/ijalalfrz/sirclo-weight-test/middleware"

	"github.com/ijalalfrz/sirclo-weight-test/response"
	"github.com/ijalalfrz/sirclo-weight-test/server"

	gctx "github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"

	"github.com/ijalalfrz/sirclo-weight-test/config"
)

var (
	cfg          *config.Config
	location     *time.Location
	indexMessage string = "Application is running properly"
)

func init() {
	cfg = config.Load()
}

func main() {
	// init logger
	logger := logrus.New()
	logger.SetFormatter(cfg.Logger.Formatter)
	logger.SetReportCaller(true)

	// init validator
	vld := validator.New()

	// set mongodb
	mc, err := mongo.NewClient(cfg.Mongodb.ClientOptions)
	if err != nil {
		logger.Fatal(err)
	}
	mca := mongodb.NewClientAdapter(mc)
	if err := mca.Connect(context.Background()); err != nil {
		logger.Fatal(err)
	}
	mdb := mca.Database(cfg.Mongodb.Database)

	// init router object
	router := mux.NewRouter()
	router.HandleFunc("/", index)

	// init domain object
	weightRepository := weight.NewWeightRepository(logger, mdb)
	weightUsecase := weight.NewWeightUsecase(weight.UsecaseProperty{
		ServiceName: cfg.Application.Name,
		Logger:      logger,
		Repository:  weightRepository,
	})

	// init http handler
	weight.NewWeightHTTPHandler(logger, vld, router, weightUsecase)

	// middleware]
	httpHandler := gctx.ClearHandler(router)
	httpHandler = middleware.Recovery(logger, httpHandler)
	httpHandler = middleware.CORS(httpHandler)

	// initiate server
	srv := server.NewServer(logger, httpHandler, cfg.Application.Port)
	srv.Start()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	<-sigterm

	// closing service for a gracefull shutdown.
	srv.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	resp := response.NewSuccessResponse(nil, response.StatOK, indexMessage)
	response.JSON(w, resp)
}
