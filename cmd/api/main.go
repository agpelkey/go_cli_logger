package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	logger "github.com/agpelkey/cli_logger"
	"github.com/agpelkey/cli_logger/internal/models"
	"github.com/agpelkey/cli_logger/internal/repository"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	DB repository.DatabaseRepo
}

func main() {

	// Set application config
	var app application

	// Connect to sqlite database
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for the CR to accurately log events\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Made by Alex Pelkey and Rishabh Bajpai")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// Parsing command line flags
	add := flag.Bool("add", false, "Add entry to the logbook")
	list := flag.Bool("list", false, "lists logbook entries")

	flag.Parse()

	// Define an logbook list
	l := &app.DB

	// Decide what to do based on the number of arguments provided
	switch {
	case *list:
		// list current log entries
		_, err := app.DB.FetchLogs()
		if err != nil {
			log.Fatal(err)
		}

	case *add:
		// flag to add logbook entry
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		app.DB.Add(t)

	default:
		// Invalid flag provided
		fmt.Println(os.Stderr, "Invalid Option")
		os.Exit(1)

	}

}

// getTask function decides where to get the description
// for a new task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
