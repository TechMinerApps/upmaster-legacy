package oauth

import (
	"net/http"

	"log"

	"gorm.io/gorm"
	oauth2gorm "src.techknowlogick.com/oauth2-gorm"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

// Client is a config for upmaster-agent
type Client struct {
	ID     string
	Secret string
	Domain string
}

// Config is the configuration required to create a Server
type Config struct {
	DB         *gorm.DB
	DBName     string
	GCInterval int
	JWTKey     []byte
	Clients    []Client /* A list of 'Client'. Since agents can shared same client credentials,
	the length of this list will probably always be one */
}

// Server is a object of OAuth Server
type Server struct {
	Handler http.Handler      // Handler can serve http requests
	Config  Config            // Config is stored here incase needed
	store   *oauth2gorm.Store // no one will likely use it, so private
}

func (s *Server) createHandler() {

	manager := manage.NewDefaultManager()

	// Token should be stored by GORM
	manager.MapTokenStorage(s.store)

	// Client is stored in memory
	clientStore := store.NewClientStore()
	for _, c := range s.Config.Clients {
		clientStore.Set(c.ID, &models.Client{
			ID:     c.ID,
			Secret: c.Secret,
			Domain: c.Domain,
		})
	}
	manager.MapClientStorage(clientStore)

	manager.MapAccessGenerate( // Hardcoded signing method, should be safe enough
		generates.NewJWTAccessGenerate("", s.Config.JWTKey, jwt.SigningMethodHS512))

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	handler := http.NewServeMux()

	handler.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	handler.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	s.Handler = handler
}

func (s *Server) createStore() {
	config := oauth2gorm.Config{
		TableName: s.Config.DBName,
	}
	s.store = oauth2gorm.NewStoreWithDB(&config, s.Config.DB, s.Config.GCInterval)
}

// NewServer generates a Server instance
func NewServer(c Config) (*Server, error) {
	var s Server
	s.Config = c
	s.createStore()
	s.createHandler()

	return &s, nil
}
