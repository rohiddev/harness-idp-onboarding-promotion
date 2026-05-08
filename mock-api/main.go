package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Option is the standard dropdown response shape used by SelectFieldFromApi.
type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// ─── Onboarding Promotion Step 1 ─────────────────────────────────────────────
// Routes: /hoover-service/mock-api/jarvis/promotion/*

var divisions = []Option{
	{Value: "Colleaguetech",     Label: "Colleaguetech"},
	{Value: "Commercial",        Label: "Commercial"},
	{Value: "CommercialConnect", Label: "CommercialConnect"},
	{Value: "Consumer",          Label: "Consumer"},
	{Value: "ConsumerConnect",   Label: "ConsumerConnect"},
	{Value: "Digital",           Label: "Digital"},
	{Value: "Edo",               Label: "Edo"},
	{Value: "Entsvc",            Label: "Entsvc"},
	{Value: "Exvendor",          Label: "Exvendor"},
}

var regions = []Option{
	{Value: "us-east-1",    Label: "us-east-1 (N. Virginia)"},
	{Value: "us-east-2",    Label: "us-east-2 (Ohio)"},
	{Value: "us-west-2",    Label: "us-west-2 (Oregon)"},
	{Value: "eu-west-1",    Label: "eu-west-1 (Ireland)"},
	{Value: "eu-central-1", Label: "eu-central-1 (Frankfurt)"},
	{Value: "eastus",       Label: "eastus (Azure East US)"},
	{Value: "eastus2",      Label: "eastus2 (Azure East US 2)"},
	{Value: "westeurope",   Label: "westeurope (Azure West Europe)"},
}

var environments = []Option{
	{Value: "dev", Label: "dev"},
	{Value: "qa",  Label: "qa"},
}

var promoteToEnvironments = []Option{
	{Value: "qa",   Label: "qa"},
	{Value: "prod", Label: "prod"},
}

var bifrostCategories = []Option{
	{Value: "Bifrost - Application Software Manual Install",           Label: "Bifrost - Application Software Manual Install"},
	{Value: "Bifrost - Application Software Partially Automated",      Label: "Bifrost - Application Software Partially Automated"},
	{Value: "Bifrost - Application Software Install Full Utilization", Label: "Bifrost - Application Software Install Full Utilization"},
	{Value: "Not Required", Label: "Not Required"},
}

var vpcOptions = []Option{
	{Value: "default",      Label: "default"},
	{Value: "custom-vpc-1", Label: "custom-vpc-1"},
	{Value: "custom-vpc-2", Label: "custom-vpc-2"},
}

// TierItem is a single row in the app/web tier tables.
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

