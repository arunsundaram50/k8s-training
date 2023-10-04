## Docker
1) Create a docker image using alpine as base
2) Install curl if it is not already there
3) Add the get-google.com.sh script 

## K8s
4) Create a k8s cronjob using the above image
5) Capture the output to google.txt

## Problem: 
There will be multiple google.txt files sitting in each container 

## Solution:
If we map the host's /Users/USERNAME/Documents/data -> /data 