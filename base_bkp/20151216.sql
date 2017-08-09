--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: weasel_auth; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA weasel_auth;


ALTER SCHEMA weasel_auth OWNER TO postgres;

--
-- Name: weasel_classifiers; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA weasel_classifiers;


ALTER SCHEMA weasel_classifiers OWNER TO postgres;

--
-- Name: weasel_main; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA weasel_main;


ALTER SCHEMA weasel_main OWNER TO postgres;

--
-- Name: weasel_storage; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA weasel_storage;


ALTER SCHEMA weasel_storage OWNER TO postgres;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = weasel_auth, pg_catalog;

--
-- Name: add_user(bigint, character varying, character varying, character varying, character varying, character varying, character varying, character varying, boolean, integer); Type: FUNCTION; Schema: weasel_auth; Owner: weasel
--

CREATE FUNCTION add_user(_organization_id bigint, _user_firstname character varying, _user_lastname character varying, _user_middlename character varying, _user_job_title character varying, _user_phone character varying, _user_email character varying, _user_password character varying, _is_admin boolean, _timezone_id integer, OUT _user_id bigint) RETURNS bigint
    LANGUAGE plpgsql
    AS $$

    declare
        _token_id varchar;
    begin

        PERFORM 1 FROM weasel_auth.users WHERE user_login = lower(_user_email) and user_password = _user_password;

        IF FOUND THEN
            RAISE EXCEPTION 'USER_EXISTS';
        END IF;

        PERFORM 1 FROM weasel_auth.users WHERE lower(user_email) = lower(_user_email);

        IF FOUND THEN
            RAISE EXCEPTION 'EMAIL_EXISTS';
        END IF;

        INSERT INTO weasel_auth.users (
            organization_id,
            user_firstname,
            user_lastname,
            user_middlename,
            user_job_title,
            user_phone,
            user_login,
            user_email,
            user_password,
            is_admin,
            is_active,
            is_deleted,
            timezone_id
        ) VALUES (
            _organization_id,
            _user_firstname,
            _user_lastname,
            _user_middlename,
            _user_job_title,
            _user_phone,
            _user_email,
            _user_email,
            _user_password,
            _is_admin,
            true,
            false,
            _timezone_id
        ) RETURNING weasel_auth.users.user_id INTO _user_id;

    end;
$$;


ALTER FUNCTION weasel_auth.add_user(_organization_id bigint, _user_firstname character varying, _user_lastname character varying, _user_middlename character varying, _user_job_title character varying, _user_phone character varying, _user_email character varying, _user_password character varying, _is_admin boolean, _timezone_id integer, OUT _user_id bigint) OWNER TO weasel;

--
-- Name: add_user(bigint, character varying, character varying, character varying, character varying, character varying, character varying, character varying, boolean, integer, bigint); Type: FUNCTION; Schema: weasel_auth; Owner: weasel
--

CREATE FUNCTION add_user(_organization_id bigint, _user_firstname character varying, _user_lastname character varying, _user_middlename character varying, _user_job_title character varying, _user_phone character varying, _user_email character varying, _user_password character varying, _is_admin boolean, _timezone_id integer, _changer_id bigint, OUT _user_id bigint) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
    
    declare
        _token_id varchar;
    begin 

        PERFORM 1 FROM weasel_auth.users WHERE user_login = lower(_user_email) and user_password = _user_password;

        IF FOUND THEN 
            RAISE EXCEPTION 'USER_EXISTS';
        END IF;

        PERFORM 1 FROM weasel_auth.users WHERE lower(user_email) = lower(_user_email);

        IF FOUND THEN 
            RAISE EXCEPTION 'EMAIL_EXISTS';
        END IF;

        INSERT INTO weasel_auth.users (
            organization_id, 
            user_firstname, 
            user_lastname, 
            user_middlename, 
            user_job_title,
            user_phone,
            user_login, 
            user_email, 
            user_password,
            is_admin,
            is_active,
            is_deleted,
            timezone_id,
            changer_id
        ) VALUES (
            _organization_id, 
            _user_firstname, 
            _user_lastname, 
            _user_middlename, 
            _user_job_title,
            _user_phone,
            _user_email, 
            _user_email, 
            _user_password,
            _is_admin,
            true,
            true,
            false,
            _timezone_id,
            _changer_id
        ) RETURNING weasel_auth.users.user_id INTO _user_id;

    end;
$$;


ALTER FUNCTION weasel_auth.add_user(_organization_id bigint, _user_firstname character varying, _user_lastname character varying, _user_middlename character varying, _user_job_title character varying, _user_phone character varying, _user_email character varying, _user_password character varying, _is_admin boolean, _timezone_id integer, _changer_id bigint, OUT _user_id bigint) OWNER TO weasel;

SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: save_classifier_item(bigint, bigint, character varying, character varying, bigint, jsonb, boolean); Type: FUNCTION; Schema: weasel_classifiers; Owner: weasel
--

CREATE FUNCTION save_classifier_item(_id bigint, _reference_id bigint, _name character varying, _alias character varying, _pid bigint, _fields jsonb, _is_group boolean) RETURNS bigint
    LANGUAGE plpgsql
    AS $$

    declare
        _version bigint;
        _nid bigint;
    begin

		if _id = 0 then
			select into _version 0;
			update weasel_classifiers.counter set total = total + 1 where reference_id = _reference_id;
		else
	    	select into _version max(ver)+1 from weasel_classifiers.items where reference_id = _reference_id and "alias" = _alias;
	    	update weasel_classifiers.items set ver = _version, updated_at = current_timestamp where id = _id;
	   end if;

	   insert into weasel_classifiers.items (reference_id, name, alias, pid, fields, is_group, created_at, parents) values
	   (_reference_id, _name, _alias, _pid, _fields, _is_group, current_timestamp,
			coalesce( (select array_to_string(array_append(string_to_array(parents, '.')::bigint[], _pid), '.') from weasel_classifiers.items where id = _pid) ,'')
		) returning id into _nid;

    	return coalesce(_nid, 0);

    end;
