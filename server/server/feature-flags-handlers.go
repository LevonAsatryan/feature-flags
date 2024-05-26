package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LevonAsatryan/feature-flags/db"

	"github.com/gin-gonic/gin"
)

type Env struct {
	ID   int    `json:"id"`
	Name string `json"name"`
}

type CreateFFBody struct {
	Name string `json:"name"`
	Env  string `json:"env"`
}

func (s *Server) CreateFFSGroup() *gin.RouterGroup {
	group := s.R.Group("/ffs")
	createMockedFeatureFlags()
	createMessages()
	createEnvs()
	group.GET("/", getAll)
	group.GET("/:id", getOne)
	group.POST("/", create)
	group.DELETE("/:id", delete)
	return group
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, &mockedFeatureFlags)
}

func create(c *gin.Context) {
	var rb CreateFFBody
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&rb)
	if err != nil {
		errorHandler(c, http.StatusBadRequest, err.Error())
	}

	env, err := checkEnv(rb.Env)
	if err != nil {
		errorHandler(c, http.StatusNotFound, err.Error())
		return
	}

	ff := db.FeatureFlag{
		ID:    len(mockedFeatureFlags),
		Name:  rb.Name,
		Env:   env.ID,
		Value: false,
	}
	mockedFeatureFlags = append(mockedFeatureFlags, ff)
	c.JSON(http.StatusOK, &ff)
}

func getOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
		return
	}
	res, _ := findFF(id)

	if res == nil {
		errorHandler(c, http.StatusNotFound, messages["ffNotFound"])
		return
	}

	c.JSON(http.StatusOK, res)
}

func delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
	}

	res, index := findFF(id)

	if res == nil {
		errorHandler(c, http.StatusNotFound, messages["ffNotFound"])
		return
	}

	mockedFeatureFlags = append(mockedFeatureFlags[:index], mockedFeatureFlags[index+1:]...)

	c.JSON(http.StatusOK, res)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func findFF(id int) (*db.FeatureFlag, int) {
	var res *db.FeatureFlag
	var index int
	for i, ff := range mockedFeatureFlags {
		if ff.ID == id {
			res = &ff
			index = i
			break
		}
	}

	return res, index
}

func errorHandler(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

var messages map[string]string = make(map[string]string)

func createMessages() {
	messages["ffNotFound"] = "Feature flag with given ID was not found"
	messages["invalidId"] = "Provided ID is not valid"
}

/**
 * Test data start
 */
var mockedFeatureFlags []db.FeatureFlag
var envs []Env

func createMockedFeatureFlags() {
	for i := 0; i < 100; i++ {
		mockedFeatureFlags = append(
			mockedFeatureFlags,
			db.FeatureFlag{
				ID:    i,
				Name:  "test" + strconv.Itoa(i),
				Value: i%2 == 0,
				Env:   0, Created_at: "0",
				Updated_at: "0",
			})
	}
}

func createEnvs() {
	envs = append(envs, Env{
		ID:   0,
		Name: "dev",
	})
	envs = append(envs, Env{
		ID:   1,
		Name: "prod",
	})
}

func checkEnv(name string) (*Env, error) {
	var ret *Env
	for _, env := range envs {
		if env.Name == name {
			ret = &env
			break
		}
	}

	if ret == nil {
		return nil, fmt.Errorf("env with name %q not found", name)
	}

	return ret, nil
}

/**
 * Test data end
 */
