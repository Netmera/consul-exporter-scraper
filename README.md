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

### Installation

#### Follow these steps to build and run the project:
 
1. Clone this repository:

   ```bash
   git clone git@github.com:Netmera/consul-exporter-scraper.git
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
   git clone git@github.com:Netmera/consul-exporter-scraper.git
   cd consul-exporter-scraper
  
2. Run the install script:

   ```bash
      ./install.sh <environment>
    ```

#### Installation with Ansible

1. Clone this repository:

   ```bash
   git clone git@github.com:Netmera/consul-exporter-scraper.git
   cd consul-exporter-scraper
    ```

2. Ensure you have Ansible installed on your local machine.

3. Adjust the variables in default/main.yaml according to your environment:
    ```bash   
      ---
      prometheus_consul_exporter_consul_adress: "http://consul:8500"
      prometheus_consul_exporter_env: "production"
      prometheus_consul_exporter_version: "v0.0.1"
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