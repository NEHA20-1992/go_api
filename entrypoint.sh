wait-for = "${DATABASE_HOST}":"${DATABASE_PORT}" --"/@"
# watch for .go file changes
CompileDemon --build="go build -o main main.go"--command=./main