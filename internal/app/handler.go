package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// GetFunc Обработчик для Get запросов.
func GetFunc(handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resLF, out := LogicGetFunc(r, handMapGet)
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
		resFL, byteRes := LogicPostFunc(w, r, handMapPost, handMapGet)
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
		resFL, byteRes := LogicPostFuncAPIShorten(handMapPost, handMapGet, w, r)
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
		resFL, byteRes := LogicGetFuncAPIUserUrls(handMapGet, w, r)
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
func GetFuncPing(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL := LogicGetFuncPing(db)
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
		resFL, byteRes := LogicPostFuncAPIShortenBatch(handMapPost, handMapGet, w, r)
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
func DeleteFuncAPIUserURLs(handMapPost map[string]string, handMapGet map[string]string, db *sql.DB, dbf string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resFL := LogicDeleteFuncAPIUserURLs(handMapPost, handMapGet, db, dbf, r)
		switch {
		case resFL == http.StatusAccepted:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
