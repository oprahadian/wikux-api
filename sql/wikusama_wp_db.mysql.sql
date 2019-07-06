set sql_mode='';

create table wiku_commentmeta
(
  meta_id    bigint unsigned auto_increment
    primary key,
  comment_id bigint unsigned default 0 not null,
  meta_key   varchar(255)              null,
  meta_value longtext                  null
)
  engine = MyISAM
  charset = utf8;

create index comment_id
  on wiku_commentmeta (comment_id);

create index meta_key
  on wiku_commentmeta (meta_key);

create table wiku_comments
(
  comment_ID           bigint unsigned auto_increment
    primary key,
  comment_post_ID      bigint unsigned default 0                     not null,
  comment_author       tinytext                                      not null,
  comment_author_email varchar(100)    default ''                    not null,
  comment_author_url   varchar(200)    default ''                    not null,
  comment_author_IP    varchar(100)    default ''                    not null,
  comment_date         datetime        default '0000-00-00 00:00:00' not null,
  comment_date_gmt     datetime        default '0000-00-00 00:00:00' not null,
  comment_content      text                                          not null,
  comment_karma        int             default 0                     not null,
  comment_approved     varchar(20)     default '1'                   not null,
  comment_agent        varchar(255)    default ''                    not null,
  comment_type         varchar(20)     default ''                    not null,
  comment_parent       bigint unsigned default 0                     not null,
  user_id              bigint unsigned default 0                     not null
)
  engine = MyISAM
  charset = utf8;

create index comment_approved_date_gmt
  on wiku_comments (comment_approved, comment_date_gmt);

create index comment_author_email
  on wiku_comments (comment_author_email);

create index comment_date_gmt
  on wiku_comments (comment_date_gmt);

create index comment_parent
  on wiku_comments (comment_parent);

create index comment_post_ID
  on wiku_comments (comment_post_ID);

create table wiku_layerslider
(
  id             int(10) auto_increment
    primary key,
  author         int(10)    default 0 not null,
  name           varchar(100)         not null,
  slug           varchar(100)         not null,
  data           mediumtext           not null,
  date_c         int(10)              not null,
  date_m         int(10)              not null,
  flag_hidden    tinyint(1) default 0 not null,
  flag_deleted   tinyint(1) default 0 not null,
  schedule_start int(10)    default 0 not null,
  schedule_end   int(10)    default 0 not null,
  flag_popup     tinyint(1) default 0 not null
)
  collate = utf8mb4_unicode_520_ci;

create table wiku_layerslider_revisions
(
  id        int(10) auto_increment
    primary key,
  slider_id int(10)           not null,
  author    int(10) default 0 not null,
  data      mediumtext        not null,
  date_c    int(10)           not null
)
  collate = utf8mb4_unicode_520_ci;

create table wiku_links
(
  link_id          bigint unsigned auto_increment
    primary key,
  link_url         varchar(255)    default ''                    not null,
  link_name        varchar(255)    default ''                    not null,
  link_image       varchar(255)    default ''                    not null,
  link_target      varchar(25)     default ''                    not null,
  link_description varchar(255)    default ''                    not null,
  link_visible     varchar(20)     default 'Y'                   not null,
  link_owner       bigint unsigned default 1                     not null,
  link_rating      int             default 0                     not null,
  link_updated     datetime        default '0000-00-00 00:00:00' not null,
  link_rel         varchar(255)    default ''                    not null,
  link_notes       mediumtext                                    not null,
  link_rss         varchar(255)    default ''                    not null
)
  engine = MyISAM
  charset = utf8;

create index link_visible
  on wiku_links (link_visible);

create table wiku_nextend2_image_storage
(
  id    int auto_increment
    primary key,
  hash  varchar(32) not null,
  image text        not null,
  value mediumtext  not null,
  constraint hash
    unique (hash)
)
  charset = utf8;

