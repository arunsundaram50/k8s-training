## Volume mount gotcha
Since our Pod is running on a Minikube node, the `minikube mount /home/asundaram/apps/docker-training/kd12-cronjob-minikube-volume-mount/data` directory should be present on the Minikube virtual machine itself, not on our local host filesystem.

Minikube runs the Kubernetes nodes inside VMs, and hence "host filesystem" refers to the filesystem of the VM, not our actual local machine where you are running Minikube.

To access files from our local filesystem inside our pod running in Minikube, you should use the `minikube mount` command to mount a local host directory into the Minikube VM. Here's how you can do this:

1. Open a new terminal window.
   
2. Run the following command to mount our local directory into the Minikube VM:

   ```sh
   minikube mount /home/asundaram/apps/docker-training/kd12-cronjob-minikube-volume-mount/data:/home/asundaram/apps/docker-training/kd12-cronjob-minikube-volume-mount/data
   ```

3. In our pod configuration YAML file, continue using the same `hostPath` configuration as before, since from the perspective of the Pod running in Minikube, it's now referring to a path inside the Minikube VM.

4. Apply our Pod configuration to create the Pod:

   ```sh
   kubectl apply -f mypod.yaml
   ```

5. Verify that the mount is working correctly by creating files in the mounted directory from our local system and checking if they appear in the `/data` directory inside our Pod.

Keep the terminal with the `minikube mount` command running. It maintains the mount and if you stop it, the mount will be disconnected. Once it is running, you should be able to see files created on our local machine appear in the Minikube VM, and vice-versa.