# Notes Keeper

A simple web app written in Go for creating to-do notes.

## Setup

1. (Optional) Set up Docker on Mac - will update Docker setup on other non-Mac setup.

   ```bash
   $ # Install Docker.app
   $ brew cask install docker
   $ # Install docker-compose
   $ brew install docker-compose
   ```
   
   Alternatively you can also install `docker-machine`.

### Development

1. To set up and run Postgres on your local machine.

   ```bash
   $ # Run Postgres container locally
   $ ./scripts/run-postgres-local.sh
   ```

1. Run the following commands to build the application.

   ```bash
   $ cd app # Let app be the working directory
   $ go build -o ../bin/notes-keeper
   ```
   
1. Run the following commands to run the application. Afterwards, launch your web browser and navigate to http://localhost:8000.

   ```bash
   $ cd app  # Let app be the working directory
   $ go run github.com/cybersamx/to-do-go/app
   ```

### Deployment

1. Run Docker Compose.

   ```bash
   $ docker-compose up --build
   ```