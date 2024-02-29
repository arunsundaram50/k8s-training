## When you are able to get into the container
To validate a certificate in a container from outside, you can use various tools and methods depending on your specific requirements. Here's a general approach using OpenSSL as an example:

1. **Access the Container**: Use `kubectl exec` to access the container running in the Kubernetes cluster. For example:
   ```sh
   kubectl exec -it <pod-name> -c <container-name> -- /bin/bash
   ```
   Replace `<pod-name>` and `<container-name>` with your actual pod and container names.

2. **Copy the Certificate**: If the certificate is stored in a file inside the container, you can use `kubectl cp` to copy it to your local machine for validation. For example:
   ```sh
   kubectl cp <pod-name>:/path/to/certificate.crt /local/path
   ```
   Replace `<pod-name>` with your actual pod name, `/path/to/certificate.crt` with the actual path inside the container, and `/local/path` with the path on your local machine where you want to save the file.

3. **Validate the Certificate**: Once you have the certificate file on your local machine, you can use OpenSSL to validate it. For example, to validate a certificate chain:
   ```sh
   openssl verify -CAfile /path/to/ca-cert.pem /local/path/certificate.crt
   ```
   Replace `/path/to/ca-cert.pem` with the path to your CA certificate file and `/local/path/certificate.crt` with the path to the certificate file you copied from the container.

This is a general approach, and the specific steps may vary depending on your environment and the tools available.

4. **Validate using a client resolving the hostname/IP match**: 
    ```sh
    curl --cacert /path/to/ca-cert.pem https://example.com --resolve example.com:443:127.0.0.1
    ```

## When you can't get into the container, but know where the cert is
It is possible to download the certificate from outside the container to validate it. Here's a general approach using `kubectl cp` to copy the certificate from a pod to your local machine for validation:

1. **Identify the Pod and Certificate Path**: First, identify the name of the pod containing the certificate and the path to the certificate file inside the pod. You can list all pods in the current namespace using `kubectl get pods`.

2. **Copy the Certificate to Your Local Machine**: Use `kubectl cp` to copy the certificate from the pod to your local machine. For example:
   ```sh
   kubectl cp <pod-name>:/path/to/certificate.crt /local/path
   ```
   Replace `<pod-name>` with the name of your pod, `/path/to/certificate.crt` with the path to the certificate file inside the pod, and `/local/path` with the path on your local machine where you want to save the file.

3. **Validate the Certificate**: Once you have the certificate file on your local machine, you can use OpenSSL or another tool to validate it. For example, to validate a certificate chain using OpenSSL:
   ```sh
   openssl verify -CAfile /path/to/ca-cert.pem /local/path/certificate.crt
   ```
   Replace `/path/to/ca-cert.pem` with the path to your CA certificate file and `/local/path/certificate.crt` with the path to the certificate file you copied from the pod.

By following these steps, you can download the certificate from a container running in a Kubernetes pod to your local machine and validate it without needing to access the container directly.

## When you don't know where the cert is, and you don't have access to the cluster
```sh
openssl s_client -showcerts -connect www.example.com:443 </dev/null 2>/dev/null | openssl x509 -outform PEM > /tmp/server.crt
```
OR using Python
```sh
python -c "import socket, ssl; cert = ssl.get_server_certificate(('www.example.com', 443)); open('server.crt', 'w').write(cert)"
```
And then run the validations on the cert
```sh
openssl verify -CAfile server.crt server.crt
```
OR
```sh
openssl x509 -in server.crt -text -noout
```
