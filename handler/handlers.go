package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.com/ironcore864/actsctrl/config"
	"gitlab.com/ironcore864/actsctrl/model"
	"gitlab.com/ironcore864/actsctrl/redisclient"
)

// IndexPostHandler handles POST request to host:port/
func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	var httpStatus int
	var responseText string
	var req model.Request

	httpStatus, responseText, body := unpackReqBody(r)
	if httpStatus != 0 {
		renderResult(httpStatus, responseText, w)
		return
	}

	httpStatus, responseText, req = unmarshalReqBody(body)
	if httpStatus != 0 {
		renderResult(httpStatus, responseText, w)
		return
	}

	httpStatus, responseText = ifExceedsLimit(req)
	renderResult(httpStatus, responseText, w)

	log.Printf("Actsctrl request processed for IP %s", req.IP)
}

func unpackReqBody(r *http.Request) (int, string, []byte) {
	var httpStatus int
	var responseText string

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if len(body) == 0 {
		httpStatus = http.StatusUnprocessableEntity
		responseText = "No body provided!"
	}

	return httpStatus, responseText, body
}

func unmarshalReqBody(body []byte) (int, string, model.Request) {
	var httpStatus int
	var responseText string

	var req model.Request
	if err := json.Unmarshal(body, &req); err != nil {
		httpStatus = http.StatusUnprocessableEntity
		responseText = err.Error()
	}
	if req.IP == "" {
		httpStatus = http.StatusUnprocessableEntity
		responseText = "No IP is provided!"
	}
	return httpStatus, responseText, req
}

func ifExceedsLimit(req model.Request) (int, string) {
	var httpStatus int
	var responseText string

	m := map[int]int{
		1:    config.Conf.ReqPerSec,
		60:   config.Conf.ReqPerMin,
		3600: config.Conf.ReqPerHour,
	}

	for interval, limit := range m {
		key := fmt.Sprintf("%s.%d", req.IP, interval)
		keyExists, err := redisclient.Exists(key)
		if err == nil {
			if keyExists == 1 {
				v, err := redisclient.Incr(key)
				if err == nil {
					if v > int64(limit) {
						responseText = "LIMIT_EXCEEDED"
					}
				}
			} else {
				redisclient.Set(key, 1, time.Second*time.Duration(interval))
			}
		} else {
			httpStatus = http.StatusInternalServerError
			responseText = "REDIS_ERROR"
		}
	}
	return httpStatus, responseText
}

func renderResult(httpStatus int, responseText string, w http.ResponseWriter) {
	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}
	if responseText == "" {
		responseText = "OK"
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	response := map[string]string{
		"res": responseText,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
