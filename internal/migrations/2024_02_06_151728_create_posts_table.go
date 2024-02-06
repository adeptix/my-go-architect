package migrations

import (
	"gitlab.com/bonch.dev/go-lib/migrator/blueprint"
	"gitlab.com/bonch.dev/go-lib/migrator/builder"
	"gitlab.com/bonch.dev/go-lib/migrator/column"
	"gitlab.com/bonch.dev/go-lib/migrator/migration"
	"gitlab.com/bonch.dev/go-lib/migrator/migrator"
)

func init() {
	migrator.AddMigrations(&M20240206151728CreatePostsTable{})
}

type M20240206151728CreatePostsTable struct {
	migration.BaseMigration
}

func (m M20240206151728CreatePostsTable) Up(b builder.Builder) {
	b.Create("posts", func(blueprint *blueprint.Blueprint) {
		blueprint.AddColumn(column.UUID, "id").Primary()
		blueprint.AddColumn(column.TimestampTZ, "created_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "updated_at").Nullable().Default("now()")
		blueprint.AddColumn(column.TimestampTZ, "deleted_at").Nullable().Index()
		blueprint.AddColumn(column.String, "title")
		blueprint.AddColumn(column.Text, "text")
		blueprint.AddColumn(column.UUID, "user_id")

		blueprint.AddForeign("user_id").RefOn("users").RefColumns("id")
	})
}

func (m M20240206151728CreatePostsTable) Down(b builder.Builder) {

}

func (m M20240206151728CreatePostsTable) Name() string {
	return "M20240206151728CreatePostsTable"
}