create table wiku_nextend2_section_storage
(
  id           int auto_increment
    primary key,
  application  varchar(20)   not null,
  section      varchar(128)  not null,
  referencekey varchar(128)  not null,
  value        mediumtext    not null,
  `system`     int default 0 not null,
  editable     int default 1 not null
)
  charset = utf8;

create index application
  on wiku_nextend2_section_storage (application, section, referencekey);

create index application_2
  on wiku_nextend2_section_storage (application, section);

create table wiku_nextend2_smartslider3_generators
(
  id      int auto_increment
    primary key,
  `group` varchar(254) not null,
  type    varchar(254) not null,
  params  text         not null
)
  charset = utf8;

create table wiku_nextend2_smartslider3_sliders
(
  id        int auto_increment
    primary key,
  alias     varchar(255)  null,
  title     varchar(100)  not null,
  type      varchar(30)   not null,
  params    mediumtext    not null,
  time      datetime      not null,
  thumbnail varchar(255)  not null,
  ordering  int default 0 not null
)
  charset = utf8;

create table wiku_nextend2_smartslider3_sliders_xref
(
  group_id  int           not null,
  slider_id int           not null,
  ordering  int default 0 not null,
  primary key (group_id, slider_id)
)
  charset = utf8;

create table wiku_nextend2_smartslider3_slides
(
  id           int auto_increment
    primary key,
  title        varchar(200) not null,
  slider       int          not null,
  publish_up   datetime     not null,
  publish_down datetime     not null,
  published    tinyint(1)   not null,
  first        int          not null,
  slide        longtext     null,
  description  text         not null,
  thumbnail    varchar(255) not null,
  params       text         not null,
  ordering     int          not null,
  generator_id int          not null
)
  charset = utf8;

create table wiku_options
(
  option_id    bigint unsigned auto_increment
    primary key,
  option_name  varchar(191) default ''    not null,
  option_value longtext                   not null,
  autoload     varchar(20)  default 'yes' not null,
  constraint option_name
    unique (option_name)
)
  engine = MyISAM
  charset = utf8;

create table wiku_participants_database
(
  id            int(6) auto_increment
    primary key,
  private_id    varchar(6)                          null,
  first_name    tinytext                            null,
  last_name     tinytext                            null,
  address       tinytext                            null,
  city          tinytext                            null,
  state         tinytext                            null,
  country       tinytext                            null,
  zip           tinytext                            null,
  phone         tinytext                            null,
  email         tinytext                            null,
  mailing_list  text                                null,
  photo         text                                null,
  website       text                                null,
  interests     text                                null,
  approved      text                                null,
  date_recorded timestamp default CURRENT_TIMESTAMP not null,
  date_updated  timestamp                           null,
  last_accessed timestamp                           null
)
  engine = MyISAM
  collate = utf8_unicode_ci;

create table wiku_participants_database_fields
(
  id                 int(3) auto_increment
    primary key,
  `order`            int(3)     default 0 not null,
  name               varchar(64)          not null,
  title              tinytext             not null,
  `default`          tinytext             null,
  `group`            varchar(64)          not null,
  help_text          text                 null,
  form_element       tinytext             null,
  `values`           longtext             null,
  options            longtext             null,
  attributes         text                 null,
  validation         tinytext             null,
  validation_message text                 null,
  display_column     int(3)     default 0 null,
  admin_column       int(3)     default 0 null,
  sortable           tinyint(1) default 0 null,
  CSV                tinyint(1) default 0 null,
  persistent         tinyint(1) default 0 null,
  signup             tinyint(1) default 0 null,
  readonly           tinyint(1) default 0 null,
  constraint name
    unique (name)
)
  engine = MyISAM
  collate = utf8_unicode_ci;

create index `group`
  on wiku_participants_database_fields (`group`);

create index `order`
  on wiku_participants_database_fields (`order`);

