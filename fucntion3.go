package gcf

import (
	"fmt"
	"github.com/nugisOrange/petback"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("post", InsertdataCoordinateGCF)
}

func InsertdataCoordinateGCF(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://nugisorange.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token, login")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://nugisorange.github.io")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, petback.GCFPostCoordinate("MONGOSTRING","PUBLICKEY", "petasal", "post", r))

}