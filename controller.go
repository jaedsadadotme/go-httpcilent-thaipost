package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) error {
	var url = "https://trackwebhook.thailandpost.co.th/post/api/v1/authenticate/token"
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", c.Request().Header.Get("Authorization"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result ResponseData
	result.Status = resp.Status
	result.StatusCode = resp.StatusCode
	json.Unmarshal(body, &result.Datas)
	return c.JSON(resp.StatusCode, result)
}

func GetTracking(c echo.Context) (err error) {
	var url = "https://trackapi.thailandpost.co.th/post/api/v1/track"
	u := new(Request)
	if err = c.Bind(u); err != nil {
		return err
	}
	convertBytes := new(bytes.Buffer)
	json.NewEncoder(convertBytes).Encode(u)

	/// http client start ///
	req, err := http.NewRequest("POST", url, convertBytes)
	req.Header.Set("Authorization", c.Request().Header.Get("Authorization"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/// http client end ///

	body, _ := ioutil.ReadAll(resp.Body)
	var result ResponseData
	result.Status = resp.Status
	result.StatusCode = resp.StatusCode
	json.Unmarshal(body, &result.Datas)
	return c.JSON(resp.StatusCode, result)
}
