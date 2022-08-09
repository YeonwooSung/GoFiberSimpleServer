if [[ $# -eq 0 ]]; then
    echo "Start app with a default (REPL) mode"
    npm run dev & go run .
elif [[ $1 == "repl" ]]; then
    echo "Run Server with REPL mode"
    npm run dev & go run .
elif [[ $1 == "compile" ]]; then
    echo "Build app"
    go build
    echo "Run Server with compiled executable file"
    npm run dev & ./dashpad
else
    echo "Invalid command line argument"
    echo "Start app with a default (REPL) mode"
    npm run dev & go run .
fi
