package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/tsi4456/gator/internal/config"
	"github.com/tsi4456/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	s := &state{cfg: &cfg, db: dbQueries}
	cList := commands{command_list: make(map[string]func(*state, command) error)}

	cList.register("login", handlerLogin)
	cList.register("register", handlerRegister)
	cList.register("reset", handlerReset)
	cList.register("users", handlerUsers)
	cList.register("agg", handlerAgg)
	cList.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cList.register("feeds", handlerFeeds)
	cList.register("follow", middlewareLoggedIn(handlerFollow))
	cList.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cList.register("following", middlewareLoggedIn(handlerFollowing))
	cList.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("No command given")
	}
	com := os.Args[1]
	args := os.Args[2:]
	if err = cList.run(s, command{com, args}); err != nil {
		log.Fatal(err)
	}
}
