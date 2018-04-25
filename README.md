# Route53 Dynamic DNS

This is a dynamic DNS client written in Go to store the external address
of the machine it's running on in Route53.

## Configuration

~/.route53-dynamic-dns/config.yaml
```yaml
hostname: pc.yourdomain.com
aws-access-key-id: XXXXXXXXX
aws-secret-access-key: XXXXXXXXX 
```

## TODO
  * Daemonize it
  * Package into debian package
