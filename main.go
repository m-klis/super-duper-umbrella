package main

import (
	"context"
	"fmt"
	"gochicoba/db"
	"gochicoba/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database := db.DatabaseInitialize()
	addr := os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: InitializeRoute(database),
	}

	go func() {
		server.ListenAndServe()
	}()

	defer Stop(server)
	log.Printf("Started server on : %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch), "in server")
	log.Println("Stopping API Server")
}

func InitializeRoute(db *gorm.DB) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.MethodNotAllowed(handler.MethodNotAllowedHandler)
	router.NotFound(handler.NotFoundHandler)

	ih := ItemHandler(db)

	router.Route("/items", func(router chi.Router) {
		router.Get("/", ih.GetAllItems)
		router.Post("/", ih.CreateItem)
		router.Route("/{itemID}", func(router chi.Router) {
			// 	router.Use(ItemContext)
			router.Get("/", ih.GetItem)
			router.Put("/", ih.UpdateItem)
			router.Delete("/", ih.DeleteItem)
		})
	})

	uh := UserHandler(db)

	router.Route("/users", func(router chi.Router) {
		router.Get("/", uh.GetAllUsers)
		router.Post("/", uh.CreateUser)
		router.Route("/{userID}", func(router chi.Router) {
			router.Get("/", uh.GetUser)
			router.Put("/", uh.UpdateUser)
			router.Delete("/", uh.DeleteUser)
		})
	})

	ub := BuyHandler(db)

	router.Route("/buys", func(router chi.Router) {
		router.Get("/", ub.GetAllBuys)
		// router.Post("/", ub.CreateBuy)
		// router.Route("/{userID}", func(router chi.Router) {
		// 	router.Get("/", ub.GetBuy)
		// 	router.Put("/", ub.UpdateBuy)
		// 	router.Delete("/", ub.DeleteBuy)
		// })
	})

	return router
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server corectly: %v\n", err)
		os.Exit(1)
	}
}
