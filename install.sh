if ! go build -o pm cmd/pm/main.go; then
    echo "Could not build pm, is go installed properly?"
else
    sudo mv pm /bin
fi