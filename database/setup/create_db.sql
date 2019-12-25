
CREATE SCHEMA if not exists adverts_schema AUTHORIZATION postgres;

CREATE EXTENSION IF not exists citext with SCHEMA adverts_schema;

create table adverts_schema.adverts
(
    id           bigserial             not null
        constraint adverts_pk
            primary key,
    title        adverts_schema.citext not null,
    description  adverts_schema.citext not null,
    pictures     text[]                not null,
    main_picture text                  not null,
    price        integer               not null,
    date_created timestamptz           not null
);

alter table adverts_schema.adverts
    owner to postgres;

create unique index adverts_id_uindex
    on adverts_schema.adverts (id);

create index price_id_idx
    on adverts_schema.adverts (price, id);

create index date_id_idx
    on adverts_schema.adverts (date_created, id);

create type adverts_schema.shortadvertdata as
    (
    id bigint,
    title adverts_schema.citext,
    main_picture text,
    price integer
    );

alter type adverts_schema.shortadvertdata owner to postgres;

create or replace function adverts_schema.getadvertslistorderedbyprice(page integer, ads_per_page integer,
                                                                       order_by_price adverts_schema.citext) returns SETOF adverts_schema.shortadvertdata
    language plpgsql
as
$$
DECLARE
    page_offset bigint = ads_per_page * (page - 1);
BEGIN
    if order_by_price = 'ASC' OR order_by_price = 'asc' THEN
        return QUERY SELECT b.id, title, main_picture, b.price
                     FROM (adverts_schema.adverts
                              JOIN (SELECT price, id
                                    FROM adverts_schema.adverts
                                    ORDER BY price ASC
                                    LIMIT ads_per_page
                                    OFFSET
                                    page_offset) as b
                                   ON b.id = adverts_schema.adverts.id)
                     ORDER BY b.price ASC;
    ELSIF order_by_price = 'DESC' OR order_by_price = 'desc' THEN
        return QUERY SELECT b.id, title, main_picture, b.price
                     FROM (adverts_schema.adverts
                              JOIN (SELECT price, id
                                    FROM adverts_schema.adverts
                                    ORDER BY price DESC
                                    LIMIT ads_per_page
                                    OFFSET
                                    page_offset) as b
                                   ON b.id = adverts_schema.adverts.id)
                     ORDER BY b.price DESC;

    end if;
END

$$;

alter function adverts_schema.getadvertslistorderedbyprice(integer, integer, adverts_schema.citext) owner to postgres;

create or replace function adverts_schema.getadvertslistorderedbydate(page integer, ads_per_page integer,
                                                                      order_by_date adverts_schema.citext) returns SETOF adverts_schema.shortadvertdata
    language plpgsql
as
$$
DECLARE
    page_offset bigint = ads_per_page * (page - 1);

BEGIN
    if order_by_date = 'ASC' OR order_by_date = 'asc' THEN
        RETURN QUERY SELECT b.id, title, main_picture, price
                     FROM (adverts_schema.adverts
                              JOIN (SELECT date_created, id
                                    FROM adverts_schema.adverts
                                    ORDER BY date_created ASC
                                    LIMIT ads_per_page
                                    OFFSET
                                    page_offset) as b
                                   ON b.id = adverts_schema.adverts.id)
                     ORDER BY b.date_created ASC;
    ELSIF order_by_date = 'DESC' or order_by_date = 'desc' THEN
        RETURN QUERY SELECT b.id, title, main_picture, price
                     FROM (adverts_schema.adverts
                              JOIN (SELECT date_created, id
                                    FROM adverts_schema.adverts
                                    ORDER BY date_created DESC
                                    LIMIT ads_per_page
                                    OFFSET
                                    page_offset) as b
                                   ON b.id = adverts_schema.adverts.id )
                     ORDER BY b.date_created DESC;
    end if;

END

$$;

alter function adverts_schema.getadvertslistorderedbydate(integer, integer, adverts_schema.citext) owner to postgres;

create or replace function adverts_schema.createadvert(title_arg adverts_schema.citext,
                                                       description_arg adverts_schema.citext, pictures_str_arg text[],
                                                       price_arg integer) returns bigint
    language plpgsql
as
$$
DECLARE
    inserted_row_id bigint;
BEGIN

    INSERT into adverts_schema.adverts (title, description, main_picture, pictures, price, date_created)
    VALUES (title_arg, description_arg, pictures_str_arg[1], pictures_str_arg, price_arg,NOW())
    RETURNING id INTO inserted_row_id;
    return inserted_row_id;
END;
$$;

alter function adverts_schema.createadvert(adverts_schema.citext, adverts_schema.citext, text[], integer) owner to postgres;

create or replace function adverts_schema.getadvert(id_arg bigint, fields text[]) returns adverts_schema.adverts
    language plpgsql
as
$$
DECLARE
    descrtiption_col_name text = 'description';
    pictures_col_name     text = 'pictures';
    result_record         adverts_schema.adverts;
    empty_array           text[];
BEGIN

    SELECT * INTO result_record from adverts_schema.adverts where id = id_arg;

    IF NOT (descrtiption_col_name = ANY (fields)) THEN
        result_record.description = '';
    end if;

    IF NOT (pictures_col_name = ANY (fields)) THEN
        result_record.pictures = empty_array;
    end if;

    return result_record;
END

$$;

alter function adverts_schema.getadvert(bigint, text[]) owner to postgres;

