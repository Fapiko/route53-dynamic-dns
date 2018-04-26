# Route53 Dynamic DNS

This is a dynamic DNS client written in Go to store the external address
of the machine it's running on in Route53.

## Configuration

/etc/route53-dynamic-dns/config.yaml
```yaml
hostname: pc.yourdomain.com
aws-access-key-id: XXXXXXXXX
aws-secret-access-key: XXXXXXXXX 
```

## TODO
  * Add license info and lint .deb release
