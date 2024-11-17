# learning-cards-app

## Running the Application

### Docker

1. **Clone the repository**

   ```bash
   git clone https://github.com/mzmbq/learning-cards-app.git
   cd learning-cards-app
   ```

2. **Setup the environemnt**

   ```bash
   mv .env.example .env
   ```

3. **Setup SSL certificates**

   a. Development: Generate with openssl

   ```bash
   openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/nginx.key -out certs/nginx.crt
   ```

   b. Production: Using Let's Encrypt certificates

   ```bash
   export DOMAIN=yourdomain.com

   sudo certbot certonly --standalone -d $DOMAIN
   sudo cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem certs/nginx.crt
   sudo cp /etc/letsencrypt/live/$DOMAIN/privkey.pem certs/nginx.key
   ```

4. **Start the application**

   ```bash
   docker compose up
   ```

5. **Access at [http://localhost:8080](http://localhost:8080)**
