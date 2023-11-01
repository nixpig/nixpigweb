## TODO

## API

### Config

- `GET` - get config for specific item
- `PATCH` - update specific config option
- `DELETE` - delete a specific config option

| Endpoint      | Verb     | Description                       |
| ------------- | -------- | --------------------------------- |
| `/config`     | `GET`    | Get all config options            |
| `/config`     | `POST`   | Create new config option          |
| `/config/:id` | `GET`    | Get specific config using `id`    |
| `/config/:id` | `PATCH`  | Update specific config using `id` |
| `/config/:id` | `DELETE` | Delete specific config using `id` |

### Meta

| Endpoint    | Verb     | Description                 |
| ----------- | -------- | --------------------------- |
| `/meta`     | `GET`    | Get all meta items          |
| `/meta`     | `POST`   | Create a new meta item      |
| `/meta/:id` | `GET`    | Get a specific meta item    |
| `/meta/:id` | `PATCH`  | Update a specific meta item |
| `/meta/:id` | `DELETE` | Delete a specific meta item |

### Post

| Endpoint    | Verb     | Description            |
| ----------- | -------- | ---------------------- |
| `/post`     | `GET`    | Get all posts          |
| `/post`     | `POST`   | Create a new post      |
| `/post/:id` | `GET`    | Get a specific post    |
| `/post/:id` | `PATCH`  | Update a specific post |
| `/post/:id` | `DELETE` | Delete a specific post |

### User

| Endpoint      | Verb     | Description            |
| ------------- | -------- | ---------------------- |
| `/user`       | `GET`    | Get all users          |
| `/user`       | `POST`   | Create a new user      |
| `/user/:id`   | `GET`    | Get a specific user    |
| `/user/:id`   | `PATCH`  | Update a specific user |
| `/user/:id`   | `DELETE` | Delete a specific user |
| `/user/login` | `POST`   | Login user             |
