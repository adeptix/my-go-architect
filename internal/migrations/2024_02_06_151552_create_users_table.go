package migrations

import (
	"gitlab.com/bonch.dev/go-lib/migrator/blueprint"
	"gitlab.com/bonch.dev/go-lib/migrator/builder"
	"gitlab.com/bonch.dev/go-lib/migrator/column"
	"gitlab.com/bonch.dev/go-lib/migrator/migration"
	"gitlab.com/bonch.dev/go-lib/migrator/migrator"
)

func init() {
	migrator.AddMigrations(&M20240206151552CreateUsersTable{})
}

type M20240206151552CreateUsersTable struct {
	migration.BaseMigration
}

func (m M20240206151552CreateUsersTable) Up(b builder.Builder) {
	b.Create("users", func(blueprint *blueprint.Blueprint) {
		blueprint.AddColumn(column.UUID, "id").Primary()
		blueprint.AddColumn(column.TimestampTZ, "created_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "updated_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "deleted_at").Nullable().Index()
		blueprint.AddColumn(column.String, "name")
		blueprint.AddColumn(column.Integer, "age")
	})
}

func (m M20240206151552CreateUsersTable) Down(b builder.Builder) {

}

func (m M20240206151552CreateUsersTable) Name() string {
	return "M20240206151552CreateUsersTable"
}
