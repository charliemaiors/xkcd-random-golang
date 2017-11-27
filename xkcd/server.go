package xkcd

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	xkcd "github.com/nishanths/go-xkcd"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"
)

var (
	router *httprouter.Router
)

func init() {
	router = httprouter.New()
	router.GET("/", handleGetRandom)
}

//RunSrv is the main routine
func RunSrv() {
	domain := viper.GetString("domain")
	certDir := viper.GetString("certdir")

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache(certDir),
	}

	server := &http.Server{
		Addr: ":443", //Different port from 443 could be hard for acme-tlsni-challenge
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
		Handler: router,
	}
	err := server.ListenAndServeTLS("", "")

	if err != nil {
		panic(err)
	}
}

func handleGetRandom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	title, url := getRandom()
	enc := json.NewEncoder(w)

	message := struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}{Title: title, URL: url}

	err := enc.Encode(&message)
	if err != nil {
		panic(err) //Should not be handled in this way
	}
}

func getRandom() (string, string) {
	client := xkcd.NewClient()
	comic, err := client.Random()
	if err != nil {
		return "error", err.Error()
	}
	return comic.Title, comic.URL
}
