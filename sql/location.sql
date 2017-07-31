create table location (
    "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),

    stock int not null default 0,
    warehouse text not null default '0',
    shelf text not null default '0',
    fllor text not null default '0',
    user_id text not null,
    goods_id uuid not null,

    create_at timestamptz not null default now(),
    update_at timestamptz not null default now()
);
