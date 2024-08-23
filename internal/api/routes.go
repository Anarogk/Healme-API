package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Anarogk/healthcare-api/configs"
	"github.com/Anarogk/healthcare-api/internal/api/handlers"
	"github.com/Anarogk/healthcare-api/internal/api/middleware"
)

type Server struct {
	config             configs.Config
	db                 *gorm.DB
	router             *gin.Engine
	appointmentHandler handlers.AppointmentHandler
	patientHandler     handlers.PatientHandler
	doctorHandler      handlers.DoctorHandler
	authHandler        handlers.AuthHandler // New handler for login and register
}

func NewServer(config configs.Config, db *gorm.DB) *Server {
	return &Server{
		config:             config,
		db:                 db,
		router:             gin.Default(),
		appointmentHandler: handlers.AppointmentHandler{DB: db},
		patientHandler:     handlers.PatientHandler{DB: db},
		doctorHandler:      handlers.DoctorHandler{DB: db},
		authHandler:        handlers.AuthHandler{DB: db, Config: config},
	}
}

func (s *Server) SetupRoutes() {
	s.router.Use(middleware.Logging())

	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware(s.config))

	// Public routes
	api.POST("/login", s.authHandler.Login)
	api.POST("/register", s.authHandler.Register)

	// Patient routes
	patients := api.Group("/patients")
	patients.Use(middleware.RoleMiddleware("patient"))
	patients.GET("", s.patientHandler.GetPatients)
	patients.GET("/:id", s.patientHandler.GetPatient)
	patients.PUT("/:id", s.patientHandler.UpdatePatient)
	patients.GET("/appointments", s.appointmentHandler.GetAppointments)
	patients.POST("/appointments", s.appointmentHandler.CreateAppointment)
	patients.PUT("/appointments/:id", s.appointmentHandler.UpdateAppointment)
	patients.DELETE("/appointments/:id", s.appointmentHandler.DeleteAppointment)

	// Doctor routes
	doctors := api.Group("/doctors")
	doctors.Use(middleware.RoleMiddleware("doctor"))
	doctors.GET("", s.doctorHandler.GetDoctors)
	doctors.GET("/:id", s.doctorHandler.GetDoctor)
	doctors.PUT("/:id", s.doctorHandler.UpdateDoctor)
	doctors.GET("/appointments", s.appointmentHandler.GetAppointments)
	doctors.PUT("/appointments/:id", s.appointmentHandler.UpdateAppointment)

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.RoleMiddleware("admin"))
	admin.POST("/patients", s.patientHandler.CreatePatient)
	admin.DELETE("/patients/:id", s.patientHandler.DeletePatient)
	admin.POST("/doctors", s.doctorHandler.CreateDoctor)
	admin.DELETE("/doctors/:id", s.doctorHandler.DeleteDoctor)
	admin.GET("/appointments", s.appointmentHandler.GetAppointments)
	admin.DELETE("/appointments/:id", s.appointmentHandler.DeleteAppointment)
}

func (s *Server) Start() error {
	s.SetupRoutes()
	return s.router.Run(":" + s.config.ServerPort)
}
