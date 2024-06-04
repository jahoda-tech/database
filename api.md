# Alarm

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_alarm&auth=access_key 
```

``` go
Name          string        
WorkplaceID   sql.NullInt32 
SqlCommand    string
MessageHeader string
MessageText   string
Recipients    string
Url           string
Pdf           string
Note          string
```

# AlarmRecord

- [x] MES
- [x] WMS

``` html
json example: http://ipaddress:port/alarm_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/alarm_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart     time.Time 
DateTimeEnd       sql.NullTime
DateTimeProcessed sql.NullTime
AlarmID           int 
Alarm             Alarm
WorkplaceID       sql.NullInt32
Workplace         Workplace
UserID            sql.NullInt32
User              User
Note              string
```

# SystemRecord

- [x] MES
- [x] WMS

``` html
json example: http://ipaddress:port/system_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/system_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DatabaseSizeInMegaBytes     float32
DatabaseGrowthInMegaBytes   float32
DiscFreeSizeInMegaBytes     float32
EstimatedDiscFreeSizeInDays float32
Note                        string
```

# State

- [x] MES
- [ ] WMS

``` go
Name  string 
Color string
State string
Note  string
```

# StateRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/state_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/state_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart time.Time 
StateID       int       
State         State
WorkplaceID   int 
Workplace     Workplace
Note          string
```

# PageCount

- [x] MES
- [x] WMS

``` go
PageName string 
Count    int
Note     string
```

# UserRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/user_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/user_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart time.Time 
DateTimeEnd   sql.NullTime
OrderRecordID int 
OrderRecord   OrderRecord
WorkplaceID   int 
Workplace     Workplace
UserID        int 
User          User
Note          string
```

# DowntimeRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/downtime_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/downtime_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart time.Time 
DateTimeEnd   sql.NullTime
WorkplaceID   int 
Workplace     Workplace
DowntimeID    int 
Downtime      Downtime
OrderRecordID sql.NullInt32 
OrderRecord   OrderRecord
UserID        sql.NullInt32 
User          User
Consumption   float32
Note          string
```

# BreakdownRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/breakdown_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/breakdown_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart time.Time 
DateTimeEnd   sql.NullTime
BreakdownID   int 
Breakdown     Breakdown
WorkplaceID   int 
Workplace     Workplace
UserID        int 
User          User
Note          string
```

# FaultRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/fault_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/fault_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTime      time.Time 
OrderRecordID int
OrderRecord   OrderRecord
FaultID       int 
Fault         Fault
WorkplaceID   int 
Workplace     Workplace
UserID        int 
User          User
Count         int
Note          string
```

# PackageRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/package_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/package_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTime      time.Time 
OrderRecordID int
OrderRecord   OrderRecord
PackageID     int 
Package       Package
WorkplaceID   int 
Workplace     Workplace
UserID        int 
User          User
Count         int
Note          string
```

# PartRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/part_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/part_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTime      time.Time 
OrderRecordID int
OrderRecord   OrderRecord
PartID        int 
Part          Part
WorkplaceID   int 
Workplace     Workplace
UserID        int 
User          User
Count         int
Note          string
```

# OrderRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/order_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/order_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTimeStart   time.Time 
DateTimeEnd     sql.NullTime
OrderID         int 
Order           Order
OperationID     int 
Operation       Operation
WorkplaceID     int 
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
```

# Operation

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_operation&auth=access_key 
```

``` go
Name    string 
OrderID int
Order   Order
Barcode string 
Note    string
```

# WorkplaceSection

- [x] MES
- [ ] WMS

``` go
Name string 
Note string
```

# Order

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_order&auth=access_key 
```

``` go
Name            string 
ProductID       sql.NullInt32
Product         Product
WorkplaceID     sql.NullInt32
Workplace       Workplace
Barcode         string 
DateTimeRequest sql.NullTime
Cavity          int
CountRequest    int
Note            string
```

# Product

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_product&auth=access_key 
```

``` go
Name             string 
Barcode          string 
Unit             sql.NullString
DownTimeDuration time.Duration
ProductTypeID    int 
Product#      ProductType
CountTypeID      int 
Count#        CountType
CycleTime        float64
Location         bool 
Image            []byte
ImageUrl         string
PurchasePrice    sql.NullFloat64
SalePrice        sql.NullFloat64
PartnerPrice     sql.NullFloat64
Fee              sql.NullFloat64
Note             string
```

# ProductType

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_product_type&auth=access_key 
```

``` go
Name string 
# string
Note string
```

# CountType

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_count_type&auth=access_key 
```

``` go
Name string 
# string
Note string
```