create table wiku_participants_database_groups
(
  id          int(3) auto_increment
    primary key,
  `order`     int(3)     default 0 not null,
  mode        varchar(64)          null,
  display     tinyint(1) default 1 null,
  admin       tinyint(1) default 0 not null,
  title       tinytext             not null,
  name        varchar(64)          not null,
  description text                 null,
  constraint name
    unique (name)
)
  engine = MyISAM
  collate = utf8_unicode_ci;

create table wiku_postmeta
(
  meta_id    bigint unsigned auto_increment
    primary key,
  post_id    bigint unsigned default 0 not null,
  meta_key   varchar(255)              null,
  meta_value longtext                  null
)
  engine = MyISAM
  charset = utf8;

create index meta_key
  on wiku_postmeta (meta_key);

create index post_id
  on wiku_postmeta (post_id);

create table wiku_posts
(
  ID                    bigint unsigned auto_increment
    primary key,
  post_author           bigint unsigned default 0                     not null,
  post_date             datetime        default '0000-00-00 00:00:00' not null,
  post_date_gmt         datetime        default '0000-00-00 00:00:00' not null,
  post_content          longtext                                      not null,
  post_title            text                                          not null,
  post_excerpt          text                                          not null,
  post_status           varchar(20)     default 'publish'             not null,
  comment_status        varchar(20)     default 'open'                not null,
  ping_status           varchar(20)     default 'open'                not null,
  post_password         varchar(255)    default ''                    not null,
  post_name             varchar(200)    default ''                    not null,
  to_ping               text                                          not null,
  pinged                text                                          not null,
  post_modified         datetime        default '0000-00-00 00:00:00' not null,
  post_modified_gmt     datetime        default '0000-00-00 00:00:00' not null,
  post_content_filtered longtext                                      not null,
  post_parent           bigint unsigned default 0                     not null,
  guid                  varchar(255)    default ''                    not null,
  menu_order            int             default 0                     not null,
  post_type             varchar(20)     default 'post'                not null,
  post_mime_type        varchar(100)    default ''                    not null,
  comment_count         bigint          default 0                     not null
)
  engine = MyISAM
  charset = utf8;

create index post_author
  on wiku_posts (post_author);

create index post_name
  on wiku_posts (post_name);

create index post_parent
  on wiku_posts (post_parent);

create index type_status_date
  on wiku_posts (post_type, post_status, post_date, ID);

create table wiku_revslider_css
(
  id       int(9) auto_increment,
  handle   text     not null,
  settings longtext null,
  hover    longtext null,
  params   longtext not null,
  advanced longtext null,
  constraint id
    unique (id)
)
  collate = utf8mb4_unicode_520_ci;

alter table wiku_revslider_css
  add primary key (id);

create table wiku_revslider_layer_animations
(
  id       int(9) auto_increment,
  handle   text not null,
  params   text not null,
  settings text null,
  constraint id
    unique (id)
)
  collate = utf8mb4_unicode_520_ci;

alter table wiku_revslider_layer_animations
  add primary key (id);

create table wiku_revslider_navigations
(
  id       int(9) auto_increment,
  name     varchar(191) not null,
  handle   varchar(191) not null,
  css      longtext     not null,
  markup   longtext     not null,
  settings longtext     null,
  constraint id
    unique (id)
);

create table wiku_revslider_settings
(
  id      int(9) auto_increment
    primary key,
  general text not null,
  params  text not null
)
  collate = utf8mb4_unicode_520_ci;

create table wiku_revslider_sliders
(
  id       int(9) auto_increment,
  title    tinytext                not null,
  alias    tinytext                null,
  params   longtext                not null,
  settings text                    null,
  type     varchar(191) default '' not null,
  constraint id
    unique (id)
)
  collate = utf8mb4_unicode_520_ci;

alter table wiku_revslider_sliders
  add primary key (id);

create table wiku_revslider_slides
(
  id          int(9) auto_increment,
  slider_id   int(9)   not null,
  slide_order int      not null,
  params      longtext not null,
  layers      longtext not null,
  settings    text     not null,
  constraint id
    unique (id)
)
  collate = utf8mb4_unicode_520_ci;

