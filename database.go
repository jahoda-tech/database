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
	WorkplaceID   sql.NullInt32 `gorm:"uniqueIndex:unique_alarm"`
	Workplace     Workplace
	SqlCommand    string
	MessageHeader string
	MessageText   string
	Recipients    string
	Url           string
	Pdf           string
	Note          string
	Data          datatypes.JSON
	ExternalId    sql.NullInt32 `gorm:"index"`
}

type AlarmRecord struct {
	gorm.Model
	DateTimeStart     time.Time `gorm:"uniqueIndex:unique_alarm_record"`
	DateTimeEnd       sql.NullTime
	DateTimeProcessed sql.NullTime
	AlarmID           int `gorm:"uniqueIndex:unique_alarm_record"`
	Alarm             Alarm
	WorkplaceID       sql.NullInt32
	Workplace         Workplace
	UserID            sql.NullInt32
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
	State string
	Note  string
	Data  datatypes.JSON
}

type StateRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_state_records_data,priority:3"`
	StateID       int       `gorm:"uniqueIndex:unique_state_records_data,priority:2"`
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
	Data     datatypes.JSON
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_user_record"`
	DateTimeEnd   sql.NullTime
	OrderRecordID int `gorm:"uniqueIndex:unique_user_record"`
	OrderRecord   OrderRecord
	WorkplaceID   int `gorm:"uniqueIndex:unique_user_record"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_user_record"`
	User          User
	Note          string
}

type DowntimeRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_downtime_record"`
	DateTimeEnd   sql.NullTime
	WorkplaceID   int `gorm:"uniqueIndex:unique_downtime_record"`
	Workplace     Workplace
	DowntimeID    int `gorm:"uniqueIndex:unique_downtime_record"`
	Downtime      Downtime
	OrderRecordID sql.NullInt32 `gorm:"uniqueIndex:unique_downtime_record"`
	OrderRecord   OrderRecord
	UserID        sql.NullInt32 `gorm:"uniqueIndex:unique_downtime_record"`
	OrderID       sql.NullInt32
	Order         Order
	OperationID   sql.NullInt32
	Operation     Operation
	ProductID     sql.NullInt32
	Product       Product
	User          User
	Consumption   float32
	TransferState string
	Note          string
	Data          datatypes.JSON
}

type ImageRecord struct {
	gorm.Model
	DateTimeStart time.Time     `gorm:"uniqueIndex:unique_downtime_record"`
	OrderRecordID sql.NullInt32 `gorm:"uniqueIndex:unique_downtime_record"`
	OrderRecord   OrderRecord
	Image         []byte
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_breakdown_record"`
	DateTimeEnd   sql.NullTime
	BreakdownID   int `gorm:"uniqueIndex:unique_breakdown_record"`
	Breakdown     Breakdown
	WorkplaceID   int `gorm:"uniqueIndex:unique_breakdown_record"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_breakdown_record"`
	User          User
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_fault_record"`
	OrderRecordID int
	OrderRecord   OrderRecord
	FaultID       int `gorm:"uniqueIndex:unique_fault_record"`
	Fault         Fault
	WorkplaceID   int `gorm:"uniqueIndex:unique_fault_record"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_fault_record"`
	User          User
	OrderID       sql.NullInt32
	Order         Order
	OperationID   sql.NullInt32
	Operation     Operation
	ProductID     sql.NullInt32
	Product       Product
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_package_record"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PackageID     int `gorm:"uniqueIndex:unique_package_record"`
	Package       Package
	WorkplaceID   int `gorm:"uniqueIndex:unique_package_record"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_package_record"`
	User          User
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_part_record"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PartID        int `gorm:"uniqueIndex:unique_part_record"`
	Part          Part
	WorkplaceID   int `gorm:"uniqueIndex:unique_part_record"`
	Workplace     Workplace
	UserID        int `gorm:"uniqueIndex:unique_part_record"`
	User          User
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart      time.Time    `gorm:"uniqueIndex:unique_order_record;index:idx_order_datetime_workplace;index:idx_order_datetime_start"`
	DateTimeEnd        sql.NullTime `gorm:"index:idx_order_datetime_end"`
	OrderID            int          `gorm:"uniqueIndex:unique_order_record"`
	Order              Order
	OperationID        int `gorm:"uniqueIndex:unique_order_record"`
	Operation          Operation
	WorkplaceID        int `gorm:"uniqueIndex:unique_order_record;index:idx_order_datetime_workplace;index:idx_order_workplace"`
	Workplace          Workplace
	UserID             sql.NullInt32 `gorm:"uniqueIndex:unique_order_record"`
	User               User
	ProductID          sql.NullInt32
	Product            Product
	WorkplaceModeID    int
	WorkplaceMode      WorkplaceMode
	WorkshiftID        int
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
	Data               datatypes.JSON
}

