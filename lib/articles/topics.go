package articles

import "weasel/app/registry"

func ListTopicsByLang(lang string) ([]Topic, error) {

	t := []Topic{}

	err := registry.Registry.Connect.Select(&t, `select topic_id, title, description, active_articles_count, image
		from weasel_articles.topics where lang = $1 and is_deleted=false`, lang)

	return t, err
}