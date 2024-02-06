package migrations

import (
	"gitlab.com/bonch.dev/go-lib/migrator/blueprint"
	"gitlab.com/bonch.dev/go-lib/migrator/builder"
	"gitlab.com/bonch.dev/go-lib/migrator/column"
	"gitlab.com/bonch.dev/go-lib/migrator/migration"
	"gitlab.com/bonch.dev/go-lib/migrator/migrator"
)

func init() {
	migrator.AddMigrations(&M20240206151742CreateLikesTable{})
}

type M20240206151742CreateLikesTable struct {
	migration.BaseMigration
}

func (m M20240206151742CreateLikesTable) Up(b builder.Builder) {
	b.Create("likes", func(blueprint *blueprint.Blueprint) {
		blueprint.AddColumn(column.UUID, "id").Primary()
		blueprint.AddColumn(column.TimestampTZ, "created_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "updated_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "deleted_at").Nullable().Index()
		blueprint.AddColumn(column.Integer, "value")
		blueprint.AddColumn(column.UUID, "user_id")
		blueprint.AddColumn(column.UUID, "post_id").Nullable()
		blueprint.AddColumn(column.UUID, "comment_id").Nullable()

		blueprint.AddForeign("user_id").RefOn("users").RefColumns("id")
		blueprint.AddForeign("post_id").RefOn("posts").RefColumns("id")
		blueprint.AddForeign("comment_id").RefOn("comments").RefColumns("id")
	})
}

func (m M20240206151742CreateLikesTable) Down(b builder.Builder) {

}

func (m M20240206151742CreateLikesTable) Name() string {
	return "M20240206151742CreateLikesTable"
}
