package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func WriteBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
}

func ParseJsonRequest(w http.ResponseWriter, r *http.Request, t interface{}) error {
	defer r.Body.Close()
	bodyData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		WriteBadRequest(w)
		return errors.New("bad request")
	}

	err = json.Unmarshal(bodyData, &t)

	if err != nil {
		WriteBadRequest(w)
		return errors.New("bad request")
	}

	return nil
}

func GetEntityId(w http.ResponseWriter, url *url.URL) (int, error) {
	pathParts := strings.Split(url.Path, "/")
	idPart := pathParts[len(pathParts)-1]
	id, err := strconv.Atoi(idPart)

	if err != nil {
		WriteBadRequest(w)
		return -1, errors.New("bad url")
	}

	return id, nil
}
