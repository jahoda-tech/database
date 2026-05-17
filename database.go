package database

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Alarm struct {
	gorm.Model
	Name          string        `gorm:"uniqueIndex:unique_alarm"`
	WorkplaceID   sql.NullInt64 `gorm:"uniqueIndex:unique_alarm;index"`
	Workplace     Workplace
	SqlCommand    string
	MessageHeader string
	MessageText   string
	Recipients    string
	Url           string
	Pdf           string
	Note          string
	Data          datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId    sql.NullInt64  `gorm:"index"`
}

type AlarmRecord struct {
	gorm.Model
	DateTimeStart     time.Time    `gorm:"uniqueIndex:unique_alarm_record;index:idx_alarm_record_brin,type:brin"`
	DateTimeEnd       sql.NullTime `gorm:"index"`
	DateTimeProcessed sql.NullTime `gorm:"index"`
	AlarmID           int          `gorm:"uniqueIndex:unique_alarm_record;index"`
	Alarm             Alarm
	WorkplaceID       sql.NullInt64 `gorm:"index;index:idx_alarm_record_open,where:date_time_end IS NULL"`
	Workplace         Workplace
	UserID            sql.NullInt64 `gorm:"index"`
	User              User
	Note              string
}

type SystemRecord struct {
	gorm.Model
	DatabaseSizeInMegaBytes     float32
	DatabaseGrowthInMegaBytes   float32
	DiscFreeSizeInMegaBytes     float32
	EstimatedDiscFreeSizeInDays float32
	Note                        string
}

type State struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex:unique_state"`
	Color string
	Type  string
	Note  string
	Data  datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type StateRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_state_records_data,priority:3;index:idx_state_record_brin,type:brin"`
	StateID       int       `gorm:"uniqueIndex:unique_state_records_data,priority:2;index"`
	State         State
	WorkplaceID   int `gorm:"uniqueIndex:unique_state_records_data,priority:1"`
	Workplace     Workplace
	Note          string
}

type PageCount struct {
	gorm.Model
	PageName string `gorm:"uniqueIndex:unique_page_count"`
	Count    int
	Note     string
	Data     datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time    `gorm:"uniqueIndex:unique_user_record;index:idx_user_record_brin,type:brin"`
	DateTimeEnd   sql.NullTime `gorm:"index"`
	OrderRecordID int          `gorm:"uniqueIndex:unique_user_record;index"`
	OrderRecord   OrderRecord
	WorkplaceID   int `gorm:"uniqueIndex:unique_user_record;index;index:idx_user_record_open,where:date_time_end IS NULL"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_user_record;index"`
	User          User
	OrderID       sql.NullInt64 `gorm:"index"`
	Order         Order
	OperationID   sql.NullInt64 `gorm:"index"`
	Operation     Operation
	ProductID     sql.NullInt64 `gorm:"index"`
	Product       Product
	Note          string
}

