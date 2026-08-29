package tiny_profile_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0001(ctx context.Context, db *sql.DB) error {
	if err := createTablePersons(ctx, db); err != nil {
		return err
	}
	if err := createIndexPersonsByID(ctx, db); err != nil {
		return err
	}

	return createIndexPersonsByExternalID(ctx, db)
}

func createTablePersons(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateTablePersons)
	if err != nil {
		return errs.NewDBMigrationError("create table persons failed", err)
	}

	return nil
}

func createIndexPersonsByID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexPersonsByID)
	if err != nil {
		return errs.NewDBMigrationError("create index persons id failed", err)
	}

	return nil
}

func createIndexPersonsByExternalID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexPersonsByExternalID)
	if err != nil {
		return errs.NewDBMigrationError("create index persons external id failed", err)
	}

	return nil
}

func down0001(ctx context.Context, db *sql.DB) error {
	if err := dropIndexPersonsByExternalID(ctx, db); err != nil {
		return err
	}
	if err := dropIndexPersonsByID(ctx, db); err != nil {
		return err
	}

	return dropTablePersons(ctx, db)
}

func dropTablePersons(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropTablePersons)
	if err != nil {
		return errs.NewDBMigrationError("drop table persons failed", err)
	}

	return nil
}

func dropIndexPersonsByID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexPersonsByID)
	if err != nil {
		return errs.NewDBMigrationError("drop index persons id failed", err)
	}

	return nil
}

func dropIndexPersonsByExternalID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexPersonsByExternalID)
	if err != nil {
		return errs.NewDBMigrationError("drop index persons external id failed", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0001, down0001)
}
