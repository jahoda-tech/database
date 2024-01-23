package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Alarm struct {
	gorm.Model
	Name          string `gorm:"index:unique"`
	WorkplaceID   sql.NullInt32
	Workplace     Workplace
	SqlCommand    string
	MessageHeader string
	MessageText   string
	Recipients    string
	Url           string
	Pdf           string
	Note          string
}

type AlarmRecord struct {
	gorm.Model
	DateTimeStart     time.Time `gorm:"index:unique,unique_alarm_data"`
	DateTimeEnd       sql.NullTime
	DateTimeProcessed sql.NullTime
	AlarmID           int `gorm:"index:unique,unique_alarm_data"`
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
	Name  string `gorm:"index:unique"`
	Color string
	State string
	Note  string
}

type StateRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"index:unique,unique_state_data"`
	StateID       int       `gorm:"index:unique,unique_state_data"`
	State         State
	WorkplaceID   int `gorm:"index:unique,unique_state_data"`
	Workplace     Workplace
	Note          string
}

type PageCount struct {
	gorm.Model
	PageName string `gorm:"index:unique"`
	Count    int
	Note     string
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"index:unique,unique_user_data"`
	DateTimeEnd   sql.NullTime
	OrderRecordID int `gorm:"index:unique,unique_user_data"`
	OrderRecord   OrderRecord
	WorkplaceID   int `gorm:"index:unique,unique_user_data"`
	Workplace     Workplace
	UserID        int `gorm:"index:unique,unique_user_data"`
	User          User
	Note          string
}

type DowntimeRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"index:unique,unique_downtime_data"`
	DateTimeEnd   sql.NullTime
	WorkplaceID   int `gorm:"index:unique,unique_downtime_data"`
	Workplace     Workplace
	DowntimeID    int `gorm:"index:unique,unique_downtime_data"`
	Downtime      Downtime
	OrderRecordID sql.NullInt32
	OrderRecord   OrderRecord
	UserID        sql.NullInt32
	User          User
	Consumption   float32
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"index:unique,unique_breakdown_data"`
	DateTimeEnd   sql.NullTime
	BreakdownID   int `gorm:"index:unique,unique_breakdown_data"`
	Breakdown     Breakdown
	WorkplaceID   int `gorm:"index:unique,unique_breakdown_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"index:unique,unique_fault_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	FaultID       int `gorm:"index:unique,unique_fault_data"`
	Fault         Fault
	WorkplaceID   int `gorm:"index:unique,unique_fault_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"index:unique,unique_package_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PackageID     int `gorm:"index:unique,unique_package_data"`
	Package       Package
	WorkplaceID   int `gorm:"index:unique,unique_package_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"index:unique,unique_part_data"`
	OrderRecordID int
	OrderRecord   OrderRecord
	PartID        int `gorm:"index:unique,unique_part_data"`
	Part          Part
	WorkplaceID   int `gorm:"index:unique,unique_part_data"`
	Workplace     Workplace
	UserID        int
	User          User
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time `gorm:"index:unique,unique_order_data"`
	DateTimeEnd     sql.NullTime
	OrderID         int `gorm:"index:unique,unique_order_data"`
	Order           Order
	OperationID     int `gorm:"index:unique,unique_order_data"`
	Operation       Operation
	WorkplaceID     int `gorm:"index:unique,unique_order_data"`
	Workplace       Workplace
	UserID          sql.NullInt32
	User            User
	WorkplaceModeID int
	WorkplaceMode   WorkplaceMode
	WorkshiftID     int
	Workshift       Workshift
	AverageCycle    float32
	Cavity          int
	CountOk         int
	CountNok        int
	Consumption     float32
	Note            string
}

type Operation struct {
	gorm.Model
	Name    string
	OrderID int
	Order   Order
	Barcode string `gorm:"index:unique"`
	Note    string
}

type WorkplaceSection struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Order struct {
	gorm.Model
	Name            string
	ProductID       sql.NullInt32
	Product         Product
	WorkplaceID     sql.NullInt32
	Workplace       Workplace
	Barcode         string `gorm:"index:unique"`
	DateTimeRequest sql.NullTime
	Cavity          int
	CountRequest    int
	Note            string
}

