🧱 WPXHost Tenant Provisioning Microservice - Development Plan

🎯 Objective

Build a Go-based microservice that provisions WordPress tenants using a shared WordPress core, with per-tenant configuration and content directories. The goal is to automate local WordPress setup for multi-tenant hosting.

⸻

✅ Phase 1: MVP Provisioning

🛠️ Milestone: Provision Local WordPress Site
•	Create CreateTenant() logic
•	Make tenant folder
•	Create wp-content/ folder
•	Generate stub files: index.php, wp-config.php, site-config.php
•	Load config using .env with godotenv
•	Support TENANT_PATH variable
•	Run via go run ./cmd/server and test with curl

⏭️ Next Steps
•	Add SHARED_CORE_PATH to .env and config
•	Create symlink: tenant/wp → SHARED_CORE_PATH
•	Add real index.php and wp-config.php templates
•	Configure NGINX (or Caddy) to point example.local to tenant folder
•	Add tenant to /etc/hosts
•	Run WP install via browser

⸻

📦 Phase 2: Expanded Provisioning
•	Add symlinks for:
•	wp-content/themes/
•	wp-content/plugins/
•	wp-content/mu-plugins/
•	Add support for shared media assets
•	Run wp-cli core install from provisioning service (optional)
•	Add tenants.json registry (or SQLite DB)

⸻

🔐 Phase 3: Admin & Security
•	Add GET /tenants endpoint
•	Add basic API key auth or HTTP basic auth
•	Validate input (existing tenant, invalid domain, etc)

⸻

🚀 Phase 4: Docker & Production
•	Dockerize provisioning service
•	Add Docker volume for tenants
•	Define .env.production for VPS
•	Move provisioned sites to /srv/wpxhost/tenants/

⸻

🧪 Testing Checklist
•	Provision via curl
•	Files appear under correct folder
•	wp symlink created properly
•	WordPress loads from browser
•	CSS/images load properly
•	Local hostname (example.local) works via /etc/hosts