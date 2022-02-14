package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func Test_12345(t *testing.T) {
	log.Println("Testing Test_12345!")
	apitest.New().
		Handler(newApp().Router).
		Get("/").
		Expect(t).
		Body(`{"data":"hello world"}`).
		Status(http.StatusOK).
		End()
}
