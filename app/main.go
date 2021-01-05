package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/controller"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/response"
	"github.com/bmf-san/goblin"
)

const timeout time.Duration = 10 * time.Second

func main() {
	threshold, _ := strconv.Atoi(os.Getenv("LOG_THRESHOLD"))
	offset, _ := strconv.Atoi(os.Getenv("LOG_TIME_ZONE_OFFSET"))
	location := time.FixedZone(os.Getenv("TIME_ZONE"), offset)

	logger := logger.NewLogger(threshold, location)
	client := api.NewClient()
	response := response.NewResponse()

	homeController := controller.NewHomeController(logger, client, response)
	postController := controller.NewPostController(logger, client, response)
	categoryController := controller.NewCategoryController(logger, client, response)
	tagController := controller.NewTagController(logger, client, response)
	commentController := controller.NewCommentController(logger, client, response)
	sitemapController := controller.NewSitemapController(logger, client, response)
	feedController := controller.NewFeedController(logger, client, response)

	r := goblin.NewRouter()

	r.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		homeController.Index(w, r)
	}))
	r.GET("/posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postController.Index(w, r)
	}))
	r.GET("/posts/:title", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postController.Show(w, r)
	}))
	r.GET("/posts/categories/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postController.IndexByCategory(w, r)
	}))
	r.GET("/posts/tags/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postController.IndexByTag(w, r)
	}))
	r.POST("/posts/:title/comments", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commentController.Store(w, r)
	}))
	r.GET("/categories", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryController.Index(w, r)
	}))
	r.GET("/tags", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tagController.Index(w, r)
	}))
	r.GET("/sitemap", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sitemapController.Index(w, r)
	}))
	r.GET("/feed", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feedController.Index(w, r)
	}))

	s := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGTERM, os.Interrupt)
	logger.Info(fmt.Sprintf("SIGNAL %d received", <-q))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Gracefully shutdown")
}
