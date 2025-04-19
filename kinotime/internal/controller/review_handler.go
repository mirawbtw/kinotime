package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"kinotime/internal/models"
	"kinotime/internal/repository"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	ReviewRepo *repository.ReviewRepository
}

func NewReviewHandler(repo *repository.ReviewRepository) *ReviewHandler {
	return &ReviewHandler{ReviewRepo: repo}
}

func (h *ReviewHandler) HandleCreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	err := h.ReviewRepo.CreateReview(c, review.UserID, review.MovieID, review.Rating, review.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review created successfully"})
}

func (h *ReviewHandler) HandleGetReviewByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	review, err := h.ReviewRepo.GetReviewByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"review": review})
}

func (h *ReviewHandler) HandleGetReviewsByMovieID(c *gin.Context) {
	movieIDStr := c.Param("movie_id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	reviews, err := h.ReviewRepo.GetReviewsByMovieID(c, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func (h *ReviewHandler) HandleUpdateReview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	err = h.ReviewRepo.UpdateReview(c, id, review.Rating, review.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

func (h *ReviewHandler) HandleDeleteReview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	err = h.ReviewRepo.DeleteReview(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
