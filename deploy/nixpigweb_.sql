--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

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
-- Name: content_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.content_ (
    id_ integer NOT NULL,
    title_ character varying(255) NOT NULL,
    subtitle_ character varying(255) NOT NULL,
    slug_ character varying(255) NOT NULL,
    body_ text,
    created_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    type_ character varying(4) DEFAULT 'post'::character varying NOT NULL,
    user_id_ integer NOT NULL,
    CONSTRAINT content__type__check CHECK (((type_)::text = ANY ((ARRAY['page'::character varying, 'post'::character varying])::text[])))
);


ALTER TABLE public.content_ OWNER TO postgres;

--
-- Name: content__id__seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.content_ ALTER COLUMN id_ ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.content__id__seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_ (
    id_ integer NOT NULL,
    username_ character varying(50) NOT NULL,
    email_ character varying(100) NOT NULL,
    password_ character varying(100) NOT NULL,
    is_admin_ boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users_ OWNER TO postgres;

--
-- Name: users__id__seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users_ ALTER COLUMN id_ ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users__id__seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: content_; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.content_ (id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_, user_id_) FROM stdin;
15	Some title	Some subtitle in here	some-title	Some body content to fill up in here.	2023-11-12 16:21:23.91082	2023-11-12 16:21:23.91082	post	1
14	Some title updated	Updated subtitle	some-title-updated	Updated body content in here.	2023-11-12 16:21:10.843748	2023-11-14 06:55:00.426472	page	1
\.


--
-- Data for Name: users_; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users_ (id_, username_, email_, password_, is_admin_) FROM stdin;
1	nixpig	test@nixpig.dev	foobarbaz	t
\.


--
-- Name: content__id__seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.content__id__seq', 15, true);


--
-- Name: users__id__seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users__id__seq', 1, true);


--
-- Name: content_ content__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.content_
    ADD CONSTRAINT content__pkey PRIMARY KEY (id_);


--
-- Name: users_ users__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_
    ADD CONSTRAINT users__pkey PRIMARY KEY (id_);


--
-- Name: content_ content__user_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.content_
    ADD CONSTRAINT content__user_id__fkey FOREIGN KEY (user_id_) REFERENCES public.users_(id_);


--
-- PostgreSQL database dump complete
--

