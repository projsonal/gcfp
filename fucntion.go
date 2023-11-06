package pasetoProPost

import (
    "fmt"
    "net/http" 

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/pasetobackend"
    // "github.com/PakArbi/backendUser"
)

func init() {
    functions.HTTP("GetDataUserFromGCF", GetDataUserFromGCF)
}

func GetDataUserFromGCF(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers for the preflight request
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }
    // Set CORS headers for the main request.
    w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
    fmt.Fprintf(w, pasetobackend.GCFPasetoTokenStr("privateKey", "MONGOVER", "GetDataUserFromGCF", "GetDataUserFromGCF", r))

}