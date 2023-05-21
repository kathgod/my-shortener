package handler

import (
	"fmt"
	"log"
	"net/http"

	lgc "urlshortener/internal/logic"
)

const (
	writeerr = "Write error"
)

// GetFunc Обработчик для Get запросов.
func GetFunc(handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resLF, out := lgc.LogicGetFunc(r, handMapGet)
		fmt.Println(out)
		switch {
		case resLF == http.StatusTemporaryRedirect:
			fmt.Println(out)
			w.Header().Set("Location", out)
			w.WriteHeader(http.StatusTemporaryRedirect)
		case resLF == http.StatusGone:
			w.WriteHeader(http.StatusGone)
		case resLF == http.StatusBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// PostFunc Обработчик Post запросов.
func PostFunc(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL, byteRes := lgc.LogicPostFunc(w, r, handMapPost, handMapGet)
		switch {
		case resFL == http.StatusCreated:
			w.WriteHeader(http.StatusCreated)
			_, err := w.Write(byteRes)
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}
		case resFL == http.StatusConflict:
			w.WriteHeader(http.StatusConflict)
			_, err := w.Write(byteRes)
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}
		case resFL == http.StatusBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// PostFuncAPIShorten Обработчик Post запросов.
func PostFuncAPIShorten(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL, byteRes := lgc.LogicPostFuncAPIShorten(handMapPost, handMapGet, w, r)
		switch {
		case resFL == http.StatusCreated:
			w.WriteHeader(http.StatusCreated)
			_, err1 := w.Write(byteRes)
			if err1 != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}
		case resFL == http.StatusConflict:
			w.WriteHeader(http.StatusConflict)
			_, err1 := w.Write(byteRes)
			if err1 != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}
		case resFL == http.StatusBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetFuncAPIUserUrls Хэндлер возвращает объект json-array, со всеми длинными и короткими URL которые создал юзер.
func GetFuncAPIUserUrls(handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL, byteRes := lgc.LogicGetFuncAPIUserUrls(handMapGet, w, r)
		switch {
		case resFL == http.StatusNoContent:
			w.WriteHeader(http.StatusNoContent)
		case resFL == http.StatusOK:
			w.Header().Set("Content-Type", "application/json")
			_, err2 := w.Write(byteRes)
			if err2 != nil {
				log.Println(writeerr)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetFuncPing Хэндлер пинга базы данных.
func GetFuncPing(DBDSN string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL := lgc.LogicGetFuncPing(DBDSN)
		switch {
		case resFL == http.StatusOK:
			w.WriteHeader(http.StatusOK)
		case resFL == http.StatusInternalServerError:
			w.WriteHeader(http.StatusOK)
		}
	}
}

// PostFuncAPIShortenBatch Хэндлер, принимающий в теле запроса множество URL для сокращения.
func PostFuncAPIShortenBatch(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL, byteRes := lgc.LogicPostFuncAPIShortenBatch(handMapPost, handMapGet, w, r)
		switch {
		case resFL == http.StatusBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		case resFL == http.StatusCreated:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, err2 := w.Write(byteRes)
			if err2 != nil {
				log.Println(writeerr)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// DeleteFuncAPIUserURLs Хэндлер, принимающая список идентификаторов сокращённых URL для удаления.
func DeleteFuncAPIUserURLs(handMapPost map[string]string, handMapGet map[string]string, dbf string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL := lgc.LogicDeleteFuncAPIUserURLs(handMapPost, handMapGet, dbf, r)
		switch {
		case resFL == http.StatusAccepted:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetFuncAPIInternalStats Хендлер для полчения всех колличества юзеров и сокращенных URL.
func GetFuncAPIInternalStats(handMapPost map[string]string, ts string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL, byteRes := lgc.LogicGetFuncAPIInternalStats(handMapPost, ts, w, r)
		switch {
		case resFL == http.StatusForbidden:
			w.WriteHeader(http.StatusForbidden)
		case resFL == http.StatusOK:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(byteRes)
			if err != nil {
				log.Println(writeerr)
				log.Println(err)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
