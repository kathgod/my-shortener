package handler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const resServError = "Server working result Error"

func Test_RandSeq(t *testing.T) {
	type want struct {
		ln  int // длина для последовательности
		res string
	}
	tests := []struct {
		name string
		arg  int
		want want
	}{
		{
			name: "test1 Length of return value",
			arg:  6,
			want: want{
				ln:  6,
				res: "123456",
			},
		},
		{
			name: "Test 2 type of return value",
			arg:  5,
			want: want{
				ln:  5,
				res: "12345",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, len(randSeq(tt.arg)), tt.want.ln, "\"randSeq() = %v, want %v\", len(randSeq(tt.arg)), tt.want.ln")
		})
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, reflect.TypeOf(randSeq(tt.arg)), reflect.TypeOf(tt.want.res), "randSeq() = %v, want %v", reflect.TypeOf(randSeq(tt.arg)), reflect.TypeOf(tt.want.res))
		})
	}
}

func TestFunc(t *testing.T) {
	type args struct {
		mPost map[string]string
		mGet  map[string]string
	}

	type want struct {
		codeP int
		codeG int
	}

	map2P := map[string]string{
		"vk.com":     "RPtDVz",
		"google.com": "XvhXrs",
		"yandex.com": "WDSMzc",
	}

	map2G := map[string]string{
		"RPtDVz": "vk.com",
		"XvhXrs": "google.com",
		"WDSMzc": "yandex.com",
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test 1",
			args: args{
				mPost: map[string]string{},
				mGet:  map[string]string{},
			},
			want: want{
				codeP: 201,
				codeG: 400,
			},
		},
		{
			name: "Test 2",
			args: args{
				mPost: map2P,
				mGet:  map2G,
			},
			want: want{
				codeP: 201,
				codeG: 307,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//Задаем тело запроса Post
			bodyP := bytes.NewBufferString("vk.com")

			//Задаем пост реквест
			requestP := httptest.NewRequest(http.MethodPost, "/", bodyP)

			// Создаем рекордер
			recP := httptest.NewRecorder()

			//Создание map для запроса Post
			var mp1P = tt.args.mPost
			var mp1G = tt.args.mGet

			//Присвоение функции хендлер с заданными параметрами
			resMfP := PostFunc(mp1P, mp1G) // Mf в имени - My function

			//Определение хендлера
			h := http.HandlerFunc(resMfP)

			//запуск сервера
			h.ServeHTTP(recP, requestP)

			//записываем результат работы сервера, через результат рекордера
			resP := recP.Result()
			err := resP.Body.Close()
			if err != nil {
				log.Println(resServError)
			}

			//Сверяем возвращаемый код
			if resP.StatusCode != tt.want.codeP {
				t.Errorf("Expected status code %d, got %d", tt.want.codeP, recP.Code)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			//Задаем Гет реквест
			requestG := httptest.NewRequest(http.MethodGet, "/XvhXrs", nil)

			//Создаем новый рекордер для Гет
			recG := httptest.NewRecorder()

			//Создание map для запроса Get
			var mp1P = tt.args.mPost
			var mp1G = tt.args.mGet

			//Присвоение функции хендлер с заданными параметрами
			resMfG := GetFunc(mp1P, mp1G)

			//Определение хендлера
			h1 := http.HandlerFunc(resMfG)

			//запуск сервера
			h1.ServeHTTP(recG, requestG)

			//записываем результат работы сервера, через результат рекордера
			resG := recG.Result()

			err := resG.Body.Close()
			if err != nil {
				log.Println(resServError)
			}

			//Сверяем возвращаемый код

			if resG.StatusCode != tt.want.codeG {
				t.Errorf("Expected status code %d, got %d", tt.want.codeG, recG.Code)
			}

			H := tt.args.mGet["XvhXrs"]

			if resG.Header.Get("Location") != H {
				t.Errorf("Expected Content-Type %s, got %s", H, resG.Header.Get("Location"))
			}

		})
	}
}
