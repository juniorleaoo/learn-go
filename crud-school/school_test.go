package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListSchools(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/schools", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var schools []School
	json.Unmarshal(recorder.Body.Bytes(), &schools)
	assert.Empty(t, schools)
}

func TestCreateSchool(t *testing.T) {
	router := setupRouter()
	newSchool := School{Name: "Fulano", Description: "Descrição", Status: true}
	body, _ := json.Marshal(newSchool)

	req, _ := http.NewRequest("POST", "/schools", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	var school School
	json.Unmarshal(recorder.Body.Bytes(), &school)
	assert.NotEmpty(t, school.ID)
	assert.Equal(t, newSchool.Name, school.Name)
	assert.Equal(t, newSchool.Description, school.Description)
	assert.Equal(t, newSchool.Status, school.Status)
}

func TestGetSchool(t *testing.T) {
	router := setupRouter()
	school, _ := insertUser(router)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/schools/%d", school.ID), nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var foundSchool School
	json.Unmarshal(recorder.Body.Bytes(), &foundSchool)
	assert.Equal(t, foundSchool.ID, school.ID)
	assert.Equal(t, foundSchool.Name, school.Name)
	assert.Equal(t, foundSchool.Description, school.Description)
	assert.Equal(t, foundSchool.Status, school.Status)
}

func TestDeleteSchool(t *testing.T) {
	router := setupRouter()
	school, _ := insertUser(router)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/schools/%d", school.ID), nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNoContent, recorder.Code)
}

func TestUpdateSchool(t *testing.T) {
	router := setupRouter()
	school, _ := insertUser(router)
	schoolUpdate := School{
		ID:          school.ID,
		Name:        "Fulano 2",
		Description: "Descrição 2",
		Status:      false,
	}
	body, _ := json.Marshal(schoolUpdate)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/schools/%d", school.ID), bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var updatedSchool School
	json.Unmarshal(recorder.Body.Bytes(), &updatedSchool)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, updatedSchool.ID, schoolUpdate.ID)
	assert.Equal(t, updatedSchool.Name, schoolUpdate.Name)
	assert.Equal(t, updatedSchool.Description, schoolUpdate.Description)
	assert.Equal(t, updatedSchool.Status, schoolUpdate.Status)
}

func insertUser(router *gin.Engine) (School, error) {
	newSchool := School{
		Name:        "Fulano",
		Description: "Descrição",
		Status:      true,
	}
	body, _ := json.Marshal(newSchool)

	req, _ := http.NewRequest("POST", "/schools", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var school School
	err := json.Unmarshal(recorder.Body.Bytes(), &school)
	return school, err
}
