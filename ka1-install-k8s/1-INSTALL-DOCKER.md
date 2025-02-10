
# Install Docker

Note: Since K8S uses `containerd`, you don't need Docker unless you want to build images, learn Docker, or for using it directly like docker-compose

If you configure K8S to use Docker one might install it in Ubuntu like so:

```bash
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install docker-ce
```