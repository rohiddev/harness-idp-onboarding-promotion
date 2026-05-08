package main

import (
	"fmt"
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

// TierItem is the shape of a single web or app tier row.
type TierItem struct {
	RepositoryName      string `json:"repositoryName"`
	OS                  string `json:"os"`
	BitbucketProjectKey string `json:"bitbucketProjectKey"`
	IsNotForPromotion   bool   `json:"isNotForPromotion"`
	IsBifrost           bool   `json:"isBifrost"`
	IsServerlessPattern bool   `json:"isServerlessPattern"`
	IsS3Only            bool   `json:"isS3Only"`
	IsPublicFacing      bool   `json:"isPublicFacing"`
	InstanceSize        string `json:"instanceSize"`
	NumOfInstances      string `json:"numOfInstances"`
	VolumeType          string `json:"volumeType"`
	VolumeSize          string `json:"volumeSize"`
	IsAutoScaling       bool   `json:"isAutoScaling"`
	IsStaticIp          bool   `json:"isStaticIp"`
	Loadbalancer        string `json:"loadbalancer"`
	IsFsx               bool   `json:"isFsx"`
	IsEfs               bool   `json:"isEfs"`
	IsElasticacheRedis  bool   `json:"isElasticacheRedis"`
}

// Register registers all routes under the provided router group.
func (m *MockPromotionController) Register(rg *gin.RouterGroup) {
	g := rg.Group("/mock-api/jarvis/promotion")
	g.GET("/health",                       m.Health)
	g.GET("/:cloudType/divisions",         m.GetDivisions)
	g.GET("/regions",                      m.GetRegions)
	g.GET("/environments",                 m.GetEnvironments)
	g.GET("/promote-to-environments",      m.GetPromoteToEnvironments)
	g.GET("/bifrost-categories",           m.GetBifrostCategories)
	g.GET("/vpc-options",                  m.GetVpcOptions)
	g.GET("/sample-apptiers",              m.GetSampleAppTiers)
	g.GET("/sample-webtiers",              m.GetSampleWebTiers)
}

func (m *MockPromotionController) MockPromotionRoutes(rg *gin.RouterGroup) {
	m.Register(rg)
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

// GET /sample-apptiers — 18 sample app tier rows for pipeline and form testing
func (m *MockPromotionController) GetSampleAppTiers(c *gin.Context) {
	tiers := []TierItem{}
	lbs := []string{"ALB", "ALB", "NLB", "ALB", "None", "None", "ALB", "ALB", "NLB", "ALB", "None", "ALB", "ALB", "NLB", "None", "ALB", "ALB", "NLB"}
	oses := []string{"Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Windows"}
	sizes := []string{"Medium", "Large", "Medium", "Large", "Medium", "Small", "Medium", "Medium", "Large", "Medium", "Small", "Medium", "Large", "Medium", "Small", "Medium", "Large", "Medium"}
	for i := 1; i <= 18; i++ {
		tiers = append(tiers, TierItem{
			RepositoryName:      fmt.Sprintf("app-tier-repo-%d", i),
			OS:                  oses[i-1],
			BitbucketProjectKey: "QNA",
			IsNotForPromotion:   i == 11,
			IsBifrost:           i != 12,
			IsServerlessPattern: i == 5 || i == 15,
			IsS3Only:            i == 6,
			IsPublicFacing:      false,
			InstanceSize:        sizes[i-1],
			NumOfInstances:      "2",
			VolumeType:          "gp3",
			VolumeSize:          "40",
			IsAutoScaling:       i%3 == 0,
			IsStaticIp:          i == 8 || i == 14,
			Loadbalancer:        lbs[i-1],
			IsFsx:               i == 7 || i == 16,
			IsEfs:               i == 2 || i == 10 || i == 16,
			IsElasticacheRedis:  i == 1 || i == 8 || i == 13 || i == 17,
		})
	}
	c.JSON(200, tiers)
}

// GET /sample-webtiers — 3 sample web tier rows for pipeline and form testing
func (m *MockPromotionController) GetSampleWebTiers(c *gin.Context) {
	tiers := []TierItem{
		{RepositoryName: "web-tier-repo-1", OS: "Linux", BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Medium", NumOfInstances: "2", VolumeType: "gp3", VolumeSize: "40", IsAutoScaling: true, IsPublicFacing: true, Loadbalancer: "ALB"},
		{RepositoryName: "web-tier-repo-2", OS: "Linux", BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Large", NumOfInstances: "3", VolumeType: "gp3", VolumeSize: "60", IsAutoScaling: true, IsPublicFacing: true, Loadbalancer: "NLB"},
		{RepositoryName: "web-tier-repo-3", OS: "Windows", BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Small", NumOfInstances: "1", VolumeType: "gp3", VolumeSize: "40", IsStaticIp: true, Loadbalancer: "None"},
	}
	c.JSON(200, tiers)
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