type DowntimeRecord struct {
	gorm.Model
	DateTimeStart time.Time    `gorm:"uniqueIndex:unique_downtime_record;index:idx_downtime_record_brin,type:brin"`
	DateTimeEnd   sql.NullTime `gorm:"index"`
	WorkplaceID   int          `gorm:"uniqueIndex:unique_downtime_record;index;index:idx_downtime_record_open,where:date_time_end IS NULL"`
	Workplace     Workplace
	DowntimeID    int `gorm:"uniqueIndex:unique_downtime_record;index"`
	Downtime      Downtime
	OrderRecordID sql.NullInt64 `gorm:"uniqueIndex:unique_downtime_record;index"`
	OrderRecord   OrderRecord
	UserID        sql.NullInt64 `gorm:"uniqueIndex:unique_downtime_record;index"`
	OrderID       sql.NullInt64 `gorm:"index"`
	Order         Order
	OperationID   sql.NullInt64 `gorm:"index"`
	Operation     Operation
	ProductID     sql.NullInt64 `gorm:"index"`
	Product       Product
	User          User
	Consumption   float32
	TransferState string
	Note          string
	Data          datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type ImageRecord struct {
	gorm.Model
	DateTimeStart time.Time     `gorm:"uniqueIndex:unique_image_record;index:idx_image_record_brin,type:brin"`
	OrderRecordID sql.NullInt64 `gorm:"uniqueIndex:unique_image_record;index"`
	OrderRecord   OrderRecord
	Image         []byte
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time    `gorm:"uniqueIndex:unique_breakdown_record;index:idx_breakdown_record_brin,type:brin"`
	DateTimeEnd   sql.NullTime `gorm:"index"`
	BreakdownID   int          `gorm:"uniqueIndex:unique_breakdown_record;index"`
	Breakdown     Breakdown
	WorkplaceID   int `gorm:"uniqueIndex:unique_breakdown_record;index;index:idx_breakdown_record_open,where:date_time_end IS NULL"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_breakdown_record;index"`
	User          User
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time     `gorm:"uniqueIndex:unique_fault_record;index:idx_fault_record_brin,type:brin"`
	OrderRecordID sql.NullInt64 `gorm:"index"`
	OrderRecord   OrderRecord
	FaultID       int `gorm:"uniqueIndex:unique_fault_record;index"`
	Fault         Fault
	WorkplaceID   int `gorm:"uniqueIndex:unique_fault_record;index"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_fault_record;index"`
	User          User
	OrderID       sql.NullInt64 `gorm:"index"`
	Order         Order
	OperationID   sql.NullInt64 `gorm:"index"`
	Operation     Operation
	ProductID     sql.NullInt64 `gorm:"index"`
	Product       Product
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_package_record;index:idx_package_record_brin,type:brin"`
	OrderRecordID int       `gorm:"index"`
	OrderRecord   OrderRecord
	PackageID     int `gorm:"uniqueIndex:unique_package_record;index"`
	Package       Package
	WorkplaceID   int `gorm:"uniqueIndex:unique_package_record;index"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_package_record;index"`
	User          User
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_part_record;index:idx_part_record_brin,type:brin"`
	OrderRecordID int       `gorm:"index"`
	OrderRecord   OrderRecord
	PartID        int `gorm:"uniqueIndex:unique_part_record;index"`
	Part          Part
	WorkplaceID   int `gorm:"uniqueIndex:unique_part_record;index"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_part_record;index"`
	User          User
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart      time.Time    `gorm:"uniqueIndex:unique_order_record;index:idx_order_datetime_workplace;index:idx_order_record_brin,type:brin"`
	DateTimeEnd        sql.NullTime `gorm:"index:idx_order_datetime_end"`
	OrderID            int          `gorm:"uniqueIndex:unique_order_record;index"`
	Order              Order
	OperationID        int `gorm:"uniqueIndex:unique_order_record;index"`
	Operation          Operation
	WorkplaceID        int `gorm:"uniqueIndex:unique_order_record;index:idx_order_datetime_workplace;index:idx_order_workplace;index:idx_order_record_open,where:date_time_end IS NULL"`
	Workplace          Workplace
	UserID             sql.NullInt64 `gorm:"uniqueIndex:unique_order_record;index"`
	User               User
	ProductID          sql.NullInt64 `gorm:"index"`
	Product            Product
	WorkplaceModeID    int `gorm:"index"`
	WorkplaceMode      WorkplaceMode
	WorkshiftID        int `gorm:"index"`
	Workshift          Workshift
	AverageCycle       float32
	Cavity             int
	CountOk            int
	CountNok           int
	Consumption        float32
	ProductionDuration time.Duration
	DowntimeDuration   time.Duration
	TransferState      string
	Note               string
	Data               datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type NoteRecord struct {
	gorm.Model
	DateTimeStart time.Time     `gorm:"uniqueIndex:unique_note_record;index:idx_note_record_brin,type:brin"`
	DateTimeEnd   sql.NullTime  `gorm:"index"`
	OrderRecordID sql.NullInt64 `gorm:"index"`
	OrderRecord   OrderRecord
	OrderID       sql.NullInt64 `gorm:"index"`
	Order         Order
	OperationID   sql.NullInt64 `gorm:"index"`
	Operation     Operation
	WorkplaceID   sql.NullInt64 `gorm:"index;index:idx_note_record_open,where:date_time_end IS NULL"`
	Workplace     Workplace
	UserID        sql.NullInt64 `gorm:"index"`
	User          User
	ProductID     sql.NullInt64 `gorm:"index"`
	Product       Product
	NoteText      string `gorm:"uniqueIndex:unique_note_record"`
	Note          string
	Data          datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Operation struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:unique_operation"`
	OrderID         int    `gorm:"index"`
	Order           Order
	ProductID       sql.NullInt64 `gorm:"index"`
	Product         Product
	Barcode         string `gorm:"uniqueIndex:unique_operation"`
	DateTimeRequest sql.NullTime
	Note            string
	Information     string
	Data            datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId      sql.NullInt64  `gorm:"index"`
}

type WorkplaceSection struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_workplace_section"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Order struct {
	gorm.Model
	Name            string        `gorm:"uniqueIndex:unique_order"`
	ProductID       sql.NullInt64 `gorm:"index"`
	Product         Product
	WorkplaceID     sql.NullInt64 `gorm:"index"`
	Workplace       Workplace
	Barcode         string `gorm:"uniqueIndex:unique_order"`
	DateTimeRequest sql.NullTime
	Cavity          int
	CountRequest    int
	Note            string
	Information     string
	Data            datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId      sql.NullInt64  `gorm:"index"`
}

type Product struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_product"`
	Barcode          string `gorm:"uniqueIndex:unique_product"`
	Unit             sql.NullString
	DowntimeDuration time.Duration
	ProductTypeID    int `gorm:"default:1;index"`
	ProductType      ProductType
	CountTypeID      int `gorm:"default:1;index"`
	CountType        CountType
	CycleTime        float64
	Location         bool `gorm:"default:false"`
	PurchasePrice    sql.NullFloat64
	SalePrice        sql.NullFloat64
	PartnerPrice     sql.NullFloat64
	Fee              sql.NullFloat64
	Image            []byte
	ImageUrl         string
	Note             string
	Information      string
	Data             datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId       sql.NullInt64  `gorm:"index"`
}

