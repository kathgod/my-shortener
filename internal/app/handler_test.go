package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_RandSeq(t *testing.T) {
	type want struct {
		ln  int
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
			if got1 := len(RandSeq(tt.arg)); got1 != tt.want.ln {
				t.Errorf("randSeq() = %v, want %v", got1, tt.want.ln)
				fmt.Println("Tes1 Fail")
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got2 := reflect.TypeOf(RandSeq(tt.arg)); got2 != reflect.TypeOf(tt.want.res) {
				t.Errorf("randSeq() = %v, want %v", got2, tt.want.res)
				fmt.Println("Tes2 Fail")
			}
		})
	}
}

func Test_Myfunc(t *testing.T) {
	type args struct {
		m_p map[string]string
		m_g map[string]string
	}

	type want struct {
		code_p int
		code_g int
	}

	map2_p := map[string]string{
		"vk.com":     "RPtDVz",
		"google.com": "XvhXrs",
		"yandex.com": "WDSMzc",
	}

	map2_g := map[string]string{
		"RPtDVz": "vk.com",
		"XvhXrs": "google.com",
		"WDSMzc": "yandex.com",
	}

	//var buf1 string

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test 1",
			args: args{
				m_p: map[string]string{},
				m_g: map[string]string{},
			},
			want: want{
				code_p: 201,
				code_g: 400,
			},
		},
		{
			name: "Test 2",
			args: args{
				m_p: map2_p,
				m_g: map2_g,
			},
			want: want{
				code_p: 201,
				code_g: 307,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//Задаем тело запроса Post
			body_p := bytes.NewBufferString("vk.com")

			//Задаем пост реквест
			request_p := httptest.NewRequest(http.MethodPost, "/", body_p)

			// Создаем рекордер
			rec_p := httptest.NewRecorder()

			//Создание map для запроса Post
			var mp1_p map[string]string = tt.args.m_p
			var mp1_g map[string]string = tt.args.m_g

			//Присвоение функции хендлер с заданными параметрами
			res_Mf_p := PostFunc(mp1_p, mp1_g)

			//Определение хендлера
			h := http.HandlerFunc(res_Mf_p)

			//запуск сервера
			h.ServeHTTP(rec_p, request_p)

			//записываем результат работы сервера, через результат рекордера
			res_p := rec_p.Result()

			//Сверяем возвращаемый код
			if res_p.StatusCode != tt.want.code_p {
				t.Errorf("Expected status code %d, got %d", tt.want.code_p, rec_p.Code)
			}

			defer res_p.Body.Close()

		})
		t.Run(tt.name, func(t *testing.T) {
			//Задаем Гет реквест
			request_g := httptest.NewRequest(http.MethodGet, "/XvhXrs", nil)

			//Создаем новый рекордер для Гет
			rec_g := httptest.NewRecorder()

			//Создание map для запроса Get
			var mp1_p map[string]string = tt.args.m_p
			var mp1_g map[string]string = tt.args.m_g

			//Присвоение функции хендлер с заданными параметрами
			res_Mf_g := GetFunc(mp1_p, mp1_g)

			//Определение хендлера
			h1 := http.HandlerFunc(res_Mf_g)

			//запуск сервера
			h1.ServeHTTP(rec_g, request_g)

			//записываем результат работы сервера, через результат рекордера
			res_g := rec_g.Result()

			//Сверяем возвращаемый код
			if res_g.StatusCode != tt.want.code_g {
				t.Errorf("Expected status code %d, got %d", tt.want.code_g, rec_g.Code)
			}

			H := tt.args.m_g["XvhXrs"]

			if res_g.Header.Get("Location") != H {
				t.Errorf("Expected Content-Type %s, got %s", H, res_g.Header.Get("Location"))
			}
		})
	}
}