alter table wiku_revslider_slides
  add primary key (id);

create table wiku_revslider_static_slides
(
  id        int(9) auto_increment,
  slider_id int(9)   not null,
  params    longtext not null,
  layers    longtext not null,
  settings  text     not null,
  constraint id
    unique (id)
)
  collate = utf8mb4_unicode_520_ci;

alter table wiku_revslider_static_slides
  add primary key (id);

create table wiku_sm_sessions
(
  session_id     bigint unsigned auto_increment,
  session_key    char(32)        not null
    primary key,
  session_value  longtext        not null,
  session_expiry bigint unsigned not null,
  constraint session_id
    unique (session_id)
)
  engine = MyISAM
  collate = utf8mb4_unicode_520_ci;

create table wiku_term_relationships
(
  object_id        bigint unsigned default 0 not null,
  term_taxonomy_id bigint unsigned default 0 not null,
  term_order       int             default 0 not null,
  primary key (object_id, term_taxonomy_id)
)
  engine = MyISAM
  charset = utf8;

create index term_taxonomy_id
  on wiku_term_relationships (term_taxonomy_id);

create table wiku_term_taxonomy
(
  term_taxonomy_id bigint unsigned auto_increment
    primary key,
  term_id          bigint unsigned default 0  not null,
  taxonomy         varchar(32)     default '' not null,
  description      longtext                   not null,
  parent           bigint unsigned default 0  not null,
  count            bigint          default 0  not null,
  constraint term_id_taxonomy
    unique (term_id, taxonomy)
)
  engine = MyISAM
  charset = utf8;

create index taxonomy
  on wiku_term_taxonomy (taxonomy);

create table wiku_termmeta
(
  meta_id    bigint unsigned auto_increment
    primary key,
  term_id    bigint unsigned default 0 not null,
  meta_key   varchar(255)              null,
  meta_value longtext                  null
)
  engine = MyISAM
  charset = utf8;

create index meta_key
  on wiku_termmeta (meta_key);

create index term_id
  on wiku_termmeta (term_id);

create table wiku_terms
(
  term_id    bigint unsigned auto_increment
    primary key,
  name       varchar(200) default '' not null,
  slug       varchar(200) default '' not null,
  term_group bigint(10)   default 0  not null
)
  engine = MyISAM
  charset = utf8;

create index name
  on wiku_terms (name);

create index slug
  on wiku_terms (slug);

create table wiku_usermeta
(
  umeta_id   bigint unsigned auto_increment
    primary key,
  user_id    bigint unsigned default 0 not null,
  meta_key   varchar(255)              null,
  meta_value longtext                  null
)
  engine = MyISAM
  charset = utf8;

create index meta_key
  on wiku_usermeta (meta_key);

create index user_id
  on wiku_usermeta (user_id);

create table wiku_users
(
  ID                  bigint unsigned auto_increment
    primary key,
  user_login          varchar(60)  default ''                    not null,
  user_pass           varchar(255) default ''                    not null,
  user_nicename       varchar(50)  default ''                    not null,
  user_email          varchar(100) default ''                    not null,
  user_url            varchar(100) default ''                    not null,
  user_registered     datetime     default '0000-00-00 00:00:00' not null,
  user_activation_key varchar(255) default ''                    not null,
  user_status         int          default 0                     not null,
  display_name        varchar(250) default ''                    not null
)
  engine = MyISAM
  charset = utf8;

create index user_email
  on wiku_users (user_email);

create index user_login_key
  on wiku_users (user_login);

create index user_nicename
  on wiku_users (user_nicename);

create table wiku_wpmm_subscribers
(
  id_subscriber bigint auto_increment
    primary key,
  email         varchar(50) not null,
  insert_date   datetime    not null
)
  engine = MyISAM
  charset = utf8;