# Part

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_part&auth=access_key 
```

``` go
Name    string 
Barcode string
Note    string
```

# Workplace

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_workplace&auth=access_key 
```

``` go
Name                       string 
Code                       string
WorkplaceModeID            int
WorkplaceMode              WorkplaceMode
Voltage                    int
PowerFactor                float32
ConsumptionTypeID          int 
Consumption#            ConsumptionType
ConsumptionImpulsesPerWatt float32
Unit                       string
Note                       string
```

# ConsumptionType

- [x] MES
- [ ] WMS

``` go
Name string 
Code string
Note string
```

# WorkplacePort

- [x] MES
- [ ] WMS

``` go
Name         string 
DevicePortID int    
DevicePort   DevicePort
StateID      sql.NullInt32
State        State
WorkplaceID  int 
Workplace    Workplace
Color        sql.NullString
CounterOK    bool 
CounterNOK   bool 
HighValue    float32
LowValue     float32
Note         string
```

# WorkplaceMode

- [x] MES
- [ ] WMS

``` go
Name             string 
DowntimeDuration time.Duration
PoweroffDuration time.Duration
Note             string
```

# WorkplaceWorkshift

- [x] MES
- [ ] WMS

``` go
WorkplaceID int 
Workplace   Workplace
WorkshiftID int 
Workshift   Workshift
Note        string
```

# Workshift

- [x] MES
- [ ] WMS

``` go
Name           string 
WorkshiftStart int
WorkshiftEnd   int
Note           string
```

# User

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_user&auth=access_key 
```

``` go
FirstName  string 
SecondName string 
UserRoleID int
UserRole   UserRole
UserTypeID int
User#   UserType
Barcode    string 
Email      string 
Password   string
Phone      string
Pin        string 
Position   string
Rfid       string 
Locale     string
Note       string
```

# UserRole

- [x] MES
- [x] WMS

``` go
Name string 
Note string
```

# UserType

- [x] MES
- [x] WMS

``` html
http://ipaddress:port/save_user_type&auth=access_key 
```

``` go
Name string 
Note string
```

# Downtime

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_downtime&auth=access_key 
```

``` go
Name           string 
DowntimeTypeID int
Downtime#   DowntimeType
Barcode        string
Color          string
Note           string
```

# DowntimeType

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_downtime_type&auth=access_key 
```

``` go
Name string 
Note string
```

# Breakdown

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_breakdown&auth=access_key 
```

``` go
Name            string 
BreakdownTypeID int
Breakdown#   BreakdownType
Barcode         string
Color           string
Note            string
```

# BreakdownType

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_breakdown_type&auth=access_key 
```

``` go
Name string 
Note string
```

# Fault

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_fault&auth=access_key 
```

``` go
Name        string 
FaultTypeID int
Fault#   FaultType
Barcode     string
Note        string
```

# FaultType

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_fault_type&auth=access_key 
```

``` go
Name string 
Note string
```

# Package

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_package&auth=access_key 
```

``` go
Name          string 
ProductID     int    
Product       Product
PackageTypeID int
Package#   PackageType
Barcode       string
Note          string
```

# PackageType

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_package_type&auth=access_key 
```

``` go
Name  string 
Count int
Note  string
```

# DevicePortType

- [x] MES
- [ ] WMS

``` go
Name string 
Note string
```

# Setting

- [x] MES
- [x] WMS

``` go
Name    string 
Value   string
Enabled bool 
Note    string
```

# DeviceType

- [x] MES
- [ ] WMS

``` go
Name string 
Note string
```

# Device

- [x] MES
- [ ] WMS

``` go
Name         string 
DeviceTypeID int    
Device#   DeviceType
Activated    bool   
IpAddress    string 
MacAddress   string
Settings     string
TypeName     string
Note         string
```

# DeviceWorkplaceRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/device_workplace_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/device_workplace_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DeviceID    int 
Device      Device
WorkplaceID int 
Workplace   Workplace
Note        string
```

# DevicePort

- [x] MES
- [ ] WMS

``` go
Name             string 
DeviceID         int    
Device           Device
DevicePortTypeID int
DevicePort#   DevicePortType
PortNumber       int
PlcData#      string
PlcDataAddress   string
Settings         string
Unit             string
Virtual          bool 
Note             string
```

# DevicePortAnalogRecord

- [x] MES
- [ ] WMS

``` go
ID           int       
DateTime     time.Time
DevicePortID int       
DevicePort   DevicePort
Data         float32

```

# DevicePortSpecialRecord

- [x] MES
- [ ] WMS

``` go
ID           int       
DateTime     time.Time 
DevicePortID int       
DevicePort   DevicePort
Data         float32
```

# DevicePortDigitalRecord

- [x] MES
- [ ] WMS

