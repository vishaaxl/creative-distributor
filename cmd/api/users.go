package main

import (
	"errors"
	"net/http"
)

func (app *application) sendOTPHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	if input.Name == "" || input.PhoneNumber == "" {
		app.sendErrorResponse(w, r, http.StatusBadRequest, errors.New("name and phone number are required"))
		return
	}

	app.sendJsonResponse(w, http.StatusOK, envelope{"success": true}, nil)
}
