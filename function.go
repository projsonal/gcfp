// package gcp

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
// 	"github.com/PakArbi/backprofile"
// )

// func init() {
// 	functions.HTTP("getData", getData)
// }

// func getData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET")
// 	fmt.Fprintf(w, backprofile.GCHandlerFunc("MongoString", "PakArbi", "Profiles"))
// }

package gcf

import (
	"fmt"
	"github.com/PakArbi/backprofile/parkiran"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("insertProfile", InsertdataProfileGCF)
}

func InsertdataProfileGCF(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token, login")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, parkiran.GCFPostParkiran("MONGOSTRING", "PakArbi", "Parkiran", r ))

}
