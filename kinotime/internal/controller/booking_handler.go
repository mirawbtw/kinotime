package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"kinotime/internal/models"
	"kinotime/internal/repository"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	BookingRepo *repository.BookingRepository
}

func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{BookingRepo: repo}
}

func (h *BookingHandler) HandleCreateBooking(c *gin.Context) {
	var bookingReq models.BookingRequest
	if err := c.ShouldBindJSON(&bookingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := h.BookingRepo.CreateBooking(c, userID.(int), bookingReq.MovieID, bookingReq.SeatsBooked, bookingReq.TotalPrice, bookingReq.Status, bookingReq.BookingTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}

func (h *BookingHandler) HandleGetBookingByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	booking, err := h.BookingRepo.GetBookingByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}

func (h *BookingHandler) HandleGetAllBookings(c *gin.Context) {
	bookings, err := h.BookingRepo.GetAllBookings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func (h *BookingHandler) HandleUpdateBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	err = h.BookingRepo.UpdateBooking(c, id, booking.SeatsBooked, booking.TotalPrice, booking.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

func (h *BookingHandler) HandleDeleteBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	err = h.BookingRepo.DeleteBooking(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
