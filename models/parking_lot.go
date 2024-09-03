package models

import "time"

type VehicleType string

const (
    Motorcycle VehicleType = "Motorcycle"
    Car        VehicleType = "Car"
    Bus        VehicleType = "Bus"
)

type ParkingLot struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string
    Motorcycles int
    Cars        int
    Buses       int
}

type Vehicle struct {
    ID       uint   `gorm:"primaryKey"`
    Type     VehicleType
    LotID    uint
    EntryAt  time.Time
    ExitAt   *time.Time
    Tariff   float64
}