``` go
ID int       
DateTime time.Time
DevicePortID int       
DevicePort DevicePort
Data int

```

# DevicePortSerialRecord

- [x] MES
- [ ] WMS

``` go
ID           int       
DateTime     time.Time 
DevicePortID int       
DevicePort   DevicePort
Data         float32
```

# Locale

- [x] MES
- [x] WMS

``` go
Name string 
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
```

# WebUserRecord

- [x] MES
- [x] WMS

``` go
UserEmail string    
WebPage   string    
DateTime  time.Time 
Note      string
```

# SummaryRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/summary_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/summary_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
WorkplaceID           int 
Workplace             Workplace
DateTime              time.Time 
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
```

# ShiftSummaryRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/shift_summary_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/shift_summary_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
WorkplaceID           int 
Workplace             Workplace
DateTime              time.Time 
WorkshiftID           int       
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
```

# Report

- [x] MES
- [x] WMS

``` go
Name string 
Url  string
Note string
```

# Bookmark

- [x] MES
- [x] WMS

``` go
Name string 
Url  string
Note string
```

# Layout

- [x] MES
- [ ] WMS

``` go
Name     string 
Image    []byte
ImageUrl string
Note     string
```

# WorkplaceSectionRecord

- [x] MES
- [ ] WMS

``` go
WorkplaceSectionID int 
WorkplaceSection   WorkplaceSection
WorkplaceID        int 
Workplace          Workplace
Note               string
```

# WebUserSettings

- [x] MES
- [x] WMS

``` go
Email string 
#  string 
Data  string
Note  string
```

# MaintenanceType

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_maintenance_type&auth=access_key 
```

``` go
Name                      string 
DayOfWeek                 sql.NullInt32
DayOfMonth                sql.NullInt32
TotalTimeFromLastRecord   sql.NullInt32
PowerOnTimeFromLastRecord sql.NullInt32
Note                      string
```

# Maintenance

- [x] MES
- [ ] WMS

``` html
http://ipaddress:port/save_maintenance&auth=access_key 
```

``` go
Name              string 
MaintenanceTypeID int
Maintenance#   MaintenanceType
Note              string
```

# MaintenanceWorkplaceRecord

- [x] MES
- [ ] WMS

``` go
MaintenanceID int 
Maintenance   Maintenance
WorkplaceID   int 
Workplace     Workplace
Note          string
```

# MaintenanceRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/maintenance_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/maintenance_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
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
```

# Checklist

- [x] MES
- [ ] WMS

``` go
Name               string 
#               string
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
User#           UserType
ProductID          sql.NullInt32
Product            Product
OrderID            sql.NullInt32
Order              Order
OperationID        sql.NullInt32
Operation          Operation
Note               string
```

# ChecklistRecord

- [x] MES
- [ ] WMS

``` html
json example: http://ipaddress:port/checklist_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/checklist_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTime       time.Time 
ChecklistID    int       
Checklist      Checklist
WorkplaceID    int 
Workplace      Workplace
UserId         int 
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
```

# Stock

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_stock&auth=access_key 
```

``` go
Name string 
Code string 
Note string
```

# StockStateRecord

- [ ] MES
- [x] WMS

``` html
json example: http://ipaddress:port/stock_state_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/stock_state_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
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
```

# BatchNumber

``` html
http://ipaddress:port/save_batch_number&auth=access_key 
```

- [ ] MES
- [x] WMS

``` go
ProductID          int 
Product            Product
Number             string 
ExpirationDuration time.Duration
Note               string
```

# SerialNumber

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_serial_number&auth=access_key 
```

``` go
ProductID int 
Product   Product
Number    string 
Note      string
```

# Company

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_company&auth=access_key 
```

``` go
Name          string 
Code          string 
Country       string 
Address       string
CompanyTypeID int
Company#   CompanyType
UserID        sql.NullInt32
User          User
Note          string
```

# CompanyType

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_company_type&auth=access_key 
```

``` go
Name string 
Note string
```

