package main

import (
	"fmt"
	"sqlgen/query"
)

func main() {

	tree := query.
		Select("players.x", "players.y").
		From(query.
			Select("stats.a", "stats.b").
			From("stats").
			As("most_recent_stats").
			Where("stats.a = 10").
			And("stats.b = 11").
			Or("stats.c = 13")).
		As("group").
		Where("players.x > 10").
		And("players.b > 11").
		And("player.c > 12").
		Or("most_recent_stats.a = 5")

	sqlString := tree.ToSql()
	fmt.Println(sqlString)
	// Prints the following:
	// SELECT
	//   players.x,
	//   players.y
	// FROM (
	//     SELECT
	//       stats.a,
	//       stats.b
	//     FROM
	//       stats AS most_recent_stats
	//     WHERE
	//       stats.a = 10
	//       AND stats.b = 11
	//       OR stats.c = 13
	// ) AS group
	// WHERE
	//   players.x > 10
	//   AND players.b > 11
	//   AND player.c > 12
	//   OR most_recent_stats.a = 5
}
