drop table IF EXISTS "segments";
create TABLE "public"."segments"
(
    "id"          SERIAL,
    "slug"        character varying(32) NOT NULL,
    CONSTRAINT "segments_ids" PRIMARY KEY ("id")
) with (oids = false);

INSERT INTO public.segments (id, slug)
VALUES (0, 'AVITO_VOICE_MESSAGES');
INSERT INTO public.segments (id, slug)
VALUES (1, 'AVITO_PERFORMANCE_VAS');
INSERT INTO public.segments (id, slug)
VALUES (2, 'AVITO_DISCOUNT_30');
ALTER SEQUENCE segments_id_seq OWNED BY public.segments.id RESTART 3;

drop table IF EXISTS "users";
create TABLE "public"."users"
(
    "id"          SERIAL,
    "user_id"     integer NOT NULL UNIQUE ,
    CONSTRAINT "users_ids" PRIMARY KEY ("id")
) with (oids = false);
alter table "public"."users"
    add constraint "uq_user_id"
        UNIQUE ("user_id");

drop table IF EXISTS "records";
create TABLE "public"."records"
(
    "user_id"        integer NOT NULL,
    "segment_id"     integer NOT NULL,
    CONSTRAINT "pk_ids" PRIMARY KEY ("user_id", "segment_id")
) with (oids = false);
alter table "public"."records"
    add constraint "fk_users_id"
        foreign key ("user_id")
            references "public"."users" ("id") on update restrict on delete cascade;
alter table "public"."records"
    add constraint "fk_segments_id"
        foreign key ("segment_id")
            references "public"."segments" ("id") on update restrict on delete cascade;