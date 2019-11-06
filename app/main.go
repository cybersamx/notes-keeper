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
	errLog    *log.Logger
	infoLog   *log.Logger
	noteModel *models.NoteModel
}

func main() {
	// Logs
	errLog := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Database
	dsn := "host=db port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		errLog.Fatal(err)
		panic(err)
	}
	defer db.Close()
	noteModel := models.NewNoteModel(db)

	// Encapsulate all the dependencies in App
	app := App{
		errLog:    errLog,
		infoLog:   infoLog,
		noteModel: noteModel,
	}

	// HTTP dynamic content
	mux := http.NewServeMux()
	mux.Handle("/", notesHandler(&app))
	mux.Handle("/editNote", editNoteHandler(&app))
	mux.Handle("/removeNote", removeNoteHandler(&app))

	// HTTP static content
	fileSrv := http.FileServer(http.Dir("../html/images/"))
	mux.Handle("/images/", http.StripPrefix("/images", fileSrv))

	// HTTP Server
	app.infoLog.Print("web server running at port ", port)
	server := http.Server{
		Addr:     ":" + port,
		Handler:  mux,
		ErrorLog: app.errLog,
	}
	err = server.ListenAndServe()
	panic(err)
}
