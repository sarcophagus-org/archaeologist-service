package ethereum

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"
)

const (
	MB = 1 << 20
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")

	log.Println("File received with header:", header)

	if header.Size > (3 * MB) {
		log.Fatal("The file sent is larger than the limit of 3MB.")
	}

	if err != nil {
		log.Fatalf("file couldnt be received: %v", err)
	}
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	/* TODO: We will be uploading the file to arweave. We may not need all these actions below */

	out, err := os.Create("/tmp/file")
	if err != nil {
		fmt.Fprintf(w, "Failed to open the file for writing")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	// the header contains useful info, like the original file name
	fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
}

func ListenForFile(filePort string) {
	server := &http.Server{Addr: ":" + filePort, Handler: http.HandlerFunc(handler)}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Error starting http server:", err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error shutting down http server:", err)
	}
}