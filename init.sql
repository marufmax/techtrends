CREATE TABLE IF NOT EXISTS job_counts
(
    id            serial constraint job_counts_pk primary key,
    category_id   int not null,
    count         int not null,
    vacancy       int,
    category_name varchar(512),
    category_type varchar(255),
    is_active     bool default true,
    crawl_date    date default current_date,
    created_at    timestamp(0) default now(),
    updated_at    timestamp(0) default now()
);

CREATE TABLE IF NOT EXISTS categories
(
    id         serial constraint categories_pk primary key,
    name       varchar(255) not null,
    type       varchar(100),
    created_at timestamp(0) default now(),
    updated_at    timestamp(0) default now()
);

create unique index if not exists job_counts_category_id_crawl_date_uindex
    on job_counts (category_id desc, crawl_date desc);




ALTER TABLE categories OWNER TO techtrend;
