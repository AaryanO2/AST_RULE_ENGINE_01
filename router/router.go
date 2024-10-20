package router

import (
	"net/http"
	"y/handlers"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./")) // serves files from the current directory
	router.Handle("/", fs)
	router.HandleFunc("/api/v1/create_rule", handlers.CreateRuleHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/combine_rules", handlers.CombineRulesHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/evaluate", handlers.EvaluateRuleHandler).Methods("POST", "OPTIONS")

	return router
}
