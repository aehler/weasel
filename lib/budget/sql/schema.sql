create table weasel_main.budget_operations(
    id               bigserial primary key,
    user_id          bigint not null DEFAULT 0,
    organization_id  bigint not null DEFAULT 0,
    sum              NUMERIC not null DEFAULT 0,
    date_op          timestamp not null default current_timestamp,
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    dims_meta        jsonb not null default '{}',
    user_meta        jsonb not null default '{}',
    tags             varchar[] not null default '{}'
)with(fillfactor = 90);

create index idx_organization_id on weasel_main.budget_operations(organization_id);
create index idx_date on  weasel_main.budget_operations(date_op);
create index idx_tags on weasel_main.budget_operations using GIN(tags);