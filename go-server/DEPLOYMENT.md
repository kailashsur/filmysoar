# Quick Deployment Guide - Existing GCP Free Tier VM

## Your Existing VM Setup
- **Type**: e2-micro (Always Free)
- **RAM**: 1GB
- **CPU**: 0.25-1 vCPU (shared)
- **Disk**: 30GB

## Step-by-Step Deployment

### 1. Connect to Your VM
```bash
# Replace with your actual instance name and zone
gcloud compute ssh YOUR_INSTANCE_NAME --zone=YOUR_ZONE
```

### 2. Create Swap Space (IMPORTANT for 1GB RAM)
```bash
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

### 3. Install Dependencies
```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Go
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install PostgreSQL, Nginx, Git, and Certbot
sudo apt install -y postgresql postgresql-contrib nginx git certbot
```

### 4. Setup PostgreSQL
```bash
# Create database
sudo -u postgres psql
CREATE DATABASE filmyfly;
CREATE USER filmyfly_user WITH ENCRYPTED PASSWORD 'YOUR_PASSWORD';
GRANT ALL PRIVILEGES ON DATABASE filmyfly TO filmyfly_user;
\q

# Upload and restore backup (from local machine)
gcloud compute scp c:\Users\Kailash\Desktop\filmyfly_go\database\supabase_full_backup.dump YOUR_INSTANCE_NAME:~/backup.dump --zone=YOUR_ZONE

# Restore (on VM)
sudo -u postgres pg_restore -d filmyfly -v ~/backup.dump --no-owner --no-acl
```

### 5. Setup Application
```bash
# Create directories
sudo mkdir -p /opt/filmyfly /var/log/filmyfly
sudo chown $USER:$USER /opt/filmyfly

# Clone repository
cd /opt/filmyfly
git clone https://github.com/YOUR_USERNAME/filmyfly-go-fiber.git .

# Create .env file
nano /opt/filmyfly/.env
# Add your environment variables (see full guide)
# Set NODE_ENV=production so admin session cookies are marked Secure

# Build application
go build -o filmyfly cmd/server/main.go
```

### 6. Create Systemd Service
```bash
# Copy service file
sudo cp filmyfly.service /etc/systemd/system/

# Start service
sudo systemctl daemon-reload
sudo systemctl enable filmyfly
sudo systemctl start filmyfly
```

### 7. Configure Nginx
```bash
# Make sure DNS points filmyfly.work and www.filmyfly.work to this VM first.

# Get the TLS certificate
sudo systemctl stop nginx
sudo certbot certonly --standalone -d filmyfly.work -d www.filmyfly.work
sudo systemctl start nginx

# Copy nginx config
sudo cp nginx.conf /etc/nginx/sites-available/filmyfly
sudo ln -s /etc/nginx/sites-available/filmyfly /etc/nginx/sites-enabled/
sudo rm /etc/nginx/sites-enabled/default
sudo nginx -t
sudo systemctl restart nginx
```

### 8. Configure Firewall
```bash
# Add firewall rules (if not exist)
gcloud compute firewall-rules create allow-http --allow tcp:80 --target-tags http-server
gcloud compute firewall-rules create allow-https --allow tcp:443 --target-tags https-server
gcloud compute instances add-tags YOUR_INSTANCE_NAME --zone=YOUR_ZONE --tags=http-server,https-server
```

### 9. Setup Auto-Deployment
```bash
# Update cloudbuild.yaml with your instance name and zone
# Then create trigger
gcloud builds triggers create github \
    --repo-name=filmyfly-go-fiber \
    --repo-owner=YOUR_GITHUB_USERNAME \
    --branch-pattern="^main$" \
    --build-config=cloudbuild.yaml
```

### 10. Get Your URL
```bash
gcloud compute instances describe YOUR_INSTANCE_NAME --zone=YOUR_ZONE --format='get(networkInterfaces[0].accessConfigs[0].natIP)'
```

Visit: `https://filmyfly.work`

## Quick Commands

```bash
# Check status
sudo systemctl status filmyfly

# View logs
sudo journalctl -u filmyfly -f

# Restart app
sudo systemctl restart filmyfly

# Check memory
free -h
```

## Important Notes

✅ **Swap space is CRITICAL** - 1GB RAM is not enough without it
✅ **PostgreSQL is optimized** for low memory in the config
✅ **Memory limits** are set in systemd service (400MB max)
✅ **Auto-deployment** triggers on push to main branch

For detailed instructions, see the full deployment plan!
