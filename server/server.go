package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

const CERT_FOLDER = ".certs"

type Server struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) NewHttpServer(sm *mux.Router) *http.Server {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("DOMAIN")),
		Cache:      autocert.DirCache(CERT_FOLDER),
	}
	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = s.getSelfSignedOrLetsEncryptCert(&certManager)
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      sm,
	}
}

func (s *Server) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = CERT_FOLDER
		}

		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			s.logger.Printf("falling back to letsencrypt due to %#v: \n", err)
			return certManager.GetCertificate(hello)
		}
		s.logger.Println("loaded self-signed certificate")
		return &certificate, err
	}
}
