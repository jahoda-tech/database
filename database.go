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
	DateTimeStart     time.Time `gorm:"uniqueIndex:unique_alarm_data"`
	DateTimeEnd       sql.NullTime
	DateTimeProcessed sql.NullTime
	AlarmID           int `gorm:"uniqueIndex:unique_alarm_data"`
	Alarm             Alarm
	WorkplaceID       int `gorm:"uniqueIndex:unique_alarm_data"`
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
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_state_data"`
	StateID       int
	State         State
	WorkplaceID   int `gorm:"uniqueIndex:unique_state_data"`
	Workplace     Workplace
	Note          string
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_user_data"`
	DateTimeEnd   sql.NullTime
	OrderRecordID int `gorm:"uniqueIndex:unique_user_data"`
	OrderRecord   OrderRecord
	WorkplaceID   int `gorm:"uniqueIndex:unique_user_data"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_user_data"`
	User          User
	Note          string
}

type DowntimeRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_downtime_data"`
	DateTimeEnd   sql.NullTime
	WorkplaceID   int `gorm:"uniqueIndex:unique_downtime_data"`
	Workplace     Workplace
	DowntimeID    int `gorm:"uniqueIndex:unique_downtime_data"`
	Downtime      Downtime
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_breakdown_data"`
	DateTimeEnd   sql.NullTime
	BreakdownID   int `gorm:"uniqueIndex:unique_breakdown_data"`
	Breakdown     Breakdown
	WorkplaceID   int `gorm:"uniqueIndex:unique_breakdown_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_fault_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	FaultID       int `gorm:"uniqueIndex:unique_fault_data"`
	Fault         Fault
	WorkplaceID   int `gorm:"uniqueIndex:unique_fault_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_package_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PackageID     int `gorm:"uniqueIndex:unique_package_data"`
	Package       Package
	WorkplaceID   int `gorm:"uniqueIndex:unique_package_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_part_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PartID        int `gorm:"uniqueIndex:unique_part_data"`
	Part          Part
	WorkplaceID   int `gorm:"uniqueIndex:unique_part_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time `gorm:"uniqueIndex:unique_order_data"`
	DateTimeEnd     sql.NullTime
	OrderID         int `gorm:"uniqueIndex:unique_order_data"`
	Order           Order
	OperationID     int `gorm:"uniqueIndex:unique_order_data"`
	Operation       Operation
	WorkplaceID     int `gorm:"uniqueIndex:unique_order_data"`
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
	Code                   string
	PowerOffPortDateTime   sql.NullTime
	ProductionPortDateTime sql.NullTime
	ProductionPortValue    sql.NullInt32
	WorkplaceSectionID     int
	WorkplaceSection       WorkplaceSection
	WorkplaceModeID        int
	WorkplaceMode          WorkplaceMode
	Note                   string
}

type WorkplacePort struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex:unique_workplace_port_data"`
	DevicePortID int    `gorm:"uniqueIndex:unique_workplace_port_data"`
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
	WorkplaceID int `gorm:"uniqueIndex:unique_workplace_workshift_data"`
	Workplace   Workplace
	WorkshiftID int `gorm:"uniqueIndex:unique_workplace_workshift_data"`
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
	FirstName  string `gorm:"uniqueIndex:username_data"`
	SecondName string `gorm:"uniqueIndex:username_data"`
	UserRoleID int
	UserRole   UserRole
	UserTypeID int
	UserType   UserType
	Barcode    string
	Email      string
	Login      string `gorm:"uniqueIndex:username_data"`
	Password   string
	Phone      string
	Pin        string
	Position   string
	Rfid       string
	Locale     string
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
	DeviceID    int `gorm:"uniqueIndex:device_workplace_data"`
	Device      Device
	WorkplaceID int `gorm:"uniqueIndex:device_workplace_data"`
	Workplace   Workplace
}

type DevicePort struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:device_port_data"`
	DeviceID         int    `gorm:"uniqueIndex:device_port_data"`
	Device           Device
	DevicePortTypeID int
	DevicePortType   DevicePortType
	PortNumber       int
	PlcDataType      string
	PlcDataAddress   string
	Settings         string
	Unit             string
	Virtual          bool `gorm:"default:false"`
	Note             string
}

type DevicePortAnalogRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_analog_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_analog_data"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_digital_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_digital_data"`
	DevicePort   DevicePort
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_serial_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_serial_data"`
	DevicePort   DevicePort
	Data         float32
}
