# InqzCLI

InqzCLI is a command-line interface tool that allows you to perform various tasks related to IP addresses and geolocation. Here are some of the things you can do with InqzCLI:

## Get your public IP address

To get your public IP address, simply run the following command:

```
inqzcli ip
```

This will make a request to the [ipify API](https://www.ipify.org/) and return your public IP address.

## Get geolocation information for an IP address

To get geolocation information for an IP address, run the following command:

```
inqzcli geolocate <ip_address>
```

Replace `<ip_address>` with the IP address you want to geolocate. This will make a request to the [ipstack API](https://ipstack.com/) and return information such as the country, city, and latitude/longitude coordinates of the IP address.

## Start an HTTP server to display your public IP address

To start an HTTP server that displays your public IP address, run the following command:

```
inqzcli serve <port>
```

Replace `<port>` with the port number you want to use for the server (default is 8080). This will start an HTTP server that listens on the specified port and displays your public IP address when accessed.

## Get help

To get help with InqzCLI, run the following command:

```
inqzcli help
```

This will display a list of available commands and their usage.

## Installation

To install InqzCLI, run the following command:

```
go get github.com/inqz/inqzcli
```

This will download and install the latest version of InqzCLI from GitHub.

## License

InqzCLI is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the LICENSE file for more information.

---