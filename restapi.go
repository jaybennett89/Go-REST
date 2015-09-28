package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/jaybennett89/Go-REST/datastore"
)

func main() {

	m := martini.Classic()

	m.Post("/api/reviews/new", handleNewReview)
	m.Get("/api/reviews/:id", handleGetReview)
	m.Post("/api/reviews/:id/update", handleUpdateReview)

	m.RunOnAddr(":6960")
}

type NewReviewData struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	GitDiff string `json:"gitdiff"`
}

func handleNewReview(httpReq *http.Request) (int, string) {

	var data NewReviewData

	decoder := json.NewDecoder(httpReq.Body)
	err := decoder.Decode(&data)
	if err != nil {
		return http.StatusBadRequest, "invalid request data"
	}

	var requestId int
	requestId, err = datastore.CreateReview(data.Title, data.Author, data.GitDiff)
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, "something went wrong"
	}

	return http.StatusOK, strconv.Itoa(requestId)
}

type ReviewData struct {
	ReviewId int      `json:"reviewId"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	GitDiffs []string `json:"gitdiffs"`
}

func handleGetReview(params martini.Params) (int, string) {

	reviewId, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, "invalid id"
	}

	title, author, diffs, err := datastore.GetReview(reviewId)
	if err != nil {
		log.Print(err)
		return http.StatusNotFound, "review not found"
	}

	result := ReviewData{reviewId, title, author, diffs}

	var jsonBytes []byte
	jsonBytes, err = json.Marshal(&result)
	if err != nil {
		return http.StatusInternalServerError, "something went wrong"
	}

	return http.StatusOK, string(jsonBytes)
}

type UpdateReviewData struct {
	GitDiff string `json:"gitdiff"`
}

func handleUpdateReview(httpReq *http.Request, params martini.Params) (int, string) {

	reviewId, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, "invalid id"
	}

	var data UpdateReviewData
	decoder := json.NewDecoder(httpReq.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return http.StatusBadRequest, "unable to parse request data"
	}

	err = datastore.UpdateReview(reviewId, data.GitDiff)
	if err != nil {
		return http.StatusInternalServerError, "something went wrong"
	}

	return http.StatusOK, "review updated"
}
