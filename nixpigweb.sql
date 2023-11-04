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
-- Name: category_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category_ (
    id integer NOT NULL,
    name_ character varying(50),
    template_id_ integer
);


ALTER TABLE public.category_ OWNER TO postgres;

--
-- Name: category__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.category_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.category__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: category_meta_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category_meta_ (
    id integer NOT NULL,
    category_id_ integer,
    meta_id_ integer
);


ALTER TABLE public.category_meta_ OWNER TO postgres;

--
-- Name: category_meta__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.category_meta_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.category_meta__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: config_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.config_ (
    id integer NOT NULL,
    name_ character varying(50),
    value_ character varying(255)
);


ALTER TABLE public.config_ OWNER TO postgres;

--
-- Name: config__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.config_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.config__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: meta_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.meta_ (
    id integer NOT NULL,
    name_ character varying(50),
    value_ character varying(255)
);


ALTER TABLE public.meta_ OWNER TO postgres;

--
-- Name: meta__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.meta_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.meta__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: page_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.page_ (
    id integer NOT NULL,
    title_ character varying(255) DEFAULT ''::character varying NOT NULL,
    body_ text DEFAULT ''::text NOT NULL,
    slug_ character varying(255) NOT NULL,
    status_ character varying(10) DEFAULT 'draft'::character varying NOT NULL,
    created_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id_ integer,
    category_id_ integer
);


ALTER TABLE public.page_ OWNER TO postgres;

--
-- Name: page__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.page_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.page__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: page_meta_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.page_meta_ (
    id integer NOT NULL,
    page_id_ integer,
    meta_id_ integer
);


ALTER TABLE public.page_meta_ OWNER TO postgres;

--
-- Name: page_meta__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.page_meta_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.page_meta__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: post_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_ (
    id integer NOT NULL,
    title_ character varying(255) DEFAULT ''::character varying NOT NULL,
    subtitle_ character varying(255) DEFAULT ''::character varying NOT NULL,
    body_ text DEFAULT ''::text NOT NULL,
    slug_ character varying(255) DEFAULT ''::character varying NOT NULL,
    status_ character varying(10) DEFAULT 'draft'::character varying NOT NULL,
    created_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id_ integer,
    category_id_ integer
);


ALTER TABLE public.post_ OWNER TO postgres;

--
-- Name: post__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.post_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.post__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: template_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.template_ (
    id integer NOT NULL,
    name_ character varying(50),
    tmpl_ character varying(255)
);


ALTER TABLE public.template_ OWNER TO postgres;

--
-- Name: template__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.template_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.template__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_ (
    id integer NOT NULL,
    username_ character varying(50) NOT NULL,
    email_ character varying(50) NOT NULL,
    is_admin_ boolean DEFAULT false NOT NULL,
    password_ character varying(255) NOT NULL,
    registered_at_ time without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_login_ time without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    role_ character varying(10) DEFAULT 'reader'::character varying NOT NULL,
    profile_ character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.user_ OWNER TO postgres;

--
-- Name: user__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.user_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_meta_; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_meta_ (
    id integer NOT NULL,
    user_id_ integer,
    meta_id_ integer
);


ALTER TABLE public.user_meta_ OWNER TO postgres;

--
-- Name: user_meta__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.user_meta_ ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_meta__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: category_ category__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_
    ADD CONSTRAINT category__pkey PRIMARY KEY (id);


--
-- Name: category_meta_ category_meta__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_meta_
    ADD CONSTRAINT category_meta__pkey PRIMARY KEY (id);


--
-- Name: config_ config__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.config_
    ADD CONSTRAINT config__pkey PRIMARY KEY (id);


--
-- Name: meta_ meta__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meta_
    ADD CONSTRAINT meta__pkey PRIMARY KEY (id);


--
-- Name: page_ page__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_
    ADD CONSTRAINT page__pkey PRIMARY KEY (id);


--
-- Name: page_meta_ page_meta__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_meta_
    ADD CONSTRAINT page_meta__pkey PRIMARY KEY (id);


--
-- Name: post_ post__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_
    ADD CONSTRAINT post__pkey PRIMARY KEY (id);


--
-- Name: template_ template__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.template_
    ADD CONSTRAINT template__pkey PRIMARY KEY (id);


--
-- Name: user_ user__email__key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_
    ADD CONSTRAINT user__email__key UNIQUE (email_);


--
-- Name: user_ user__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_
    ADD CONSTRAINT user__pkey PRIMARY KEY (id);


--
-- Name: user_ user__username__key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_
    ADD CONSTRAINT user__username__key UNIQUE (username_);


--
-- Name: user_meta_ user_meta__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_meta_
    ADD CONSTRAINT user_meta__pkey PRIMARY KEY (id);


--
-- Name: category_ category__template_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_
    ADD CONSTRAINT category__template_id__fkey FOREIGN KEY (template_id_) REFERENCES public.template_(id);


--
-- Name: category_meta_ category_meta__category_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_meta_
    ADD CONSTRAINT category_meta__category_id__fkey FOREIGN KEY (category_id_) REFERENCES public.category_(id);


--
-- Name: category_meta_ category_meta__meta_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_meta_
    ADD CONSTRAINT category_meta__meta_id__fkey FOREIGN KEY (meta_id_) REFERENCES public.meta_(id);


--
-- Name: page_ page__category_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_
    ADD CONSTRAINT page__category_id__fkey FOREIGN KEY (category_id_) REFERENCES public.category_(id);


--
-- Name: page_ page__user_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_
    ADD CONSTRAINT page__user_id__fkey FOREIGN KEY (user_id_) REFERENCES public.user_(id);


--
-- Name: page_meta_ page_meta__meta_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_meta_
    ADD CONSTRAINT page_meta__meta_id__fkey FOREIGN KEY (meta_id_) REFERENCES public.meta_(id);


--
-- Name: page_meta_ page_meta__page_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.page_meta_
    ADD CONSTRAINT page_meta__page_id__fkey FOREIGN KEY (page_id_) REFERENCES public.page_(id);


--
-- Name: post_ post__category_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_
    ADD CONSTRAINT post__category_id__fkey FOREIGN KEY (category_id_) REFERENCES public.category_(id);


--
-- Name: post_ post__user_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_
    ADD CONSTRAINT post__user_id__fkey FOREIGN KEY (user_id_) REFERENCES public.user_(id);


--
-- Name: user_meta_ user_meta__meta_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_meta_
    ADD CONSTRAINT user_meta__meta_id__fkey FOREIGN KEY (meta_id_) REFERENCES public.meta_(id);


--
-- Name: user_meta_ user_meta__user_id__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_meta_
    ADD CONSTRAINT user_meta__user_id__fkey FOREIGN KEY (user_id_) REFERENCES public.user_(id);


--
-- PostgreSQL database dump complete
--

