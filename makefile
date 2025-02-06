# Add these to your existing Makefile

# Database migrations
migrate-up:
	migrate -path migrations -database "postgresql://comin_owner:Ye5rfjcIB7FX@ep-flat-shadow-a8onelva.eastus2.azure.neon.tech/comin?sslmode=require" up

migrate-down:
	migrate -path migrations -database "postgresql://comin_owner:Ye5rfjcIB7FX@ep-flat-shadow-a8onelva.eastus2.azure.neon.tech/comin?sslmode=require" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Force specific version
migrate-force:
	migrate -path migrations -database "postgresql://comin_owner:Ye5rfjcIB7FX@ep-flat-shadow-a8onelva.eastus2.azure.neon.tech/comin?sslmode=require" force $(version)

# Go to specific version
migrate-goto:
	migrate -path migrations -database "postgresql://comin_owner:Ye5rfjcIB7FX@ep-flat-shadow-a8onelva.eastus2.azure.neon.tech/comin?sslmode=require" goto $(version)

# Show current version
migrate-version:
	migrate -path migrations -database "postgresql://comin_owner:Ye5rfjcIB7FX@ep-flat-shadow-a8onelva.eastus2.azure.neon.tech/comin?sslmode=require" version