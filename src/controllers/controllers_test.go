package controllers

import (
	"booleans/mock"
	"booleans/services"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type responseType struct {
	Name string `json:"key"`
	Val  bool   `json:"value"`
	UUID string `json:"id"`
}

func TestNewBoolean(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	b := services.Boolean{
		Name: "ash",
		Val:  true,
		UUID: "68df2cbb-a432-4b35-8a99-2cf3de9b243c",
	}
	mockRepo.EXPECT().AddToDB("ash", true).Return(b)

	router := gin.Default()
	router.POST("/", MyController.NewBoolean)

	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
		"key": "ash",
		"value": true
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.Name, response.Name)
	assert.Equal(t, b.Val, response.Val)
	assert.Equal(t, b.UUID, response.UUID)
}

func TestNewBooleanEmptyKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	b := services.Boolean{
		Val:  true,
		UUID: "68df2cbb-a432-4b35-8a99-2cf3de9b243c",
	}
	mockRepo.EXPECT().AddToDB("", true).Return(b)

	router := gin.Default()
	router.POST("/", MyController.NewBoolean)

	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
		"value": true
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.Name, response.Name)
	assert.Equal(t, b.Val, response.Val)
	assert.Equal(t, b.UUID, response.UUID)
}

func TestNewBooleanEmptyValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/", MyController.NewBoolean)

	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
		"key": "ash"
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewBooleanIncorrectDatatype(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/", MyController.NewBoolean)

	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
		"key": "ash",
		"value": 420
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBoolean(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	b := services.Boolean{
		Name: "ash",
		Val:  true,
		UUID: "68df2cbb-a432-4b35-8a99-2cf3de9b243c",
	}
	mockRepo.EXPECT().
		GetFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c").
		Return(b, nil)

	router := gin.Default()
	router.GET("/:id", MyController.GetBoolean)

	req, err := http.NewRequest("GET", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.Name, response.Name)
	assert.Equal(t, b.Val, response.Val)
	assert.Equal(t, b.UUID, response.UUID)
}

func TestGetBooleanInvalidId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/:id", MyController.GetBoolean)

	req, err := http.NewRequest("GET", "/68dfcbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBooleanNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	mockRepo.EXPECT().
		GetFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c").
		Return(services.Boolean{}, errors.New("record not found"))

	router := gin.Default()
	router.GET("/:id", MyController.GetBoolean)

	req, err := http.NewRequest("GET", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateBoolean(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	b := services.Boolean{
		Name: "bash",
		Val:  false,
		UUID: "68df2cbb-a432-4b35-8a99-2cf3de9b243c",
	}
	mockRepo.EXPECT().UpdateInDB("bash", false, "68df2cbb-a432-4b35-8a99-2cf3de9b243c").Return(b, nil)

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"key": "bash",
		"value": false
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.Name, response.Name)
	assert.Equal(t, b.Val, response.Val)
	assert.Equal(t, b.UUID, response.UUID)
}

func TestUpdateBooleanInvalidId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68dfcbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"key": "bash",
		"value": false
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateBooleanEmptyKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	b := services.Boolean{
		Name: "ash",
		Val:  false,
		UUID: "68df2cbb-a432-4b35-8a99-2cf3de9b243c",
	}
	mockRepo.EXPECT().UpdateInDB("", false, "68df2cbb-a432-4b35-8a99-2cf3de9b243c").Return(b, nil)

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"value": false
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.Name, response.Name)
	assert.Equal(t, b.Val, response.Val)
	assert.Equal(t, b.UUID, response.UUID)
}

func TestUpdateBooleanEmptyValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"key": "bash"
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateBooleanIncorrectDatatype(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"key": "bash",
		"value": [1,2,3]
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateBooleanNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	mockRepo.EXPECT().
		UpdateInDB("bash", false, "68df2cbb-a432-4b35-8a99-2cf3de9b243c").
		Return(services.Boolean{}, errors.New("record not found"))

	router := gin.Default()
	router.PATCH("/:id", MyController.UpdateBoolean)

	req, err := http.NewRequest("PATCH", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", strings.NewReader(`{
		"key": "bash",
		"value": false
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBoolean(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	mockRepo.EXPECT().DeleteFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c").Return(nil)

	router := gin.Default()
	router.DELETE("/:id", MyController.DeleteBoolean)

	req, err := http.NewRequest("DELETE", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestDeleteBooleanInvalidId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.DELETE("/:id", MyController.DeleteBoolean)

	req, err := http.NewRequest("DELETE", "/68dfcbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteBooleanNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	services.MyRepo = mockRepo

	mockRepo.EXPECT().
		DeleteFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c").
		Return(errors.New("record not found"))

	router := gin.Default()
	router.DELETE("/:id", MyController.DeleteBoolean)

	req, err := http.NewRequest("DELETE", "/68df2cbb-a432-4b35-8a99-2cf3de9b243c", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
