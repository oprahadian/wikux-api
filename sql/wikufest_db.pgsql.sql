--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3 (Debian 10.3-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: app_user; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.app_user (
    id integer NOT NULL,
    password character varying(255),
    fullname character varying(255),
    email character varying(255),
    created_date timestamp without time zone,
    is_active boolean,
    class_name character varying(50)
);


ALTER TABLE public.app_user OWNER TO wiku;

--
-- Name: app_user_id_seq; Type: SEQUENCE; Schema: public; Owner: wiku
--

CREATE SEQUENCE public.app_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.app_user_id_seq OWNER TO wiku;

--
-- Name: app_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wiku
--

ALTER SEQUENCE public.app_user_id_seq OWNED BY public.app_user.id;


--
-- Name: course; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.course (
    id integer NOT NULL,
    is_private_class boolean,
    course_name character varying(255),
    description text,
    additional_link text,
    is_active boolean,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.course OWNER TO wiku;

--
-- Name: course_id_seq; Type: SEQUENCE; Schema: public; Owner: wiku
--

CREATE SEQUENCE public.course_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_id_seq OWNER TO wiku;

--
-- Name: course_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wiku
--

ALTER SEQUENCE public.course_id_seq OWNED BY public.course.id;


--
-- Name: course_session; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.course_session (
    id integer NOT NULL,
    course_id integer,
    room_name character varying(255),
    room_location character varying(255),
    quota integer,
    available_quota integer,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    is_active boolean,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.course_session OWNER TO wiku;

--
-- Name: course_session_enrollment; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.course_session_enrollment (
    course_session_id integer,
    app_user_id integer,
    created_date timestamp without time zone
);


ALTER TABLE public.course_session_enrollment OWNER TO wiku;

--
-- Name: course_session_id_seq; Type: SEQUENCE; Schema: public; Owner: wiku
--

CREATE SEQUENCE public.course_session_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_session_id_seq OWNER TO wiku;

--
-- Name: course_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wiku
--

ALTER SEQUENCE public.course_session_id_seq OWNED BY public.course_session.id;


--
-- Name: course_x_instructor; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.course_x_instructor (
    course_id integer,
    instructor_id integer
);


ALTER TABLE public.course_x_instructor OWNER TO wiku;

--
-- Name: data_peserta_wikufest_2018; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.data_peserta_wikufest_2018 (
    id character varying(100),
    nis character varying(100),
    email character varying(100),
    password character varying(100),
    fullname character varying(100),
    created_date character varying(100),
    is_active character varying(100),
    class_name character varying(100)
);


ALTER TABLE public.data_peserta_wikufest_2018 OWNER TO wiku;

--
-- Name: user_coupon; Type: TABLE; Schema: public; Owner: wiku
--

CREATE TABLE public.user_coupon (
    id integer NOT NULL,
    app_user_id integer,
    coupon_name character varying(50),
    coupon_code character varying(100),
    is_redeemed boolean,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.user_coupon OWNER TO wiku;

--
-- Name: user_coupon_id_seq; Type: SEQUENCE; Schema: public; Owner: wiku
--

CREATE SEQUENCE public.user_coupon_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_coupon_id_seq OWNER TO wiku;

--
-- Name: user_coupon_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wiku
--

ALTER SEQUENCE public.user_coupon_id_seq OWNED BY public.user_coupon.id;


--
-- Name: app_user id; Type: DEFAULT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.app_user ALTER COLUMN id SET DEFAULT nextval('public.app_user_id_seq'::regclass);


--
-- Name: course id; Type: DEFAULT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.course ALTER COLUMN id SET DEFAULT nextval('public.course_id_seq'::regclass);


--
-- Name: course_session id; Type: DEFAULT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.course_session ALTER COLUMN id SET DEFAULT nextval('public.course_session_id_seq'::regclass);


--
-- Name: user_coupon id; Type: DEFAULT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.user_coupon ALTER COLUMN id SET DEFAULT nextval('public.user_coupon_id_seq'::regclass);


--
-- Name: app_user app_user_pkey; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.app_user
    ADD CONSTRAINT app_user_pkey PRIMARY KEY (id);


--
-- Name: app_user app_user_un; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.app_user
    ADD CONSTRAINT app_user_un UNIQUE (email);


--
-- Name: course course_pkey; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.course
    ADD CONSTRAINT course_pkey PRIMARY KEY (id);


--
-- Name: course_session_enrollment course_session_enrollment_un; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.course_session_enrollment
    ADD CONSTRAINT course_session_enrollment_un UNIQUE (course_session_id, app_user_id);


--
-- Name: course_session course_session_pkey; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.course_session
    ADD CONSTRAINT course_session_pkey PRIMARY KEY (id);


--
-- Name: user_coupon user_coupon_pkey; Type: CONSTRAINT; Schema: public; Owner: wiku
--

ALTER TABLE ONLY public.user_coupon
    ADD CONSTRAINT user_coupon_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--
