package main

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	// d := []byte{0x01, 0xff, 0x3a, 0xcd}
	log := setupLogger("")
	str := "68656c6c6f"

	dec, err := hex.DecodeString(str)
	if err != nil {
		fmt.Printf("err decod str: %v", err)
		os.Exit(1)
	}

	// s := hex.EncodeToString(d)

	log.Info("mtd DecodeString: %s", slog.Attr{Key: "hex", Value: slog.StringValue(string(dec))})

}

func setupLogger(confApp string) *slog.Logger {
	var log *slog.Logger
	switch confApp {
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
