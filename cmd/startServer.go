package cmd

import (
	"encoding/json"
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
	startServer.Flags().StringVarP(&port, "port", "p", ":5000", "Port")
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

	headerMapping := r.Header
	var bodyMapping map[string]string
	bodyStr := ""

	if err := json.NewDecoder(r.Body).Decode(&bodyMapping); err != nil && err.Error() != "EOF" {
		fmt.Println("Error in decoding body:", err.Error())
	}

	body, err := json.Marshal(bodyMapping)
	if err != nil {
		fmt.Println("Error in encoding body:", err.Error())
	}
	bodyStr = string(body)

	header, err := json.Marshal(headerMapping)
	if err != nil {
		fmt.Println("Error in parsing headers:", err.Error())
	}
	if path != "" {
		output := RunCommand(path, bodyStr, string(header))
		fmt.Fprintln(w, output)
	} else {
		fmt.Fprint(w, "Hello there")
	}
}
