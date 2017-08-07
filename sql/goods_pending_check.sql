create table goods_pending_check (
      "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),
      isbn text not null,
      num int not null default 0,
      user_id text not null,
      warehouse text not null default '*',
      shelf text not null default '*',
      floor text not null default '*',
      create_at timestamptz not null default now(),
      update_at timestamptz not null default now()
);
