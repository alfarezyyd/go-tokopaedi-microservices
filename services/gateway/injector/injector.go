package injector

import (
	"fmt"
	"go-tokopaedi-microservices/configs"
	"go-tokopaedi-microservices/pkg/exception"
	"go-tokopaedi-microservices/pkg/logger"
	"go-tokopaedi-microservices/pkg/middleware"
	"go-tokopaedi-microservices/services/gateway/routes"
	"go-tokopaedi-microservices/services/gateway/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

func NewDatabaseConnection(databaseCredentials *configs.DatabaseCredentials) *gorm.DB {
	databaseConnectionInstance := configs.NewDatabaseConnection(databaseCredentials)
	return databaseConnectionInstance.GetDatabaseConnection()
}

func NewRedisInstance(redisConfig configs.RedisConfig) *configs.RedisInstance {
	redisInstance, err := configs.InitRedisInstance(redisConfig)
	if err != nil {
		panic(err)
	}
	return redisInstance
}

func InitRedisConfig(viperConfig *viper.Viper) configs.RedisConfig {
	redisHostName, redisPassword, parsedRedisDatabaseIndex := configs.LoadRedisConfigFromEnvironment(viperConfig)
	return configs.NewRedisConfig(redisHostName, redisPassword, int(parsedRedisDatabaseIndex))
}

func NewValidator(gormDatabase *gorm.DB) (*validator.Validate, universalTranslator.Translator) {
	return configs.InitializeValidator(gormDatabase)
}

// NewViperConfig --- Provider untuk Viper config ---
func NewViperConfig() *viper.Viper {
	viperConfig := viper.New()
	viperConfig.SetConfigFile("./services/gateway/.env")
	viperConfig.SetConfigType("env")
	viperConfig.AutomaticEnv()
	if err := viperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	return viperConfig
}

func NewGrpcClientRegistry() *configs.GrpcClientRegistry {
	grpcClientRegistry := configs.NewGrpcClientRegistry()
	mapOfConnection := make(map[string]configs.GrpcConnectionClient)
	mapOfConnection["user"] = configs.GrpcConnectionClient{
		EndpointTarget: "localhost:10001",
		DialOption:     grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	for serviceName, grpcConnectionClient := range mapOfConnection {
		err := grpcClientRegistry.Connect(serviceName, grpcConnectionClient.EndpointTarget, grpcConnectionClient.DialOption)
		if err != nil {
			logger.Debug(fmt.Sprintf("Failed to connect to GRPC service %s", serviceName))
		}
	}
	return grpcClientRegistry
}

// NewDatabaseCredentials --- Provider untuk Database Credentials ---
func NewDatabaseCredentials(viperConfig *viper.Viper) *configs.DatabaseCredentials {
	return &configs.DatabaseCredentials{
		DatabaseHost:     viperConfig.GetString("DATABASE_HOST"),
		DatabasePort:     viperConfig.GetString("DATABASE_PORT"),
		DatabaseName:     viperConfig.GetString("DATABASE_NAME"),
		DatabasePassword: viperConfig.GetString("DATABASE_PASSWORD"),
		DatabaseUsername: viperConfig.GetString("DATABASE_USERNAME"),
	}
}

// NewGinEngine --- Provider untuk Gin Engine ---
func NewGinEngine() (*gin.Engine, *gin.RouterGroup) {
	gin.SetMode(gin.DebugMode)
	ginEngine := gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(exception.Interceptor())
	ginEngineRoot := ginEngine.Group("/")
	ginEngineRoot.Use(middleware.RequestMetaMiddleware())

	return ginEngine, ginEngineRoot
}

var CoreModule = fx.Module("coreModule", fx.Provide(
	NewViperConfig,
	NewDatabaseCredentials,
	NewDatabaseConnection,
	NewValidator,
	NewGinEngine,
	InitRedisConfig,
	NewRedisInstance,
	NewGrpcClientRegistry,
))

var ApplicationRoutesModule = fx.Module("applicationRoutes",
	fx.Provide(
		routes.NewPublicRoutes,
		routes.NewAuthenticationRoutes,
		routes.NewProtectedRoutes,
		func(
			ginEngine *gin.Engine,
			publicRoutes *routes.PublicRoutes,
			authenticationRoutes *routes.AuthenticationRoutes,
			protectedRoutes *routes.ProtectedRoutes,
		) *routes.ApplicationRoutes {
			return routes.NewApplicationRoutes(ginEngine, publicRoutes, authenticationRoutes, protectedRoutes)
		},
	),
	fx.Invoke(func(applicationRoutes *routes.ApplicationRoutes) {
		applicationRoutes.Setup()
	}),
)

var UserModule = fx.Module("userFeature",
	fx.Provide(fx.Annotate(user.NewService, fx.As(new(user.Service)))),
	fx.Provide(fx.Annotate(user.NewHandler, fx.As(new(user.Handler)))),
)
