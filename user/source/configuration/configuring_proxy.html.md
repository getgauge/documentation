---
title: "Configuring Proxy"
---

# Configuring Proxy

Gauge connects to internet for downloading plugins, templates, etc. If you are behind a proxy, you will have to configure the proxy settings so that Gauge connects to internet via the proxy server.

### Without Authentication
If authentication is not required, set the environment variable `HTTP_PROXY` to proxy server URL.

```
export HTTP_PROXY=http://10.0.2.2:5678
```

### With Authentication
If authentication is not required, set the environment variable `HTTP_PROXY` to proxy server URL along with the credentials.

```
export HTTP_PROXY=http://username:password@10.0.2.2:5678
```
