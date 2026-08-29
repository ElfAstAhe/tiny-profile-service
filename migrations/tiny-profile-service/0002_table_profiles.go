package tiny_profile_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0002(ctx context.Context, db *sql.DB) error {
	if err := createTableProfiles(ctx, db); err != nil {
		return err
	}
	if err := createIndexProfilesByID(ctx, db); err != nil {
		return err
	}
	if err := createIndexProfilesByUserID(ctx, db); err != nil {
		return err
	}

	return createIndexProfileByPerson(ctx, db)
}

func createTableProfiles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateTableProfiles)
	if err != nil {
		return errs.NewDBMigrationError("create table profiles failed", err)
	}

	return nil
}

func createIndexProfilesByID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexProfilesByID)
	if err != nil {
		return errs.NewDBMigrationError("create index profiles by id failed", err)
	}

	return nil
}

func createIndexProfilesByUserID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexProfilesByUserID)
	if err != nil {
		return errs.NewDBMigrationError("create index profiles by user id failed", err)
	}

	return nil
}

func createIndexProfileByPerson(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexProfilesByPersonID)
	if err != nil {
		return errs.NewDBMigrationError("create index profiles by person id failed", err)
	}

	return nil
}

func down0002(ctx context.Context, db *sql.DB) error {
	if err := dropIndexProfilesByID(ctx, db); err != nil {
		return err
	}
	if err := dropIndexProfilesByUserID(ctx, db); err != nil {
		return err
	}
	if err := dropIndexProfilesByPersonID(ctx, db); err != nil {
		return err
	}

	return dropTableProfiles(ctx, db)
}

func dropTableProfiles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropTableProfiles)
	if err != nil {
		return errs.NewDBMigrationError("drop table profiles failed", err)
	}

	return nil
}

func dropIndexProfilesByID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexProfilesByID)
	if err != nil {
		return errs.NewDBMigrationError("drop index profiles by id failed", err)
	}

	return nil
}

func dropIndexProfilesByUserID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexProfilesByUserID)
	if err != nil {
		return errs.NewDBMigrationError("drop index profiles by user id failed", err)
	}

	return nil
}

func dropIndexProfilesByPersonID(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexProfilesByPersonID)
	if err != nil {
		return errs.NewDBMigrationError("drop index profiles by person id failed", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0002, down0002)
}
