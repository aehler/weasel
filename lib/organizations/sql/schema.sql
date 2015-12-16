create table weasel_main.organizations(
    id               bigserial primary key,
    user_id          bigint not null DEFAULT 0,
    organization_id  bigint not null DEFAULT 0 references weasel_auth.organizations (organization_id),
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    meta_info        jsonb not null default '{}',
    inn              VARCHAR(12) not null,
    kpp              VARCHAR(12) not null
)with(fillfactor = 90);

create index idx_org_organization_id on weasel_main.organizations(organization_id);
create UNIQUE index idx_org_inn_kpp on weasel_main.organizations(inn, kpp);