type Product struct {
	gorm.Model
	Name             string `gorm:"index:unique"`
	Barcode          string
	Unit             sql.NullString
	DownTimeDuration time.Duration
	ProductTypeID    int `gorm:"default:1"`
	ProductType      ProductType
	CountTypeID      int `gorm:"default:1"`
	CountType        CountType
	CycleTime        float64
	Location         bool
	Image            []byte
	ImageUrl         string
	PurchasePrice    sql.NullFloat64
	SalePrice        sql.NullFloat64
	PartnerPrice     sql.NullFloat64
	Fee              sql.NullFloat64
	Note             string
}

type ProductType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Type string
	Note string
}

type CountType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Type string
	Note string
}

type Part struct {
	gorm.Model
	Name    string `gorm:"index:unique"`
	Barcode string
	Note    string
}

type Workplace struct {
	gorm.Model
	Name                       string `gorm:"index:unique"`
	Code                       string
	WorkplaceModeID            int
	WorkplaceMode              WorkplaceMode
	Voltage                    int
	PowerFactor                float32
	ConsumptionTypeID          int
	ConsumptionType            ConsumptionType
	ConsumptionImpulsesPerWatt float32
	Unit                       string
	Note                       string
}

type ConsumptionType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Code string
	Note string
}

type WorkplacePort struct {
	gorm.Model
	Name         string `gorm:"index:unique,unique_workplace_port_data"`
	DevicePortID int    `gorm:"index:unique,unique_workplace_port_data"`
	DevicePort   DevicePort
	StateID      sql.NullInt32
	State        State
	WorkplaceID  int `gorm:"index:unique,unique_workplace_port_data"`
	Workplace    Workplace
	Color        sql.NullString
	CounterOK    bool
	CounterNOK   bool
	HighValue    float32
	LowValue     float32
	Note         string
}

type WorkplaceMode struct {
	gorm.Model
	Name             string `gorm:"index:unique"`
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
	Note        string
}

type Workshift struct {
	gorm.Model
	Name           string `gorm:"index:unique"`
	WorkshiftStart int
	WorkshiftEnd   int
	Note           string
}

type User struct {
	gorm.Model
	FirstName  string `gorm:"index:unique,username_data"`
	SecondName string `gorm:"index:unique,username_data"`
	UserRoleID int
	UserRole   UserRole
	UserTypeID int
	UserType   UserType
	Barcode    string
	Email      string `gorm:"index:unique,username_data"`
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
	Name string `gorm:"index:unique"`
	Note string
}

type UserType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Downtime struct {
	gorm.Model
	Name           string `gorm:"index:unique"`
	DowntimeTypeID int
	DowntimeType   DowntimeType
	Barcode        string
	Color          string
	Note           string
}

type DowntimeType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Breakdown struct {
	gorm.Model
	Name            string `gorm:"index:unique"`
	BreakdownTypeID int
	BreakdownType   BreakdownType
	Barcode         string
	Color           string
	Note            string
}

type BreakdownType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Fault struct {
	gorm.Model
	Name        string `gorm:"index:unique"`
	FaultTypeID int
	FaultType   FaultType
	Barcode     string
	Note        string
}

type FaultType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Package struct {
	gorm.Model
	Name          string `gorm:"index:unique"`
	ProductID     int
	Product       Product
	PackageTypeID int
	PackageType   PackageType
	Barcode       string
	Note          string
}

type PackageType struct {
	gorm.Model
	Name  string `gorm:"index:unique"`
	Count int
	Note  string
}

