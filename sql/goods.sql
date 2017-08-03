create table goods (
      "id"              UUID PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),

      /* book info */
      book_id text not null,
      isbn text not null,
      book_no text default '',

      /* goods info */
      status int not null default 0,
      user_id text not null,
      remark text default '',
      stock int not null default 0,

      create_at timestamptz not null default now(),
      update_at timestamptz not null default now()
);
