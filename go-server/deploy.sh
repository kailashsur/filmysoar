#!/bin/bash
set -e

echo "🚀 Deploying FilmyFly to GCP..."

# Configuration
ZONE="us-central1-a"
INSTANCE="filmyfly-server"
APP_DIR="/opt/filmyfly"

# Build the application
echo "📦 Building application..."
GOOS=linux GOARCH=amd64 go build -o filmyfly cmd/server/main.go

# Copy binary to server
echo "📤 Uploading binary..."
gcloud compute scp filmyfly ${INSTANCE}:${APP_DIR}/ --zone=${ZONE}

# Copy views directory
echo "📤 Uploading views..."
gcloud compute scp --recurse views ${INSTANCE}:${APP_DIR}/ --zone=${ZONE}

# Copy public directory
echo "📤 Uploading public assets..."
gcloud compute scp --recurse public ${INSTANCE}:${APP_DIR}/ --zone=${ZONE}

# Restart application
echo "🔄 Restarting application..."
gcloud compute ssh ${INSTANCE} --zone=${ZONE} \
    --command="sudo systemctl restart filmyfly"

# Check status
echo "✅ Checking application status..."
gcloud compute ssh ${INSTANCE} --zone=${ZONE} \
    --command="sudo systemctl status filmyfly --no-pager"

echo "🎉 Deployment complete!"
echo "🌐 Access your application at: http://$(gcloud compute instances describe ${INSTANCE} --zone=${ZONE} --format='get(networkInterfaces[0].accessConfigs[0].natIP)')"
