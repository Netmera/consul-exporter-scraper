<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Consul Exporter Scraper</h3>

  <p align="center">
    Consul Exporter Scraper is a tool that automatically discovers exporters running on a machine and registers them with Consul for easy service discovery.
  </p>
</div>

## Getting Started

Follow these steps to get started with using `consul-exporter-scraper`.

### Prerequisites

The following prerequisites are required for the project to be used:

- `curl` command
- A user account with `sudo` privileges

# Exporter Configuration

This configuration file (`exporter.yaml`) is used to define the exporters to be scraped by the Prometheus Consul Exporter. You can specify the details of each exporter, including its name, port, and export type.

- **consuladdresses**: The addresses of the Consul services to connect to.

- **exporters**: A list of exporters to be scraped. Each exporter object should contain the following fields:
  - `name`: The name of the exporter.
  - `port`: The port number on which the exporter is running.
  - `exporttype`: The type of service being exported.

To make changes to the exporter list or add/remove exporters, you can modify the `exporter.yaml` file located at `/etc/consul-exporter-scraper/exporter.yaml`.

Example `exporter.yaml` file:

```yaml
{
    "consuladdresses": ["your_consul_address_1", "your_consul_address_2"],
    "exporters": [
        {"name": "Mongo Exporter", "port": 9216, "exporttype": "mongodb"},
        {"name": "Postgresql Exporter", "port": 9187, "exporttype": "postgresql"},
        {"name": "Kubernetes Cert Exporter", "port": 9117, "exporttype": "kubernetes"},
        {"name": "Nginx Log Exporter", "port": 4040, "exporttype": "nginx"},
        {"name": "Nginx Exporter", "port": 9113, "exporttype": "nginx"},
        {"name": "Kafka Exporter", "port": 7072, "exporttype": "kafka"},
        {"name": "Kafka Consumer Group Exporter", "port": 9093, "exporttype": "kafka"},
        {"name": "Cassandra Exporter", "port": 9999, "exporttype": "cassandra"},
        {"name": "Blackbox Exporter", "port": 9115, "exporttype": "blackbox"},
        {"name": "Node Exporter", "port": 9100, "exporttype": "node"}
    ]
}
```
### Installation

#### Follow these steps to build and run the project:
 
1. Clone this repository:

   ```bash
   git clone https://github.com/Netmera/consul-exporter-scraper.git
   cd consul-exporter-scraper
  
2. Install dependencies:

   ```bash
      go mod tidy
      go mod download
    ```

3. Build the project
   ```bash
     env GOOS=linux GOARCH=amd64 go build .
    ```

4. Run the built binary:

   ```bash
     ./consul-exporter-scraper -environment=<environment>
    ```

#### Installation with install.sh

1. Clone this repository:

   ```bash
   git clone https://github.com/Netmera/consul-exporter-scraper.git
   cd consul-exporter-scraper
  
2. Run the install script:

   ```bash
      ./install.sh <environment>
    ```

#### Installation with Ansible

1. Clone this repository:

   ```bash
   git clone https://github.com/Netmera/consul-exporter-scraper.git
   cd consul-exporter-scraper
    ```

2. Ensure you have Ansible installed on your local machine.

3. Adjust the variables in default/main.yaml according to your environment:
    ```bash   
      ---
      consul_exporter_scraper_consul_adress: "http://consul:8500"
      consul_exporter_scraper_env: "production"
      consul_exporter_scraper_version: "v0.0.1"
   ```

4. Run the Ansible playbook:
    ```bash   
   ansible-playbook -i your_inventory_file install.yml
   ```


<!-- LICENSE -->
## License

Distributed under the APACHE-2.0 License. See `LICENSE` for more information.


[contributors-shield]: https://img.shields.io/github/contributors/Netmera/consul-exporter-scraper?style=for-the-badge
[contributors-url]: https://github.com/Netmera/consul-exporter-scraper/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Netmera/consul-exporter-scraper?style=for-the-badge
[forks-url]: https://github.com/Netmera/consul-exporter-scraper/network/members
[stars-shield]: https://img.shields.io/github/stars/Netmera/consul-exporter-scraper?style=for-the-badge
[stars-url]: https://github.com/Netmera/consul-exporter-scraper/stargazers
[issues-shield]: https://img.shields.io/github/issues/Netmera/consul-exporter-scraper?style=for-the-badge
[issues-url]: https://github.com/Netmera/consul-exporter-scraper/issues
[license-shield]: https://img.shields.io/github/license/Netmera/consul-exporter-scraper?style=for-the-badge
[license-url]: https://github.com/Netmera/consul-exporter-scraper/blob/main/LICENSE