type DevicePortType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"index:unique"`
	Value   string
	Enabled bool
	Note    string
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Note string
}

type Device struct {
	gorm.Model
	Name         string `gorm:"index:unique"`
	DeviceTypeID int
	DeviceType   DeviceType
	Activated    bool
	IpAddress    string `gorm:"index:unique"`
	MacAddress   string
	Settings     string
	TypeName     string
	Note         string
}

type DeviceWorkplaceRecord struct {
	gorm.Model
	DeviceID    int `gorm:"index:unique,device_workplace_data"`
	Device      Device
	WorkplaceID int `gorm:"index:unique,device_workplace_data"`
	Workplace   Workplace
	Note        string
}

type DevicePort struct {
	gorm.Model
	Name             string `gorm:"index:unique,device_port_data"`
	DeviceID         int    `gorm:"index:unique,device_port_data"`
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
	DateTime     time.Time `gorm:"index:unique,unique_analog_data;index:analog_data_date_time"`
	DevicePortID int       `gorm:"index:unique,unique_analog_data;index:analog_data_device_port_id"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortSpecialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"index:unique,unique_special_data;index:special_data_date_time"`
	DevicePortID int       `gorm:"index:unique,unique_special_data;index:special_data_device_port_id"`
	DevicePort   DevicePort
	Data         float32
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"index:unique,unique_digital_data;index:digital_data_date_time"`
	DevicePortID int       `gorm:"index:unique,unique_digital_data;index:digital_data_device_port_id"`
	DevicePort   DevicePort
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"index:unique,unique_serial_data;index:serial_data_date_time"`
	DevicePortID int       `gorm:"index:unique,unique_serial_data;index:serial_data_device_port_id"`
	DevicePort   DevicePort
	Data         float32
}

type Locale struct {
	gorm.Model
	Name string `gorm:"index:unique"`
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
}

type WebUserRecord struct {
	gorm.Model
	UserEmail string
	WebPage   string
	DateTime  time.Time
	Note      string
}

type SummaryRecord struct {
	gorm.Model
	WorkplaceID           int `gorm:"index:unique,unique_summary_data"`
	Workplace             Workplace
	DateTime              time.Time `gorm:"index:unique,unique_summary_data"`
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
	WorkplaceID           int `gorm:"index:unique,unique_summary_data"`
	Workplace             Workplace
	DateTime              time.Time `gorm:"index:unique,unique_summary_data"`
	WorkshiftID           int       `gorm:"index:unique,unique_summary_data"`
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
	Name string `gorm:"index:unique"`
	Url  string
	Note string
}

type Bookmark struct {
	gorm.Model
	Name string `gorm:"index:unique"`
	Url  string
	Note string
}

type Layout struct {
	gorm.Model
	Name     string `gorm:"index:unique"`
	Image    []byte
	ImageUrl string
	Note     string
}

type WorkplaceSectionRecord struct {
	gorm.Model
	WorkplaceSectionID int
	WorkplaceSection   WorkplaceSection
	WorkplaceID        int
	Workplace          Workplace
	Note               string
}

type WebUserSettings struct {
	gorm.Model
	Email string
	Type  string
	Data  string
	Note  string
}

type MaintenanceType struct {
	gorm.Model
	Name                      string `gorm:"index:unique"`
	DayOfWeek                 sql.NullInt32
	DayOfMonth                sql.NullInt32
	TotalTimeFromLastRecord   sql.NullInt32
	PowerOnTimeFromLastRecord sql.NullInt32
	Note                      string
}

type Maintenance struct {
	gorm.Model
	Name              string `gorm:"index:unique"`
	MaintenanceTypeID int
	MaintenanceType   MaintenanceType
	Note              string
}

type MaintenanceWorkplaceRecord struct {
	gorm.Model
	MaintenanceID int
	Maintenance   Maintenance
	WorkplaceID   int
	Workplace     Workplace
	Note          string
}

type MaintenanceRecord struct {
	gorm.Model
	MaintenanceID   int
	Maintenance     Maintenance
	DateTime        time.Time
	UserID          int
	User            User
	WorkplaceID     int
	Workplace       Workplace
	UserNote        string
	MaintenanceNote string
	ControlUserID   sql.NullInt32
	ControlUser     User
	ControlDateTime sql.NullTime
	Note            string
}

type Checklist struct {
	gorm.Model
	Name               string `gorm:"index:unique"`
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
}

type ChecklistRecord struct {
	gorm.Model
	DateTime       time.Time `gorm:"index:unique,unique_checklist_data"`
	ChecklistID    int       `gorm:"index:unique,unique_checklist_data"`
	Checklist      Checklist
	WorkplaceID    int `gorm:"index:unique,unique_checklist_data"`
	Workplace      Workplace
	UserId         int `gorm:"index:unique,unique_checklist_data"`
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
	Name string `gorm:"index:unique,unique_stock"`
	Code string `gorm:"index:unique,unique_stock"`
	Note string
}

type StockStateRecord struct {
	gorm.Model
	StockID         int
	Stock           Stock
	ProductID       int
	Product         Product
	SerialNumberID  sql.NullInt32
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt32
	BatchNumber     BatchNumber
	StockLocationID sql.NullInt32
	StockLocation   StockLocation
	Count           sql.NullInt32
	Volume          sql.NullFloat64
	Note            string
}

type BatchNumber struct {
	gorm.Model
	ProductID          int `gorm:"index:unique,unique_batch_number"`
	Product            Product
	Number             string `gorm:"index:unique,unique_batch_number"`
	ExpirationDuration time.Duration
	Note               string
}

type SerialNumber struct {
	gorm.Model
	ProductID int `gorm:"index:unique,unique_serial_number"`
	Product   Product
	Number    string `gorm:"index:unique,unique_serial_number"`
	Note      string
}

type Company struct {
	gorm.Model
	Name          string `gorm:"index:unique,unique_company"`
	Code          string `gorm:"index:unique,unique_company"`
	Country       string `gorm:"index:unique,unique_company"`
	Address       string
	CompanyTypeID int
	CompanyType   CompanyType
	UserID        sql.NullInt32
	User          User
	Note          string
}

func (Company) TableName() string {
	return "companies"
}

type CompanyType struct {
	gorm.Model
	Type string `gorm:"index:unique,unique_company_type"`
	Note string
}

type StockLocation struct {
	gorm.Model
	Name      string `gorm:"index:unique,unique_stock_location"`
	StockID   int    `gorm:"index:unique,unique_stock_location"`
	Stock     Stock
	MaxCount  sql.NullInt32
	MaxVolume sql.NullFloat64
	Note      string
}

type StockOrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time
	StockID         int
	Stock           Stock
	CompanyID       int
	Company         Company
	RecordTypeID    int
	RecordType      RecordType
	ProductID       int
	Product         Product
	ExternalID      int
	DateTimeEnd     sql.NullTime
	StockLocationID sql.NullInt32
	StockLocation   StockLocation
	SerialNumberID  sql.NullInt32
	SerialNumber    SerialNumber
	BatchNumberID   sql.NullInt32
	BatchNumber     BatchNumber
	Count           sql.NullInt32
	Volume          sql.NullFloat64
	CanChange       bool
	Completed       bool
	Note            string
}

type StockRecord struct {
	gorm.Model
	DateTime           time.Time `gorm:"index:unique,unique_stock_record"`
	UserID             int       `gorm:"index:unique,unique_stock_record"`
	User               User
	RecordTypeID       int `gorm:"index:unique,unique_stock_record"`
	RecordType         RecordType
	CompanyID          int `gorm:"index:unique,unique_stock_record"`
	Company            Company
	StockInID          sql.NullInt32
	StockIn            Stock
	StockOutID         sql.NullInt32
	StockOut           Stock
	StockOrderRecordID sql.NullInt32
	StockOrderRecord   StockOrderRecord
	Closed             bool
	Note               string
}

type StockRecordItem struct {
	gorm.Model
	StockRecordId      int
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
	Count              sql.NullInt32
	Volume             sql.NullFloat64
	Note               string
}

type RecordType struct {
	gorm.Model
	Name string
	Type string
	Note sql.NullString
}

type ProductPackageRecord struct {
	gorm.Model
	PackageBarcode string
	SerialNumberID sql.NullInt32
	SerialNumber   SerialNumber
	BatchNumberID  sql.NullInt32
	BatchNumber    BatchNumber
}

type Holiday struct {
	gorm.Model
	Date        time.Time
	Name        string
	Holiday     bool
	HolidayName string
}
