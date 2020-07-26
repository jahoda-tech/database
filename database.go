package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Alarm struct {
	gorm.Model
	Name          string `gorm:"unique"`
	WorkplaceID   int
	Workplace     Workplace
	SqlCommand    string
	MessageHeader string
	MessageText   string
	Recipients    string
	Url           string
	Pdf           string
}

type AlarmRecord struct {
	gorm.Model
	DateTimeStart     time.Time
	DateTimeEnd       sql.NullTime
	DateTimeProcessed sql.NullTime
	Duration          time.Duration
	AlarmID           int
	Alarm             Alarm
	WorkplaceID       int
	Workplace         Workplace
}

type SystemRecord struct {
	gorm.Model
	DatabaseSizeInMegaBytes     float32
	DatabaseGrowthInMegaBytes   float32
	DiscFreeSizeInMegaBytes     float32
	EstimatedDiscFreeSizeInDays float32
}

type State struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Color string
	Note  string
}

type StateRecord struct {
	gorm.Model
	DateTimeStart time.Time
	StateID       int
	State         State
	WorkplaceID   int
	Workplace     Workplace
	Note          string
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time
	DateTimeEnd   sql.NullTime
	OrderRecordID int
	OrderRecord   OrderRecord
	UserID        int
	User          User
	Note          string
}

type DownTimeRecord struct {
	gorm.Model
	DateTimeStart time.Time
	DateTimeEnd   sql.NullTime
	DeviceID      int
	Device        Device
	WorkplaceID   int
	Workplace     Workplace
	OrderRecordID int
	OrderRecord   OrderRecord
	DowntimeID    int
	Downtime      Downtime
	UserID        int
	User          User
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time
	DateTimeEnd   sql.NullTime
	BreakdownID   int
	Breakdown     Breakdown
	DeviceID      int
	Device        Device
	OrderRecordID int
	OrderRecord   OrderRecord
	UserID        int
	User          User
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time
	DeviceID      int
	Device        Device
	OrderRecordID int
	OrderRecord   OrderRecord
	FaultID       int
	Fault         Fault
	UserID        int
	User          User
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time
	DeviceID      int
	Device        Device
	OrderRecordID int
	OrderRecord   OrderRecord
	PackageID     int
	Package       Package
	UserID        int
	User          User
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time
	DeviceID      int
	Device        Device
	OrderRecordID int
	OrderRecord   OrderRecord
	PartID        int
	Part          Part
	UserID        int
	User          User
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time
	DateTimeEnd     sql.NullTime
	DeviceID        int
	Device          Device
	OrderID         int
	Order           Order
	OperationID     int
	Operation       Operation
	WorkplaceID     int
	Workplace       Workplace
	WorkplaceModeID int
	WorkplaceMode   WorkplaceMode
	WorkshiftID     int
	Workshift       Workshift
	AverageCycle    float32
	Cavity          int
	CountOk         int
	CountNok        int
	Note            string
}

type Operation struct {
	gorm.Model
	Name    string `gorm:"unique"`
	OrderID int
	Order   Order
	Barcode int
	Note    string
}

type WorkplaceSection struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}
type Order struct {
	gorm.Model
	Name            string `gorm:"unique"`
	ProductID       sql.NullInt32
	Product         Product
	WorkplaceID     sql.NullInt32
	Workplace       Workplace
	Barcode         int
	DateTimeRequest sql.NullTime
	Cavity          int
	CountRequest    int
	Note            string
}
type Product struct {
	gorm.Model
	Name             string `gorm:"unique"`
	Barcode          int
	CycleTime        int
	DownTimeDuration time.Duration
	Note             string
}

type Part struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Barcode int
	Note    string
}

type Workplace struct {
	gorm.Model
	Name                   string `gorm:"unique"`
	WorkplaceSectionID     int
	WorkplaceSection       WorkplaceSection
	Code                   string
	PowerOffPortDateTime   sql.NullTime
	ProductionPortDateTime sql.NullTime
	ProductionPortValue    sql.NullInt32
	Note                   string
}

type WorkplacePort struct {
	gorm.Model
	Name         string
	DevicePortID int
	DevicePort   DevicePort
	StateID      int
	State        State
	WorkplaceID  int
	Workplace    Workplace
	Color        string
	CounterOK    bool
	CounterNOK   bool
	HighValue    float32
	LowValue     float32
	Note         string
}

type WorkplaceMode struct {
	gorm.Model
	Name             string `gorm:"unique"`
	DowntimeDuration time.Duration
	PoweroffDuration time.Duration
	Note             string
}
type WorkplaceWorkshift struct {
	gorm.Model
	WorkplaceID int
	Workplace   Workplace
	WorkshiftID int
	Workshift   Workshift
}
type Workshift struct {
	gorm.Model
	Name           string `gorm:"unique"`
	WorkshiftStart time.Time
	WorkshiftEnd   time.Time
	Note           string
}

type User struct {
	gorm.Model
	FirstName  string
	SecondName string
	UserRoleID int
	UserRole   UserRole
	UserTypeID int
	UserType   UserType
	Barcode    string
	Email      string
	Login      string
	Password   string
	Phone      string
	Pin        string
	Position   string
	Rfid       string
	Note       string
}

type UserRole struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type UserType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Downtime struct {
	gorm.Model
	Name           string `gorm:"unique"`
	DowntimeTypeID int
	DowntimeType   DowntimeType
	Barcode        string
	Color          string
	Note           string
}

type DowntimeType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Breakdown struct {
	gorm.Model
	Name            string `gorm:"unique"`
	BreakdownTypeID int
	BreakdownType   BreakdownType
	Barcode         string
	Color           string
	Note            string
}

type BreakdownType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Fault struct {
	gorm.Model
	Name        string `gorm:"unique"`
	FaultTypeID int
	FaultType   FaultType
	Barcode     string
	Note        string
}

type FaultType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Package struct {
	gorm.Model
	Name          string `gorm:"unique"`
	OrderID       int
	Order         Order
	PackageTypeID int
	PackageType   PackageType
	Barcode       string
	Note          string
}

type PackageType struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Count int
	Note  string
}

type DevicePortType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   string
	Enabled bool
	Note    string
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Device struct {
	gorm.Model
	Name         string `gorm:"unique"`
	DeviceTypeID int
	DeviceType   DeviceType
	Activated    bool
	IpAddress    string `gorm:"unique"`
	MacAddress   string
	Settings     string
	TypeName     string
	Note         string
}

type DeviceWorkplaceRecord struct {
	gorm.Model
	DeviceID    int
	Device      Device
	WorkplaceID int
	Workplace   Workplace
}

type DevicePort struct {
	gorm.Model
	Name             string
	DeviceID         int
	Device           Device
	DevicePortTypeID int
	DevicePortType   DevicePortType
	PortNumber       int
	PlcDataType      string
	PlcDataAddress   string
	Settings         string
	Unit             string
	Virtual          bool
	Note             string
}

type DevicePortAnalogRecord struct {
	ID           int       `gorm:"primary_key"`
	DateTime     time.Time `gorm:"unique_index:unique_analog_data"`
	DevicePortID int       `gorm:"unique_index:unique_analog_data"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primary_key"`
	DateTime     time.Time `gorm:"unique_index:unique_digital_data"`
	DevicePortID int       `gorm:"unique_index:unique_digital_data"`
	DevicePort   DevicePort
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primary_key"`
	DateTime     time.Time `gorm:"unique_index:unique_serial_data"`
	DevicePortID int       `gorm:"unique_index:unique_serial_data"`
	DevicePort   DevicePort
	Data         float32
}
