# go-url-shortener
- Generate a short URL and a qrcode image.
- Redirect a short URL  to its original URL.

https://golang-url-shortener.onrender.com/

# Approach
Just a simplest way to generate short url.
To avoid data collisions, generate a random string and use a unique key constraint in the database.

# Todo
- Use redis to improve performance.
- Add shprt URL expired time.
- Change the way of generating short url.
