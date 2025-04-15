package commands

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestConnection(cmd *cobra.Command, args []string) {
	// Get config path
	cdir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("❌ Failed to get user config dir:", err)
		os.Exit(1)
	}
	configFile := filepath.Join(cdir, "schemasnap", "config.yaml")

	// Setup Viper to read it
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("❌ Failed to read config:", err)
		os.Exit(1)
	}

	// Get values
	host := viper.GetString("host")
	port := viper.GetString("port")
	user := viper.GetString("user")
	password := viper.GetString("password")
	dbname := viper.GetString("database")

	// Build DSN
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Try to connect
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("❌ Failed to open DB:", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("❌ Database unreachable:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Successfully connected to the database!")
}
