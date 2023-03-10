package v1

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	/*
		Stores server infomation and access to database interface
	*/
	listenAddr string
	store      Storage
}

type ApiFunc func(http.ResponseWriter, *http.Request) error // handler signature

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"success"`
}

func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *ApiServer) Run() {
	/*
		Url controller using gorilla mux
	*/
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/auth/register", makeHttpHandler(s.handleSignUpUser))
	router.HandleFunc("/api/v1/auth/login", makeHttpHandler(s.handleLoginUser))
	router.HandleFunc("/api/v1/otp/generate", makeHttpHandler(s.handleGenerateOTP))
	router.HandleFunc("/api/v1/otp/verify", makeHttpHandler(s.handleVerifyOTP))
	router.HandleFunc("/api/v1/otp/validate", makeHttpHandler(s.handleValidateOtp))
	log.Println("Starting the server at port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