# StockLocation

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_stock_location&auth=access_key 
```

``` go
Name      string 
StockID   int    
Stock     Stock
MaxCount  sql.NullInt32
MaxVolume sql.NullFloat64
Note      string
```

# StockOrderRecord

- [ ] MES
- [x] WMS

``` html
http://ipaddress:port/save_stock_order_record&auth=access_key 
```

``` go
DateTimeStart   time.Time 
StockID         int       
Stock           Stock
CompanyID       int 
Company         Company
RecordTypeID    int 
Record#      RecordType
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
```

# StockRecord

- [ ] MES
- [x] WMS

``` html
json example: http://ipaddress:port/stock_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/stock_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
DateTime           time.Time 
UserID             int       
User               User
RecordTypeID       int 
Record#         RecordType
CompanyID          int 
Company            Company
StockInID          sql.NullInt32
StockIn            Stock
StockOutID         sql.NullInt32
StockOut           Stock
StockOrderRecordID sql.NullInt32
StockOrderRecord   StockOrderRecord
Closed             bool 
Note               string
```

# StockRecordItem

- [ ] MES
- [x] WMS

``` html
json example: http://ipaddress:port/stock_record_items?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/stock_record_items?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
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
```

# RecordType

- [ ] MES
- [x] WMS

``` go
Name string
Type string 
Note string
```

# ProductPackageRecord

- [ ] MES
- [x] WMS

``` html
json example: http://ipaddress:port/product_package_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=access_key
xml example:  http://ipaddress:port/product_package_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=xml&auth=access_key
```

``` go
PackageBarcode string 
SerialNumberID sql.NullInt32
SerialNumber   SerialNumber
BatchNumberID  sql.NullInt32
BatchNumber    BatchNumber
Note           string
```

# Holiday

- [x] MES
- [x] WMS

``` go
Date        time.Time 
Name        string
Holiday     bool 
HolidayName string
Note        string
```

---

## Example reading data

Reading downtime records

``` powershell
http://ipaddress:port/downtime_records?date_from=2023-04-02T10:40:00&date_to=2023-08-04T14:40:00&format=json&auth=FT7YCAYBAEDUY2LDMVXHGZIB76BAAAIDAECEIYLUMEAQUAABAFJAD74EAAAQCUYB76CAAAAABL7YGBIBAL7YMAAAAD72T74CAE7HWITTN5THI53BOJSSEORCIRUXG4DMMF4SAV3FMJJWK4TWNFRWKIRMEJRXK43UN5WWK4RCHIREEYLSORSWG2BMEBZS44RON4XCE7IBGEBCWKLCP7S6O5IRQNL2EEYHDKS556YXRSFZJFOLKOUFI2SEBPUEKMKPIGDBZUCBWWEEKEHWO4DEIDZ4AEYQEKRS42UASJOXRW3ORAKN73OVJOOND2TVVKS3QM5BCK3ODFZFRU7DZUIZZOVS2453WLW2J2UCHFL2IQAA====
```

Example result

``` json
{
  "Result": "OK, data limited to 1000 latest records in selected datetime",
  "Data": [
    {
      "ID": 1,
      "CreatedAt": "2023-05-13T15:04:49.136+02:00",
      "UpdatedAt": "2023-05-16T16:04:59.339+02:00",
      "DateTimeStart": "2023-04-27T17:05:06.554+02:00",
      "DateTimeEnd": {
        "Time": "2023-05-19T17:05:09.255+02:00",
        "Valid": true
      },
      "WorkplaceID": 1,
      "WorkplaceName": "Stroj",
      "DowntimeID": 1,
      "DowntimeName": "Nespecifikovaný prostoj",
      "OrderRecordID": {
        "Int32": 1,
        "Valid": true
      },
      "UserID": {
        "Int32": 1,
        "Valid": true
      },
      "UserName": "Admin Admin",
      "Note": "111"
    }
  ]
}
```

## Example creating data

Creating new stock location

``` powershell
$body = @{
"Name"="Main"
"StockId"=1
"MaxCount"=1
"MaxVolume"=13
"Note"=""

} | ConvertTo-Json

$header = @{
"Accept"="application/json"
"Content-Type"="application/json"
}

Invoke-RestMethod -Uri "http://localhost:93/save_stock_location&auth=FT7YCAYBAEDUY2LDMVXHGZIB76BAAAIDAECEIYLUMEAQUAABAFJAD74EAAAQCUYB76CAAAAABL7YGBIBAL7YMAAAAD72T74CAE7HWITTN5THI53BOJSSEORCIRUXG4DMMF4SAV3FMJJWK4TWNFRWKIRMEJRXK43UN5WWK4RCHIREEYLSORSWG2BMEBZS44RON4XCE7IBGEBCWKLCP7S6O5IRQNL2EEYHDKS556YXRSFZJFOLKOUFI2SEBPUEKMKPIGDBZUCBWWEEKEHWO4DEIDZ4AEYQEKRS42UASJOXRW3ORAKN73OVJOOND2TVVKS3QM5BCK3ODFZFRU7DZUIZZOVS2453WLW2J2UCHFL2IQAA====" -Method 'Post' -Body $body -Headers $header | ConvertTo-HTML
```

Example result

``` json
{
  "Result": "OK",
  "Type": "Udated",
  "CreatedAt": "2023-05-13T15:04:49.136+02:00",
}
```