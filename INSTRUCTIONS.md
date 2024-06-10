1. CREATE YOUR CODE //
2. MOVE TO YOUR CODES DIRECTORY
3. RUN YOUR CODE : docker run --rm -v "$PWD":/app -w /app golang:latest go run main.go && go version

--> to check Container GoLang Version --> docker run --rm -v "$PWD":/app -w /app golang:latest go version 

Explanation

--> docker run: This command runs a new container.
--> --rm: This option automatically removes the container after it exits.
--> -v "$PWD":/app: This option mounts your current directory ($PWD) to the /app directory inside the container.
--> -w /app: This option sets the working directory inside the container to /app.
--> golang
--> : This specifies the Go Docker image to use.
--> go run main.go: This is the command that runs your Go code inside the container.


#build an image that has terraform installed
docker build -t terraform-alpine .

#run image just built and run terraform code
 docker run --rm -it -v "$PWD":/terraform_code terraform-alpine

#inside image RUN terraform apply -auto-approve 

#Write the output to a json file
terraform output -json > output.json


WRITE INISDE THE CONTAINER:

1.terraform_output=$(terraform output -json)
//echo $terraform_output
2.value_array=$(echo "$terraform_output" | jq -r '.summarized_cidr_blocks.value')
//echo $value_array
3.echo $value_array > ips.json
//cat ips.json
4.jq . ips.json > tmp.json && mv tmp.json ips.json




2.ips=$(echo "$terraform_output" | jq -r '.summarized_cidr_blocks.value[]')
//echo $ips
3.formatted_ips=$(echo "$ips" | jq -cR '[inputs]')
//echo formatted_ips
4.echo "$formatted_ips" > ips.json
5.jq . ips.json > tmp.json && mv tmp.json ips.json

value_array=$(echo "$terraform_output" | jq -r '.summarized_cidr_blocks.value')

# Print the extracted value array
echo "$value_array"