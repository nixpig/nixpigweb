## TODO

1. Create `Admin` middleware which verifies the current user token has a `role` of `admin` before accessing `Config` routes.

## API

### Public (anybody can access)

### Protected (must be logged in to access)

### Private (must be logged in and admin role to access)

#### `/config`

- `GET` - get all config options
- `POST` - create new config option

#### `/config/:id`

- `GET` - get config for specific item
- `PATCH` - update specific config option
- `DELETE` - delete a specific config option