$$;


ALTER FUNCTION weasel_classifiers.save_classifier_item(_id bigint, _reference_id bigint, _name character varying, _alias character varying, _pid bigint, _fields jsonb, _is_group boolean) OWNER TO weasel;

SET search_path = weasel_main, pg_catalog;

--
-- Name: get_plan_rows(bigint, bigint); Type: FUNCTION; Schema: weasel_main; Owner: weasel
--

CREATE FUNCTION get_plan_rows(_oid bigint, _period_id bigint) RETURNS SETOF record
    LANGUAGE plpgsql
    AS $$

begin

return query
	with ref_trunc as ( select id,
		jsonb_array_elements(dims_meta)->>'ReferenceAlias' as rkey,
		jsonb_array_elements(dims_meta)->>'Label' as rlabel,
		jsonb_array_elements(dims_meta)->>'Value' as rvalue
		from budget_plan
		)
	
	select distinct on (budget_plan.id) budget_operations.id, sum, date_op, dims_meta, user_meta from budget_plan
	left join ref_trunc on ref_trunc.id = budget_plan.id
	where organization_id = _oid
	and 
		ref_trunc.rkey = 'period' and ref_trunc.rvalue = _period_id::text
	order by date_op desc;
		
    end;
$$;


ALTER FUNCTION weasel_main.get_plan_rows(_oid bigint, _period_id bigint) OWNER TO weasel;

--
-- Name: save_budget_operation(bigint, bigint, bigint, numeric, text, jsonb, jsonb); Type: FUNCTION; Schema: weasel_main; Owner: weasel
--

CREATE FUNCTION save_budget_operation(_id bigint, _user_id bigint, _organization_id bigint, _sum numeric, _date_op text, _user_meta jsonb, _dims_meta jsonb) RETURNS bigint
    LANGUAGE plpgsql
    AS $$

    declare
        _nid bigint;
        dt TIMESTAMP;
    begin

    select into dt _date_op::TIMESTAMP;

    update weasel_main.budget_operations set "sum" = _sum, date_op = dt, updated_at = current_timestamp,
    user_meta = _user_meta, dims_meta = _dims_meta
        where id = _id RETURNING id INTO _nid;

	   insert into weasel_main.budget_operations (user_id, organization_id, "sum", date_op, user_meta, dims_meta)
	   select _user_id, _organization_id, _sum, dt, _user_meta, _dims_meta
	   where NOT EXISTS(select 1 from weasel_main.budget_operations where id = _id)
			returning id into _nid;

    	return coalesce(_nid, 0);

    end;
$$;


ALTER FUNCTION weasel_main.save_budget_operation(_id bigint, _user_id bigint, _organization_id bigint, _sum numeric, _date_op text, _user_meta jsonb, _dims_meta jsonb) OWNER TO weasel;

SET search_path = weasel_auth, pg_catalog;

--
-- Name: organization_id; Type: SEQUENCE; Schema: weasel_auth; Owner: weasel
--

CREATE SEQUENCE organization_id
    START WITH 1024
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE organization_id OWNER TO weasel;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: organizations; Type: TABLE; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

CREATE TABLE organizations (
    organization_id bigint DEFAULT nextval('organization_id'::regclass) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    organization_name character varying DEFAULT ''::character varying NOT NULL
);


ALTER TABLE organizations OWNER TO weasel;

--
-- Name: user_id; Type: SEQUENCE; Schema: weasel_auth; Owner: weasel
--

CREATE SEQUENCE user_id
    START WITH 1024
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id OWNER TO weasel;

--
-- Name: users; Type: TABLE; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

