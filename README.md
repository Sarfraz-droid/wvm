## WVM - Web Version Manager

reverse-proxy made with go that handles your website having different versions.

This app let's you manage and handle multiple website versions as per requirements.

## Overall Architecture (in the demo)

![image](https://github.com/user-attachments/assets/9368fc04-3e1f-4fa3-9380-bac7f34aa421)

## Code Structure & Steps

1. Fork the repo by `git clone https://github.com/Sarfraz-droid/wvm.git`.
2. You can find `sample.config.yaml` in the root directory.
3. Expose `sample.config.yaml` file path to environment variable `MVN_CONFIG_FILE` for the service to fetch.
4. Modify the middleware at `overrides/middleware/main.go` as per the requirements.

## CLI
### Updating the config

You can update the config by running the CLI command `go run cmd/main.go reload`.