type ProductType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_product_type"`
	Type       string
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type CountType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_count_type"`
	Type string
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Part struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_part"`
	Barcode    string
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type InformationRecord struct {
	gorm.Model
	Information string        `gorm:"uniqueIndex:unique_information_record"`
	OrderID     sql.NullInt64 `gorm:"uniqueIndex:unique_information_record;index"`
	Order       Order
	OperationID sql.NullInt64 `gorm:"uniqueIndex:unique_information_record;index"`
	Operation   Operation
	Note        string
	Data        datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Workplace struct {
	gorm.Model
	Name                       string `gorm:"uniqueIndex:unique_workplace"`
	Code                       string
	WorkplaceModeID            int `gorm:"index"`
	WorkplaceMode              WorkplaceMode
	Phases                     int
	Voltage                    int
	PowerFactor                float32
	ConsumptionTypeID          int `gorm:"default:1;index"`
	ConsumptionType            ConsumptionType
	ConsumptionImpulsesPerWatt float32
	Unit                       string
	Note                       string
	Data                       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId                 sql.NullInt64  `gorm:"index"`
}

type ConsumptionType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_consumption_type"`
	Code string
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type WorkplacePort struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex:unique_workplace_port"`
	DevicePortID int    `gorm:"uniqueIndex:unique_workplace_port"`
	DevicePort   DevicePort
	StateID      sql.NullInt64 `gorm:"index"`
	State        State
	WorkplaceID  int `gorm:"uniqueIndex:unique_workplace_port"`
	Workplace    Workplace
	Color        sql.NullString
	CounterOK    bool `gorm:"default:false"`
	CounterNOK   bool `gorm:"default:false"`
	HighValue    float32
	LowValue     float32
	Note         string
	Data         datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type WorkplaceMode struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_workplace_mode"`
	DowntimeDuration time.Duration
	PoweroffDuration time.Duration
	Note             string
	Data             datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId       sql.NullInt64  `gorm:"index"`
}

type WorkplaceWorkshift struct {
	gorm.Model
	WorkplaceID int `gorm:"uniqueIndex:unique_workplace_workshift"`
	Workplace   Workplace
	WorkshiftID int `gorm:"uniqueIndex:unique_workplace_workshift"`
	Workshift   Workshift
	Note        string
	Data        datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Workshift struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex:unique_workshift"`
	WorkshiftStart int
	WorkshiftEnd   int
	Note           string
	Data           datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId     sql.NullInt64  `gorm:"index"`
}

