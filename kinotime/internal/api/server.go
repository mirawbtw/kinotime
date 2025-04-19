package api

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kinotime/internal/configs"
	"kinotime/internal/controller"
	"kinotime/internal/middleware"
	"kinotime/internal/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg         *configs.Config
	srv         *http.Server
	userRepo    *repository.UserRepository
	movieRepo   *repository.MovieRepository
	bookingRepo *repository.BookingRepository
	reviewRepo  *repository.ReviewRepository
}

func NewServer(cfg *configs.Config, db *sql.DB) *Server {
	engine := gin.Default()

	server := &Server{
		cfg:         cfg,
		userRepo:    repository.NewUserRepository(db),
		movieRepo:   repository.NewMovieRepository(db),
		bookingRepo: repository.NewBookingRepository(db),
		reviewRepo:  repository.NewReviewRepository(db),
		srv: &http.Server{
			Addr:    cfg.Port,
			Handler: engine,
		},
	}

	server.setRoutes()

	return server
}

func (s *Server) Start() error {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start " + err.Error())
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	return nil
}

func (s *Server) setRoutes() {
	engine := s.srv.Handler.(*gin.Engine)

	engine.Use(middleware.CORSMiddleware())

	// Public routes
	engine.POST("/login", func(c *gin.Context) {
		controller.HandleLogin(c, s.userRepo)
	})

	engine.POST("/register", func(c *gin.Context) {
		controller.HandleRegister(c, s.userRepo)
	})

	// Movie routes (public)
	movieHandler := controller.NewMovieHandler(s.movieRepo)
	engine.GET("/movie/:id", movieHandler.HandleGetMovieByID)
	engine.GET("/movies", movieHandler.HandleGetAllMovies)

	// Review routes (public viewing)
	reviewHandler := controller.NewReviewHandler(s.reviewRepo)
	engine.GET("/review/:id", reviewHandler.HandleGetReviewByID)
	engine.GET("/reviews/movie/:movie_id", reviewHandler.HandleGetReviewsByMovieID)

	// Private routes (JWT protected)
	private := engine.Group("/private")
	private.Use(controller.JWTMiddleware())

	{
		private.GET("/profile", func(c *gin.Context) {
			controller.HandleProfile(c, s.userRepo)
		})

		// Movie management (admin only)
		private.POST("/movie", movieHandler.HandleCreateMovie)
		private.PUT("/movie/:id", movieHandler.HandleUpdateMovie)
		private.DELETE("/movie/:id", movieHandler.HandleDeleteMovie)

		// Booking routes (protected)
		bookingHandler := controller.NewBookingHandler(s.bookingRepo)
		private.POST("/booking", bookingHandler.HandleCreateBooking)
		private.GET("/booking/:id", bookingHandler.HandleGetBookingByID)
		private.GET("/bookings", bookingHandler.HandleGetAllBookings)
		private.PUT("/booking/:id", bookingHandler.HandleUpdateBooking)
		private.DELETE("/booking/:id", bookingHandler.HandleDeleteBooking)

		// Review routes (only for creating/updating reviews)
		private.POST("/review", reviewHandler.HandleCreateReview)
		private.PUT("/review/:id", reviewHandler.HandleUpdateReview)
		private.DELETE("/review/:id", reviewHandler.HandleDeleteReview)
	}
}

func (s *Server) Shutdown() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server failed to shutdown " + err.Error())
	}
}
