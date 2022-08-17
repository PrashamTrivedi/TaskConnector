package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var configFile string
var startServer = &cobra.Command{
	Use:   "startServer",
	Short: "Start Server",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	addConfigCmd.Flags().StringVarP(&configFile, "configFile", "c", "", "Config File")
	RootCmd.AddCommand(getConfig)
}

func StartServer() {
	mux := defaultMux()
	http.ListenAndServe(":5000", mux)
	fmt.Println("Starting the server: http://localhost:5000")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("key")
	if path != "" {
		output := RunCommand(path)
		fmt.Fprintln(w, output)
	} else {
		fmt.Fprint(w, "Hello there")
	}
}
