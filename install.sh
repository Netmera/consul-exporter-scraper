#!/bin/bash

# Determine the latest version
LATEST_VERSION=$(curl -sL https://github.com/Netmera/prometheus-consul-exporter/releases/latest | grep -o 'tag/v[0-9]\+\.[0-9]\+\.[0-9]\+' | awk -F '/' '{print $2}' | uniq)

# If LATEST_VERSION is empty, print error message and exit the script
if [ -z "$LATEST_VERSION" ]; then
    echo "Error: Latest version not found."
    exit 1
fi

# Construct the download URL
DOWNLOAD_URL="https://github.com/Netmera/prometheus-consul-exporter/releases/download/${LATEST_VERSION}/prometheus-consul-exporter_${LATEST_VERSION}_linux_amd64.tar.gz"

# Perform the download
echo "Download URL: $DOWNLOAD_URL"
wget -nv -O /tmp/prometheus-consul-exporter.tar.gz "${DOWNLOAD_URL}"

# Check for errors during the download by examining the exit code
if [ $? -ne 0 ]; then
    echo "Error: File could not be downloaded."
    exit 1
fi

# Extract the downloaded file
tar xzf /tmp/prometheus-consul-exporter.tar.gz -C /usr/local/bin --strip-components=1

# Check if the file copying operation was successful
if [ $? -ne 0 ]; then
    echo "Error: File could not be copied."
    exit 1
fi

# Create the systemd service file
cat <<EOF > /etc/systemd/system/prometheus-consul-exporter.service
[Unit]
Description=Prometheus Consul Exporter
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/prometheus-consul-exporter
Restart=always

[Install]
WantedBy=multi-user.target
EOF

# Create the systemd timer file
cat <<EOF > /etc/systemd/system/prometheus-consul-exporter.timer
[Unit]
Description=Run prometheus-consul-exporter every hour

[Timer]
OnCalendar=hourly
Persistent=true

[Install]
WantedBy=timers.target
EOF

# Reload systemd, enable the service, and start it
systemctl daemon-reload

# Remove temporary files
rm -rf /tmp/prometheus-consul-exporter.tar.gz
rm -rf /tmp/prometheus-consul-exporter
rm -rf /tmp/configs
rm -rf /tmp/README.md
rm -rf /tmp/LICENSE

# Additional steps
# Create /etc/prometheus-consul-exporter directory and copy exporter.yaml from GitHub
mkdir -p /etc/prometheus-consul-exporter
mv /tmp/configs/exporter.yaml /etc/prometheus-consul-exporter/exporter.yaml
mkdir -p /var/log/prometheus-consul-exporter

# Display success message
echo "Installation successful. Prometheus Consul Exporter has been installed and configured."
echo "Please edit the configuration file /etc/prometheus-consul-exporter/exporter.yaml to suit your environment."
echo "Then start the service using the following commands:"
echo "systemctl start prometheus-consul-exporter"
echo "systemctl enable prometheus-consul-exporter"
echo "systemctl start prometheus-consul-exporter.timer"
echo "systemctl enable prometheus-consul-exporter.timer"