type Operation struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:unique_operation"`
	OrderID         int
	Order           Order
	ProductID       sql.NullInt32
	Product         Product
	Barcode         string `gorm:"uniqueIndex:unique_operation"`
	DateTimeRequest sql.NullTime
	Note            string
	Information     string
	Data            datatypes.JSON
	ExternalId      sql.NullInt32 `gorm:"index"`
}

type WorkplaceSection struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:workplace_section"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Order struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:unique_order"`
	ProductID       sql.NullInt32
	Product         Product
	WorkplaceID     sql.NullInt32
	Workplace       Workplace
	Barcode         string `gorm:"uniqueIndex:unique_order"`
	DateTimeRequest sql.NullTime
	Cavity          int
	CountRequest    int
	Note            string
	Information     string
	Data            datatypes.JSON
	ExternalId      sql.NullInt32 `gorm:"index"`
}

type Product struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_product"`
	Barcode          string `gorm:"uniqueIndex:unique_product, uniqueIndex:unique_barcode"`
	Unit             sql.NullString
	DownTimeDuration time.Duration
	ProductTypeID    int `gorm:"default:1"`
	ProductType      ProductType
	CountTypeID      int `gorm:"default:1"`
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
	Data             datatypes.JSON
	ExternalId       sql.NullInt32 `gorm:"index"`
}

type ProductType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_product_type"`
	Type       string
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type CountType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_count_type"`
	Type string
	Note string
	Data datatypes.JSON
}

type Part struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_part"`
	Barcode    string
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type InformationRecord struct {
	gorm.Model
	Information string        `gorm:"uniqueIndex:unique_information_record"`
	OrderID     sql.NullInt32 `gorm:"uniqueIndex:unique_information_record"`
	Order       Order
	OperationID sql.NullInt32 `gorm:"uniqueIndex:unique_information_record"`
	Operation   Operation
	Note        string
	Data        datatypes.JSON
}

type Workplace struct {
	gorm.Model
	Name                       string `gorm:"uniqueIndex:unique_workplace"`
	Code                       string
	WorkplaceModeID            int
	WorkplaceMode              WorkplaceMode
	Voltage                    int
	PowerFactor                float32
	ConsumptionTypeID          int `gorm:"default:1"`
	ConsumptionType            ConsumptionType
	ConsumptionImpulsesPerWatt float32
	Unit                       string
	Note                       string
	Data                       datatypes.JSON
	ExternalId                 sql.NullInt32 `gorm:"index"`
}

type ConsumptionType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_consumption_type"`
	Code string
	Note string
	Data datatypes.JSON
}

type WorkplacePort struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex:unique_workplace_port"`
	DevicePortID int    `gorm:"uniqueIndex:unique_workplace_port"`
	DevicePort   DevicePort
	StateID      sql.NullInt32
	State        State
	WorkplaceID  int `gorm:"uniqueIndex:unique_workplace_port"`
	Workplace    Workplace
	Color        sql.NullString
	CounterOK    bool `gorm:"default:false"`
	CounterNOK   bool `gorm:"default:false"`
	HighValue    float32
	LowValue     float32
	Note         string
	Data         datatypes.JSON
}

type WorkplaceMode struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_workplace_mode"`
	DowntimeDuration time.Duration
	PoweroffDuration time.Duration
	Note             string
	Data             datatypes.JSON
	ExternalId       sql.NullInt32 `gorm:"index"`
}

type WorkplaceWorkshift struct {
	gorm.Model
	WorkplaceID int `gorm:"uniqueIndex:unique_workplace_workshift"`
	Workplace   Workplace
	WorkshiftID int `gorm:"uniqueIndex:unique_workplace_workshift"`
	Workshift   Workshift
	Note        string
	Data        datatypes.JSON
}

type Workshift struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex:unique_workshift"`
	WorkshiftStart int
	WorkshiftEnd   int
	Note           string
	Data           datatypes.JSON
	ExternalId     sql.NullInt32 `gorm:"index"`
}

