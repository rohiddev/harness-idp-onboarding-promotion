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
	divisions            []Option
	regions              []Option
	environments         []Option
	promoteToEnvironments []Option
	bifrostCategories    []Option
	vpcOptions           []Option
}

// NewMockPromotionController creates a new controller with mock data.
func NewMockPromotionController() *MockPromotionController {
	return &MockPromotionController{

		divisions: []Option{
			{Value: "Colleaguetech",     Label: "Colleaguetech"},
			{Value: "Commercial",        Label: "Commercial"},
			{Value: "CommercialConnect", Label: "CommercialConnect"},
			{Value: "Consumer",          Label: "Consumer"},
			{Value: "ConsumerConnect",   Label: "ConsumerConnect"},
			{Value: "Digital",           Label: "Digital"},
			{Value: "Edo",               Label: "Edo"},
			{Value: "Entsvc",            Label: "Entsvc"},
			{Value: "Exvendor",          Label: "Exvendor"},
		},

		regions: []Option{
			{Value: "us-east-1",    Label: "us-east-1 (N. Virginia)"},
			{Value: "us-east-2",    Label: "us-east-2 (Ohio)"},
			{Value: "us-west-2",    Label: "us-west-2 (Oregon)"},
			{Value: "eu-west-1",    Label: "eu-west-1 (Ireland)"},
			{Value: "eu-central-1", Label: "eu-central-1 (Frankfurt)"},
			{Value: "eastus",       Label: "eastus (Azure East US)"},
			{Value: "eastus2",      Label: "eastus2 (Azure East US 2)"},
			{Value: "westeurope",   Label: "westeurope (Azure West Europe)"},
		},

		// promoted-from environments (where the app currently lives)
		environments: []Option{
			{Value: "dev", Label: "dev"},
			{Value: "qa",  Label: "qa"},
		},

		// promote-to environments (target)
		promoteToEnvironments: []Option{
			{Value: "qa",   Label: "qa"},
			{Value: "prod", Label: "prod"},
		},

		bifrostCategories: []Option{
			{Value: "Bifrost - Application Software Manual Install",           Label: "Bifrost - Application Software Manual Install"},
			{Value: "Bifrost - Application Software Partially Automated",      Label: "Bifrost - Application Software Partially Automated"},
			{Value: "Bifrost - Application Software Install Full Utilization", Label: "Bifrost - Application Software Install Full Utilization"},
			{Value: "Not Required", Label: "Not Required"},
		},

		vpcOptions: []Option{
			{Value: "default",      Label: "default"},
			{Value: "custom-vpc-1", Label: "custom-vpc-1"},
			{Value: "custom-vpc-2", Label: "custom-vpc-2"},
		},
	}
}

// Register registers all routes under the provided router group.
func (m *MockPromotionController) Register(rg *gin.RouterGroup) {
	g := rg.Group("/mock-api/jarvis/promotion")
	g.GET("/health",                  m.Health)
	g.GET("/divisions",               m.GetDivisions)
	g.GET("/regions",                 m.GetRegions)
	g.GET("/environments",            m.GetEnvironments)
	g.GET("/promote-to-environments", m.GetPromoteToEnvironments)
	g.GET("/bifrost-categories",      m.GetBifrostCategories)
	g.GET("/vpc-options",             m.GetVpcOptions)
}

func (m *MockPromotionController) Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok", "service": "jarvis-promotion-mock-api"})
}

func (m *MockPromotionController) GetDivisions(c *gin.Context) {
	c.JSON(200, m.divisions)
}

func (m *MockPromotionController) GetRegions(c *gin.Context) {
	c.JSON(200, m.regions)
}

// GET /environments — promoted-from (where the app currently lives)
func (m *MockPromotionController) GetEnvironments(c *gin.Context) {
	c.JSON(200, m.environments)
}

// GET /promote-to-environments — target environment for this promotion
func (m *MockPromotionController) GetPromoteToEnvironments(c *gin.Context) {
	c.JSON(200, m.promoteToEnvironments)
}

func (m *MockPromotionController) GetBifrostCategories(c *gin.Context) {
	c.JSON(200, m.bifrostCategories)
}

// GET /vpc-options — used for both internalVpc and ingressVpc
func (m *MockPromotionController) GetVpcOptions(c *gin.Context) {
	c.JSON(200, m.vpcOptions)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

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

	hooverGroup := r.Group("/hoover-service")
	NewMockPromotionController().Register(hooverGroup)

	log.Printf("Jarvis Promotion Mock API running on :8084")
	log.Printf("Routes: /hoover-service/mock-api/jarvis/promotion/*")
	if err := r.Run(":8084"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
