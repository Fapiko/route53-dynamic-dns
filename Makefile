build-linux64:
	rm -rf target
	mkdir target
	GOARCH=amd64 GOOS=linux && go build -o target/route53-dynamic-dns github.com/fapiko/route53-dynamic-dns

deb: build-linux64
	cp -r build/deb target/
	mkdir -p target/deb/usr/sbin
	chmod 0755 target/route53-dynamic-dns
	mv target/route53-dynamic-dns target/deb/usr/sbin
	find ./target/deb -type d | xargs chmod 0755
	fakeroot dpkg-deb --build target/deb
	mv target/deb.deb target/route53-dynamic-dns_0.1.0_amd64.deb

test-deb: deb
	cp -r build/docker target/
	cp target/route53-dynamic-dns_0.1.0_amd64.deb target/docker/
	docker build -t route53-dynamic-dns target/docker

shell-test-deb:
	docker run -it route53-dynamic-dns /bin/bash
