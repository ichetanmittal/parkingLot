package main

import (
    "parking-lot-service/db"
    "parking-lot-service/handlers"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // Initialize Database
    db.Init()

    // Routes
    e.GET("/free-slots", handlers.GetFreeSlots)
    e.POST("/park", handlers.ParkVehicle)
    e.POST("/unpark", handlers.UnparkVehicle)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}