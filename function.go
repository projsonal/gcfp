package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/PakArbi/goqrcode"
)

func init() {
	functions.HTTP("PostQrCode", generateQrCode)
}

func generateQrCode(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.vaidiq.cloud")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.vaidiq.cloud")
	fmt.Fprintf(w, goqrcode.generateQrCode())

}

//insertParkiranNPM

package gcf

import (
	"fmt"
	"github.com/PakArbi/pakarbibackend"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("insertParkiranEmail", PostParkiran)
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
	fmt.Fprintf(w, pakarbibackend.GCFInsertParkiranEmail("publickey", "MONGOCONNSTRINGENV", "PakArbiApp", "user", "parkiran", r))

}

//DELETE GCF
package gcf

import (
	"fmt"
	"github.com/PakArbi/pakarbibackend"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("deleteParkiranNPM", PostParkiran)
}

func PostParkiran(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token, Login")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, pakarbibackend.GCFDeleteParkiranNPM("publickey", "MONGOCONNSTRINGENV", "PakArbiApp", "user", "parkiran", r))

}

//LOGINUSERNPM OR LOGINUSEREMAIL
package gcf

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/PakArbi/pakarbibackend"
)

func init() {
    functions.HTTP("LoginUserEmail", backGCFkPost)
}

func backGCFkPost(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers for the preflight request
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io, https://pakarbi.vaidiq.cloud")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }
    // Set CORS headers for the main request.
    w.Header().Set("Access-Control-Allow-Origin", "https://pakarbi.github.io,https://pakarbi.vaidiq.cloud")
    fmt.Fprintf(w, pakarbibackend.LoginUserEmail("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "PakArbiApp", "user", r))

}