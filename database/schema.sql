--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-05-28 10:02:44

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 226 (class 1259 OID 16474)
-- Name: application_data_structures; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application_data_structures (
    id bigint NOT NULL,
    application_id bigint NOT NULL,
    version integer DEFAULT 0 NOT NULL,
    start_byte integer NOT NULL,
    end_byte integer NOT NULL,
    data_type_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.application_data_structures OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16473)
-- Name: application_data_structures_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.application_data_structures_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.application_data_structures_id_seq OWNER TO postgres;

--
-- TOC entry 4833 (class 0 OID 0)
-- Dependencies: 225
-- Name: application_data_structures_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_data_structures_id_seq OWNED BY public.application_data_structures.id;


--
-- TOC entry 222 (class 1259 OID 16456)
-- Name: applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.applications (
    id bigint NOT NULL,
    name character varying NOT NULL,
    company_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.applications OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16455)
-- Name: applications_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.applications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.applications_id_seq OWNER TO postgres;

--
-- TOC entry 4834 (class 0 OID 0)
-- Dependencies: 221
-- Name: applications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.applications_id_seq OWNED BY public.applications.id;


--
-- TOC entry 220 (class 1259 OID 16447)
-- Name: companies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.companies (
    id bigint NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.companies OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16446)
-- Name: companies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.companies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.companies_id_seq OWNER TO postgres;

--
-- TOC entry 4835 (class 0 OID 0)
-- Dependencies: 219
-- Name: companies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.companies_id_seq OWNED BY public.companies.id;


--
-- TOC entry 224 (class 1259 OID 16465)
-- Name: data_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.data_types (
    id bigint NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.data_types OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16464)
-- Name: data_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.data_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.data_types_id_seq OWNER TO postgres;

--
-- TOC entry 4836 (class 0 OID 0)
-- Dependencies: 223
-- Name: data_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.data_types_id_seq OWNED BY public.data_types.id;


--
-- TOC entry 216 (class 1259 OID 16406)
-- Name: unpiped_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.unpiped_data (
    user_id bigint NOT NULL,
    application character varying NOT NULL,
    created_at time without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id integer NOT NULL,
    data bytea NOT NULL
);


ALTER TABLE public.unpiped_data OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16422)
-- Name: unpiped_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.unpiped_data_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.unpiped_data_id_seq OWNER TO postgres;

--
-- TOC entry 4837 (class 0 OID 0)
-- Dependencies: 218
-- Name: unpiped_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.unpiped_data_id_seq OWNED BY public.unpiped_data.id;


--
-- TOC entry 215 (class 1259 OID 16399)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    email character varying NOT NULL,
    apikey character varying NOT NULL,
    id integer NOT NULL,
    created_at time without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    company_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16413)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4838 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4670 (class 2604 OID 16477)
-- Name: application_data_structures id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_data_structures ALTER COLUMN id SET DEFAULT nextval('public.application_data_structures_id_seq'::regclass);


--
-- TOC entry 4666 (class 2604 OID 16459)
-- Name: applications id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications ALTER COLUMN id SET DEFAULT nextval('public.applications_id_seq'::regclass);


--
-- TOC entry 4664 (class 2604 OID 16450)
-- Name: companies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.companies_id_seq'::regclass);


--
-- TOC entry 4668 (class 2604 OID 16468)
-- Name: data_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_types ALTER COLUMN id SET DEFAULT nextval('public.data_types_id_seq'::regclass);


--
-- TOC entry 4663 (class 2604 OID 16423)
-- Name: unpiped_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.unpiped_data ALTER COLUMN id SET DEFAULT nextval('public.unpiped_data_id_seq'::regclass);


--
-- TOC entry 4659 (class 2604 OID 16414)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4684 (class 2606 OID 16480)
-- Name: application_data_structures application_data_structures_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_data_structures
    ADD CONSTRAINT application_data_structures_pkey PRIMARY KEY (id);


--
-- TOC entry 4680 (class 2606 OID 16463)
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- TOC entry 4678 (class 2606 OID 16454)
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- TOC entry 4682 (class 2606 OID 16472)
-- Name: data_types data_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_types
    ADD CONSTRAINT data_types_pkey PRIMARY KEY (id);


--
-- TOC entry 4676 (class 2606 OID 16430)
-- Name: unpiped_data unpiped_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.unpiped_data
    ADD CONSTRAINT unpiped_data_pkey PRIMARY KEY (id);


--
-- TOC entry 4674 (class 2606 OID 16421)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2024-05-28 10:02:45

--
-- PostgreSQL database dump complete
--

