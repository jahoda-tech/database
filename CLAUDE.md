# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library package (`github.com/jahoda-tech/database`) that defines GORM model structs for a Manufacturing Execution System (MES). It contains no binaries, no CLI, and no tests — it is imported by other services as a shared schema definition.

Target database is **PostgreSQL** (consumers use `gorm.io/driver/postgres`).

## Development Workflow

There is no build step, test suite, or CI pipeline. The workflow is:
1. Edit struct definitions in `database.go`
2. Run `go vet ./...` to check for errors
3. Commit, tag a new version, and push
4. Consuming services update with `go get github.com/jahoda-tech/database@latest`

Schema changes have wide blast radius across consuming services. Type-widening, nullability tightening, or column renames are breaking changes — bump major version and coordinate consumer deploys.

## Architecture

**Single file:** All ~70 GORM struct definitions live in `database.go`.

**Two struct patterns:**

- **Master/config entities** embed `gorm.Model` (provides ID, CreatedAt, UpdatedAt, DeletedAt with soft-delete):
  ```go
  type Workplace struct {
      gorm.Model
      Name string `gorm:"uniqueIndex:unique_workplace"`
      // ...
  }
  ```

- **High-volume time-series records** use a manual `ID int` primary key with no soft-delete overhead:
  ```go
  type DevicePortAnalogRecord struct {
      ID       int       `gorm:"primaryKey"`
      DateTime time.Time `gorm:"uniqueIndex:...;index:idx_analog_record_brin,type:brin"`
      // ...
  }
  ```

  Most `*Record` tables currently still embed `gorm.Model` for historical reasons; only `DevicePort{Analog,Digital,Special,Serial}Record` follow the manual-ID pattern.

**Recurring field conventions:**
- `Note string` — human annotation, on most structs
- `Data datatypes.JSON` with `gorm:"type:jsonb;index:,type:gin"` — flexible JSONB blob with GIN index for `@>` queries
- `ExternalId sql.NullInt64` with `gorm:"index"` — external system integration key
- FK pairs: `WorkplaceID int` + `Workplace Workplace` (GORM preload pattern)
- Composite unique indexes: `gorm:"uniqueIndex:name,priority:N"`
- Nullable fields use `sql.NullInt64`, `sql.NullFloat64`, `sql.NullTime`, `sql.NullString`

**Index conventions:**
- Every FK column carries `gorm:"index"` (GORM does NOT auto-index FKs). Composite unique only covers the leading column for standalone lookups; add explicit `index` for non-leading FKs.
- Time-series tables get a BRIN index on the primary DateTime/DateTimeStart column: `index:idx_<table>_brin,type:brin`. BRIN is ~kilobytes vs B-tree gigabytes for append-only data.
- "Open record" lookup tables (those with `DateTimeEnd sql.NullTime`) get a partial index on `WorkplaceID where date_time_end IS NULL` for fast current-state queries. Pattern: `index:idx_<table>_open,where:date_time_end IS NULL`.
- `DateTimeEnd` columns on closed-record lookups get `gorm:"index"`.

**One custom TableName override:** `Company` maps to `"companies"` (overrides GORM's default `companys`).

## Dependencies

- `gorm.io/gorm` — ORM framework
- `gorm.io/datatypes` — provides `datatypes.JSON` field type (rendered as PG `jsonb` via column tag)
- PostgreSQL is the target database (consumers import `gorm.io/driver/postgres`)
