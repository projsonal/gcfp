package gcf

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/pasetobackend"
)

func init() {
    functions.HTTP("LoginUserNPM", backGCFkPost)
}

func backGCFkPost(w http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf(w, pasetobackend.LoginUserNPM("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbi", "user", r))



    // emailResult, errEmail := pakarbibackend.LoginUserEmail("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbi", "user", r)
    // if errEmail != nil {
    //     // Handle the error, log, or respond accordingly
    //     http.Error(w, "Internal server error", http.StatusInternalServerError)
    //     return
    // }
    // fmt.Fprintf(w, emailResult)
    
    // npmResult, errNPM := pasetobackend.LoginUserNPM("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbi", "user", r)
    // if errNPM != nil {
    //     // Handle the error, log, or respond accordingly
    //     http.Error(w, "Internal server error", http.StatusInternalServerError)
    //     return
    // }
    // fmt.Fprintf(w, npmResult)
    
      
}





//codq qr

package gcf

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/backparkir"
)

func init() {
	functions.HTTP("postCodeQR", PostParkiran)
}

func PostParkiran(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token, Login")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, backparkir.GCFPostParkiran("MONGOSTRINGENV", "PakArbi", "CodeQR", "r" ))

}


package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/PakArbi/backparkir"
)

func init() {
	functions.HTTP("post", CodeQR)
}

func CodeQR(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io/tester, pakarbi.vaidiq.me/tester")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io/tester, pakarbi.vaidiq.me/tester")
	fmt.Fprintf(w, backparkir.GCFPostParkiran("MONGOCONNSTRINGENV", "PakArbi", "CodeQR", r))

}


// UPDATE DATA PARKIRAN
package gcf

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/pakarbibackend"
)

func init() {
    functions.HTTP("updateDataParkiran", backGCFkPost)
}

func backGCFkPost(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers for the preflight request
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
        w.Header().Set("Access-Control-Allow-Methods", "PUT")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }
    // Set CORS headers for the main request.
    w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io,https://pakarbi.vaidiq.cloud")
    fmt.Fprintf(w, pakarbibackend.GCFUUpdateParkiranNPM("publickey", "MONGOCONNSTRINGENV", "PakArbiApp", "user", "parkiran", r))

}