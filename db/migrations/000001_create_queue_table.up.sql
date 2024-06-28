create schema if not exists queues;

create table if not exists queues.mail_queue
(
    id             uuid        default gen_random_uuid() not null primary key,
    created_at     timestamptz default current_timestamp not null,
    started_at     timestamptz                           null,
    locked_until   timestamptz                           null,
    processed_at   timestamptz                           null,
    consumed_count integer     default 0                 not null,
    error_detail   text                                  null,
    payload        jsonb                                 not null,
    metadata       jsonb                                 not null
    );

create index if not exists mail_queue_created_at_idx on queues.mail_queue (created_at);
create index if not exists mail_queue_processed_at_null_idx on queues.mail_queue (processed_at) WHERE (processed_at IS NULL);

