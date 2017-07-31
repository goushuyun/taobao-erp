create table goods (
      "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),

      book_id text not null,
      status int not null default 0,
      user_id text not null,

      create_at timestamptz not null default now(),
      update_at timestamptz not null default now()
);
