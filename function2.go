package gcf

import (
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/pakarbibackend"
)

func init() {
    functions.HTTP("loginUserEmail", loginWithEmail)
    functions.HTTP("loginUserNPM", loginWithNPM)
}

func loginWithEmail(w http.ResponseWriter, r *http.Request) {
    handleCORS(w, r)
    fmt.Fprintf(w, pakarbibackend.LoginUserEmail("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbi", "user", r))
}

func loginWithNPM(w http.ResponseWriter, r *http.Request) {
    handleCORS(w, r)
    fmt.Fprintf(w, pakarbibackend.LoginUserNPM("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbi", "user", r))
}

func handleCORS(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io")
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }
}
