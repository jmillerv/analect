# Build the binary
build:
    go build -o analect && chmod +x ./analect

# Build mobile app
build-mobile:
    fyne-cross android -app-id com.jmillerv.analect

# Run the binary
run:
    ./analect