package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Option is the standard response shape for all dropdown endpoints.
type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// MockPromotionController handles all Mock API endpoints for onboarding promotion.
type MockPromotionController struct {
	divisions         []Option
	regions           []Option
	environments      []Option
	bifrostCategories []Option
}

// NewMockPromotionController creates a new controller with mock data.
func NewMockPromotionController() *MockPromotionController {
	return &MockPromotionController{
		divisions: []Option{
			{Value: "Cps",      Label: "Cps"},
			{Value: "Sec",      Label: "Sec"},
			{Value: "Standard", Label: "Standard"},
		},

		// Regions — expand when more are confirmed from JARVIS
		regions: []Option{
			{Value: "eus2", Label: "eus2"},
		},

		environments: []Option{
			{Value: "dev", Label: "dev"},
			{Value: "qa",  Label: "qa"},
		},

		bifrostCategories: []Option{
			{Value: "Bifrost - Application Software Manual install",          Label: "Bifrost - Application Software Manual install"},
			{Value: "Bifrost - Application Software Partially Automated",     Label: "Bifrost - Application Software Partially Automated"},
			{Value: "Bifrost - Application Software install Full utilization", Label: "Bifrost - Application Software install Full utilization"},
			{Value: "Not Required", Label: "Not Required"},
		},
	}
}

// MockPromotionRoutes registers all routes under the /mock-api/jarvis/promotion group.
func (m *MockPromotionController) MockPromotionRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/mock-api/jarvis/promotion")
	g.GET("/health",             m.Health)
	g.GET("/divisions",          m.GetDivisions)
	g.GET("/regions",            m.GetRegions)
	g.GET("/environments",       m.GetEnvironments)
	g.GET("/bifrost-categories", m.GetBifrostCategories)
}

// GET /mock-api/jarvis/promotion/health
func (m *MockPromotionController) Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok", "service": "jarvis-promotion-mock-api"})
}

// GET /mock-api/jarvis/promotion/divisions
// Returns all deployment divisions.
func (m *MockPromotionController) GetDivisions(c *gin.Context) {
	c.JSON(200, m.divisions)
}

// GET /mock-api/jarvis/promotion/regions
// Returns all available promoted-from regions.
func (m *MockPromotionController) GetRegions(c *gin.Context) {
	c.JSON(200, m.regions)
}

// GET /mock-api/jarvis/promotion/environments
// Returns all available promoted-from environments.
func (m *MockPromotionController) GetEnvironments(c *gin.Context) {
	c.JSON(200, m.environments)
}

// GET /mock-api/jarvis/promotion/bifrost-categories
// Returns all Bifrost application categories.
func (m *MockPromotionController) GetBifrostCategories(c *gin.Context) {
	c.JSON(200, m.bifrostCategories)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS — allow all origins (Harness IDP proxy + local preview)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	controller := NewMockPromotionController()

	// Routes registered under /hoover-service to match proxy pathRewrite:
	//   proxy/platform-api/hoover-service/mock-api/jarvis/promotion/* -> /hoover-service/mock-api/jarvis/promotion/*
	hooverGroup := r.Group("/hoover-service")
	controller.MockPromotionRoutes(hooverGroup)

	log.Printf("Jarvis Promotion Mock API running on :8084")
	log.Printf("Routes: /hoover-service/mock-api/jarvis/promotion/*")
	if err := r.Run(":8084"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
