package handler

import (
	"context"
	"net/http"
	"strings"

	pb "urlshortener/internal/app/proto"
	lgc "urlshortener/internal/logic"
)

// HandlerMapGet Мапа для храненения длинных юрл.
var HandlerMapGet map[string]string

// HandlerMapPost Мапа для хранения коротких юрл.
var HandlerMapPost map[string]string

// UserServer структура для реализации RPC.
type UserServer struct {
	pb.UnimplementedMyServiceServer
}

// GetFuncRPC Хендлер для RPC аналогичен GetFunc для HTTP-сервера.
func (s *UserServer) GetFuncRPC(ctx context.Context, req *pb.GetFuncRequest) (*pb.GetFuncResponse, error) {
	request, _ := http.NewRequest("GET", req.Url.Url, nil)
	resLF, out := lgc.LogicGetFunc(request, HandlerMapGet)

	response := &pb.GetFuncResponse{}
	switch {
	case resLF == http.StatusTemporaryRedirect:
		response.Out.Out = out
		response.Status = http.StatusTemporaryRedirect
	case resLF == http.StatusGone:
		response.Status = http.StatusGone
	case resLF == http.StatusBadRequest:
		response.Status = http.StatusBadRequest
	}

	return response, nil
}

// PostFuncRPC Хендлер для RPC аналогичен PostFunc для HTTP-сервера.
func (s *UserServer) PostFuncRPC(ctx context.Context, req *pb.PostFuncRequest) (*pb.PostFuncResponse, error) {
	request, _ := http.NewRequest("POST", req.Addresss, strings.NewReader(req.Longurl.Url))
	var wr http.ResponseWriter
	resFL, byteRes := lgc.LogicPostFunc(wr, request, HandlerMapPost, HandlerMapGet)

	response := &pb.PostFuncResponse{}
	switch {
	case resFL == http.StatusCreated:
		response.Status = http.StatusCreated
		response.Out = byteRes
	case resFL == http.StatusConflict:
		response.Status = http.StatusConflict
		response.Out = byteRes
	case resFL == http.StatusBadRequest:
		response.Status = http.StatusBadRequest
	default:
		response.Status = http.StatusInternalServerError
	}

	return response, nil
}

// GetFuncPingRPC Хендлер для RPC аналогичен GetFuncPing для HTTP-сервера.
func GetFuncPingRPC(ctx context.Context, req *pb.GetFuncPingRequest) (*pb.GetFuncPingResponse, error) {
	resFL := lgc.LogicGetFuncPing(lgc.ResHandParam.DataBaseDSN)
	response := &pb.GetFuncPingResponse{}
	switch {
	case resFL == http.StatusOK:
		response.Status = http.StatusOK
	case resFL == http.StatusInternalServerError:
		response.Status = http.StatusInternalServerError
	}
	return response, nil
}
