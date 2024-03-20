package main

import (
	"context"
	"fmt"
	"io"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)
	e.GET("/static", file)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func file(c echo.Context) error {
	return c.File("Logo-Colour-Transparant-1.png")
}
func upload(c echo.Context) error {
	// Read form fields
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dir := "uploads/"
	dirUser := dir + id

	if _, err := os.Stat(dirUser); os.IsNotExist(err) {
		fmt.Println(dirUser, "does not exist")
		os.Mkdir(dirUser, 0644)

	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
	// Destination
	dst, err := os.Create(dirUser + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}
