
# TCP Port Scanner

A simple Go-based TCP port scanner that allows you to scan specific or all TCP ports on a given host. It identifies whether the port is open and attempts to extract the service type and version based on the banner received from the service running on the open port.

## Features

- **Scan Specific Port**: Check if a specific port is open on a target host and retrieve the service details.
- **Scan All Ports**: Optionally scan all ports (1-65535) on a target host.
- **Service Identification**: Identifies the service type and version if available in the banner.
- **Timeouts**: Supports connection and read timeouts for better control.

## Requirements

- Go 1.x or higher.

## Installation

1. Clone the repository or download the source code:

   ```bash
   git clone https://github.com/yourusername/tcp-port-scanner.git
   ```

2. Navigate to the project folder:

   ```bash
   cd tcp-port-scanner
   ```

3. Build the application:

   ```bash
   go build
   ```

## Usage

### 1. Scan a Specific Port

To scan a specific port on a given host, use the following command:

```bash
go run main.go -h <HOSTNAME> -p <PORT>
```

- `<HOSTNAME>`: The target IP address or hostname.
- `<PORT>`: The port number to scan.

Example:

```bash
go run main.go -h 192.168.1.1 -p 80
```

This will check if port 80 is open on the target `192.168.1.1`.

### 2. Scan All Ports

To scan all ports (1-65535) on a given host, use the following command:

```bash
go run main.go -h <HOSTNAME> -Ap
```

- `<HOSTNAME>`: The target IP address or hostname.

Example:

```bash
go run main.go -h 192.168.1.1 -Ap
```

This will scan all ports from 1 to 65535 on the target `192.168.1.1`.

## Example Output

### When scanning all ports:

```text
Scan report from 192.168.1.1
Port    Status    Service    Version
80      OPEN      HTTP       1.1
443     OPEN      HTTPS      1.2
6379    OPEN      Redis      7.0.15
```

### When scanning a specific port:

```text
<========================>
192.168.1.1:80
<========================>
Port 80 OPEN
<========================>
```

## Troubleshooting

If the program does not return results for some ports, make sure the target server's firewall allows incoming connections on those ports. You can verify this using other tools such as `nmap`.

### Issues you may encounter:

1. **No Banner or Unknown Service**: Sometimes the service banner may not be received, or the program might fail to decode it. You can check if the port is actually open using `nmap` to verify.
   
2. **Port Timeout**: If you're scanning a large range of ports, some ports might time out. You can try adjusting the timeout value in the code for longer waits on slow ports.

3. **Firewall or Security Software**: Firewalls or security software on the target host might block scan attempts. In such cases, adjust your scanning strategy or verify access with the system administrator.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
