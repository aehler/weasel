create table weasel_classifiers.references(
    id               bigserial primary key,
    name             varchar(500) not null,
    alias            varchar(500) not null,
    organization_id  bigint not null,
    blocked          boolean not null default false,
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    items_meta       jsonb not null default '{}'
)with(fillfactor = 90);

create index idx_nm on weasel_classifiers.references(LOWER(alias));
create index idx_nm on weasel_classifiers.references(LOWER(name));
create index idx_ref_org_id on weasel_classifiers.references(organization_id);

create table weasel_classifiers.counter(
  reference_id bigint references weasel_classifiers."references" (id) primary key,
  total bigint not null default 0
) with(fillfactor=50);


create table weasel_classifiers.items(
    id               bigserial primary key,
    reference_id     bigint references weasel_classifiers."references" (id),
    name             varchar(500) not null,
    alias            varchar(500) not null,
    ver              integer not null default 0,
    is_group         boolean not null default false,
    parents          varchar(500) not null default '',
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    pid              bigint not null default 0,
    fields           jsonb not null default '{}'
)with(fillfactor = 90);

create index idx_inm on weasel_classifiers.items(LOWER(name));
create index idx_iparents on weasel_classifiers.items(LOWER(parents));
create index idx_irefid on weasel_classifiers.items(reference_id);
create index idx_iver on weasel_classifiers.items(ver, reference_id);
create index idx_ipid on weasel_classifiers.items(pid);