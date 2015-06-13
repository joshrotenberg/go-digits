// loginapp is an example web app with Login with Digits (phone number).
package main

import (
	"fmt"
	dgtsAPI "github.com/dghubble/go-digits/digits"
	dgtsLogin "github.com/dghubble/go-digits/login"
	"github.com/dghubble/sessions"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	digitsConsumerKey = "YOUR_DIGITS_CONSUMER_KEY"
	sessionName       = "loginapp-session"
	sessionSecret     = "example cookie signing secret"
	sessionUserKey    = "digitsID"
)

// sessionStore encodes and decodes session data stored in signed cookies
var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// digitsService provides a login handler for validation and account retrieval
// 1. Create a Digits Login Service struct with your Digits Consumer Key
var digitsService = dgtsLogin.NewLoginService(digitsConsumerKey)

// New returns a new ServeMux with app routes.
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/profile", requireLogin(http.HandlerFunc(profileHandler)))
	mux.HandleFunc("/logout", logoutHandler)
	// 2. Register a Digits LoginHandler to receive POST from JS after login
	mux.Handle("/digits_login", digitsService.LoginHandlerFunc(successHandler, dgtsLogin.ErrorHandler))
	return mux
}

// successHandler issues a cookie session upon successful Digits login
func successHandler(w http.ResponseWriter, req *http.Request, account *dgtsAPI.Account) {
	// 3. Implement successHandler to issue some form of session or write to db
	session := sessionStore.New(sessionName)
	session.Values[sessionUserKey] = account.ID
	session.Values["phoneNumber"] = account.PhoneNumber
	session.Save(w)
	http.Redirect(w, req, "/profile", http.StatusFound)
}

// homeHandler shows a login page or a user profile page if authenticated.
func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	if isAuthenticated(req) {
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	page, _ := ioutil.ReadFile("home.html")
	fmt.Fprintf(w, string(page))
}

// profileHandler shows protected user content.
func profileHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, `<p>You are logged in!</p><form action="/logout" method="post"><input type="submit" value="Logout"></form>`)
}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, sessionStore.New(sessionName))
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// requireLogin redirects unauthenticated users to the login route.
func requireLogin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if !isAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

// isAuthenticated returns true if the user has a signed session cookie.
func isAuthenticated(req *http.Request) bool {
	if _, err := sessionStore.Get(req, sessionName); err == nil {
		return true
	}
	return false
}

// main creates and starts a Server listening.
func main() {
	const address = "localhost:8080"
	log.Printf("Starting Server listening on %s\n", address)
	err := http.ListenAndServe(address, New())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}