CREATE TABLE users (
    user_id bigint DEFAULT nextval('user_id'::regclass) NOT NULL,
    user_password character varying(64) NOT NULL,
    user_firstname character varying(64) NOT NULL,
    user_lastname character varying(64) NOT NULL,
    user_middlename character varying(64) NOT NULL,
    user_groups bigint[] DEFAULT '{}'::bigint[] NOT NULL,
    user_roles bigint[] DEFAULT '{}'::bigint[] NOT NULL,
    user_login character varying(128) NOT NULL,
    user_job_title character varying(255) DEFAULT ''::character varying NOT NULL,
    user_phone character varying DEFAULT ''::character varying NOT NULL,
    user_email character varying(128) NOT NULL,
    email_is_confirmed boolean DEFAULT false NOT NULL,
    password_expiration_date date DEFAULT ('now'::text)::date NOT NULL,
    organization_id bigint NOT NULL,
    timezone_id integer DEFAULT 0 NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    is_active boolean DEFAULT false NOT NULL,
    is_deleted boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE users OWNER TO weasel;

SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: counter; Type: TABLE; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE TABLE counter (
    reference_id bigint NOT NULL,
    total bigint DEFAULT 0 NOT NULL
)
WITH (fillfactor=50);


ALTER TABLE counter OWNER TO weasel;

--
-- Name: items; Type: TABLE; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE TABLE items (
    id bigint NOT NULL,
    reference_id bigint,
    name character varying(500) NOT NULL,
    alias character varying(500) NOT NULL,
    ver integer DEFAULT 0 NOT NULL,
    is_group boolean DEFAULT false NOT NULL,
    parents character varying(500) DEFAULT '0'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    pid bigint DEFAULT 0 NOT NULL,
    fields jsonb DEFAULT '{}'::jsonb NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE items OWNER TO weasel;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: weasel_classifiers; Owner: weasel
--

CREATE SEQUENCE items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE items_id_seq OWNER TO weasel;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_classifiers; Owner: weasel
--

ALTER SEQUENCE items_id_seq OWNED BY items.id;


--
-- Name: references; Type: TABLE; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE TABLE "references" (
    id bigint NOT NULL,
    name character varying(500) NOT NULL,
    alias character varying(500) NOT NULL,
    organization_id bigint NOT NULL,
    blocked boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    items_meta jsonb DEFAULT '{}'::jsonb NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE "references" OWNER TO weasel;

--
-- Name: references_id_seq; Type: SEQUENCE; Schema: weasel_classifiers; Owner: weasel
--

CREATE SEQUENCE references_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE references_id_seq OWNER TO weasel;

--
-- Name: references_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_classifiers; Owner: weasel
--

ALTER SEQUENCE references_id_seq OWNED BY "references".id;


SET search_path = weasel_main, pg_catalog;

--
-- Name: budget_operations; Type: TABLE; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE TABLE budget_operations (
    id bigint NOT NULL,
    user_id bigint DEFAULT 0 NOT NULL,
    organization_id bigint DEFAULT 0 NOT NULL,
    sum numeric DEFAULT 0 NOT NULL,
    date_op timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    dims_meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    user_meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    tags text[] DEFAULT '{}'::text[] NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE budget_operations OWNER TO weasel;

--
-- Name: budget_operations_id_seq; Type: SEQUENCE; Schema: weasel_main; Owner: weasel
--

CREATE SEQUENCE budget_operations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE budget_operations_id_seq OWNER TO weasel;

--
-- Name: budget_operations_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_main; Owner: weasel
--

ALTER SEQUENCE budget_operations_id_seq OWNED BY budget_operations.id;


--
-- Name: budget_plan; Type: TABLE; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE TABLE budget_plan (
    id bigint NOT NULL,
    user_id bigint DEFAULT 0 NOT NULL,
    organization_id bigint DEFAULT 0 NOT NULL,
    sum numeric DEFAULT 0 NOT NULL,
    date_op timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    dims_meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    user_meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    tags character varying[] DEFAULT '{}'::character varying[] NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE budget_plan OWNER TO weasel;

--
-- Name: budget_plan_id_seq; Type: SEQUENCE; Schema: weasel_main; Owner: weasel
--

CREATE SEQUENCE budget_plan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE budget_plan_id_seq OWNER TO weasel;

--
-- Name: budget_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_main; Owner: weasel
--

ALTER SEQUENCE budget_plan_id_seq OWNED BY budget_plan.id;


--
-- Name: eklps_horses; Type: TABLE; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE TABLE eklps_horses (
    id bigint NOT NULL,
    name character varying(500) NOT NULL,
    born integer NOT NULL,
    owner character varying(500) NOT NULL,
    breeder character varying(500) DEFAULT now() NOT NULL,
    sex character varying(128) NOT NULL,
    height jsonb,
    wins_total numeric,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE eklps_horses OWNER TO weasel;

--
-- Name: eklps_horses_id_seq; Type: SEQUENCE; Schema: weasel_main; Owner: weasel
--

CREATE SEQUENCE eklps_horses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE eklps_horses_id_seq OWNER TO weasel;

--
-- Name: eklps_horses_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_main; Owner: weasel
--

ALTER SEQUENCE eklps_horses_id_seq OWNED BY eklps_horses.id;


--
-- Name: organizations; Type: TABLE; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE TABLE organizations (
    id bigint NOT NULL,
    user_id bigint DEFAULT 0 NOT NULL,
    organization_id bigint DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    meta_info jsonb DEFAULT '{}'::jsonb NOT NULL,
    inn character varying(12) NOT NULL,
    kpp character varying(12) NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE organizations OWNER TO weasel;

--
-- Name: organizations_id_seq; Type: SEQUENCE; Schema: weasel_main; Owner: weasel
--

CREATE SEQUENCE organizations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE organizations_id_seq OWNER TO weasel;

--
-- Name: organizations_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_main; Owner: weasel
--

ALTER SEQUENCE organizations_id_seq OWNED BY organizations.id;


SET search_path = weasel_storage, pg_catalog;

--
-- Name: files; Type: TABLE; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

CREATE TABLE files (
    id bigint NOT NULL,
    entity character varying(256) NOT NULL,
    entity_id bigint NOT NULL,
    name character varying(500) NOT NULL,
    file_path character varying(1000) NOT NULL,
    organization_id bigint NOT NULL,
    owner_user jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    items_meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    version_id bigint DEFAULT 0 NOT NULL,
    bucket bigint DEFAULT 0 NOT NULL,
    md5_sum character(32) DEFAULT ''::bpchar NOT NULL,
    content_type character varying(1000) DEFAULT ''::character varying NOT NULL
)
WITH (fillfactor=90);


ALTER TABLE files OWNER TO weasel;

--
-- Name: files_id_seq; Type: SEQUENCE; Schema: weasel_storage; Owner: weasel
--

CREATE SEQUENCE files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE files_id_seq OWNER TO weasel;

--
-- Name: files_id_seq; Type: SEQUENCE OWNED BY; Schema: weasel_storage; Owner: weasel
--

ALTER SEQUENCE files_id_seq OWNED BY files.id;


SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: id; Type: DEFAULT; Schema: weasel_classifiers; Owner: weasel
--

ALTER TABLE ONLY items ALTER COLUMN id SET DEFAULT nextval('items_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: weasel_classifiers; Owner: weasel
--

ALTER TABLE ONLY "references" ALTER COLUMN id SET DEFAULT nextval('references_id_seq'::regclass);


SET search_path = weasel_main, pg_catalog;

--
-- Name: id; Type: DEFAULT; Schema: weasel_main; Owner: weasel
--

ALTER TABLE ONLY budget_operations ALTER COLUMN id SET DEFAULT nextval('budget_operations_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: weasel_main; Owner: weasel
--

ALTER TABLE ONLY budget_plan ALTER COLUMN id SET DEFAULT nextval('budget_plan_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: weasel_main; Owner: weasel
--

ALTER TABLE ONLY eklps_horses ALTER COLUMN id SET DEFAULT nextval('eklps_horses_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: weasel_main; Owner: weasel
--

ALTER TABLE ONLY organizations ALTER COLUMN id SET DEFAULT nextval('organizations_id_seq'::regclass);


SET search_path = weasel_storage, pg_catalog;

--
-- Name: id; Type: DEFAULT; Schema: weasel_storage; Owner: weasel
--

ALTER TABLE ONLY files ALTER COLUMN id SET DEFAULT nextval('files_id_seq'::regclass);


SET search_path = weasel_auth, pg_catalog;

--
-- Name: organization_id; Type: SEQUENCE SET; Schema: weasel_auth; Owner: weasel
--

SELECT pg_catalog.setval('organization_id', 1024, false);


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: weasel_auth; Owner: weasel
--

COPY organizations (organization_id, created_at, organization_name) FROM stdin;
1	2015-09-01 16:54:58.173604	Основная организация
\.


--
-- Name: user_id; Type: SEQUENCE SET; Schema: weasel_auth; Owner: weasel
--

SELECT pg_catalog.setval('user_id', 1027, true);


--
-- Data for Name: users; Type: TABLE DATA; Schema: weasel_auth; Owner: weasel
--

COPY users (user_id, user_password, user_firstname, user_lastname, user_middlename, user_groups, user_roles, user_login, user_job_title, user_phone, user_email, email_is_confirmed, password_expiration_date, organization_id, timezone_id, is_admin, is_active, is_deleted, created_at, updated_at) FROM stdin;
1025	f639273dd133c1a620a7f57a0112568d187a82c7		asdf		{}	{}	asdf	job_title		asdf	f	2015-09-01	1	1	t	t	f	2015-09-01 16:55:03.477796	2015-09-01 16:55:03.477796
1026	922194c40c9a4d599ae44eab7ef2d447d5bcdf99	Админ	Админов	Админович	{}	{}	aehler@yandex.ru	job_title		aehler@yandex.ru	f	2015-09-03	1	1	t	t	f	2015-09-03 16:21:39.625593	2015-09-03 16:21:39.625593
1027	f639273dd133c1a620a7f57a0112568d187a82c7				{}	{}		job_title			f	2015-09-04	1	1	t	t	f	2015-09-04 15:42:38.227792	2015-09-04 15:42:38.227792
\.


SET search_path = weasel_classifiers, pg_catalog;

--
-- Data for Name: counter; Type: TABLE DATA; Schema: weasel_classifiers; Owner: weasel
--

COPY counter (reference_id, total) FROM stdin;
3	0
4	0
2	2
1	4
5	5
7	4
9	4
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: weasel_classifiers; Owner: weasel
--

COPY items (id, reference_id, name, alias, ver, is_group, parents, created_at, updated_at, pid, fields) FROM stdin;
6	2	Возврат задолженности	fede0b67faff295500f2970abf72298d111ec25d	0	f	1	2015-09-17 15:17:52.708744	2015-09-17 15:17:52.708744	1	{"pid": "1", "title": "Возврат задолженности", "is_cost": "false", "is_group": "false"}
7	2	Расходы на поддержание жизнедеятельности	365638e2012bbdf9ac89dafb4371d154a0f34e98	1	t	5	2015-09-17 17:03:37.940121	2015-09-17 17:53:12.966866	5	{"pid": "5", "title": "Расходы на поддержание жизнедеятельности", "is_cost": "true", "is_group": "true"}
8	2	Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности	365638e2012bbdf9ac89dafb4371d154a0f34e98	0	f	5	2015-09-17 17:53:12.966866	2015-09-17 17:53:12.966866	5	{"title": "Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности", "is_cost": "true, true", "is_group": "true, false"}
11	1	Важный	89fe1668e84e73fdc8fcd68a954d3e570b251e84	0	f		2015-09-17 17:59:43.172269	2015-09-17 17:59:43.172269	0	{"pid": "0", "title": "Важный"}
1	2	Доходы	eb92d669f5c4593a3b9b9475ec30e13e414fed78	0	t		2015-09-16 14:46:35.928543	2015-09-16 14:46:35.928543	0	{"pid": "0", "title": "Доходы", "is_cost": "false", "is_group": "true", "is_deleted": "false"}
3	2	Заработная плата	54c7f21909fa05a59114fd91e86b5b1fe14ea02f	0	f	1	2015-09-16 15:43:15.257041	2015-09-16 15:43:15.257041	1	{"pid": "1", "title": "Заработная плата", "is_cost": "false", "is_group": "false", "is_deleted": "false"}
12	1	Нормальный	4fb3acf5191475a5488142ed7c5d237e4d3bc2d8	0	f		2015-09-17 17:59:55.787565	2015-09-17 17:59:55.787565	0	{"pid": "0", "title": "Нормальный"}
5	2	Расходы	04fff035d530318f02adc582d6aa4640420ccea3	0	t		2015-09-16 15:47:09.552155	2015-09-16 15:47:09.552155	0	{"pid": "0", "title": "Расходы", "is_cost": "true", "is_group": "true", "is_deleted": "false"}
13	1	Не важный	ffdb11115900cef2d8e4f6329f9318cd68f834a8	0	f		2015-09-17 18:00:03.641717	2015-09-17 18:00:03.641717	0	{"pid": "0", "title": "Не важный"}
14	1	Срочный	f923c9c56f2062fb34a2e0397d500851b777ac69	0	f		2015-09-17 18:00:10.037553	2015-09-17 18:00:10.037553	0	{"pid": "0", "title": "Срочный"}
15	5	Карта зарплатная ГПБ	41260e3e0d7917c062a9806d6998afbe05ff0413	0	f		2015-09-17 18:05:37.387622	2015-09-17 18:05:37.387622	0	{"pid": "0", "title": "Карта зарплатная ГПБ"}
16	5	Карта зарплатная Промсвязбанк	3ffb773545b003c164c5b6f6b8bd22c3270dd7ce	0	f		2015-09-17 18:05:51.383968	2015-09-17 18:05:51.383968	0	{"pid": "0", "title": "Карта зарплатная Промсвязбанк"}
17	5	Карта крединная	4629125ca9372c04aef1b9ca9772a1601b77fffa	0	f		2015-09-17 18:06:00.240242	2015-09-17 18:06:00.240242	0	{"pid": "0", "title": "Карта крединная"}
18	5	Карта дебетовая Сбербанк	a9124675fb7a5a302ef342aba14ad6a1daab0470	0	f		2015-09-17 18:06:22.414914	2015-09-17 18:06:22.414914	0	{"pid": "0", "title": "Карта дебетовая Сбербанк"}
19	5	Наличные	5827b3cd7bfadda4e12d6ddef96992b0b8df006b	0	f		2015-09-17 18:06:30.377454	2015-09-17 18:06:30.377454	0	{"pid": "0", "title": "Наличные"}
20	7	RUB	134eddf39ef541315e07e1b63b22f8b4144d8dbe	0	f		2015-09-22 15:51:42.46625	2015-09-22 15:51:42.46625	0	{"pid": "0", "code": "643", "title": "RUB"}
21	7	USD	d63a586540c41dc3ba5f46f53e2b23d9ecdbaaef	0	f		2015-09-22 15:52:33.629232	2015-09-22 15:52:33.629232	0	{"pid": "0", "code": "840", "title": "USD"}
22	7	EUR	624207a4c9df7fec62a081439aa94483788148a3	0	f		2015-09-22 15:53:00.793095	2015-09-22 15:53:00.793095	0	{"pid": "0", "code": "978", "title": "EUR"}
23	7	PLN	44c819893e21a7edff9123cd3ea3962bb4954211	0	f		2015-09-22 15:53:15.286227	2015-09-22 15:53:15.286227	0	{"pid": "0", "code": "985", "title": "PLN"}
24	9	Год 2015	2781471ba2dfd32662b9374c8dc09f0136c4574d	0	t		2015-10-20 17:03:19.276346	2015-10-20 17:03:19.276346	0	{"pid": "0", "title": "Год 2015", "date_to": "31.12.2015", "is_group": "true", "date_from": "01.01.2015"}
25	9	Окрябрь 2015	0ed75a0326d05a11d5c2da324f7e6b52376316c0	0	f	24	2015-10-20 17:03:46.593102	2015-10-20 17:03:46.593102	24	{"pid": "24", "title": "Окрябрь 2015", "date_to": "31.10.2015", "is_group": "false", "date_from": "01.10.2015"}
26	9	Ноябрь 2015	a42fdffaee71834c7e1439ae095d76177e6f0846	0	f	24	2015-10-20 17:04:04.324899	2015-10-20 17:04:04.324899	24	{"pid": "24", "title": "Ноябрь 2015", "date_to": "30.11.2015", "is_group": "false", "date_from": "01.11.2015"}
27	9	Декабрь 2015	c2f772cde9b705b65e8b63ee8c3d347d00a5f6e0	0	f	24	2015-10-20 17:04:21.484147	2015-10-20 17:04:21.484147	24	{"pid": "24", "title": "Декабрь 2015", "date_to": "31.12.2015", "is_group": "false", "date_from": "01.12.2015"}
\.


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: weasel_classifiers; Owner: weasel
--

SELECT pg_catalog.setval('items_id_seq', 27, true);


--
-- Data for Name: references; Type: TABLE DATA; Schema: weasel_classifiers; Owner: weasel
--

COPY "references" (id, name, alias, organization_id, blocked, created_at, updated_at, items_meta) FROM stdin;
2	Статьи затрат	cost_items	1	f	2015-09-08 19:03:06.272854	2015-09-08 19:03:06.272854	{"type": "tree", "fields": {"cost": "bool", "title": "text"}}
1	Приоритет платежа	importance	1	f	2015-09-08 19:02:02.964436	2015-09-08 19:02:02.964436	{"type": "plain", "fields": {"title": "text"}}
3	Проект	project	1	f	2015-09-08 19:03:29.179313	2015-09-08 19:03:29.179313	{"type": "plain", "fields": {"title": "text"}}
4	ЦФО	zfo	1	f	2015-09-14 13:28:24.843479	2015-09-14 13:28:24.843479	{"type": "tree", "fields": {"title": "text"}}
5	Счета	accounts	1	f	2015-09-17 18:03:23.002926	2015-09-17 18:03:23.002926	{}
7	Валюты	currency	1	f	2015-09-22 15:42:40.910828	2015-09-22 15:42:40.910828	{}
9	Периоды	periods	1	f	2015-10-20 17:02:25.343689	2015-10-20 17:02:25.343689	{}
\.


--
-- Name: references_id_seq; Type: SEQUENCE SET; Schema: weasel_classifiers; Owner: weasel
--

SELECT pg_catalog.setval('references_id_seq', 9, true);


SET search_path = weasel_main, pg_catalog;

--
-- Data for Name: budget_operations; Type: TABLE DATA; Schema: weasel_main; Owner: weasel
--

COPY budget_operations (id, user_id, organization_id, sum, date_op, created_at, updated_at, dims_meta, user_meta, tags) FROM stdin;
1	1026	1	1	0001-01-01 00:00:00	2015-09-23 17:50:26.446296	2015-09-23 17:50:26.446296	[{"Label": "EUR", "Value": 22, "Options": [{"n": "EUR", "v": 22}, {"n": "PLN", "v": 23}, {"n": "RUB", "v": 20}, {"n": "USD", "v": 21}], "ReferenceAlias": "currency", "ReferenceLabel": "Валюты"}, {"Label": "Доходы", "Value": 1, "Options": [{"n": "Доходы", "v": 1}, {"n": "Возврат задолженности", "v": 6}, {"n": "Заработная плата", "v": 3}, {"n": "Расходы", "v": 5}, {"n": "Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности", "v": 8}], "ReferenceAlias": "cost_items", "ReferenceLabel": "Статьи затрат"}, {"Label": "Важный", "Value": 11, "Options": [{"n": "Важный", "v": 11}, {"n": "Не важный", "v": 13}, {"n": "Нормальный", "v": 12}, {"n": "Срочный", "v": 14}], "ReferenceAlias": "importance", "ReferenceLabel": "Приоритет платежа"}, {"Label": "Наличные", "Value": 19, "Options": [{"n": "Карта дебетовая Сбербанк", "v": 18}, {"n": "Карта зарплатная ГПБ", "v": 15}, {"n": "Карта зарплатная Промсвязбанк", "v": 16}, {"n": "Карта крединная", "v": 17}, {"n": "Наличные", "v": 19}], "ReferenceAlias": "accounts", "ReferenceLabel": "Счета"}, {"Label": "", "Value": 0, "Options": [], "ReferenceAlias": "project", "ReferenceLabel": "Проект"}]	{"a": true, "e": "aehler@yandex.ru", "i": 1026, "l": "aehler@yandex.ru", "oi": 1, "uf": "Админ", "ul": "Админов", "um": "Админович", "adm": true}	{}
2	1026	1	123123123	2015-09-23 00:00:00	2015-09-23 17:51:44.75716	2015-09-23 17:51:44.75716	[{"Label": "EUR", "Value": 22, "Options": [{"n": "EUR", "v": 22}, {"n": "PLN", "v": 23}, {"n": "RUB", "v": 20}, {"n": "USD", "v": 21}], "ReferenceAlias": "currency", "ReferenceLabel": "Валюты"}, {"Label": "Доходы", "Value": 1, "Options": [{"n": "Доходы", "v": 1}, {"n": "Возврат задолженности", "v": 6}, {"n": "Заработная плата", "v": 3}, {"n": "Расходы", "v": 5}, {"n": "Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности", "v": 8}], "ReferenceAlias": "cost_items", "ReferenceLabel": "Статьи затрат"}, {"Label": "Важный", "Value": 11, "Options": [{"n": "Важный", "v": 11}, {"n": "Не важный", "v": 13}, {"n": "Нормальный", "v": 12}, {"n": "Срочный", "v": 14}], "ReferenceAlias": "importance", "ReferenceLabel": "Приоритет платежа"}, {"Label": "Карта дебетовая Сбербанк", "Value": 18, "Options": [{"n": "Карта дебетовая Сбербанк", "v": 18}, {"n": "Карта зарплатная ГПБ", "v": 15}, {"n": "Карта зарплатная Промсвязбанк", "v": 16}, {"n": "Карта крединная", "v": 17}, {"n": "Наличные", "v": 19}], "ReferenceAlias": "accounts", "ReferenceLabel": "Счета"}, {"Label": "", "Value": 0, "Options": [], "ReferenceAlias": "project", "ReferenceLabel": "Проект"}]	{"a": true, "e": "aehler@yandex.ru", "i": 1026, "l": "aehler@yandex.ru", "oi": 1, "uf": "Админ", "ul": "Админов", "um": "Админович", "adm": true}	{}
3	1026	1	83	0001-01-01 00:00:00	2015-10-16 10:49:21.360684	2015-10-16 10:49:21.360684	[{"Label": "EUR", "Value": 22, "Options": [{"n": "EUR", "v": 22}, {"n": "PLN", "v": 23}, {"n": "RUB", "v": 20}, {"n": "USD", "v": 21}], "ReferenceAlias": "currency", "ReferenceLabel": "Валюты"}, {"Label": "Доходы", "Value": 1, "Options": [{"n": "Доходы", "v": 1}, {"n": "Возврат задолженности", "v": 6}, {"n": "Заработная плата", "v": 3}, {"n": "Расходы", "v": 5}, {"n": "Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности", "v": 8}], "ReferenceAlias": "cost_items", "ReferenceLabel": "Статьи затрат"}, {"Label": "Важный", "Value": 11, "Options": [{"n": "Важный", "v": 11}, {"n": "Не важный", "v": 13}, {"n": "Нормальный", "v": 12}, {"n": "Срочный", "v": 14}], "ReferenceAlias": "importance", "ReferenceLabel": "Приоритет платежа"}, {"Label": "Карта дебетовая Сбербанк", "Value": 18, "Options": [{"n": "Карта дебетовая Сбербанк", "v": 18}, {"n": "Карта зарплатная ГПБ", "v": 15}, {"n": "Карта зарплатная Промсвязбанк", "v": 16}, {"n": "Карта крединная", "v": 17}, {"n": "Наличные", "v": 19}], "ReferenceAlias": "accounts", "ReferenceLabel": "Счета"}, {"Label": "", "Value": 0, "Options": [], "ReferenceAlias": "project", "ReferenceLabel": "Проект"}]	{"a": true, "e": "aehler@yandex.ru", "i": 1026, "l": "aehler@yandex.ru", "oi": 1, "uf": "Админ", "ul": "Админов", "um": "Админович", "adm": true}	{}
4	1026	1	50	2015-10-16 00:00:00	2015-10-16 10:49:45.498503	2015-10-16 10:49:45.498503	[{"Label": "RUB", "Value": 20, "Options": [{"n": "EUR", "v": 22}, {"n": "PLN", "v": 23}, {"n": "RUB", "v": 20}, {"n": "USD", "v": 21}], "ReferenceAlias": "currency", "ReferenceLabel": "Валюты"}, {"Label": "Расходы", "Value": 5, "Options": [{"n": "Доходы", "v": 1}, {"n": "Возврат задолженности", "v": 6}, {"n": "Заработная плата", "v": 3}, {"n": "Расходы", "v": 5}, {"n": "Расходы на поддержание жизнедеятельности, Расходы на поддержание жизнедеятельности", "v": 8}], "ReferenceAlias": "cost_items", "ReferenceLabel": "Статьи затрат"}, {"Label": "Нормальный", "Value": 12, "Options": [{"n": "Важный", "v": 11}, {"n": "Не важный", "v": 13}, {"n": "Нормальный", "v": 12}, {"n": "Срочный", "v": 14}], "ReferenceAlias": "importance", "ReferenceLabel": "Приоритет платежа"}, {"Label": "Наличные", "Value": 19, "Options": [{"n": "Карта дебетовая Сбербанк", "v": 18}, {"n": "Карта зарплатная ГПБ", "v": 15}, {"n": "Карта зарплатная Промсвязбанк", "v": 16}, {"n": "Карта крединная", "v": 17}, {"n": "Наличные", "v": 19}], "ReferenceAlias": "accounts", "ReferenceLabel": "Счета"}, {"Label": "", "Value": 0, "Options": [], "ReferenceAlias": "project", "ReferenceLabel": "Проект"}]	{"a": true, "e": "aehler@yandex.ru", "i": 1026, "l": "aehler@yandex.ru", "oi": 1, "uf": "Админ", "ul": "Админов", "um": "Админович", "adm": true}	{}
\.


--
-- Name: budget_operations_id_seq; Type: SEQUENCE SET; Schema: weasel_main; Owner: weasel
--

SELECT pg_catalog.setval('budget_operations_id_seq', 4, true);


--
-- Data for Name: budget_plan; Type: TABLE DATA; Schema: weasel_main; Owner: weasel
--

COPY budget_plan (id, user_id, organization_id, sum, date_op, created_at, updated_at, dims_meta, user_meta, tags) FROM stdin;
\.


--
-- Name: budget_plan_id_seq; Type: SEQUENCE SET; Schema: weasel_main; Owner: weasel
--

SELECT pg_catalog.setval('budget_plan_id_seq', 1, false);


--
-- Data for Name: eklps_horses; Type: TABLE DATA; Schema: weasel_main; Owner: weasel
--

COPY eklps_horses (id, name, born, owner, breeder, sex, height, wins_total, updated_at) FROM stdin;
\.


--
-- Name: eklps_horses_id_seq; Type: SEQUENCE SET; Schema: weasel_main; Owner: weasel
--

SELECT pg_catalog.setval('eklps_horses_id_seq', 1, false);


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: weasel_main; Owner: weasel
--

COPY organizations (id, user_id, organization_id, created_at, updated_at, meta_info, inn, kpp) FROM stdin;
\.


--
-- Name: organizations_id_seq; Type: SEQUENCE SET; Schema: weasel_main; Owner: weasel
--

SELECT pg_catalog.setval('organizations_id_seq', 1, false);


SET search_path = weasel_storage, pg_catalog;

--
-- Data for Name: files; Type: TABLE DATA; Schema: weasel_storage; Owner: weasel
--

COPY files (id, entity, entity_id, name, file_path, organization_id, owner_user, created_at, updated_at, items_meta, version_id, bucket, md5_sum, content_type) FROM stdin;
\.


--
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: weasel_storage; Owner: weasel
--

SELECT pg_catalog.setval('files_id_seq', 1, false);


SET search_path = weasel_auth, pg_catalog;

--
-- Name: organizations_pkey; Type: CONSTRAINT; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (organization_id);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: counter_pkey; Type: CONSTRAINT; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY counter
    ADD CONSTRAINT counter_pkey PRIMARY KEY (reference_id);


--
-- Name: items_pkey; Type: CONSTRAINT; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: references_pkey; Type: CONSTRAINT; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY "references"
    ADD CONSTRAINT references_pkey PRIMARY KEY (id);


SET search_path = weasel_main, pg_catalog;

--
-- Name: budget_operations_pkey; Type: CONSTRAINT; Schema: weasel_main; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY budget_operations
    ADD CONSTRAINT budget_operations_pkey PRIMARY KEY (id);


--
-- Name: budget_plan_pkey; Type: CONSTRAINT; Schema: weasel_main; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY budget_plan
    ADD CONSTRAINT budget_plan_pkey PRIMARY KEY (id);


--
-- Name: eklps_horses_pkey; Type: CONSTRAINT; Schema: weasel_main; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY eklps_horses
    ADD CONSTRAINT eklps_horses_pkey PRIMARY KEY (id);


--
-- Name: organizations_pkey; Type: CONSTRAINT; Schema: weasel_main; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


SET search_path = weasel_storage, pg_catalog;

--
-- Name: files_pkey; Type: CONSTRAINT; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

ALTER TABLE ONLY files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


SET search_path = weasel_auth, pg_catalog;

--
-- Name: idx_user_member_in; Type: INDEX; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_user_member_in ON users USING btree (organization_id);


--
-- Name: uidx_users_email; Type: INDEX; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

CREATE UNIQUE INDEX uidx_users_email ON users USING btree (lower((user_email)::text));


--
-- Name: uidx_users_password; Type: INDEX; Schema: weasel_auth; Owner: weasel; Tablespace: 
--

CREATE UNIQUE INDEX uidx_users_password ON users USING btree (user_password, user_login);


SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: idx_inm; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_inm ON items USING btree (lower((name)::text));


--
-- Name: idx_iparents; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_iparents ON items USING btree (lower((parents)::text));


--
-- Name: idx_ipid; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_ipid ON items USING btree (pid);


--
-- Name: idx_irefid; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_irefid ON items USING btree (reference_id);


--
-- Name: idx_iver; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_iver ON items USING btree (ver, reference_id);


--
-- Name: idx_nm; Type: INDEX; Schema: weasel_classifiers; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_nm ON "references" USING btree (lower((alias)::text));


SET search_path = weasel_main, pg_catalog;

--
-- Name: idx_date; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_date ON budget_operations USING btree (date_op);


--
-- Name: idx_org_inn_kpp; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE UNIQUE INDEX idx_org_inn_kpp ON organizations USING btree (inn, kpp);


--
-- Name: idx_org_organization_id; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_org_organization_id ON organizations USING btree (organization_id);


--
-- Name: idx_organization_id; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_organization_id ON budget_operations USING btree (organization_id);


--
-- Name: idx_p_date; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_p_date ON budget_plan USING btree (date_op);


--
-- Name: idx_p_organization_id; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_p_organization_id ON budget_plan USING btree (organization_id);


--
-- Name: idx_p_tags; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_p_tags ON budget_plan USING gin (tags);


--
-- Name: idx_tags; Type: INDEX; Schema: weasel_main; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_tags ON budget_operations USING gin (tags);


SET search_path = weasel_storage, pg_catalog;

--
-- Name: idx_f_ent_ent_id; Type: INDEX; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_f_ent_ent_id ON files USING btree (entity, entity_id);


--
-- Name: idx_f_org_id; Type: INDEX; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_f_org_id ON files USING btree (organization_id);


--
-- Name: idx_f_version; Type: INDEX; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_f_version ON files USING btree (version_id);


--
-- Name: idx_fctype; Type: INDEX; Schema: weasel_storage; Owner: weasel; Tablespace: 
--

CREATE INDEX idx_fctype ON files USING btree (lower((content_type)::text));


SET search_path = weasel_auth, pg_catalog;

--
-- Name: users_organization_id_fkey; Type: FK CONSTRAINT; Schema: weasel_auth; Owner: weasel
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(organization_id);


SET search_path = weasel_classifiers, pg_catalog;

--
-- Name: items_reference_id_fkey; Type: FK CONSTRAINT; Schema: weasel_classifiers; Owner: weasel
--

ALTER TABLE ONLY items
    ADD CONSTRAINT items_reference_id_fkey FOREIGN KEY (reference_id) REFERENCES "references"(id);


SET search_path = weasel_main, pg_catalog;

--
-- Name: organizations_organization_id_fkey; Type: FK CONSTRAINT; Schema: weasel_main; Owner: weasel
--

ALTER TABLE ONLY organizations
    ADD CONSTRAINT organizations_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES weasel_auth.organizations(organization_id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;

--
-- Name: weasel_auth; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA weasel_auth FROM PUBLIC;
REVOKE ALL ON SCHEMA weasel_auth FROM postgres;
GRANT ALL ON SCHEMA weasel_auth TO postgres;
GRANT ALL ON SCHEMA weasel_auth TO weasel;


--
-- Name: weasel_classifiers; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA weasel_classifiers FROM PUBLIC;
REVOKE ALL ON SCHEMA weasel_classifiers FROM postgres;
GRANT ALL ON SCHEMA weasel_classifiers TO postgres;
GRANT ALL ON SCHEMA weasel_classifiers TO weasel;


--
-- Name: weasel_main; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA weasel_main FROM PUBLIC;
REVOKE ALL ON SCHEMA weasel_main FROM postgres;
GRANT ALL ON SCHEMA weasel_main TO postgres;
GRANT ALL ON SCHEMA weasel_main TO weasel;


--
-- Name: weasel_storage; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA weasel_storage FROM PUBLIC;
REVOKE ALL ON SCHEMA weasel_storage FROM postgres;
GRANT ALL ON SCHEMA weasel_storage TO postgres;
GRANT ALL ON SCHEMA weasel_storage TO weasel;


--
-- PostgreSQL database dump complete
--

