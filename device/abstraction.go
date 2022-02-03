package device

import (
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Sensor struct {
	gorm.Model
	Measure string `json:"measure"`
	Name string `json:"name"`
	Value float64 `json:"value"`
	DeviceID uint `json:"id"`
}

type Device struct {
	gorm.Model
	Type string `json:"type"`
	Serial string `json:"serial"`
	Sensors []Sensor`json:"sensors"`
}


func (s *Device) UpdateSensor(unit, name string, value float64) {
	sensor := Sensor{Measure: unit, Name: name, Value: value, DeviceID: s.ID}
	_ = db.FirstOrCreate(&sensor, "device_id = ? AND name = ? AND measure = ?", s.ID, name, unit)
	db.Model(&sensor).Update("Value", value)
}

func (s *Device) GenerateDash() string {
	return ""
}

var db *gorm.DB = initModule()

var cache map[uint]interface{} = make(map[uint]interface{})

func initModule() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("devices.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&Device{}, &Sensor{})
	if err != nil {
		panic("Failed to migrate database")
	}
	var devices []Device
	_ = db.Find(&devices)
	for _, dev := range(devices) {
		cache[dev.ID] = dev
	}
	return db
}

func GetByID(id uint) Device {
	var device Device
	db.Find(&device, id)
	return device
}

func GetOrCreate(model string, serial string) Device{
	device := Device{Type: model, Serial: serial}
	_ = db.FirstOrCreate(&device, "type = ? AND serial = ?", model, serial)
	cache[device.ID] = &device
	return device
}

func GetAll() string {
	cacheJson, err := json.Marshal(cache)
	if err != nil {
		return ""
	}
	return string(cacheJson)
}
