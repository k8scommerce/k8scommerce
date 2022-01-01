# Cert Setup

```sh
docker run -it ubuntu bash
apt update && apt install openssl -y
```

### https://letsencrypt.org/docs/certificates-for-localhost/
```sh
openssl req -x509 -out local-registry.crt -keyout local-registry.key \
-newkey rsa:2048 -nodes -sha256 \
-subj '/CN=local-registry' -extensions EXT -config <( \
printf "[dn]\nCN=local-registry\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:local-registry\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

copy the crt and key files 

install the crt file in Keychain Access

double click the crt file in the keychain access and trust all access


```sh
docker run -it --entrypoint htpasswd -v $PWD/auth:/auth -w /auth registry:2.7.0 -Bbc /auth/htpasswd admin password
```

```sh
kubectl create secret generic tls-certs --from-file=certFile=certs/local-registry.crt --from-file=keyFile=certs/local-registry.key
```

username: admin
password: password
```sh
kubectl create secret -n flux-system docker-registry registry-credentials \
--docker-server="local-registry:5000" \
--docker-username=admin \
--docker-password='$2y$05$s813FqUpaomywVVXhi16Dei4YVBmo3TBe1k2Iz.j9uryic9W9AWte'
```


flux create secret tls local-registry-cert --ca-file ./certs/local-registry.pem -n flux-system




openssl x509 -inform der -in certificate.cer -out certificate.pem


# Brew Install

```sh
A CA file has been bootstrapped using certificates from the system
keychain. To add additional certificates, place .pem files in
  /usr/local/etc/openssl@3/certs

and run
  /usr/local/opt/openssl@3/bin/c_rehash

openssl@3 is keg-only, which means it was not symlinked into /usr/local,
because macOS provides LibreSSL.

If you need to have openssl@3 first in your PATH, run:
  echo 'export PATH="/usr/local/opt/openssl@3/bin:$PATH"' >> ~/.zshrc

For compilers to find openssl@3 you may need to set:
  export LDFLAGS="-L/usr/local/opt/openssl@3/lib"
  export CPPFLAGS="-I/usr/local/opt/openssl@3/include"

For pkg-config to find openssl@3 you may need to set:
  export PKG_CONFIG_PATH="/usr/local/opt/openssl@3/lib/pkgconfig"
```