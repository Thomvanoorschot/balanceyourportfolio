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
  KINDE_ISSUER_URL     = "https://balanceyourportfolio-test.eu.kinde.com"
  KINDE_CLIENT_ID      = "d24c661bb681479f83f49765d6dafbd3"
  KINDE_CLIENT_SECRET  = "IlmsubGMBZ4vUrGoUVzKOza2cmzi1aiY2YcHjYww8Zo9mPfBzCDa"
  KINDE_AUTH_WITH_PKCE = "false"
  PROTO_FILES_LOCATION = "/app/proto/main.proto"
}
