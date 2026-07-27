package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"https://github.com/aguspurnomomz/jurnalmotorik-be/config"
	"https://github.com/aguspurnomomz/jurnalmotorik-be/models"
)

// Request Structs untuk Validasi Payload JSON
type RegisterRequest struct {
	Username       string  `json:"username" binding:"required"`
	Fullname       string  `json:"fullname" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	Password       string  `json:"password" binding:"required,min=6"`
	WhatsappNumber string  `json:"whatsapp_number"`
	ChildName      string  `json:"child_name" binding:"required"`
	ChildBirthDate string  `json:"child_birth_date" binding:"required"` // Format: YYYY-MM-DD
	ChildGender    string  `json:"child_gender" binding:"required"`    // "L" / "P"
	ChildHeight    float64 `json:"child_height"`
	ChildWeight    float64 `json:"child_weight"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 1. REGISTER PARENT & FIRST CHILD
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enkripsi Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password"})
		return
	}

	// Transaksi Database (Simpan User & Child sekaligus)
	tx := config.DB.Begin()

	newUser := models.User{
		Username:       req.Username,
		Fullname:       req.Fullname,
		Email:          req.Email,
		PasswordHash:   string(hashedPassword),
		WhatsappNumber: req.WhatsappNumber,
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau Email sudah terdaftar"})
		return
	}

	newChild := models.Child{
		UserID:    newUser.ID,
		Name:      req.ChildName,
		BirthDate: req.ChildBirthDate,
		Gender:    req.ChildGender,
		HeightCm:  req.ChildHeight,
		WeightKg:  req.ChildWeight,
	}

	if err := tx.Create(&newChild).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftarkan data anak"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil!",
		"data": gin.H{
			"user":  newUser,
			"child": newChild,
		},
	})
}

// 2. LOGIN PARENT
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Preload("Children").Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	// Verifikasi Password Hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil!",
		"data":    user,
	})
}