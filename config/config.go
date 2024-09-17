package config

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/viper"
)

// You can configure you app config here
type Config struct {
	AppName    string
	ApiURL     string
	AWSRegion  string
	DBPassword string
}

var AppConfig Config

func InitConfig() {
	// Set default values and load from config.yaml or environment variables
	viper.SetDefault("app.name", "CLI app!")
	viper.SetDefault("aws.region", "eu-central-1")

	// Load from config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // Optionally add config paths
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %s", err)
	}

	// Bind environment variables
	viper.AutomaticEnv()

	// Populate configuration values
	AppConfig = Config{
		AppName:   viper.GetString("app.name"),
		AWSRegion: viper.GetString("aws.region"),
		ApiURL:    viper.GetString("api.url"),
	}

	// Fetch secrets from AWS Secrets Manager
	// You can replace this with your required secrets and add more

	// dbPassword, err := getSecret("db-password")
	// if err != nil {
	// 	log.Fatalf("Failed to get secret from AWS Secrets Manager: %s", err)
	// }
	// AppConfig.DBPassword = dbPassword
}

// Helper function to get secret from AWS Secrets Manager
func getSecret(secretName string) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("aws.region")),
	}))
	svc := secretsmanager.New(sess)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret: %w", err)
	}

	return *result.SecretString, nil
}
