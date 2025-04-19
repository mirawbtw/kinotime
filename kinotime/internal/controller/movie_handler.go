package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"kinotime/internal/models"
	"kinotime/internal/repository"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieRepo *repository.MovieRepository
}

func NewMovieHandler(repo *repository.MovieRepository) *MovieHandler {
	return &MovieHandler{MovieRepo: repo}
}

func (h *MovieHandler) HandleCreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	err := h.MovieRepo.CreateMovie(c, movie.Title, movie.PosterUrl, movie.Genre, movie.Description, movie.Year, movie.Actors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie created successfully"})
}

func (h *MovieHandler) HandleGetMovieByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	movie, err := h.MovieRepo.GetMovieByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (h *MovieHandler) HandleGetAllMovies(c *gin.Context) {
	movies, err := h.MovieRepo.GetAllMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (h *MovieHandler) HandleUpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	err = h.MovieRepo.UpdateMovie(c, id, movie.Title, movie.PosterUrl, movie.Genre, movie.Description, movie.Year, movie.Actors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func (h *MovieHandler) HandleDeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	err = h.MovieRepo.DeleteMovie(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