type User struct {
	gorm.Model
	FirstName  string `gorm:"uniqueIndex:unique_user"`
	SecondName string `gorm:"uniqueIndex:unique_user"`
	UserRoleID int
	UserRole   UserRole
	UserTypeID int
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
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type UserRole struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_user_role"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type UserType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_user_type"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Downtime struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex:unique_downtime"`
	DowntimeTypeID int
	DowntimeType   DowntimeType
	Barcode        string
	Color          string
	Note           string
	Data           datatypes.JSON
	ExternalId     sql.NullInt32 `gorm:"index"`
}

type DowntimeType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_downtime_type"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Breakdown struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:unique_breakdown"`
	BreakdownTypeID int
	BreakdownType   BreakdownType
	Barcode         string
	Color           string
	Note            string
	Data            datatypes.JSON
	ExternalId      sql.NullInt32 `gorm:"index"`
}

type BreakdownType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_breakdown_type"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Fault struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex:unique_fault"`
	FaultTypeID int
	FaultType   FaultType
	Barcode     string
	Note        string
	Data        datatypes.JSON
	ExternalId  sql.NullInt32 `gorm:"index"`
}

type FaultType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_fault_type"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Package struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex:unique_package"`
	ProductID     int    `gorm:"uniqueIndex:unique_package"`
	Product       Product
	PackageTypeID int
	PackageType   PackageType
	Barcode       string
	Note          string
	Data          datatypes.JSON
	ExternalId    sql.NullInt32 `gorm:"index"`
}

type PackageType struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_package_type"`
	Count      int
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type DevicePortType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_device_port_type"`
	Note string
	Data datatypes.JSON
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex:unique_settings"`
	Value   string
	Enabled bool `gorm:"default:true"`
	Note    string
	Data    datatypes.JSON
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_device_type"`
	Note string
	Data datatypes.JSON
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
	Data         datatypes.JSON
	ExternalId   sql.NullInt32 `gorm:"index"`
}

type DeviceWorkplaceRecord struct {
	gorm.Model
	DeviceID    int `gorm:"uniqueIndex:device_workplace_record"`
	Device      Device
	WorkplaceID int `gorm:"uniqueIndex:device_workplace_record"`
	Workplace   Workplace
	Note        string
}

type DevicePort struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:device_port"`
	DeviceID         int    `gorm:"uniqueIndex:device_port"`
	Device           Device
	DevicePortTypeID int
	DevicePortType   DevicePortType
	PortNumber       int
	PlcDataType      string
	PlcDataAddress   string
	Settings         string
	Unit             string
	Virtual          bool    `gorm:"default:false"`
	Threshold        float64 `gorm:"default:-999999"`
	Note             string
	Data             datatypes.JSON
}

type DevicePortAnalogRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_analog_records_data,priority:2"`
	DevicePortID int       `gorm:"uniqueIndex:unique_analog_records_data,priority:1"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortSpecialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_special_records_data,index,priority:2"`
	DevicePortID int       `gorm:"uniqueIndex:unique_special_records_data,index,priority:1"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_digital_records_data,priority:2"`
	DevicePortID int       `gorm:"uniqueIndex:unique_digital_records_data,priority:1"`
	DevicePort   DevicePort
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_serial_records_data,priority:2"`
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
	Data datatypes.JSON
}

type WebUserRecord struct {
	gorm.Model
	UserEmail string    `gorm:"uniqueIndex:unique_web_user_record"`
	WebPage   string    `gorm:"uniqueIndex:unique_web_user_record"`
	DateTime  time.Time `gorm:"uniqueIndex:unique_web_user_record"`
	Note      string
}

type SummaryRecord struct {
	gorm.Model
	WorkplaceID           int `gorm:"uniqueIndex:unique_summary_record"`
	Workplace             Workplace
	DateTime              time.Time `gorm:"uniqueIndex:unique_summary_record"`
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
	DateTime              time.Time `gorm:"uniqueIndex:unique_shift_summary_record"`
	WorkshiftID           int       `gorm:"uniqueIndex:unique_shift_summary_record"`
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
	Data datatypes.JSON
}

type Bookmark struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_bookmark"`
	Url  string
	Note string
	Data datatypes.JSON
}

type Layout struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex:unique_layout"`
	Image    []byte
	ImageUrl string
	Note     string
	Data     datatypes.JSON
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
	DayOfWeek                 sql.NullInt32
	DayOfMonth                sql.NullInt32
	TotalTimeFromLastRecord   sql.NullInt32
	PowerOnTimeFromLastRecord sql.NullInt32
	Note                      string
	Data                      datatypes.JSON
	ExternalId                sql.NullInt32 `gorm:"index"`
}

