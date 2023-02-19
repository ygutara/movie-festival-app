# movie-festival-app
This application allows admin to create, read, and update movie,
and user can watch, vote, and unvote

# Features
+ AUTH
+ CINEMA

+ ADMIN
- Create new movie records
- View a list of movie records
- Update movie record information

+ USER
- Watch movie records
- Vote movie records
- Unvote movie information

# Technologies
- Backend: Golang
- Database: PostgreSQL

# Installation
- Clone this repository to your local machine
- Install Go and PostgreSQL
- Open terminal and navigate to the application directory
- Run the command go mod download to install dependencies
- Run 'cd /app/config' and edit config.go to change database's configuration 
- Run 'cd /app/migration' and run 'go run migration.go'
- username = admin, password = admin 
- Back to the application directory 
- Run the command 'go run main.go'
- Api Application can be accessed at 'http://localhost:8080'

# Installation using Docker Compose
- Open terminal and navigate to the application directory
- run 'docker compose create'
- run 'docker compose start'

