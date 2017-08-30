create table taobao_csv_record_item (
      "id"                  UUID PRIMARY KEY        NOT NULL        DEFAULT gen_random_uuid(),

      goods_id              text                    NOT NULL,
      taobao_csv_record_id  text                    NOT NULL ,

      create_at timestamptz not null default now(),
      update_at timestamptz not null default now()
);
CREATE UNIQUE INDEX IF NOT EXISTS taobao_csv_record_item_goods_and_csv ON  taobao_csv_record_item(goods_id,taobao_csv_record_id);