func getPromotionSampleAppTiers(c *gin.Context) {
	lbs   := []string{"ALB", "ALB", "NLB", "ALB", "None", "None", "ALB", "ALB", "NLB", "ALB", "None", "ALB", "ALB", "NLB", "None", "ALB", "ALB", "NLB"}
	oses  := []string{"Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Linux", "Windows", "Linux", "Linux", "Linux", "Windows"}
	sizes := []string{"Medium", "Large", "Medium", "Large", "Medium", "Small", "Medium", "Medium", "Large", "Medium", "Small", "Medium", "Large", "Medium", "Small", "Medium", "Large", "Medium"}
	tiers := make([]TierItem, 0, 18)
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

func getPromotionSampleWebTiers(c *gin.Context) {
	tiers := []TierItem{
		{RepositoryName: "web-tier-repo-1", OS: "Linux",   BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Medium", NumOfInstances: "2", VolumeType: "gp3", VolumeSize: "40", IsAutoScaling: true,  IsPublicFacing: true,  Loadbalancer: "ALB"},
		{RepositoryName: "web-tier-repo-2", OS: "Linux",   BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Large",  NumOfInstances: "3", VolumeType: "gp3", VolumeSize: "60", IsAutoScaling: true,  IsPublicFacing: true,  Loadbalancer: "NLB"},
		{RepositoryName: "web-tier-repo-3", OS: "Windows", BitbucketProjectKey: "QNA", IsBifrost: true, InstanceSize: "Small",  NumOfInstances: "1", VolumeType: "gp3", VolumeSize: "40", IsStaticIp:   true,  IsPublicFacing: false, Loadbalancer: "None"},
	}
	c.JSON(200, tiers)
}

func registerPromotionRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/jarvis/promotion")
	g.GET("/health",                  func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok", "service": "jarvis-promotion"}) })
	g.GET("/:cloudType/divisions",    func(c *gin.Context) { c.JSON(200, divisions) })
	g.GET("/regions",                 func(c *gin.Context) { c.JSON(200, regions) })
	g.GET("/environments",            func(c *gin.Context) { c.JSON(200, environments) })
	g.GET("/promote-to-environments", func(c *gin.Context) { c.JSON(200, promoteToEnvironments) })
	g.GET("/bifrost-categories",      func(c *gin.Context) { c.JSON(200, bifrostCategories) })
	g.GET("/vpc-options",             func(c *gin.Context) { c.JSON(200, vpcOptions) })
	g.GET("/sample-apptiers",         getPromotionSampleAppTiers)
	g.GET("/sample-webtiers",         getPromotionSampleWebTiers)
}

// ─── Onboarding Day 2 Step 1 ─────────────────────────────────────────────────
// Routes: /hoover-service/mock-api/jarvis/day2/step1/*

var day2Divisions = []Option{
	{Value: "Colleaguetech",     Label: "Colleaguetech"},
	{Value: "Commercial",        Label: "Commercial"},
	{Value: "Commercialconnect", Label: "Commercialconnect"},
	{Value: "Consumer",          Label: "Consumer"},
	{Value: "Consumerconnect",   Label: "Consumerconnect"},
	{Value: "Digital",           Label: "Digital"},
	{Value: "Edo",               Label: "Edo"},
	{Value: "Embackup",          Label: "Embackup"},
	{Value: "Entsvc",            Label: "Entsvc"},
	{Value: "Exvendor",          Label: "Exvendor"},
}

var day2Regions = []Option{
	{Value: "us-east-1",    Label: "us-east-1 (N. Virginia)"},
	{Value: "us-east-2",    Label: "us-east-2 (Ohio)"},
	{Value: "us-west-2",    Label: "us-west-2 (Oregon)"},
	{Value: "eu-west-1",    Label: "eu-west-1 (Ireland)"},
	{Value: "eu-central-1", Label: "eu-central-1 (Frankfurt)"},
	{Value: "eastus",       Label: "eastus (Azure East US)"},
	{Value: "eastus2",      Label: "eastus2 (Azure East US 2)"},
	{Value: "westeurope",   Label: "westeurope (Azure West Europe)"},
}

var day2Environments = []Option{
	{Value: "dev",  Label: "dev"},
	{Value: "qa",   Label: "qa"},
	{Value: "prod", Label: "prod"},
}

func registerDay2Step1Routes(rg *gin.RouterGroup) {
	g := rg.Group("/jarvis/day2/step1")
	g.GET("/health",       func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok", "service": "jarvis-day2-step1"}) })
	g.GET("/divisions",    func(c *gin.Context) { c.JSON(200, day2Divisions) })
	g.GET("/regions",      func(c *gin.Context) { c.JSON(200, day2Regions) })
	g.GET("/environments", func(c *gin.Context) { c.JSON(200, day2Environments) })
}

// ─── Innovation Lab Onboarding ────────────────────────────────────────────────
// Routes: /hoover-service/mock-api/jarvis/innovation-lab/*

var innovationLabPocTypes = []Option{
	{Value: "aws-learning",    Label: "aws-learning"},
	{Value: "aws-service-poc", Label: "aws-service-poc"},
	{Value: "cloud-poc",       Label: "cloud-poc"},
	{Value: "db-poc",          Label: "db-poc"},
	{Value: "genai-poc",       Label: "genai-poc"},
	{Value: "iam-poc",         Label: "iam-poc"},
	{Value: "network-poc",     Label: "network-poc"},
	{Value: "security-poc",    Label: "security-poc"},
	{Value: "vendor-poc",      Label: "vendor-poc"},
	{Value: "other-poc",       Label: "other-poc"},
}

func registerInnovationLabRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/jarvis/innovation-lab")
	g.GET("/health",    func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok", "service": "jarvis-innovation-lab"}) })
	g.GET("/poc-types", func(c *gin.Context) { c.JSON(200, innovationLabPocTypes) })
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors())

	// All routes under /hoover-service/mock-api to match Harness IDP proxy pathRewrite:
	//   proxy/platform-api/hoover-service/mock-api/* -> /hoover-service/mock-api/*
	base := r.Group("/hoover-service/mock-api")

	registerPromotionRoutes(base)
	registerDay2Step1Routes(base)
	registerInnovationLabRoutes(base)

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	port := ":8084"
	log.Printf("Jarvis Mock API running on %s", port)
	log.Printf("Promotion:      /hoover-service/mock-api/jarvis/promotion/*")
	log.Printf("Day 2 Step 1:   /hoover-service/mock-api/jarvis/day2/step1/*")
	log.Printf("Innovation Lab: /hoover-service/mock-api/jarvis/innovation-lab/*")
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
