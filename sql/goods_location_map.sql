create table goods_location_map (
     "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),

     location_id text not null,
     goods_id text not null,
     stock int not null default 0,
     user_id text null,

     create_at timestamptz not null default now(),
     update_at timestamptz not null default now()
);
