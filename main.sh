if [[ $# -eq 0 ]]; then
    echo "Start app with a default mode"
    go run .
elif [[ $1 == "compile" ]]; then
    echo "Build app"
    go build
    echo "Run Server with compiled executable file"
    ./dashpad
else
    echo "Invalid command line argument"
    echo "Start app with a default mode"
    go run .
fi
