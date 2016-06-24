package main

//Session manager
import (
	"sync"
	"fmt"
	"io"
	"encoding/base64"
	"crypto/rand"
	"net/url"
	"net/http"
	"time"
)

var provides = make(map[string]Provider)

//This design takes its roots from the database/sql/driver , which defines the interface first, then registers specific
// structures when we want to use it. The following code is the internal implementation of a session register function.

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil, it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

//We know that we can save sessions in many ways including in memory, the file system or directly into the database. We
//need to define a Provider interface in order to represent the underlying structure of our session manager
type Provider interface {
	//SessionInit implements the initialization of a session, and returns new a session if it succeeds.
	SessionInit(sid string) (Session, error)

	//SessionRead returns a session represented by the corresponding sid. Creates a new session and returns it if it does
	//not already exist.
	SessionRead(sid string) (Session, error)

	//SessionDestroy given an sid, deletes the corresponding session.
	SessionDestroy(sid string) error

	//SessionGC deletes expired session variables according to maxLifeTime .
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{} //get session value
	Delete(key interface{}) error //delete session value
	SessionID() string //back current sessionID
}

type Manager struct {
	cookieName string //private cookiename
	lock sync.Mutex // protects session
	provider Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// Session ids are for identifying users of web applications, so they must be unique
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//We need to allocate or get an existing session in order to validate user operations. The SessionStart function is for
//checking if any there are any sessions related to the current user, creating a new session non are found.
//The SessionStart function returns a variable that implements a session interface
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//Destroy sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// Session manager delete a session.
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}