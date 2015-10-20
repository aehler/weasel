create table weasel_storage.files (
    id               bigserial primary key,
    entity           varchar(256) not null,
    entity_id        bigint not null,
    name             varchar(500) not null,
    file_path        varchar(1000) not null,
    organization_id  bigint not null,
    owner_user       jsonb not null,
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    items_meta       jsonb not null default '{}',
    version_id       bigint not null default 0,
    bucket           bigint not null default 0,
    md5_sum          char(32) not null default '',
    content_type     varchar(1000) not null default ''
)with(fillfactor = 90);

create index idx_fctype on weasel_storage.files(LOWER(content_type));
create index idx_f_org_id on weasel_storage.files(organization_id);
create index idx_f_ent_ent_id on weasel_storage.files(entity, entity_id);
create index idx_f_version on weasel_storage.files(version_id);