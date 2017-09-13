create table weasel_articles.articles(
	id bigserial not null primary key,
	title varchar(255) not null default '',
	lang char(2) not null,
	main_image varchar(1024) not null default '',
	author_id bigint not null references weasel_auth.users(user_id),
	topics bigint[] not null,
	"body" text not null,
	is_deleted boolean not null default false,
	is_moderated boolean not null default false,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	viewed_count bigint not null default 0,
	hashtags tsvector not null default '{}',
	lang_links jsonb not null
);

create table weasel_articles.topics(
	topic_id bigserial not null primary key,
	title varchar(255) not null,
	lang char(2) not null,
	active_articles_count bigint not null default 0,
	total_articles_count bigint not null default 0,
	is_deleted boolean not null default false
);

create index idx_weasel_articles_articles_topics on weasel_articles.articles (topics);
create index idx_weasel_articles_articles_hashtags on weasel_articles.articles using gin (hashtags);
create index idx_weasel_articles_articles_active on weasel_articles.articles (is_deleted, is_moderated);
create index idx_weasel_articles_topics_active on weasel_articles.topics (is_deleted);
create index idx_weasel_articles_topics_lang on weasel_articles.topics (lang);