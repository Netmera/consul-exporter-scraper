#!/bin/bash

# Set environment
environment=$1

# Determine the latest version
LATEST_VERSION=$(curl -sL https://github.com/Netmera/consul-exporter-scraper/releases/latest | grep -o 'tag/v[0-9]\+\.[0-9]\+\.[0-9]\+' | awk -F '/' '{print $2}' | uniq)

# If LATEST_VERSION is empty, print error message and exit the script
if [ -z "$LATEST_VERSION" ]; then
    echo -e "\e[91mError:\e[0m Latest version not found."
    exit 1
fi

# Construct the download URL
DOWNLOAD_URL="https://github.com/Netmera/consul-exporter-scraper/releases/download/${LATEST_VERSION}/consul-exporter-scraper_${LATEST_VERSION}_linux_amd64.tar.gz"

# Perform the download
echo -e "\e[94mStep 1:\e[0m Downloading Consul Exporter Scraper...\n"
echo "Download URL: $DOWNLOAD_URL"
wget -nv -O /tmp/consul-exporter-scraper.tar.gz "${DOWNLOAD_URL}"

# Check for errors during the download by examining the exit code
if [ $? -ne 0 ]; then
    echo -e "\e[91mError:\e[0m File could not be downloaded."
    exit 1
fi

# Extract the downloaded file
echo -e "\n\e[94mStep 2:\e[0m Extracting downloaded files...\n"
tar xzf /tmp/consul-exporter-scraper.tar.gz -C /tmp

# Check if the file copying operation was successful
if [ $? -ne 0 ]; then
    echo -e "\e[91mError:\e[0m File could not be copied."
    exit 1
fi

# Create the systemd service file
echo -e "\n\e[94mStep 3:\e[0m Creating systemd service file...\n"
cat <<EOF > /etc/systemd/system/consul-exporter-scraper.service
[Unit]
Description=Consul Exporter Scraper
After=network.target
Wants=consul-exporter-scraper.timer

[Service]
Type=oneshot
ExecStart=/usr/local/bin/consul-exporter-scraper -environment=${environment}

[Install]
WantedBy=multi-user.target
EOF

# Create the systemd timer file
echo -e "\n\e[94mStep 4:\e[0m Creating systemd timer file...\n"
cat <<EOF > /etc/systemd/system/consul-exporter-scraper.timer
[Unit]
Description=Consul Exporter Scraper Timer
Requires=consul-exporter-scraper.service

[Timer]
Unit=consul-exporter-scraper.service
OnCalendar=*-*-* *:00:00
Persistent=true

[Install]
WantedBy=timers.target
EOF

# Reload systemd, enable the service, and start it
echo -e "\n\e[94mStep 5:\e[0m Reloading systemd and enabling services...\n"
systemctl daemon-reload

# Additional steps
echo -e "\n\e[94mStep 6:\e[0m Additional steps...\n"
# Create /etc/consul-exporter-scraper directory and copy exporter.yaml from GitHub
mkdir -p /etc/consul-exporter-scraper
mv /tmp/configs/exporter.yaml /etc/consul-exporter-scraper/exporter.yaml
mv /tmp/consul-exporter-scraper /usr/local/bin/consul-exporter-scraper


# Remove temporary files
echo -e "\n\e[94mStep 7:\e[0m Cleaning up temporary files...\n"
rm -rf /tmp/consul-exporter-scraper.tar.gz
rm -rf /tmp/consul-exporter-scraper
rm -rf /tmp/configs
rm -rf /tmp/README.md
rm -rf /tmp/LICENSE

# Display success message
echo -e "\e[92mInstallation successful.\e[0m Consul Exporter Scraper has been installed and configured."
echo -e "\e[91mPlease edit the configuration file /etc/consul-exporter-scraper/exporter.yaml to suit your environment."
echo -e "\e[91mThen start the service using the following commands:"
echo -e "\e[91msystemctl start consul-exporter-scraper"
echo -e "\e[91msystemctl enable consul-exporter-scraper"
echo -e "\e[91msystemctl start consul-exporter-scraper.timer"
echo -e "\e[91msystemctl enable consul-exporter-scraper.timer"
