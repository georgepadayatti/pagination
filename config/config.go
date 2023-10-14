package config

import (
	"github.com/spf13/viper"
)

// Collection
type Collection struct {
	Name string `mapstructure:"name"`
}

// Database
type Database struct {
	User       string     `mapstructure:"user"`
	Password   string     `mapstructure:"password"`
	Name       string     `mapstructure:"name"`
	Collection Collection `mapstructure:"collection"`
}

// ConfigInterface
type ConfigInterface interface {
	GetDatabase() Database
}

// Config holds the structure of the configuration
type Config struct {
	Database Database `mapstructure:"database"`
}

// GetDatabaseUser returns the database user
func (c *Config) GetDatabaseUser() string {
	return c.Database.User
}

// GetDatabasePassword returns the database password
func (c *Config) GetDatabasePassword() string {
	return c.Database.Password
}

// GetDatabaseName returns the database name
func (c *Config) GetDatabaseName() string {
	return c.Database.Name
}

// GetCollectionName returns the collection name
func (c *Config) GetCollectionName() string {
	return c.Database.Collection.Name
}

var AppConfig Config

// LoadConfig
func LoadConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
