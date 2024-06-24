package main

import (
	"log"
	"net"
	s "root/server"
	"root/server/trie"
	"sync"
	"time"
)

type Database struct {
	users      map[User]trie.Trie[string]
	authorized map[string]struct{}
	sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{}
}

type User struct {
	name     string
	password string
}

func main() {
	r, db := s.NewRouter(), NewDatabase()

	r.HandlerFunc("QUIT", Quit)
	r.HandlerFunc("USER", db.User)
	r.HandlerFunc("INFO", Info)

	if err := r.Listen("localhost:5000"); err != nil {
		log.Println(err)
	}
}

func (db *Database) User(c net.Conn, args string) {
	if len(args) != 2 {
		return
	}

	user := User{
		name: args,
	}

	db.RLock()
	if _, exist := db.users[user]; !exist {
		return
	}
	db.RUnlock()

	db.Lock()
	db.authorized[c.RemoteAddr().String()] = struct{}{}
	db.Unlock()
}

func Info(c net.Conn, args string) {

}

func Quit(c net.Conn, args string) {
	defer c.Close()
	c.Write([]byte("log out: " + time.Now().String()))
}
