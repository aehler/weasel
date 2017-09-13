package articles

type Topic struct {
	ID uint `db:"topic_id"`
	Title string `db:"title"`
	Description string `db:"description"`
	ArticlesCount uint `db:"active_articles_count"`
	Image string `db:"image"`
}
