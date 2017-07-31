create table location (
    "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),

    warehouse text not null default '0',
    shelf text not null default '0',
    floor text not null default '0',
    user_id text not null,

    create_at timestamptz not null default now(),
    update_at timestamptz not null default now()
);