type User struct {
	gorm.Model
	FirstName  string `gorm:"uniqueIndex:unique_user"`
	SecondName string `gorm:"uniqueIndex:unique_user"`
	UserRoleID int    `gorm:"index"`
	UserRole   UserRole
	UserTypeID int `gorm:"index"`
	UserType   UserType
	Barcode    string `gorm:"index"`
	Email      string `gorm:"uniqueIndex:unique_user"`
	Password   string
	Phone      string
	Pin        string `gorm:"index"`
	Position   string
	Rfid       string `gorm:"index"`
	Locale     string
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type UserRole struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_user_role"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type UserType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_user_type"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Downtime struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex:unique_downtime"`
	DowntimeTypeID int    `gorm:"index"`
	DowntimeType   DowntimeType
	Barcode        string
	Color          string
	Note           string
	Data           datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId     sql.NullInt64  `gorm:"index"`
}

type DowntimeType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_downtime_type"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Breakdown struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:unique_breakdown"`
	BreakdownTypeID int    `gorm:"index"`
	BreakdownType   BreakdownType
	Barcode         string
	Color           string
	Note            string
	Data            datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId      sql.NullInt64  `gorm:"index"`
}

type BreakdownType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_breakdown_type"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Fault struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex:unique_fault"`
	FaultTypeID int    `gorm:"index"`
	FaultType   FaultType
	Barcode     string
	Note        string
	Data        datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId  sql.NullInt64  `gorm:"index"`
}

type FaultType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_fault_type"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Package struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex:unique_package"`
	ProductID     int    `gorm:"uniqueIndex:unique_package"`
	Product       Product
	PackageTypeID int `gorm:"index"`
	PackageType   PackageType
	Barcode       string
	Note          string
	Data          datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId    sql.NullInt64  `gorm:"index"`
}

type PackageType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_package_type"`
	Count      int
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type DevicePortType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_device_port_type"`
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex:unique_settings"`
	Value   string
	Enabled bool `gorm:"default:true"`
	Note    string
	Data    datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_device_type"`
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Device struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex:unique_device"`
	DeviceTypeID int    `gorm:"uniqueIndex:unique_device"`
	DeviceType   DeviceType
	Activated    bool   `gorm:"default:false"`
	IpAddress    string `gorm:"uniqueIndex:unique_device"`
	MacAddress   string
	Settings     string
	TypeName     string
	Note         string
	Data         datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId   sql.NullInt64  `gorm:"index"`
}

type DeviceWorkplaceRecord struct {
	gorm.Model
	DeviceID    int `gorm:"uniqueIndex:unique_device_workplace_record"`
	Device      Device
	WorkplaceID int `gorm:"uniqueIndex:unique_device_workplace_record"`
	Workplace   Workplace
	Note        string
}

type DevicePort struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_device_port"`
	DeviceID         int    `gorm:"uniqueIndex:unique_device_port"`
	Device           Device
	DevicePortTypeID int `gorm:"index"`
	DevicePortType   DevicePortType
	PortNumber       int
	PlcDataType      string
	PlcDataAddress   string
	Settings         string
	Unit             string
	Virtual          bool    `gorm:"default:false"`
	Threshold        float64 `gorm:"default:-999999"`
	Note             string
	Data             datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type DevicePortAnalogRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_analog_records_data,priority:2;index:idx_analog_record_brin,type:brin"`
	DevicePortID int       `gorm:"uniqueIndex:unique_analog_records_data,priority:1"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortSpecialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_special_records_data,priority:2;index:idx_special_record_brin,type:brin"`
	DevicePortID int       `gorm:"uniqueIndex:unique_special_records_data,priority:1"`
	DevicePort   DevicePort
	Data         string
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_digital_records_data,priority:2;index:idx_digital_record_brin,type:brin"`
	DevicePortID int       `gorm:"uniqueIndex:unique_digital_records_data,priority:1"`
	DevicePort   DevicePort
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_serial_records_data,priority:2;index:idx_serial_record_brin,type:brin"`
	DevicePortID int       `gorm:"uniqueIndex:unique_serial_records_data,priority:1"`
	DevicePort   DevicePort
	Data         float32
}

type Locale struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_locale"`
	CsCZ string
	DeDE string
	EnUS string
	EsES string
	FrFR string
	ItIT string
	PlPL string
	PtPT string
	SkSK string
	RuRU string
	UkUA string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type WebUserRecord struct {
	gorm.Model
	UserEmail string    `gorm:"uniqueIndex:unique_web_user_record"`
	WebPage   string    `gorm:"uniqueIndex:unique_web_user_record"`
	DateTime  time.Time `gorm:"uniqueIndex:unique_web_user_record;index:idx_web_user_record_brin,type:brin"`
	Note      string
}

