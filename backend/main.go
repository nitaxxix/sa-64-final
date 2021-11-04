package main

import (
	"github.com/nitaxxix/sa-64-final/controller"

	"github.com/nitaxxix/sa-64-final/entity"

	"github.com/nitaxxix/sa-64-final/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// Appoint
			protected.GET("/appoints", controller.ListAppoint)
			protected.POST("/appoint", controller.CreateAppoint)

			// Insurance
			protected.GET("/insrs", controller.ListInsurance)
			protected.POST("/insr", controller.CreateInsurance)

			// Job
			protected.GET("/jobs", controller.ListJob)

			protected.POST("/job", controller.CreateJob)

			// MedicalProduct
			protected.GET("/medical_products", controller.ListMedicalProduct)
			protected.POST("/medical_product", controller.CreateMedicalProduct)

			// MedRecord
			protected.GET("/api/MedRec", controller.ListMedRecord)
			protected.POST("/api/submit", controller.CreateMedRecord)

			// Patient
			protected.GET("/patients", controller.ListPatient)
			protected.POST("/patient", controller.CreatePatient)

			// RemedyType
			protected.GET("/remedy_types", controller.ListRemedyType)
			protected.POST("/remedy_type", controller.CreateRemedyType)

			// Role
			protected.GET("/roles", controller.ListRole)
			protected.POST("/role", controller.CreateRole)

			// Screening
			protected.GET("/screenings", controller.ListScreening)
			protected.POST("/screening", controller.CreateScreening)

			// Sex
			protected.GET("/sexs", controller.ListSex)
			protected.POST("/sex", controller.CreateSex)

			// Treatment
			protected.POST("/treatmentRecord", controller.CreateTreatment)

			// User
			protected.GET("/users", controller.ListUser)
			protected.GET("/user/dentist/:id", controller.GetUserDentist)
			protected.GET("/user/dentistass", controller.GetUserDentistass)
			protected.GET("/user/nurse", controller.GetUserNurse)
			protected.GET("/user/pharmacist", controller.GetUserPharmacist)
			protected.GET("/user/financial", controller.GetUserFinancial)
			protected.POST("/user", controller.CreateUser)

		}
	}

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
