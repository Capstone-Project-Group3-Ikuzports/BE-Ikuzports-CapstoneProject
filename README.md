
<!-- ABOUT THE PROJECT -->
<br/>
<div align="center">
<!--  mengarah ke repo  -->
  <a href="https://fe-ikuzports-capstone-project-1.vercel.app/">
    <img src="images/logo.png" width="200px">
  </a>

  <h3 align="center">Ikuzports Apps</h3>

  <p align="center">
    Capstone Project Alterra Academy Backend Batch 13
    <br />
      <a href="https://github.com/Capstone-Project-Group3-Ikuzports/BE-Ikuzports-CapstoneProject/issues">Report Bug</a>
    ·
       <a href="https://github.com/Capstone-Project-Group3-Ikuzports/BE-Ikuzports-CapstoneProject/issues">Request Feature</a>
    <br />
  </p>
</div>

## 💻 About The Project

Ikuzports is an apps that helps users to carry out hobbies with other people, join an event and join the club. Many of us have a hobby but sometimes we do not have a partner or team to carry it out. This apps will connect every user with other users to carry out hobbies and even create their own club. 

## Feature in Ikuzports

  <!--- feature USER
   --->
<div>
      <details>
<summary>🙎 Users</summary>

In users feature, there are a feature to Create, Read, Update, Delete, Login for users here. You can sign in with your google account as well. 

<div>

- CRUD User
- User can login 
- User can see the event list and the detail of each event
- User can search event by name, sports category and city
- User can join a event, create an event and see their own event
- User can see the club list and the detail of each club
- User can search club by name, sports category and city
- User can join a club, create a club and see their own club
- User can see the activity and galery after join a club
- User can chat with another member of club
- User can see the prouct list in store
- User can search product by name, product category and city
- User can sell and buy product in store
- User can see the history of their transaction

</details>
<div>
      <details>
<summary>👱‍♂️ Guest</summary>

<div>
  
- Guest can see all event in homepage
- Guest can see the list of club
- Guest can see the products in store
- Guest can register 

</details>

<div>
      <details>
<summary>👨‍💻 Users (Club Owner) </summary>

<div>
  
- Club owner can approve member request from user
- Club owner can delete member from their club
- Club owner can create, update or delete club activity
- Club owner can create, update, or delete image in their club

</details>

### 🛠 &nbsp;Build App & Database

<a href="https://code.visualstudio.com/">![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white) </a>
<a href="https://www.mysql.com/">![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white) </a>
<a href="https://go.dev/"> ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)</a>
<a href="https://s3.console.aws.amazon.com/">![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)</a>
<a href="https://hub.docker.com/">![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)</a>
<a href="https://www.cloudflare.com/">![Cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)</a>
<a href="https://jwt.io/">![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)</a>
<a href="https://swagger.io/">![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)</a>
<a href="https://www.postman.com/">![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)</a>
<a href="https://console.cloud.google.com/">![Google Cloud Platform](https://img.shields.io/badge/Google%20Cloud%20-%234285F4.svg?&style=for-the-badge&logo=google-cloud&logoColor=white)</a>
<a href="https://git-scm.com/">![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)</a>
<a href="https://github.com/">![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)</a>

## 🗃️ ERD

<img src="images/ERD_Ikuzports.png">

## Open Api

 <a href="https://app.swaggerhub.com/apis-docs/RAMADINAAINIRIZKI/Ikuzports/1.0.0#/"> Click here to redirect to swagger </a>

## Getting Started

### Installation
1. Clone the repo
   ```bash
   git clone https://github.com/Capstone-Project-Group3-Ikuzports/BE-Ikuzports-CapstoneProject.git
   ```
2. Move to project directory:
    ```bash
    cd BE-Ikuzports-CapstoneProject
    ```
3. Create your environment file
    ```bash
    touch .env
    ```
4. Fill environment file
   ```bash
    export DB_USERNAME="YOUR DB USERNAME"
    export DB_PASSWORD="YOUR DB PASSWORD"
    export DB_HOST="YOUR HOST ADDRESS"
    export DB_PORT="3306"
    export DB_NAME="YOUR DB NAME" 
    export SERVER_PORT="YOUR PORT"
    export MIDTRANS_CLIENT_KEY="YOUR MIDTRANS CLIENT KEY"
    export MIDTRANS_SERVER_KEY="YOUR MIDTRANS SERVER KEY"
    export AWS_REGION="YOUR AWS REGION"
    export ACCESS_KEY_IAM="YOUR ACCESS KEY IAM"
    export SECRET_KEY_IAM="YOUR SECRET KEY IAM"
    export AWS_BUCKET_NAME="YOUR AWS BUCKET NAME"
    export EMAIL_FROM="YOUR EMAIL ADDRESS FOR SENDING EMAIL"
    export EMAIL_PASSWORD="YOUR EMAIL PASSWORD FOR SMPT SEND EMAIL"
    export JWT_SECRET="YOU STRING SECRET FOR JWT"
    export GOOGLE_OAUTH_CLIENT_ID="YOUR GOOGLE OAUTH CLIENT ID FOR LOGIN OAUTH"
    " >> .env
   ```
3. Run Project
    ```bash
    source .env && go run main.go
    ```

Contributor :
<br>

[![GitHub Achmad Qizwini](https://img.shields.io/github/followers/achmadqizwini?label=Aqiz&style=social)](https://github.com/achmadqizwini)


[![GitHub Ramadina Ainirizki](https://img.shields.io/github/followers/ramadinaainirizqi?label=Ramadina&style=social)](https://github.com/ramadinaainirizqi)
<br>

Mentor :
<br>

[![GitHub Fakhry Ihsan](https://img.shields.io/github/followers/iffakhry?label=FakhryIkhsan&Ikhsan&style=social)](https://github.com/iffakhry)


<p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: 2023 | Alterra Academy BE13 :fire:</p>
</h3>