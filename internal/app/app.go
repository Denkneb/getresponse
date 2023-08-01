package app

import (
	"context"
	// "errors"
	// "getresponse/internal/handlers"
	"net/http"
	// "os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	// "github.com/joho/godotenv"
)

var DB *pgxpool.Pool

type App struct {
	Config Config
	Router *chi.Mux
	DB     *pgxpool.Pool
}

type Config struct {
	DatabaseUrl string
}

// func (config *Config) ConfigLoad(filename string) error {
// 	err := godotenv.Load(filename)
// 	if err != nil {
// 		return errors.New("Error loading " + filename)
// 	}

// 	config.DatabaseUrl = os.Getenv("DATABASE_URL")
// 	if config.DatabaseUrl == "" {
// 		return errors.New("missing DATABASE_URL")
// 	}

// 	return nil
// }

func NewApp() App {
	return App{}
}

func (app *App) Setup() error {
	poolConfig, err := pgxpool.ParseConfig(app.Config.DatabaseUrl)
	if err != nil {
		return err
	}
	DB, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}
	app.DB = DB

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Post("/v1/webhook", handlers.GetresponseEvent(DB))

	app.Router = r

	return nil
}

func (app *App) Run() error {
	return http.ListenAndServe(":3000", app.Router)
}
