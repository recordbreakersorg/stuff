create table if not exists files (
    id bigserial primary key,
    mime text not null,
    file_size bigint not null
);
