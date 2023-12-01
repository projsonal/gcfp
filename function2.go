package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/PakArbi/backprofile/parkiran"
)

func init() {
	functions.HTTP("get", petaPediaGet)
}

func petaPediaGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://gis5syahid.github.io")
	fmt.Fprintf(w, parkiran.GCFHandler("MONGOSTRING", "PakArbi", "parkiran"))
}