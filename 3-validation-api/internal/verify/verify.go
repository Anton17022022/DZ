package verify

import (
	"3-validation-api/configs"
	"3-validation-api/internal/temporarydb"
	"3-validation-api/pkg/req"
	"3-validation-api/pkg/res"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	*configs.Config
	Db temporarydb.TemporaryDb
}

type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps, Db temporarydb.TemporaryDb) {
	handler := &VerifyHandler{
		Config: deps.Config,
		Db:     Db,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandelBody[SendRequest](&w, r, &handler.StatusResponce)
		if err != nil {
			res.Json(&w, err.Error(), handler.StatusResponce.StatusCodeBadRequest)
		} else {

			hash, err := handler.Db.RegisteryAcc(&body.Email)
			if err != nil {
				res.Json(&w, err.Error(), handler.StatusResponce.StatusCodeBadRequest)
			}

			e := &email.Email{
				To:      []string{body.Email},
				From:    fmt.Sprintf("AP <%s>", handler.Config.Verify.Email),
				Subject: "Verify email",
				Text:    []byte(fmt.Sprintf("http://localhost:8081/verify/%s", hash)),
			}
			err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Config.Verify.Email, handler.Config.Verify.Password, handler.Config.Verify.Host))

			if err != nil {
				res.Json(&w, err.Error(), handler.StatusResponce.StatusCodeBadRequest)
			} else {

				response := SendResponse{}
				res.Json(&w, response, handler.StatusResponce.StatusCodeOk)
			}
		}
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) >= 3 {
			hash := parts[2]
			res.Json(&w, handler.Db.CheckAccountRegistration(hash), handler.Config.StatusResponce.StatusCodeOk)
		} else {
			res.Json(&w, "Not enough arguments. Check request.", handler.Config.StatusResponce.StatusCodeBadRequest)
		}
	}
}
