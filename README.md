# Project Contents

## backend

- [go.dockerfile](backend/go.dockerfile)
- [go.mod](backend/go.mod)
- [go.sum](backend/go.sum)
- [main.go](backend/main.go)
- [main_test.go](backend/main_test.go)

### controllers

- [auth_controller.go](backend/controllers/auth_controller.go)
- [feedback_controller.go](backend/controllers/feedback_controller.go)
- [folder_controller.go](backend/controllers/folder_controller.go)
- [quote_controller.go](backend/controllers/quote_controller.go)
- [user_controller.go](backend/controllers/user_controller.go)

### database

- [db.go](backend/database/db.go)
- [errors.go](backend/database/errors.go)
- [init.go](backend/database/init.go)
- [quote_db.go](backend/database/quote_db.go)
- [user_db.go](backend/database/user_db.go)

### middleware

- [auth.go](backend/middleware/auth.go)
- [authenticate.go](backend/middleware/authenticate.go)
- [authorize.go](backend/middleware/authorize.go)
- [error_handling.go](backend/middleware/error_handling.go)
- [logging.go](backend/middleware/logging.go)
- [rate_limiting.go](backend/middleware/rate_limiting.go)
- [request_logging.go](backend/middleware/request_logging.go)
- [security_headers.go](backend/middleware/security_headers.go)
- [tracing.go](backend/middleware/tracing.go)

### migrations

- [001_create_users_table.go](backend/migrations/001_create_users_table.go)
- [002_create_categories_table.go](backend/migrations/002_create_categories_table.go)
- [003_create_tags_table.go](backend/migrations/003_create_tags_table.go)
- [004_create_quotes_table.go](backend/migrations/004_create_quotes_table.go)
- [005_create_feedback_table.go](backend/migrations/005_create_feedback_table.go)
- [006_create_literary_works_table.go](backend/migrations/006_create_literary_works_table.go)
- [007_create_folders_table.go](backend/migrations/007_create_folders_table.go)
- [008_create_likes_table.go](backend/migrations/008_create_likes_table.go)

### models

- [category.go](backend/models/category.go)
- [feedback.go](backend/models/feedback.go)
- [folder.go](backend/models/folder.go)
- [like.go](backend/models/like.go)
- [lw.go](backend/models/lw.go)
- [models.go](backend/models/models.go)
- [quote.go](backend/models/quote.go)
- [tag.go](backend/models/tag.go)
- [user.go](backend/models/user.go)

### routes

- [routes.go](backend/routes/routes.go)

### utils

- [jwt.go](backend/utils/jwt.go)
- [utils.go](backend/utils/utils.go)


