package lilypad

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func newEnvCmd() *cobra.Command {

	envCmd := &cobra.Command{
		Use:     "env",
		Short:   "Update env vars from dopler.",
		Long:    "Update env vars from dopler.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			// options, err := optionsfactory.ProcessSolverOptions(options)
			// if err != nil {
			// 	return err
			// }
			// return runSolver(cmd, options)
			// dbHost := os.Getenv("POSTGRES_HOST")
			if err := godotenv.Load(); err != nil {
				fmt.Println("Error loading .env file: %v", err)
			}

			// fmt.Println("os.Environ()", os.Environ())
			godotenv.Load()
			var myEnv map[string]string
			myEnv, err := godotenv.Read()
			// myEnv["POSTGRES_HOST"] = "localhost"
			//myEnv.append
			// fmt.Println("myEnv", myEnv)
			// env, err := godotenv.Unmarshal("POSTGRES_USER=mkyong")
			if err != nil {
				fmt.Println("Error writing to .env file:", err)

			}

			godotenv.Write(myEnv, ".env")
			conf := ""
			if len(args) == 1 {
				conf = args[0]
			}
			if conf == "docker" {
				if myEnv["METRICS_HOST"] == "" {
					myEnv["METRICS_HOST"] = "metrics"
				}
				if myEnv["POSTGRES_HOST"] == "" {
					myEnv["POSTGRES_HOST"] = "postgres"
				}
				if myEnv["DIR"] == "" {
					myEnv["DIR"] = "/app"
				}
				if myEnv["GETH"] == "" {
					myEnv["GETH"] = "geth2"
				}
				// if myEnv["WEB3_RPC_URL"] == "" {
				myEnv["WEB3_RPC_URL"] = "ws://geth2:8546"
				myEnv["NETWORK"] = "local_docker_network"

				// }
				// ="ws://localhost:8546"

			} else {
				if myEnv["METRICS_HOST"] == "" {
					myEnv["METRICS_HOST"] = "localhost"
				}
				if myEnv["POSTGRES_HOST"] == "" {
					myEnv["POSTGRES_HOST"] = "localhost"
				}
			}
			if myEnv["POSTGRES_USER"] == "" {
				myEnv["POSTGRES_USER"] = "lp"
			}
			if myEnv["POSTGRES_PASSWORD"] == "" {
				myEnv["POSTGRES_PASSWORD"] = "password" //todo: instead generate random password
			}
			if myEnv["POSTGRES_DB"] == "" {
				myEnv["POSTGRES_DB"] = "mydb"
			}
			godotenv.Write(myEnv, ".env")
			// lilymetrics.LogMetric("env", "test")
			// fmt.Println("updating env for host", os.Getenv("POSTGRES_HOST"))
			return nil
		},
	}

	// optionsfactory.AddSolverCliFlags(solverCmd, &options);

	return envCmd
}
