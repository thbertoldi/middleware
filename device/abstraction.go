package device

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Sensor struct {
	gorm.Model
	Measure string
	Name string
	Value float64
	DeviceID uint
}

type Device struct {
	gorm.Model
	Type string
	Serial string
	Sensors []Sensor
}


func (s *Device) UpdateSensor(unit, name string, value float64) {
	sensor := Sensor{Measure: unit, Name: name, Value: value, DeviceID: s.ID}
	_ = db.FirstOrCreate(&sensor, "device_id = ? AND name = ? AND measure = ?", s.ID, name, unit)
	db.Model(&sensor).Update("Value", value)
}

var db *gorm.DB = initModule()

func initModule() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("devices.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&Device{}, &Sensor{})
	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func GetOrCreate(model string, serial string) Device{
	device := Device{Type: model, Serial: serial}
	_ = db.FirstOrCreate(&device, "type = ? AND serial = ?", model, serial)
	return device
}
