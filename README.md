.
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.modTest
├── README.md
├── structure.txt
└── backend
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── config
    │   └── config.go
    ├── core
    │   ├── api
    │   │   ├── handlers
    │   │   │   ├── auth
    │   │   │   │   ├── login_handler.go
    │   │   │   │   └── register_handler.go
    │   │   │   ├── business
    │   │   │   │   ├── business_management_handler.go
    │   │   │   │   ├── business_profile_handler.go
    │   │   │   │   └── business_retrieval_handler.go
    │   │   │   ├── events
    │   │   │   │   ├── event_likes_handler.go
    │   │   │   │   ├── event_management_handler.go
    │   │   │   │   └── event_retrieval_handler.go
    │   │   │   ├── file
    │   │   │   │   ├── delete_file_handler.go
    │   │   │   │   ├── get_file_url_handler.go
    │   │   │   │   └── upload_file_handler.go
    │   │   │   └── user
    │   │   │       ├── user_management_handler.go
    │   │   │       ├── user_profile_handler.go
    │   │   │       └── user_retrieval_handler.go
    │   │   ├── middleware
    │   │   │   ├── auth_middleware.go
    │   │   │   ├── cors_middleware.go
    │   │   │   ├── logging_middleware.go
    │   │   │   ├── role_middleware.go
    │   │   │   └── upload_middleware.go
    │   │   ├── request
    │   │   │   ├── auth
    │   │   │   │   └── auth_request.go
    │   │   │   ├── business
    │   │   │   │   └── business_request.go
    │   │   │   ├── event
    │   │   │   │   └── event.request.go
    │   │   │   └── user
    │   │   │       └── user.request.go
    │   │   ├── response
    │   │   │   ├── response.go
    │   │   │   ├── auth
    │   │   │   │   └── auth_response.go
    │   │   │   ├── business
    │   │   │   │   └── business_response.go
    │   │   │   ├── event
    │   │   │   │   └── event.response.go
    │   │   │   └── user
    │   │   │       └── user_response.go
    │   │   ├── routes
    │   │   │   ├── routes.go
    │   │   │   ├── auth
    │   │   │   │   └── auth_routes.go
    │   │   │   ├── business
    │   │   │   │   └── business_routes.go
    │   │   │   ├── event
    │   │   │   │   └── event_routes.go
    │   │   │   ├── file
    │   │   │   │   └── file_routes.go
    │   │   │   ├── pay
    │   │   │   │   └── payment_routes.txt
    │   │   │   ├── ticket
    │   │   │   │   └── ticket_routes.txt
    │   │   │   └── user
    │   │   │       └── user_routes.go
    │   │   ├── models
    │   │   │   ├── association.go
    │   │   │   ├── association_membership.go
    │   │   │   ├── business_user.go
    │   │   │   ├── category.go
    │   │   │   ├── educational_institution.go
    │   │   │   ├── event.go
    │   │   │   ├── event_like.go
    │   │   │   ├── event_option.go
    │   │   │   ├── event_view.go
    │   │   │   ├── password_reset.go
    │   │   │   ├── payment.go
    │   │   │   ├── payment_transaction.go
    │   │   │   ├── point_history.go
    │   │   │   ├── role.go
    │   │   │   ├── school_membership.go
    │   │   │   ├── studibox_transaction.go
    │   │   │   ├── tag.go
    │   │   │   ├── ticket.go
    │   │   │   ├── user.go
    │   │   │   └── user_role.go
    │   │   ├── services
    │   │   │   ├── auth
    │   │   │   │   └── auth_service.go
    │   │   │   ├── business
    │   │   │   │   ├── business_management_service.go
    │   │   │   │   ├── business_profile_service.go
    │   │   │   │   ├── business_retrieval_service.go
    │   │   │   │   └── business_service.go
    │   │   │   ├── event
    │   │   │   │   ├── event_likes_service.go
    │   │   │   │   ├── event_management_service.go
    │   │   │   │   ├── event_retrieval_service.go
    │   │   │   │   └── event_service.go
    │   │   │   ├── storage
    │   │   │   │   ├── s3_storage.go
    │   │   │   │   └── storage_service.go
    │   │   │   └── user
    │   │   │       ├── user_management_service.go
    │   │   │       ├── user_profile_service.go
    │   │   │       ├── user_retrieval_service.go
    │   │   │       └── user_service.go
    │   │   ├── stores
    │   │   │   ├── business
    │   │   │   │   ├── association_membership_store.go
    │   │   │   │   ├── association_store.go
    │   │   │   │   ├── business_user_store.go
    │   │   │   │   ├── educational_institution_store.go
    │   │   │   │   └── school_membership_store.go
    │   │   │   ├── event
    │   │   │   │   ├── category_store.go
    │   │   │   │   ├── event_like_store.go
    │   │   │   │   ├── event_option_store.go
    │   │   │   │   ├── event_store.go
    │   │   │   │   ├── event_view_store.go
    │   │   │   │   └── tag_store.go
    │   │   │   ├── payment
    │   │   │   │   ├── payment_store.go
    │   │   │   │   ├── payment_transaction_store.go
    │   │   │   │   ├── point_history_store.go
    │   │   │   │   ├── studibox_transaction_store.go
    │   │   │   │   └── ticket_store.go
    │   │   │   ├── user
    │   │   │   │   ├── password_reset_store.go
    │   │   │   │   ├── role_store.go
    │   │   │   │   ├── user_role_store.go
    │   │   │   │   └── user_store.go
    │   │   ├── templates
    │   │   │   └── index.html
    │   │   └── utils
    │   │       ├── convert.go
    │   │       └── jwt.go
    ├── database
    │   └── db.go
    └── pkg
        ├── cache
        │   └── cache.go
        └── client
            └── client.go