type Maintenance struct {
	gorm.Model
	Name              string `gorm:"uniqueIndex:unique_maintenance"`
	Information       string
	MaintenanceTypeID int
	MaintenanceType   MaintenanceType
	Image             []byte
	ImageUrl          string
	Note              string
	Data              datatypes.JSON
	ExternalId        sql.NullInt32 `gorm:"index"`
}

type Place struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex:unique_place"`
	Information string
	Note        string
	Data        datatypes.JSON
	ExternalId  sql.NullInt32 `gorm:"index"`
}

type MaintenanceWorkplaceRecord struct {
	gorm.Model
	MaintenanceID         int `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Maintenance           Maintenance
	WorkplaceID           sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Workplace             Workplace
	UserID                sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	User                  User
	PlaceID               sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_workplace_record"`
	Place                 Place
	AdditionalInformation string
	StartDate             time.Time
	EndDate               time.Time
	IntervalDays          int
	Note                  string
}

type MaintenanceRecord struct {
	gorm.Model
	MaintenanceID     int `gorm:"uniqueIndex:unique_maintenance_record"`
	Maintenance       Maintenance
	RequestedDateTime time.Time `gorm:"uniqueIndex:unique_maintenance_record"`
	RequestedUserID   sql.NullInt32
	RequestedUser     User
	DateTimeStart     sql.NullTime
	DateTimeEnd       sql.NullTime
	AlarmRecordID     sql.NullInt32
	AlarmRecord       AlarmRecord
	UserID            sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_record"`
	User              User
	WorkplaceID       sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_record"`
	Workplace         Workplace
	PlaceID           sql.NullInt32 `gorm:"uniqueIndex:unique_maintenance_record"`
	Place             Place
	MaintenanceNote   string
	Status            string
	Cost              float32
	ControlUserID     sql.NullInt32
	ControlUser       User
	ControlDateTime   sql.NullTime
	Image             []byte
	ImageUrl          string
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
	Repeat             sql.NullInt32
	WorkplaceID        sql.NullInt32
	Workplace          Workplace
	Image              []byte
	ImageUrl           string
	WorkplaceSectionID sql.NullInt32
	WorkplaceSection   WorkplaceSection
	UserTypeID         sql.NullInt32
	UserType           UserType
	ProductID          sql.NullInt32
	Product            Product
	OrderID            sql.NullInt32
	Order              Order
	OperationID        sql.NullInt32
	Operation          Operation
	Note               string
	Data               datatypes.JSON
}

type ChecklistRecord struct {
	gorm.Model
	DateTime       time.Time `gorm:"uniqueIndex:unique_checklist_record"`
	ChecklistID    int       `gorm:"uniqueIndex:unique_checklist_record"`
	Checklist      Checklist
	WorkplaceID    int `gorm:"uniqueIndex:unique_checklist_record"`
	Workplace      Workplace
	UserID         int `gorm:"uniqueIndex:unique_checklist_record"`
	User           User
	OrderID        int
	Order          Order
	OperationID    sql.NullInt32
	Operation      Operation
	ProductID      int
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
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type StockStateRecord struct {
	gorm.Model
	StockID         int `gorm:"uniqueIndex:unique_stock_state_record"`
	Stock           Stock
	ProductID       int `gorm:"uniqueIndex:unique_stock_state_record"`
	Product         Product
	SerialNumberID  sql.NullInt32 `gorm:"uniqueIndex:unique_stock_state_record"`
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt32 `gorm:"uniqueIndex:unique_stock_state_record"`
	BatchNumber     BatchNumber
	StockLocationID sql.NullInt32 `gorm:"uniqueIndex:unique_stock_state_record"`
	StockLocation   StockLocation
	Count           sql.NullInt32
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
	Data               datatypes.JSON
	ExternalId         sql.NullInt32 `gorm:"index"`
}

type SerialNumber struct {
	gorm.Model
	ProductID  int `gorm:"uniqueIndex:unique_serial_number"`
	Product    Product
	Number     string `gorm:"uniqueIndex:unique_serial_number"`
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type Company struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex:unique_company"`
	Code          string `gorm:"uniqueIndex:unique_company"`
	Country       string `gorm:"uniqueIndex:unique_company"`
	Address       string
	CompanyTypeID int
	CompanyType   CompanyType
	UserID        sql.NullInt32
	User          User
	Note          string
	Data          datatypes.JSON
	ExternalId    sql.NullInt32 `gorm:"index"`
}

func (Company) TableName() string {
	return "companies"
}

type CompanyType struct {
	gorm.Model
	Type string `gorm:"uniqueIndex:unique_company_type"`
	Note string
	Data datatypes.JSON
}

type StockLocation struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex:unique_stock_location"`
	StockID    int    `gorm:"uniqueIndex:unique_stock_location"`
	Stock      Stock
	MaxCount   sql.NullInt32
	MaxVolume  sql.NullFloat64
	Note       string
	Data       datatypes.JSON
	ExternalId sql.NullInt32 `gorm:"index"`
}

type StockOrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time `gorm:"uniqueIndex:unique_stock_order_record"`
	StockID         int       `gorm:"uniqueIndex:unique_stock_order_record"`
	Stock           Stock
	CompanyID       int `gorm:"uniqueIndex:unique_stock_order_record"`
	Company         Company
	RecordTypeID    int `gorm:"uniqueIndex:unique_stock_order_record"`
	RecordType      RecordType
	ProductID       int
	Product         Product
	DateTimeEnd     sql.NullTime
	StockLocationID sql.NullInt32
	StockLocation   StockLocation
	SerialNumberID  sql.NullInt32
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt32
	BatchNumber     BatchNumber
	Count           sql.NullInt32
	Volume          sql.NullFloat64
	CanChange       bool `gorm:"default:false"`
	Completed       bool `gorm:"default:false"`
	Note            string
	Data            datatypes.JSON
	ExternalId      sql.NullInt32 `gorm:"index"`
}

type StockRecord struct {
	gorm.Model
	DateTime           time.Time `gorm:"uniqueIndex:unique_stock_record"`
	UserID             int       `gorm:"uniqueIndex:unique_stock_record"`
	User               User
	RecordTypeID       int `gorm:"uniqueIndex:unique_stock_record"`
	RecordType         RecordType
	CompanyID          int `gorm:"uniqueIndex:unique_stock_record"`
	Company            Company
	StockInID          sql.NullInt32
	StockIn            Stock
	StockOutID         sql.NullInt32
	StockOut           Stock
	StockOrderRecordID sql.NullInt32
	StockOrderRecord   StockOrderRecord
	Closed             bool `gorm:"default:false"`
	Note               string
	ExternalId         sql.NullInt32 `gorm:"index"`
}

type StockRecordItem struct {
	gorm.Model
	StockRecordID      int
	StockRecord        StockRecord
	ProductID          int
	Product            Product
	BatchNumberID      sql.NullInt32
	BatchNumber        BatchNumber
	SerialNumberID     sql.NullInt32
	SerialNumber       SerialNumber
	StockLocationInID  sql.NullInt32
	StockLocationIn    StockLocation
	StockLocationOutID sql.NullInt32
	StockLocationOut   StockLocation
	StockOrderRecordID sql.NullInt32
	StockOrderRecord   StockOrderRecord
	Count              sql.NullInt32
	Volume             sql.NullFloat64
	Note               string
}

type RecordType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_record_type"`
	Type string
	Note string
	Data datatypes.JSON
}

type ProductPackageRecord struct {
	gorm.Model
	PackageBarcode string `gorm:"uniqueIndex:unique_product_package_record"`
	SerialNumberID sql.NullInt32
	SerialNumber   SerialNumber
	BatchNumberID  sql.NullInt32
	BatchNumber    BatchNumber
	Note           string
}

type Holiday struct {
	gorm.Model
	Date        time.Time `gorm:"uniqueIndex:unique_holiday"`
	CountryCode string    `gorm:"uniqueIndex:unique_holiday"`
	Name        string
	Holiday     bool `gorm:"default:false"`
	HolidayName string
	Note        string
	Data        datatypes.JSON
}

type FileRecord struct {
	gorm.Model
	Name        string        `gorm:"uniqueIndex:unique_product_file_record,priority:4"`
	ProductId   sql.NullInt32 `gorm:"uniqueIndex:unique_product_file_record,priority:1"`
	Product     Product
	OrderId     sql.NullInt32 `gorm:"uniqueIndex:unique_product_file_record,priority:2"`
	Order       Order
	OperationId sql.NullInt32 `gorm:"uniqueIndex:unique_product_file_record,priority:3"`
	Operation   Operation
	Url         string
	Note        string
}
