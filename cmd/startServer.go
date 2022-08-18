package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var port string
var startServer = &cobra.Command{
	Use:   "startServer",
	Short: "Start Server",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func init() {
	addConfigCmd.Flags().StringVarP(&port, "port", "p", ":5000", "Port")
	RootCmd.AddCommand(startServer)
}

func StartServer() {
	mux := defaultMux()
	http.ListenAndServe(port, mux)
	fmt.Printf("Starting the server: http://localhost:%s", port)
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
