create table users_taobao_setting (
      "id"                  UUID PRIMARY KEY        NOT NULL        DEFAULT gen_random_uuid(),
      user_id               text                    NOT NULL,

      discount              int                     NOT NULL           default 0,
      supplemental_fee      int                     NOT NULL           default 0,
      province              text                    NOT NULL           default '',
      city                  text                    NOT NULL           default '',
      express_template      text                    NOT NULL           default '',
      pingyou_fee           int                     NOT NULL           default 0,
      express_fee           int                     NOT NULL           default 0,
      ems_fee               int                     NOT NULL           default 0,
      reduce_stock_style    text                    NOT NULL           default '1',
      product_title         text                    NOT NULL           default '',
      product_describe      text                    NOT NULL           default '',
      create_at             timestamptz             not null default now(),
      update_at             timestamptz             not null default now()
);


CREATE UNIQUE INDEX IF NOT EXISTS users_taobao_setting_user ON  users_taobao_setting(user_id);