type SummaryRecord struct {
	gorm.Model
	WorkplaceID           int `gorm:"uniqueIndex:unique_summary_record"`
	Workplace             Workplace
	DateTime              time.Time `gorm:"uniqueIndex:unique_summary_record;index:idx_summary_record_brin,type:brin"`
	Production            time.Duration
	Downtime              time.Duration
	PowerOff              time.Duration
	CountOk               int
	CountNok              int
	CountFail             int
	Consumption           float32
	ProductionConsumption float32
	DowntimeConsumption   float32
	Note                  string
}

type ShiftSummaryRecord struct {
	gorm.Model
	WorkplaceID           int `gorm:"uniqueIndex:unique_shift_summary_record"`
	Workplace             Workplace
	DateTime              time.Time `gorm:"uniqueIndex:unique_shift_summary_record;index:idx_shift_summary_record_brin,type:brin"`
	WorkshiftID           int       `gorm:"uniqueIndex:unique_shift_summary_record;index"`
	Workshift             Workshift
	Production            time.Duration
	Downtime              time.Duration
	PowerOff              time.Duration
	CountOk               int
	CountNok              int
	CountFail             int
	Consumption           float32
	ProductionConsumption float32
	DowntimeConsumption   float32
	Note                  string
}

type Report struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_report"`
	Url  string
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Bookmark struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_bookmark"`
	Url  string
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type Layout struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex:unique_layout"`
	Image    []byte
	ImageUrl string
	Note     string
	Data     datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type WorkplaceSectionRecord struct {
	gorm.Model
	WorkplaceSectionID int `gorm:"uniqueIndex:unique_workplace_section_record"`
	WorkplaceSection   WorkplaceSection
	WorkplaceID        int `gorm:"uniqueIndex:unique_workplace_section_record"`
	Workplace          Workplace
	Note               string
}

type WebUserSettings struct {
	gorm.Model
	Email string `gorm:"uniqueIndex:unique_web_user_settings"`
	Type  string `gorm:"uniqueIndex:unique_web_user_settings"`
	Data  string
	Note  string
}

type MaintenanceType struct {
	gorm.Model
	Name                      string `gorm:"uniqueIndex:unique_maintenance_type"`
	DayOfWeek                 sql.NullInt64
	DayOfMonth                sql.NullInt64
	TotalTimeFromLastRecord   sql.NullInt64
	PowerOnTimeFromLastRecord sql.NullInt64
	Note                      string
	Data                      datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId                sql.NullInt64  `gorm:"index"`
}

type Maintenance struct {
	gorm.Model
	Name              string `gorm:"uniqueIndex:unique_maintenance"`
	Information       string
	MaintenanceTypeID int
	MaintenanceType   MaintenanceType
	Image             []byte
	ImageUrl          string
	RequestedValue    sql.NullFloat64
	RequestedUnit     string
	Note              string
	Data              datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId        sql.NullInt64  `gorm:"index"`
}

type Place struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex:unique_place"`
	Information string
	Note        string
	Data        datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId  sql.NullInt64  `gorm:"index"`
}

type MaintenanceWorkplaceRecord struct {
	gorm.Model
	MaintenanceID         int `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Maintenance           Maintenance
	WorkplaceID           sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Workplace             Workplace
	UserID                sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	User                  User
	PlaceID               sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Place                 Place
	AdditionalInformation string
	StartDate             sql.NullTime
	EndDate               sql.NullTime
	IntervalDays          int
	Note                  string
}

type MaintenanceRecord struct {
	gorm.Model
	MaintenanceID     int `gorm:"uniqueIndex:unique_maintenance_record;index"`
	Maintenance       Maintenance
	RequestedDateTime time.Time     `gorm:"uniqueIndex:unique_maintenance_record;index:idx_maintenance_record_brin,type:brin"`
	RequestedUserID   sql.NullInt64 `gorm:"index"`
	RequestedUser     User
	DateTimeStart     sql.NullTime  `gorm:"index"`
	DateTimeEnd       sql.NullTime  `gorm:"index"`
	AlarmRecordID     sql.NullInt64 `gorm:"index"`
	AlarmRecord       AlarmRecord
	UserID            sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_record;index"`
	User              User
	WorkplaceID       sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_record;index"`
	Workplace         Workplace
	PlaceID           sql.NullInt64 `gorm:"uniqueIndex:unique_maintenance_record;index"`
	Place             Place
	MaintenanceNote   string
	Status            string
	Cost              float32
	ControlUserID     sql.NullInt64 `gorm:"index"`
	ControlUser       User
	ControlDateTime   sql.NullTime
	Image             []byte
	ImageUrl          string
	MeasuredValue     sql.NullFloat64
	Note              string
}

