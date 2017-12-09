create table taobao_csv_record (
      "id"              UUID PRIMARY KEY        NOT NULL        DEFAULT gen_random_uuid(),
      user_id               text                    NOT NULL,
      discount              int                     NOT NULL           default 0,
      supplemental_fee      int                     NOT NULL           default 0,
      province              text                    NOT NULL           default '',
      city                  text                    NOT NULL           default '',
      express_template      text                    NOT NULL           default '',
      pingyou_fee           int                     NOT NULL           default 0,
      express_fee           int                     NOT NULL           default 0,
      ems_fee               int                     NOT NULL           default 0,
      reduce_stock_style    text                    NOT NULL           default '1',    -- 1拍下减库存  2付款减库存

      total_num             int                     NOT NULL           default 0,
      success_num           int                     NOT NULL           default 0,
      product_title         text                    NOT NULL           default '',
      product_describe      text                    NOT NULL           default '',
      filepath              text                    NOT NULL           default '',

      search_isbn           text                    NOT NULL           default '',
      search_title          text                    NOT NULL           default '',
      search_publisher      text                    NOT NULL           default '',
      search_compare        text                    NOT NULL           default '',
      search_stock          int                     NOT NULL           default 0,
      search_author         text                    NOT NULL           default '',

      status                smallint                NOT NULL           default 1,    -- 1 进行中  2 完成   3失败
      file_url              text                    NOT NULL           default '',
      summary               text                    NOT NULL           default '',
      complete_at           timestamptz,
      create_at             timestamptz             not null            default now(),
      update_at             timestamptz             not null            default now()
);
