# Receipt-processor
A simple receipt processor

# Setting Up a Go Project

To set up a Go project from GitHub on another system, follow these steps:

## 1. Install Go

Ensure Go is installed on the system. Download it from the [official Go website](https://go.dev/doc/install).

## 2. Clone the Repository

Clone the repository to your local system with:

```sh
git clone repository_url
a. SSH : git clone git@github.com:Rtannu/receipt-processor.git
b. HTTPS : git clone https://github.com/Rtannu/receipt-processor.git
```
## 3. Navigate to the Project Directory

After cloning, navigate to the project directory:

```sh
cd path_to_your_project
ex: cd /your_project_path/receipt-processor
```

## 4. Download Dependencies

With the `go.mod` and `go.sum` files in your project, you can download all the necessary dependencies by running the following command in your terminal:

```sh
go mod download
```

## 5. Build the Project (Optional)

To compile the project and make sure everything is set up correctly, you can build it using the following command in your terminal:

```sh
go build
```
## 6. Run the Project

You can run the project directly by executing the following command in your terminal:

```sh
go run .
```
Server starts to run on url:

```sh
http://localhost:8080
```
## 7. Curl of Path: /receipts/process 

```sh
curl --location 'http://localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

## 7. Curl of Path: /receipts/{id}/points 

```sh
curl --location 'http://localhost:8080/receipts/de1ac000-9bad-4e12-95d5-f1f3aee00578/points'
```
## 8. Postman Screenshot of Request and Response of above Curl

a. Screenshot of Path: /receipts/process 
![image](https://github.com/Rtannu/receipt-processor/assets/40348395/14faf11c-6dbe-4268-8b35-94df1bde8a7b)


b. Screenshot of Path: /receipts/{id}/points
![image](https://github.com/Rtannu/receipt-processor/assets/40348395/22ae8600-238b-4872-a0c1-95f991fd14d1)

 


