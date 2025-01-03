digital_ocean_token = "dop_v1_d2fe6b7829de88e8e4a14bc3aa400870310fa53a2ae297fbeeae6d99098c54d6"
environment         = "staging"
project_description = "The balanceyourportfolio project"
project_name        = "balanceyourportfolio"
frontend_domain     = "balanceyourportfolio.com"
letsencrypt_email   = "thomvanoorschot94@gmail.com"

// Backend
backend_component_name = "backend"
backend_image          = "registry.digitalocean.com/balanceyourportfolio/backend"
backend_port           = 8080
backend_vars = {
  DB_CONNECTION_STRING = "user=postgres.mlunrkdivnylgkxgudws password=shY74qI5zbbWhEMU dbname=postgres host=aws-0-eu-central-1.pooler.supabase.com"
}

// Frontend
frontend_component_name = "frontend"
frontend_image          = "registry.digitalocean.com/balanceyourportfolio/frontend"
frontend_port           = 3000
frontend_vars = {
  KINDE_ISSUER_URL               = "https://auth.balanceyourportfolio.com"
  KINDE_CLIENT_ID                = "f290c7662ce94d0ca5328c7567691db1"
  KINDE_CLIENT_SECRET            = "K2ltHxGiE5FsfIeoSIcpO4trNyFd0fh9wPL1XPQiadQifPgtGVG"
  KINDE_REDIRECT_URL             = "https://balanceyourportfolio.com/api/auth/kinde_callback"
  KINDE_POST_LOGOUT_REDIRECT_URL = "https://balanceyourportfolio.com/fund/overview"
  KINDE_POST_LOGIN_REDIRECT_URL  = "https://balanceyourportfolio.com/fund/overview"
  KINDE_AUTH_WITH_PKCE           = false
  PROTO_FILES_LOCATION           = "/app/proto/main.proto"
}