type Checklist struct {
	gorm.Model
	Name               string `gorm:"uniqueIndex:unique_checklist"`
	Type               string
	Possibilities      string
	Text               string
	Start              string
	StartInterval      int
	Repeat             sql.NullInt64
	WorkplaceID        sql.NullInt64 `gorm:"index"`
	Workplace          Workplace
	Image              []byte
	ImageUrl           string
	WorkplaceSectionID sql.NullInt64 `gorm:"index"`
	WorkplaceSection   WorkplaceSection
	UserTypeID         sql.NullInt64 `gorm:"index"`
	UserType           UserType
	ProductID          sql.NullInt64 `gorm:"index"`
	Product            Product
	OrderID            sql.NullInt64 `gorm:"index"`
	Order              Order
	OperationID        sql.NullInt64 `gorm:"index"`
	Operation          Operation
	Note               string
	Data               datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type ChecklistRecord struct {
	gorm.Model
	DateTime       time.Time `gorm:"uniqueIndex:unique_checklist_record;index:idx_checklist_record_brin,type:brin"`
	ChecklistID    int       `gorm:"uniqueIndex:unique_checklist_record;index"`
	Checklist      Checklist
	WorkplaceID    int `gorm:"uniqueIndex:unique_checklist_record;index"`
	Workplace      Workplace
	UserID         int `gorm:"uniqueIndex:unique_checklist_record;index"`
	User           User
	OrderID        sql.NullInt64 `gorm:"index"`
	Order          Order
	OperationID    sql.NullInt64 `gorm:"index"`
	Operation      Operation
	ProductID      int `gorm:"index"`
	Product        Product
	ResultNumber   sql.NullFloat64
	ResultText     sql.NullString
	ResultOption   sql.NullString
	ResultDateTime sql.NullTime
	Note           string
}

type Stock struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_stock"`
	Code       string `gorm:"uniqueIndex:unique_stock"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type StockStateRecord struct {
	gorm.Model
	StockID         int `gorm:"uniqueIndex:unique_stock_state_record"`
	Stock           Stock
	ProductID       int `gorm:"uniqueIndex:unique_stock_state_record"`
	Product         Product
	SerialNumberID  sql.NullInt64 `gorm:"uniqueIndex:unique_stock_state_record"`
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt64 `gorm:"uniqueIndex:unique_stock_state_record"`
	BatchNumber     BatchNumber
	StockLocationID sql.NullInt64 `gorm:"uniqueIndex:unique_stock_state_record"`
	StockLocation   StockLocation
	Count           sql.NullInt64
	Volume          sql.NullFloat64
	Note            string
}

type BatchNumber struct {
	gorm.Model
	ProductID          int `gorm:"uniqueIndex:unique_batch_number"`
	Product            Product
	Number             string `gorm:"uniqueIndex:unique_batch_number"`
	ExpirationDuration time.Duration
	Note               string
	Data               datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId         sql.NullInt64  `gorm:"index"`
}

type SerialNumber struct {
	gorm.Model
	ProductID  int `gorm:"uniqueIndex:unique_serial_number"`
	Product    Product
	Number     string `gorm:"uniqueIndex:unique_serial_number"`
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type Company struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex:unique_company"`
	Code          string `gorm:"uniqueIndex:unique_company"`
	Country       string `gorm:"uniqueIndex:unique_company"`
	Address       string
	CompanyTypeID int `gorm:"index"`
	CompanyType   CompanyType
	UserID        sql.NullInt64 `gorm:"index"`
	User          User
	Note          string
	Data          datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId    sql.NullInt64  `gorm:"index"`
}

func (Company) TableName() string {
	return "companies"
}

type CompanyType struct {
	gorm.Model
	Type string `gorm:"uniqueIndex:unique_company_type"`
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type StockLocation struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_stock_location"`
	StockID    int    `gorm:"uniqueIndex:unique_stock_location"`
	Stock      Stock
	MaxCount   sql.NullInt64
	MaxVolume  sql.NullFloat64
	Note       string
	Data       datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId sql.NullInt64  `gorm:"index"`
}

type StockOrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time `gorm:"uniqueIndex:unique_stock_order_record;index:idx_stock_order_record_brin,type:brin"`
	StockID         int       `gorm:"uniqueIndex:unique_stock_order_record;index"`
	Stock           Stock
	CompanyID       int `gorm:"uniqueIndex:unique_stock_order_record;index"`
	Company         Company
	RecordTypeID    int `gorm:"uniqueIndex:unique_stock_order_record;index"`
	RecordType      RecordType
	ProductID       int `gorm:"index"`
	Product         Product
	DateTimeEnd     sql.NullTime  `gorm:"index"`
	StockLocationID sql.NullInt64 `gorm:"index"`
	StockLocation   StockLocation
	SerialNumberID  sql.NullInt64 `gorm:"index"`
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt64 `gorm:"index"`
	BatchNumber     BatchNumber
	Count           sql.NullInt64
	Volume          sql.NullFloat64
	CanChange       bool `gorm:"default:false"`
	Completed       bool `gorm:"default:false"`
	Note            string
	Data            datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId      sql.NullInt64  `gorm:"index"`
}

type StockRecord struct {
	gorm.Model
	DateTime           time.Time `gorm:"uniqueIndex:unique_stock_record;index:idx_stock_record_brin,type:brin"`
	UserID             int       `gorm:"uniqueIndex:unique_stock_record;index"`
	User               User
	RecordTypeID       int `gorm:"uniqueIndex:unique_stock_record;index"`
	RecordType         RecordType
	CompanyID          int `gorm:"uniqueIndex:unique_stock_record;index"`
	Company            Company
	StockInID          sql.NullInt64 `gorm:"index"`
	StockIn            Stock
	StockOutID         sql.NullInt64 `gorm:"index"`
	StockOut           Stock
	StockOrderRecordID sql.NullInt64 `gorm:"index"`
	StockOrderRecord   StockOrderRecord
	Closed             bool `gorm:"default:false"`
	Note               string
	Data               datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
	ExternalId         sql.NullInt64  `gorm:"index"`
}

type StockRecordItem struct {
	gorm.Model
	StockRecordID      int `gorm:"index"`
	StockRecord        StockRecord
	ProductID          int `gorm:"index"`
	Product            Product
	BatchNumberID      sql.NullInt64 `gorm:"index"`
	BatchNumber        BatchNumber
	SerialNumberID     sql.NullInt64 `gorm:"index"`
	SerialNumber       SerialNumber
	StockLocationInID  sql.NullInt64 `gorm:"index"`
	StockLocationIn    StockLocation
	StockLocationOutID sql.NullInt64 `gorm:"index"`
	StockLocationOut   StockLocation
	StockOrderRecordID sql.NullInt64 `gorm:"index"`
	StockOrderRecord   StockOrderRecord
	Count              sql.NullInt64
	Volume             sql.NullFloat64
	Note               string
}

type RecordType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_record_type"`
	Type string
	Note string
	Data datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type ProductPackageRecord struct {
	gorm.Model
	PackageBarcode string        `gorm:"uniqueIndex:unique_product_package_record"`
	SerialNumberID sql.NullInt64 `gorm:"index"`
	SerialNumber   SerialNumber
	BatchNumberID  sql.NullInt64 `gorm:"index"`
	BatchNumber    BatchNumber
	Note           string
}

type Holiday struct {
	gorm.Model
	Date        time.Time `gorm:"uniqueIndex:unique_holiday"`
	CountryCode string    `gorm:"uniqueIndex:unique_holiday"`
	Name        string
	IsHoliday   bool `gorm:"default:false"`
	HolidayName string
	Note        string
	Data        datatypes.JSON `gorm:"type:jsonb;index:,type:gin"`
}

type FileRecord struct {
	gorm.Model
	Name        string        `gorm:"uniqueIndex:unique_product_file_record,priority:4"`
	ProductID   sql.NullInt64 `gorm:"uniqueIndex:unique_product_file_record,priority:1"`
	Product     Product
	OrderID     sql.NullInt64 `gorm:"uniqueIndex:unique_product_file_record,priority:2"`
	Order       Order
	OperationID sql.NullInt64 `gorm:"uniqueIndex:unique_product_file_record,priority:3"`
	Operation   Operation
	Url         string
	Note        string
}
