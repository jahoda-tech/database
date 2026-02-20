# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library package (`github.com/jahoda-tech/database`) that defines GORM model structs for a Manufacturing Execution System (MES). It contains no binaries, no CLI, and no tests — it is imported by other services as a shared schema definition.

## Development Workflow

There is no build step, test suite, or CI pipeline. The workflow is:
1. Edit struct definitions in `database.go`
2. Run `go vet ./...` to check for errors
3. Commit, tag a new version, and push
4. Consuming services update with `go get github.com/jahoda-tech/database@latest`

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
      DateTime time.Time `gorm:"uniqueIndex:..."`
      // ...
  }
  ```

**Recurring field conventions:**
- `Note string` — human annotation, on most structs
- `Data datatypes.JSON` — flexible JSON blob on most master entities
- `ExternalId sql.NullInt32` with `gorm:"index"` — external system integration key
- FK pairs: `WorkplaceID int` + `Workplace Workplace` (GORM preload pattern)
- Composite unique indexes: `gorm:"uniqueIndex:name,priority:N"`
- Nullable fields use `sql.NullInt32`, `sql.NullFloat64`, `sql.NullTime`

**One custom TableName override:** `Company` maps to `"companies"` (overrides GORM's default `companys`).

## Dependencies

- `gorm.io/gorm` — ORM framework
- `gorm.io/datatypes` — provides `datatypes.JSON` field type
- MySQL is the target database (indirect dep on `gorm.io/driver/mysql`)
