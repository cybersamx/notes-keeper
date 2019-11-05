package main

import (
	"github.com/cybersamx/to-do-go/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

const (
	port = "8000"
)

type App struct {
	errLog	*log.Logger
	infoLog *log.Logger
	noteModel *models.NoteModel
}

func main() {
	// Initialize the dependencies

	// Logs
	errLog := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Database
	dsn := "host=localhost port=5432 user=app-user dbname=app password=not-secure-pwd sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		errLog.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
	noteModel := models.NewNoteModel(db)

	// Encapsulate all the dependencies in App
	app := App{
		errLog: errLog,
		infoLog: infoLog,
		noteModel: noteModel,
	}

	// HTTP
	mux := http.NewServeMux()

	mux.Handle("/", notesHandler(&app))
	mux.Handle("/createNote", createNoteHandler(&app))

	app.infoLog.Print("web server running at port ", port)
	server := http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ErrorLog:          app.errLog,
	}
	err = server.ListenAndServe()
	errLog.Fatal(err)
}
