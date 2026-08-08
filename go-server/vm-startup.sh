#!/bin/bash
# VM Initialization Script for FilmyFly
# This script runs on first boot of the GCP VM

set -e

echo "Starting VM initialization..."

# Update system
apt-get update
apt-get upgrade -y

# Install Go
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
rm go1.21.6.linux-amd64.tar.gz

# Install PostgreSQL
apt-get install -y postgresql postgresql-contrib

# Install Nginx
apt-get install -y nginx

# Install Git
apt-get install -y git

# Create application directory
mkdir -p /opt/filmyfly
mkdir -p /var/log/filmyfly

# Create www-data user if not exists
id -u www-data &>/dev/null || useradd -r -s /bin/false www-data

# Set permissions
chown -R www-data:www-data /opt/filmyfly
chown -R www-data:www-data /var/log/filmyfly

echo "VM initialization complete!"
