#!/bin/bash

# Set environment
environment=$1

# Determine the latest version
LATEST_VERSION=$(curl -sL https://github.com/Netmera/prometheus-consul-exporter/releases/latest | grep -o 'tag/v[0-9]\+\.[0-9]\+\.[0-9]\+' | awk -F '/' '{print $2}' | uniq)

# If LATEST_VERSION is empty, print error message and exit the script
if [ -z "$LATEST_VERSION" ]; then
    echo -e "\e[91mError:\e[0m Latest version not found."
    exit 1
fi

# Construct the download URL
DOWNLOAD_URL="https://github.com/Netmera/prometheus-consul-exporter/releases/download/${LATEST_VERSION}/prometheus-consul-exporter_${LATEST_VERSION}_linux_amd64.tar.gz"

# Perform the download
echo -e "\e[94mStep 1:\e[0m Downloading Prometheus Consul Exporter...\n"
echo "Download URL: $DOWNLOAD_URL"
wget -nv -O /tmp/prometheus-consul-exporter.tar.gz "${DOWNLOAD_URL}"

# Check for errors during the download by examining the exit code
if [ $? -ne 0 ]; then
    echo -e "\e[91mError:\e[0m File could not be downloaded."
    exit 1
fi

# Extract the downloaded file
echo -e "\n\e[94mStep 2:\e[0m Extracting downloaded files...\n"
tar xzf /tmp/prometheus-consul-exporter.tar.gz -C /tmp

# Check if the file copying operation was successful
if [ $? -ne 0 ]; then
    echo -e "\e[91mError:\e[0m File could not be copied."
    exit 1
fi

# Create the systemd service file
echo -e "\n\e[94mStep 3:\e[0m Creating systemd service file...\n"
cat <<EOF > /etc/systemd/system/prometheus-consul-exporter.service
[Unit]
Description=Prometheus Consul Exporter
After=network.target
Wants=prometheus-consul-exporter.timer

[Service]
Type=oneshot
ExecStart=/usr/local/bin/prometheus-consul-exporter -environment=${environment}

[Install]
WantedBy=multi-user.target
EOF

# Create the systemd timer file
echo -e "\n\e[94mStep 4:\e[0m Creating systemd timer file...\n"
cat <<EOF > /etc/systemd/system/prometheus-consul-exporter.timer
[Unit]
Description=Prometheus Consul Exporter Timer
Requires=prometheus-consul-exporter.service

[Timer]
Unit=prometheus-consul-exporter.service
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
# Create /etc/prometheus-consul-exporter directory and copy exporter.yaml from GitHub
mkdir -p /etc/prometheus-consul-exporter
mv /tmp/configs/exporter.yaml /etc/prometheus-consul-exporter/exporter.yaml
mv /tmp/prometheus-consul-exporter /usr/local/bin/prometheus-consul-exporter
mkdir -p /var/log/prometheus-consul-exporter


# Remove temporary files
echo -e "\n\e[94mStep 7:\e[0m Cleaning up temporary files...\n"
rm -rf /tmp/prometheus-consul-exporter.tar.gz
rm -rf /tmp/prometheus-consul-exporter
rm -rf /tmp/configs
rm -rf /tmp/README.md
rm -rf /tmp/LICENSE

# Display success message
echo -e "\e[92mInstallation successful.\e[0m Prometheus Consul Exporter has been installed and configured."
echo -e "\e[91mPlease edit the configuration file /etc/prometheus-consul-exporter/exporter.yaml to suit your environment."
echo -e "\e[91mThen start the service using the following commands:"
echo -e "\e[91msystemctl start prometheus-consul-exporter"
echo -e "\e[91msystemctl enable prometheus-consul-exporter"
echo -e "\e[91msystemctl start prometheus-consul-exporter.timer"
echo -e "\e[91msystemctl enable prometheus-consul-exporter.timer"
